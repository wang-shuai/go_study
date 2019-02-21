package common

import (
	"github.com/olebedev/config"
	"os"
	"fmt"
)

var (
	CfgMgr *config.Config
)

func init() {
	var err error
	pd, err := os.Getwd()
	CfgMgr, err = config.ParseYamlFile(pd + "/conf/app.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(CfgMgr)
}
