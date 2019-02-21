package main

import (
	"KiRpc/services"
	"fmt"
	"gitlab.dev.daikuan.com/platform/servicecenter/lib/go/servicecenter/pkg"
)

func main() {
	pkg.ServiceManagerInstance.SimpleRegister(services.NewIKiRpcServiceThriftProcessor(&services.KiRpcService{}), func() {
		fmt.Println("KiRpc start success")
	})

	service := services.KiRpcService{}
	data, err := service.GetKiLicenseCities()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	group, err := service.GetKiLicenseCityGroups()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(group)
}
