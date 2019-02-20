package loadFile

import (
	"ELuckyMoneyServer/src/UselessHelper"
	"encoding/json"
	"fmt"
)

type sysConf struct {
	LogDir      string //日志文件夹
	LogLevel    int32  //日志输出级别
	MaxLogLen_m uint64 //日志文件最大大小
	HttpHost    string //http服务开启域名
	// Fakehungrymen []string
	// Realhungrymen string //要领取红包的cookie
}

var PrivateConf sysConf

var HttpMonan string
var SpareHttpMonan string
var RegionForm map[string]string

func Loadconf() bool {
	fmt.Println("Loading configuration.")
	cpath := UselessHelper.GetConfPath("./config.json")
	fmt.Println("Loading configuration.path:" + cpath)
	bdat := UselessHelper.GetAllFileData(cpath)

	err := json.Unmarshal(bdat, &PrivateConf)
	if err != nil {
		fmt.Println(err)
		return false
	}
	PrivateConf.LogDir = UselessHelper.GetConfPath(PrivateConf.LogDir)
	PrivateConf.MaxLogLen_m = PrivateConf.MaxLogLen_m * 1024 * 1024
	fmt.Println("HttpHost address: " + PrivateConf.HttpHost)
	// if len(PrivateConf.Fakehungrymen) < 10 {
	// 	fmt.Println("Fakehungrymen length is not enough")
	// 	return false
	// }
	return true
}
