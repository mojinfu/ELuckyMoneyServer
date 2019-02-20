package main

import (
	"ELuckyMoneyServer/src/UselessHelper"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rlds/rlog"
)

func getPictureType(base64Arr0 string) string {
	defer UselessHelper.RecoverMedicine("getPictureType")
	ArrA := strings.Split(base64Arr0, "/")
	if len(ArrA) != 2 {
		return "jpeg"
	}
	ArrB := strings.Split(ArrA[1], ";")
	if len(ArrB) != 2 {
		return "jpeg"
	}
	return ArrB[0]
}
func (this *ChangePhoneRequestStruct) converjson() []byte {
	body, err := json.Marshal(this)
	if err != nil {
		rlog.V(1).Info("func converjson error:" + err.Error())
		return nil
	}
	return body
}
func (this *OpenItSMSRespStruct) converjson() []byte {
	body, err := json.Marshal(this)
	if err != nil {
		rlog.V(1).Info("func converjson error:" + err.Error())
		return nil
	}
	return body
}
func (this *VipCheckSMSRespStruct) converjson() []byte {
	body, err := json.Marshal(this)
	if err != nil {
		rlog.V(1).Info("func converjson error:" + err.Error())
		return nil
	}
	return body
}
func (this *VipSendSMSRespStruct) converjson() []byte {
	body, err := json.Marshal(this)
	if err != nil {
		rlog.V(1).Info("func converjson error:" + err.Error())
		return nil
	}
	return body
}
func (this *SendSMSRequestStruct) converjson() []byte {
	body, err := json.Marshal(this)
	if err != nil {
		rlog.V(1).Info("func converjson error:" + err.Error())
		return nil
	}
	return body
}
func (this *BindPhoneCheckSMSRequestStuct) converjson() []byte {
	body, err := json.Marshal(this)
	if err != nil {
		rlog.V(1).Info("func converjson error:" + err.Error())
		return nil
	}
	return body
}
func DeleteMapAfter1Min(key string) {
	time.Sleep(time.Minute * 5)
	_, isOK := SMSUserMap[key]
	if isOK {
		fmt.Println("删 短信session")
		delete(SMSUserMap, key)
	}
}
func GetLastUserListFilePath(wookPath string) string {
	myLastUserListFilePath := wookPath + "users.json"
	myLastUserListFileTime, _ := time.Parse("2006-01-02 15:04:05", "2018-09-01 15:04:05")
	filepath.Walk(wookPath, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if fi.IsDir() { // 忽略目录
			return nil
		}
		Arr1 := strings.Split(fi.Name(), "@")
		if len(Arr1) != 2 {
			return nil
		}
		Arr2 := strings.Split(Arr1[1], ".")
		if len(Arr2) != 2 {
			return nil
		}
		if Arr2[1] != "json" {
			return nil
		}
		myTime, err := time.Parse("2006-01-02", Arr2[0])
		if err != nil {
			return nil
		}
		if myTime.After(myLastUserListFileTime) {
			myLastUserListFileTime = myTime
			myLastUserListFilePath = wookPath + fi.Name()
		} else {
			return nil
		}
		return nil
	})
	return myLastUserListFilePath
}
func (this *UserListStruct) LoadData() bool {
	fmt.Println("Loading LoadData.")
	cpath := UselessHelper.GetConfPath(GetLastUserListFilePath("./"))
	fmt.Println("Loading LoadData.path:" + cpath)
	bdat := UselessHelper.GetAllFileData(cpath)
	err := json.Unmarshal(bdat, &this.UserList)
	if err != nil {
		cpath := UselessHelper.GetConfPath(GetLastUserListFilePath("../../bin/"))
		fmt.Println("Loading LoadData.path:" + cpath)
		bdat := UselessHelper.GetAllFileData(cpath)
		err := json.Unmarshal(bdat, &this.UserList)
		if err != nil {
			fmt.Println(err)
			return false
		} else {
			this.SeedFilePath = "../../bin/"
		}
	} else {
		this.SeedFilePath = "./"
	}

	return true
}
func (this *UserListStruct) OutputData() string {
	myData, err := json.Marshal(this.UserList)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	UselessHelper.SaveReplaceFile(this.SeedFilePath+"users@"+time.Now().Format("2006-01-02")+".json", myData)
	return string(myData)
}
func (this *UserListStruct) GetEmptyPhoneUser() *UserStruct {
	for index := range this.UserList {
		if len(this.UserList[index].PhoneInfo.Phone) == 0 {
			myLastSMSTimesamp, err := time.Parse("2006-01-02 15:04:05", this.UserList[index].CookieInfo.LastSMSTimesamp)
			if err != nil {
				return this.UserList[index]
			}
			if myLastSMSTimesamp.Add(time.Minute * 2).After(time.Now()) {
				continue
			}
			return this.UserList[index]
		}
	}
	return nil
}
func (this *UserListStruct) GetUserByUserName(myUserName string) *UserStruct {
	for index := range this.UserList {
		if this.UserList[index].UserName == myUserName {
			return this.UserList[index]
		}
	}
	return nil
}
