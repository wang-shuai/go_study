package services

import (
	"KiRpc/common"
	"log"
	"github.com/garyburd/redigo/redis"
	"github.com/pquerna/ffjson/ffjson"
	"sort"
	"strings"
)

type KiRpcService struct {
}

func (service *KiRpcService) GetKiLicenseCities() (r []*ViewKiLicenseCity, err error) {
	//该数据从redis获取  缓存数据由服务提前压入并根据Excel修改时间完成更新
	r = []*ViewKiLicenseCity{}
	conn := common.Pool.Get()
	data, err := redis.Bytes(conn.Do("get", "license_cities"))
	if err != nil {
		log.Fatal("从redis获取license_cities失败：", err)
		return r, err
	}

	var ret []*ViewKiLicenseCity
	err = ffjson.Unmarshal(data, &ret)
	if err != nil {
		log.Fatal("反序列化ViewKiLicenseCityGroup失败：", err)
		return r, err
	}

	return ret, err
}

func (service *KiRpcService) GetKiLicenseCityGroups() (r []*ViewKiLicenseCityGroup, err error) {
	data, err := service.GetKiLicenseCities()
	r = []*ViewKiLicenseCityGroup{}
	gmap := make(map[byte][]*ViewKiLicenseCity)
	for _, city := range data {
		letter := city.CitySpell[0]
		gmap[letter] = append(gmap[letter], city)
	}
	for key, val := range gmap {
		g := &ViewKiLicenseCityGroup{GroupName: strings.ToUpper(string(key)), Cities: val}
		r = append(r, g)
	}
	sort.Slice(r, func(i, j int) bool {
		return strings.Compare(r[i].GroupName, r[j].GroupName) < 0
	})
	return r, err
}
