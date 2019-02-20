package main

import (
	"ELuckyMoneyServer/src/UselessDownload"
	"ELuckyMoneyServer/src/UselessHelper"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"io/ioutil"
	"net/http"

	"github.com/rlds/rlog"
)

type VipSendSMSRequestStruct struct {
	Phone string
}
type VipSendSMSRespStruct struct {
	Message       string
	Code          int64
	ValidateToken string
}

func VipBindPhoneSendSMS(w http.ResponseWriter, r *http.Request) {
	defer UselessHelper.RecoverMedicine("VipBindPhoneSendSMS")

	myVipSendSMSResp := &VipSendSMSRespStruct{}
	err := r.ParseForm()
	if err != nil {
		myVipSendSMSResp.Message = "解析Url时:" + err.Error()
		w.Write([]byte(myVipSendSMSResp.converjson()))
		rlog.V(1).Info(myVipSendSMSResp.Message)
		rlog.V(1).Info("_______________________________________________")
		return
	}
	askBody, ioutilerr := ioutil.ReadAll(r.Body)
	if r.Body != nil {
		defer r.Body.Close()
	}
	if ioutilerr != nil {
		myVipSendSMSResp.Message = "ioutilerr:" + err.Error()
		w.Write([]byte(myVipSendSMSResp.converjson()))
		rlog.V(1).Info(myVipSendSMSResp.Message)
		rlog.V(1).Info("_______________________________________________")
		return
	}
	var VipSendSMSJson string = bytes.NewBuffer(askBody).String()
	myVipSendSMS := &VipSendSMSRequestStruct{}
	err = json.Unmarshal([]byte(VipSendSMSJson), &myVipSendSMS)
	if err != nil {
		myVipSendSMSResp.Message = "Unmarshal:" + VipSendSMSJson + "--" + err.Error()
		w.Write([]byte(myVipSendSMSResp.converjson()))
		rlog.V(1).Info(myVipSendSMSResp.Message)
		rlog.V(1).Info("_______________________________________________")
		return
	}
	if len(myVipSendSMS.Phone) != 11 {
		myVipSendSMSResp.Message = "手机号码：" + myVipSendSMS.Phone + "位数不对哦"
		w.Write([]byte(myVipSendSMSResp.converjson()))
		rlog.V(1).Info("_______________________________________________")
		return
	}
	myUser := PublicUserList.GetEmptyPhoneUser() //saya
	//myUser := PublicUserList.GetUserByUserName("颜半仙")
	if myUser == nil {
		myVipSendSMSResp.Message = "没有可用的小号了"
		w.Write([]byte(myVipSendSMSResp.converjson()))
		rlog.V(1).Info("_______________________________________________")
		return
	} else {
		sendSMSResp, err := myUser.BindPhoneSendSMS(
			&SendSMSRequestStruct{Mobile: myVipSendSMS.Phone})
		if err != nil {
			myVipSendSMSResp.Message = "发送手机短信时:" + err.Error()
			w.Write([]byte(myVipSendSMSResp.converjson()))
			rlog.V(1).Info("_______________________________________________")
			return
		}
		if sendSMSResp.Name == "NEED_CAPTCHA" {
			myCapRespJson, _, myCookies := UselessDownload.DownloadPOSTJsonCollectCookie(
				myEOrigin+`/restapi/eus/v3/captchas`,
				`{"captcha_str":"`+myVipSendSMS.Phone+`"}`,
				myUser.CookieInfo.Cookie)
			for index := range myCookies {
				myUser.ChangeCookieKeyValue(myCookies[index].Name, myCookies[index].Value)
			}
			fmt.Println("myCapRespJson:", myCapRespJson)
			myCap := &CapRespStruct{}
			err := json.Unmarshal([]byte(myCapRespJson), &myCap)
			base64Arr := strings.Split(myCap.CaptchaImage, ",")
			if err != nil || len(myCap.CaptchaImage) == 0 || len(base64Arr) != 2 {
				w.Write([]byte("图片验证码请求出错：" + err.Error() + myCapRespJson))
				return
			}
			ddd, _ := base64.StdEncoding.DecodeString(base64Arr[1])
			UselessHelper.SaveReplaceFile("./rider."+getPictureType(base64Arr[0]), ddd)
			//打码
			Rider := ""
			fmt.Println("Please enter Rider: ")
			fmt.Scanln(&Rider)
			fmt.Println("Rider:", Rider)
			//打码
			sendSMSResp, err = myUser.BindPhoneSendSMS(
				&SendSMSRequestStruct{
					Mobile:       myVipSendSMS.Phone,
					CaptchaValue: Rider,
					CaptchaHash:  myCap.CaptchaHash,
				})

			if err != nil {
				myVipSendSMSResp.Message = "发送手机短信时:" + err.Error()
				w.Write([]byte(myVipSendSMSResp.converjson()))
				rlog.V(1).Info("_______________________________________________")
				return
			} else if sendSMSResp.Name == "NEED_CAPTCHA" {
				myVipSendSMSResp.Message = "验证码识别 发送手机短信时:" + sendSMSResp.Message
				w.Write([]byte(myVipSendSMSResp.converjson()))
				rlog.V(1).Info("_______________________________________________")
				return
			}
		}
		myVipSendSMSResp.ValidateToken = sendSMSResp.ValidateToken
		w.Write([]byte(myVipSendSMSResp.converjson()))
		rlog.V(1).Info("_______________________________________________")
		return
	}
}

type VipCheckSMSRequestStruct struct {
	Phone         string
	ValidateToken string
	ValidateCode  string
	PhoneOwner    string
}
type VipCheckSMSRespStruct struct {
	Message string
	Code    int64
}

func VipBindPhoneCheckSMS(w http.ResponseWriter, r *http.Request) {
	defer UselessHelper.RecoverMedicine("VipBindPhoneCheckSMS")
	myVipCheckSMSResp := &VipCheckSMSRespStruct{}
	err := r.ParseForm()
	if err != nil {
		myVipCheckSMSResp.Message = "解析Url时:" + err.Error()
		w.Write([]byte(myVipCheckSMSResp.converjson()))
		rlog.V(1).Info(myVipCheckSMSResp.Message)
		rlog.V(1).Info("_______________________________________________")
		return
	}
	askBody, ioutilerr := ioutil.ReadAll(r.Body)
	if r.Body != nil {
		defer r.Body.Close()
	}
	if ioutilerr != nil {
		myVipCheckSMSResp.Message = "ioutilerr:" + err.Error()
		w.Write([]byte(myVipCheckSMSResp.converjson()))
		rlog.V(1).Info(myVipCheckSMSResp.Message)
		rlog.V(1).Info("_______________________________________________")
		return
	}
	var VipCheckSMSJson string = bytes.NewBuffer(askBody).String()
	myVipCheckSMS := &VipCheckSMSRequestStruct{}
	err = json.Unmarshal([]byte(VipCheckSMSJson), &myVipCheckSMS)
	if err != nil {
		myVipCheckSMSResp.Message = "Unmarshal:" + err.Error()
		w.Write([]byte(myVipCheckSMSResp.converjson()))
		rlog.V(1).Info(myVipCheckSMSResp.Message)
		rlog.V(1).Info("_______________________________________________")
		return
	}
	if len(myVipCheckSMS.Phone) != 11 {
		myVipCheckSMSResp.Message = "手机号码：" + myVipCheckSMS.Phone + "位数不对哦"
		w.Write([]byte(myVipCheckSMSResp.converjson()))
		rlog.V(1).Info("_______________________________________________")
		return
	}
	if len(myVipCheckSMS.ValidateToken) == 0 {
		myVipCheckSMSResp.Message = "ValidateToken" + myVipCheckSMS.Phone + "位数不对哦"
		w.Write([]byte(myVipCheckSMSResp.converjson()))
		rlog.V(1).Info("_______________________________________________")
		return
	}
	myCheckSMSRequestStruct, err := CheckSMS(myVipCheckSMS)
	if err != nil {
		myVipCheckSMSResp.Message = err.Error()
	} else {
		myVipCheckSMSResp.Message = myCheckSMSRequestStruct.Message
		if myCheckSMSRequestStruct.UserID != 0 {
			myVipCheckSMSResp.Message = "ok"
		}
	}
	w.Write([]byte(myVipCheckSMSResp.converjson()))
	rlog.V(1).Info("_______________________________________________")
	return
}

type OpenItRequestStruct struct {
	Phone string
	EUrl  string
}
type OpenItSMSRespStruct struct {
	Message string
	Code    int64
}

func OpenIt(w http.ResponseWriter, r *http.Request) {
	defer UselessHelper.RecoverMedicine("VipBindPhoneCheckSMS")
	myOpenItSMSResp := &OpenItSMSRespStruct{}
	err := r.ParseForm()
	if err != nil {
		myOpenItSMSResp.Message = "解析Url时:" + err.Error()
		w.Write([]byte(myOpenItSMSResp.converjson()))
		rlog.V(1).Info(myOpenItSMSResp.Message)
		rlog.V(1).Info("_______________________________________________")
		return
	}
	askBody, ioutilerr := ioutil.ReadAll(r.Body)
	if r.Body != nil {
		defer r.Body.Close()
	}
	if ioutilerr != nil {
		myOpenItSMSResp.Message = "ioutilerr:" + err.Error()
		w.Write([]byte(myOpenItSMSResp.converjson()))
		rlog.V(1).Info(myOpenItSMSResp.Message)
		rlog.V(1).Info("_______________________________________________")
		return
	}
	var VipCheckSMSJson string = bytes.NewBuffer(askBody).String()
	myOpenItRequest := &OpenItRequestStruct{}
	err = json.Unmarshal([]byte(VipCheckSMSJson), &myOpenItRequest)
	if err != nil {
		myOpenItSMSResp.Message = "Unmarshal:" + err.Error()
		w.Write([]byte(myOpenItSMSResp.converjson()))
		rlog.V(1).Info(myOpenItSMSResp.Message)
		rlog.V(1).Info("_______________________________________________")
		return
	}
	if len(myOpenItRequest.Phone) != 11 {
		myOpenItSMSResp.Message = "手机号码：" + myOpenItRequest.Phone + "位数不对哦"
		w.Write([]byte(myOpenItSMSResp.converjson()))
		rlog.V(1).Info("_______________________________________________")
		return
	}
	myUser := GetUserByPhone(myOpenItRequest.Phone)
	if myUser == nil {
		myOpenItSMSResp.Message = "用户并未注册，清注册后使用。（受饿了么红包限制）"
		w.Write([]byte(myOpenItSMSResp.converjson()))
		rlog.V(1).Info("_______________________________________________")
		return
	}
	myE, err := NewELuckyMoney(myOpenItRequest.EUrl)
	if err != nil {
		myOpenItSMSResp.Message = "解析红包链接时发生错误：" + err.Error()
		w.Write([]byte(myOpenItSMSResp.converjson()))
		rlog.V(1).Info("_______________________________________________")
		return
	}
	err = myUser.HelpMeOpen(myE)
	if err != nil {
		myOpenItSMSResp.Message = err.Error()
	} else {
		myOpenItSMSResp.Message = "ok"
	}

	w.Write([]byte(myOpenItSMSResp.converjson()))
	rlog.V(1).Info("_______________________________________________")
	return
}
func GetUserByPhone(phone string) *UserStruct {
	if len(phone) == 0 {
		return nil
	}
	for index := range PublicUserList.UserList {
		if PublicUserList.UserList[index].PhoneInfo.Phone == phone {
			return PublicUserList.UserList[index]
		} else {
			continue
		}
	}
	return nil
}
