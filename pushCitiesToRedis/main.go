package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"pushCitiesToRedis/redisops"
	//"strings"
	"encoding/json"
	"os"
	"os/signal"
	"pushCitiesToRedis/emailops"
	"syscall"
	"time"
	"github.com/olebedev/config"
	"pushCitiesToRedis/flog"
	"strings"
)

type cityinfo struct {
	CitySpell   string
	BuyCarCity  string
	LicenseCity string
}

func main() {
	base, err := os.Getwd()

	cfg, err := config.ParseYamlFile(base + "/config/logconfig.yaml")
	fname,_ := cfg.String("logname")
	fpath,_ := cfg.String("logpath")
	flevel,_:= cfg.Int("loglevel")
	flog.Init(fname,fpath,flevel)

	p := base + "/cities.xlsx"
	err = readFile(p)
	if err != nil {
		fmt.Printf("压数据错误：%v \n", err)
	}
	forever := make(chan os.Signal)
	go checkFile(p)

	signal.Notify(forever, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGKILL)

	<-forever
}

func readFile(path string) error {
	xlsx, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//fmt.Println(xlsx.GetSheetName(1))
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows(xlsx.GetSheetName(1))

	var data []cityinfo

	for i, row := range rows {
		if i == 0 {
			continue //首行标题
		}
		if len(strings.TrimSpace(row[0])) > 0 {
			data = append(data, cityinfo{row[0], row[1], row[2]})
		}
	}
	//fmt.Println("实体：", data)
	err = pushIntoRedis("license_cities", data)
	// 发送邮件
	var emailBody string
	if err != nil {
		emailBody = fmt.Sprintf("更新购车上牌城市报错：%v", err)
		flog.Errorf(emailBody)
	} else {
		emailBody = fmt.Sprintf("更新购车上牌城市完成，条数：%d\n", len(rows))
	}
	go emailops.SendMail(emailBody)
	return err
}

func pushIntoRedis(key string, data interface{}) error {
	conn := redisops.Pool.Get()
	defer conn.Close()

	val, err := json.Marshal(data)
	_, err = conn.Do("set", key, val)
	//_, err = conn.Do("EXPIRE", key, 2 * 60 * 60)  //2 hour   // 不需要过期，心跳去更新服务
	if err != nil {
		body :=	fmt.Sprintf("set key:%s val:%s 错误：%s", key, val, err)
		flog.Errorf(body)

		return err
	}
	flog.Infof(string(val))

	return nil
}

func checkFile(path string) {
	for {
		time.Sleep(30 * time.Minute) //休息30min 然后检查文件是否更新过
		fileInfo, _ := os.Stat(path)
		now := time.Now()
		filetime := fileInfo.ModTime()
		//小于30分钟 则更新内存
		if now.Sub(filetime) < 30*time.Minute {
			readFile(path)
		} else {
			flog.Infof("30分钟内，文件没有更新过")
		}
	}
}
