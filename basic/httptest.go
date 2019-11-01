package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

//{"code":200,"data":{"cityId":201,"cityName":"北京","provinceId":2,"provinceName":"北京"},"message":"","duration":0}

type City struct {
	CityId       int    `json:"cityId"`
	CityName     string `json:"cityName"`
	ProvinceId   int    `json:"provinceId"`
	ProvinceName string `json:"provinceName"`
}

type Location struct {
	Code     int    `json:"code"`
	Data     City   `json:"data"`
	Message  string `json:"message"`
	Duration int    `json:"duration"`
}

func main() {
	//router := gin.Default()
	//
	//router.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "Hello World")
	//})
	//router.Run(":8000")

	//http://api.map.baidu.com/geocoder?output=xml&amp;location=39.93,116.32,&amp;key=4bc250f30872d677dab56e41766dbe5a

	cnt := 0;
	for i := 0; i < 1000; i++ {
		//latitude := strconv.FormatFloat(39.99, 'f', 2, 64)
		//longitude := strconv.FormatFloat(116.32, 'f', 2, 64)
		//fmt.Println(latitude,",",longitude)
		res, err := http.Get("https://apiwx.taoche.com/api/Generic/GetLocation") // ?latitude=" + latitude + "&longitude=" + longitude + "&t=" + strconv.Itoa(i))
		if err != nil {
			fmt.Printf("获取定位信息失败：%v \n\r", err)

			//res, err = http.Get("http://api.map.baidu.com/geocoder?output=xml&location="+latitude+","+longitude+"&key=4bc250f30872d677dab56e41766dbe5a")
			//if err !=nil {
			//	fmt.Println("百度定位出错",err)
			//}
			//body, _ := ioutil.ReadAll(res.Body)
			//fmt.Println(string(body))
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
		var loc Location
		json.Unmarshal(body, &loc)
		if (loc.Code == 200) {
			//fmt.Println("获取城市定位成功：", loc)
		} else {
			//fmt.Println("获取城市定位失败：", loc)
			fmt.Println("header:",res.Header,"\n body:",string(body))
			cnt++
		}
	}
	fmt.Println("错误总数：", cnt)

}
