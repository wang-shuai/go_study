package emailops

import (
	"crypto/tls"
	"fmt"
	"github.com/olebedev/config"
	"gopkg.in/gomail.v2"
	"os"
	"strings"
	"pushCitiesToRedis/flog"
)

var (
	cfg *config.Config
)

func init() {
	var err error
	pd, err := os.Getwd()
	cfg, err = config.ParseYamlFile(pd + "/config/emailconfig.yaml")
	checkErr(err)
}

func SendMail(body string) {
	m := gomail.NewMessage()

	//设置host
	host, err := cfg.String("email.host")
	checkErr(err)
	port, err := cfg.Int("email.port")
	checkErr(err)
	//设置账户
	username, err := cfg.String("email.username")
	checkErr(err)
	pwd, err := cfg.String("email.password")
	checkErr(err)

	//设置from发件人
	from, err := cfg.String("email.from")
	checkErr(err)
	//设置to 收件人
	to, err := cfg.String("email.to")
	checkErr(err)
	tos := []string{"wangshuai@yxqiche.com"}
	if len(to) > 0 {
		tos = strings.Split(to, ",")
	}

	m.SetHeader("From", from)
	m.SetHeader("To", tos...)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "更新上牌城市Redis")
	m.SetBody("text/html", body)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(host, port, username, pwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("邮件发送成功！")
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		flog.Fatalf("发送邮件配置信息异常：%v",err)
	}
}
