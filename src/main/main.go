package main

import (
	"ELuckyMoneyServer/src/UselessHelper"
	. "ELuckyMoneyServer/src/loadFile"
	"net/http"
	"time"

	"github.com/rlds/rlog"
)

type UserListStruct struct {
	UserList     []*UserStruct
	SeedFilePath string
}

var SMSUserMap map[string]*UserStruct
var PublicUserList UserListStruct

func main() {
	SMSUserMap = make(map[string]*UserStruct)
	if Loadconf() && PublicUserList.LoadData() {
		UselessHelper.MkAlldir(PrivateConf.LogDir)
		rlog.LogInit(PrivateConf.LogLevel, PrivateConf.LogDir, PrivateConf.MaxLogLen_m, 1)

		profServeMux := http.NewServeMux()
		profServeMux.HandleFunc("/OpenIt", OpenIt)
		profServeMux.HandleFunc("/Vip/BindPhoneSendSMS", VipBindPhoneSendSMS)
		profServeMux.HandleFunc("/Vip/BindPhoneCheckSMS", VipBindPhoneCheckSMS)
		server := &http.Server{
			Addr:           PrivateConf.HttpHost,
			Handler:        profServeMux,
			ReadTimeout:    100 * time.Second,
			WriteTimeout:   100 * time.Second,
			MaxHeaderBytes: 4 << 20,
		}
		server.SetKeepAlivesEnabled(false)
		rlog.V(1).Info("即将启动服务:[" + PrivateConf.HttpHost + "]")
		server.ListenAndServe()
	}
}
