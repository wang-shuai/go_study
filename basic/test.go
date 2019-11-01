package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Tttt struct {
	A int
	B string
}

type Entity_Cacheucarbasicinfo_Full struct {

	// AlternateContact
	AlternateContact string
	// BrandID
	BrandID string `json:"BrandID int"`
	// BuyCarDate
	BuyCarDate string
	// CalcCarPrice
	CalcCarPrice string
	// CarCityId
	CarCityId string `json:"CarCityId int"`
	// CardCode
	CardCode string
	// CarDistrictId
	CarDistrictId string `json:"CarDistrictId int"`
	// CarDownTime
	CarDownTime string
	// CardRangeID
	CardRangeID string `json:"CarDistrictId int"`
	// CarID
	CarID int
	// CarProviceId
	CarProviceId int
	// CarPublishTime
	CarPublishTime string
	// CarSetting
	CarSetting string
	// 车辆来源1级分类（1：个人车源 2：车商通(商家)车源  3：一线通  4：其他）
	CarSource1L int
	// 车辆来源2级分类(1：个人自主上传  5：4S店  6：专业公司  9：其他商家(经纪人)  11：非注册个人用户)
	CarSource2L int
	// CarSourceStatus
	CarSourceStatus int
	// CarType
	CarType int
	// CheckedTime
	CheckedTime string
	// Color
	Color string
	// CompleteRate
	CompleteRate int
	// 联系方式
	Contact string
	// Coordinate
	Coordinate string
	// CreateTime
	CreateTime string
	// DisplayPrice
	DisplayPrice float64
	// DriveLicensePic
	DriveLicensePic string
	// DrivingMileage
	DrivingMileage int
	// ExamineExpireDate
	ExamineExpireDate string
	// Exhaust
	Exhaust string
	// ExhaustValue
	ExhaustValue float64
	// Fee
	Fee float64
	// FirstPicTrue
	FirstPicTrue int
	// GearBoxType
	GearBoxType int
	// InnerColor
	InnerColor string
	// InsuranceExpireDate
	InsuranceExpireDate string
	// IsAuthenticated
	IsAuthenticated int
	// IsCards
	IsCards int
	// IsChangePrice
	IsChangePrice int
	// IsDealerRecommend
	IsDealerRecommend int
	// 是否包含过户
	IsIncTransfer int
	// IsLicensed
	IsLicensed int
	// IsNeglect
	IsNeglect int
	// IsTransferToOld
	IsTransferToOld int
	// IsVideoCar
	IsVideoCar int
	// LandMarkBuilding
	LandMarkBuilding string
	// LicenseCityId
	LicenseCityId int
	// LicenseCode
	LicenseCode string
	// LicenseProviceId
	LicenseProviceId int
	// 联系人
	LinkMan string
	// 联系人性别（1:男  2:女）
	LinkManSex int
	// 看车地址
	LookPlace string
	// MainBrandID
	MainBrandID int
	// MaintainRecord
	MaintainRecord int
	// picturecount
	picturecount int
	// ProducerID
	ProducerID int
	// RegUserId
	RegUserId int
	// Source
	Source int
	// StateDescription
	StateDescription string
	// StatusModifyTime
	StatusModifyTime string
	// SystemId
	SystemId int
	// UBUserId
	UBUserId int
	// UcarID
	UcarID int
	// UcarSerialNumber
	UcarSerialNumber string
	// UcarStatus
	UcarStatus int
	// UserID
	UserID int
	// UserType
	UserType int
	// VinCode
	VinCode string
	// 微信二维码
	WebchatUrl string
	// 微信号
	WechatCode string
	// 此车源图片路径
	ImgUrls []string
	// 主品牌名称(为空表示主品牌不存在)
	MBrand_Name string
	// 主品牌构造 Url 的拼音
	MBrand_UrlSpell string
	// 车系显示名称(为空表示车系不存在)
	Serial_ShowName string
	// 车系拼音
	Serial_AllSpell string
	// 车辆级别名称
	CarLevelName string
	// 车辆级别类型
	CarLevelType int
	// 变速箱名称
	GearBoxName string
	// 车型名称(为空表示车型不存在)
	Basic_Name string
	// 车型年款
	CarYear int
	// 车型信息中的排量
	Basic_Exhaust string
	// 发动机增压方式(涡轮增压，涡轮机械双增压，双涡轮增压，机械增压，待查，无)
	AddPressType string
	// 平均油耗
	Consumption float64
	// 新车厂商指导价(单位：万元)
	NewCarPrice float64
	// 手机扫描二维码指向的链接
	QrCodeUrl string
	// 车源亮点，返回的每个字符串中可能含有占位符 string，对应当前页面中的二级域名名称,如 www 或 beijing
	// <para>占位符在显示时需要过滤掉</para>
	CarHighLights []string
	// 经销商服务
	DealerService []string
	// 是否代办过户(是否包含过户费使用 IsIncTransfer 属性)
	IsTransferAgent bool
	// 小编推荐(专家推荐)
	Recommend string
	// 广告语
	Slogan string
	// 是否疑似商家
	IsDealerCar bool
	// 联系电话
	Tell string
	// 联系电话2
	Tell2 string
	// 联系人2
	LinkMan2 string
	// 微信号2
	WechatCode2 string
	// 是否保障车:0非保障车；1保障车；-1所有
	IsWarranty int
	// 保障服务类型:1原厂质保;2延长质保;4免费过户
	WarrantyTypes int
	// 是否维保：1维保；0非维保；-1所有
	IsShowMr int
	// 原生电话1
	ProtogTel string
	// 原生电话2
	ProtogTel1 string
	// 淘车认证标志
	IsDealerAuthorized int
}

func main() {

	str := strings.Join([]string{"228117",
		"228116",
		"228115",
		"228114",
		"228113",
		"228112",
		"228111",
		"228110",
		"227078",
		"227076",
		"227075",
		"227074",
		"227073",
		"227072",
		"227071",
		"226866",
		"224258",
		"224257",
		"223984",
		"223983",
		"223927",
		"223926",
		"223925",
		"223521",
		"223520",
		"223519",
		"223456",
		"223248",
		"223247",
		"223013",
		"223012",
		"223011",
		"222708",
		"221888",
		"221887",
		"221886",
		"221885",
		"221046",
		"220864",
		"215312",
		"215311",
		"215191",
		"214957",
		"202907",
		"202432",
		"202431",
		"199434",
		"198355",
		"196867",
		"196866",
		"196864",
		"196082",
		"196081",
		"196080",
		"196077",
		"196076",
		"194314",
		"194312",
		"194311",
		"192394",
		"192393",
		"187516",
		"187515",
		"187514",
		"187513",
		"187512",
		"187511",
		"187510",
		"187489",
		"187488",
		"187487",
		"182840",
		"182296",
		"181659",
		"180805",
		"180562",
		"180561",
		"180560",
		"180559",
		"180558",
		"179533",
		"179532",
		"179531",
		"179530",
		"179529",
		"177894",
		"177620",
		"176855",
		"176854",
		"176853",
		"176850",
		"176849",
		"176482",
		"176481",
		"176480",
		"170573",
		"170572",
		"170357",
		"170356",
		"169581",
		"169580",
		"169579",
		"169578",
		"167581",
		"167580",
		"167579",
		"167578",
		"167577",
		"167572",
		"167571",
		"167570",
		"163946",
		"156204"}, ",")

	fmt.Println(str)

	t := Tttt{32, "ewer"}
	fmt.Println(t)

	str1 := `{"AlternateContact":null,"BrandID":1679,"BuyCarDate":"2013-05-01T00:00:00","CalcCarPrice":"","CarCityId":507,"CardCode":null,"CarDistrictId":0,"CarDownTime":null,"CardRangeID":null,"CarID":103636,"CarProviceId":5,"CarPublishTime":"2018-01-04T15:02:06.983","CarSetting":null,"CarSource1L":2,"CarSource2L":9,"CarSourceStatus":0,"CarType":0,"CheckedTime":"2017-12-08T11:36:10.187","Color":"深灰色","CompleteRate":95,"Contact":"18316056086","Coordinate":null,"CreateTime":"2017-12-08T11:33:59.55","DisplayPrice":7.3000,"DriveLicensePic":"","DrivingMileage":40000,"ExamineExpireDate":"2019-05-01T00:00:00","Exhaust":"","ExhaustValue":1.6,"Fee":null,"FirstPicTrue":1,"GearBoxType":4,"InnerColor":"","InsuranceExpireDate":"2018-05-01T00:00:00","IsAuthenticated":0,"IsCards":0,"IsChangePrice":0,"IsDealerRecommend":null,"IsIncTransfer":1,"IsLicensed":1,"IsNeglect":1,"IsTransferToOld":0,"IsVideoCar":0,"LandMarkBuilding":null,"LicenseCityId":507,"LicenseCode":null,"LicenseProviceId":5,"LinkMan":"许小姐","LinkManSex":0,"LookPlace":"汕头市龙湖区华山路","MainBrandID":30,"MaintainRecord":1,"picturecount":10,"ProducerID":10046,"RegUserId":null,"Source":0,"StateDescription":"新到13骊威劲锐版1.6自动，车况精品，首付只要3万开走。","StatusModifyTime":"2018-01-11T15:41:06.563","SystemId":null,"UBUserId":462529,"UcarID":20377726,"UcarSerialNumber":"Dealer17120814414","UcarStatus":4,"UserID":225737,"UserType":22,"VinCode":"","WebchatUrl":"","WechatCode":"","ImgUrls":["02171f8l00.jpg","02171f8l3e.jpg","02171f8l38.jpg","02171f8l3r.jpg","02171f8l1x.jpg","02171f8l2z.jpg","02171f8l3d.jpg","02171f8l3u.jpg","02171f8l3o.jpg","02171f8l3y.jpg"],"MBrand_Name":"日产","MBrand_UrlSpell":"nissan","Serial_ShowName":"骊威","Serial_AllSpell":"liwei","CarLevelName":"小型","CarLevelType":4,"GearBoxName":"CVT无级变速","Basic_Name":"13款 劲锐 1.6L CVT XL 舒适版","CarYear":2013,"Basic_Exhaust":"1.6","AddPressType":"","Consumption":6.10,"NewCarPrice":10.78,"QrCodeUrl":null,"CarHighLights":["<a href='http://{0}.taoche.com/buycar/pgesbxcdz4a/' target='_blank'>儿童锁</a>"],"DealerService":[],"IsTransferAgent":false,"Recommend":null,"Slogan":null,"IsDealerCar":false,"Tell":"18316056086","Tell2":null,"LinkMan2":null,"WechatCode2":null,"IsWarranty":0,"WarrantyTypes":0,"IsShowMr":0,"ProtogTel":"18316056086","ProtogTel1":null,"IsDealerAuthorized":0} `

	bt := []byte(str1)
	var item Entity_Cacheucarbasicinfo_Full
	err := json.Unmarshal(bt, &item)
	// fmt.Println("bt：", bt)
	// fmt.Println("str:", string(bt))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		//fmt.Printf("%+v", item)
	}

	datamap := make(map[int]Entity_Cacheucarbasicinfo_Full)
	datamap[item.UcarID] = item

	fmt.Println(datamap)
}
