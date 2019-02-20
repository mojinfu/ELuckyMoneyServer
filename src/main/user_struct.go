package main

import (
	"ELuckyMoneyServer/src/UselessDownload"
	"ELuckyMoneyServer/src/UselessHelper"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/rlds/rlog"
)

const SplitLine string = "-----------------------------------------------"

func getUser() *UserStruct {
	myUser := &UserStruct{
		UserName: "颜半仙",
		PhoneInfo: PhoneInfoStruct{
			Phone:          "",
			PhoneOwner:     "",
			HasLuckedToday: false,
		},
		CookieInfo: CookieInfoStruct{
			Cookie:             "SID=fmO8EPeH1aVXdQswrLaySC9bcUpOqzx6mshg; USERID=4339802; track_id=1536291426|69395bec3a8e693fe23b6564c998fa34785f8a8f38b7dc5335|fe2368f9da9d8aefbce7ddc40a2a1cbe; _utrace=7781a14c03acacba170d423ddb2a0c0e_2018-08-09; snsInfo[wx2a416286e96100ed]=%7B%22city%22%3A%22%E6%9D%AD%E5%B7%9E%22%2C%22country%22%3A%22%E4%B8%AD%E5%9B%BD%22%2C%22eleme_key%22%3A%22d7bd99eaf4bdaec6ba96220d99c6776e%22%2C%22headimgurl%22%3A%22http%3A%2F%2Fthirdwx.qlogo.cn%2Fmmopen%2Fvi_32%2FQ0j4TwGTfTJYk0mY3BuCX8sfBkQwxAqvdniawPnqJXNx2JuXNibMicvzdjz7BgfsibYcJgXjYLBLxgviakO9rBL2c0g%2F132%22%2C%22language%22%3A%22zh_CN%22%2C%22nickname%22%3A%22%E9%A2%9C%E8%89%B21934%22%2C%22openid%22%3A%22oEGLvjgfRqfHoYNHx0IxY1EwjLMg%22%2C%22privilege%22%3A%5B%5D%2C%22province%22%3A%22%E6%B5%99%E6%B1%9F%22%2C%22sex%22%3A1%2C%22unionid%22%3A%22o_PVDuIDTHd67WEarVE7vWW22XdM%22%2C%22name%22%3A%22%E9%A2%9C%E8%89%B21934%22%2C%22avatar%22%3A%22http%3A%2F%2Fthirdwx.qlogo.cn%2Fmmopen%2Fvi_32%2FQ0j4TwGTfTJYk0mY3BuCX8sfBkQwxAqvdniawPnqJXNx2JuXNibMicvzdjz7BgfsibYcJgXjYLBLxgviakO9rBL2c0g%2F132%22%7D; ubt_ssid=r1mrxh8blwtcd5gbl54qsacysewvl860_2018-08-09; perf_ssid=qek2wdc82t10tbh0p7ja67rna74bg6eg_2018-08-09",
			CookieOwner:        "mojinfu Wechat",
			OfferTimesEveryDay: 4,
			OfferTimesToday:    0,
		},
	}
	return myUser
}

type UserStruct struct {
	UserName   string
	PhoneInfo  PhoneInfoStruct
	CookieInfo CookieInfoStruct
}
type PhoneInfoStruct struct {
	Phone          string
	PhoneOwner     string
	HasLuckedToday bool
}
type CookieInfoStruct struct {
	Cookie             string
	CookieOwner        string
	OfferTimesEveryDay int64
	OfferTimesToday    int64
	LastSMSTimesamp    string
	elemeKey           string
	openID             string
}

func (this *UserStruct) ClearPhoneInfo() {
	rlog.V(1).Info("用户:" + this.PhoneInfo.Phone + "@" + this.PhoneInfo.PhoneOwner + "好像已经取消注册了")
	this.PhoneInfo.Phone = ""
	this.PhoneInfo.PhoneOwner = ""
	this.PhoneInfo.HasLuckedToday = false
}
func (this *UserStruct) OpenLuckyMoney(myELuckyMoney *ELuckyMoneyStruct) (*luckyMoneyRespStruct, error) {
	var myLuckydoor luckyMoneyRequestStruct
	myLuckydoor.Device_id = ""
	myLuckydoor.Hardware_id = ""
	myLuckydoor.Method = "phone"
	myLuckydoor.Group_sn = myELuckyMoney.sn
	myLuckydoor.Phone = this.PhoneInfo.Phone
	myLuckydoor.Platform = 0
	myLuckydoor.Sign = this.GetElemeKey()
	myLuckydoor.Track_id = this.GetCookieTrackID()
	//myLuckydoor.Track_id = this.GetCookieTrackID()
	myLuckydoor.Unionid = "fuck"
	myLuckydoor.Weixin_avatar = "http://a3.topitme.com/d/08/dc/1131099938aecdc08do.jpg"
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// if len(this.UserName) == 0 {
	// 	myLuckydoor.Weixin_username = "摸金校尉" + fmt.Sprintf("%d", r.Intn(1000))
	// } else {
	myLuckydoor.Weixin_username = this.UserName + this.PhoneInfo.PhoneOwner
	// }
	fmt.Println("myLuckydoor:", string(myLuckydoor.converjson()))
	this.CookieInfo.OfferTimesToday++
	luckydoorJson, _ := UselessDownload.Download_POST_Json(
		fmt.Sprintf(myEOrigin+`/restapi/marketing/promotion/weixin/%s`, this.GetOpenID()),
		string(myLuckydoor.converjson()),
		myELuckyMoney.sn,
		// `"track_id=1536291426|69395bec3a8e693fe23b6564c998fa34785f8a8f38b7dc5335|fe2368f9da9d8aefbce7ddc40a2a1cbe; SID=hdYzMMFQlxaqktFcaZdF0lUWcuy8EnLXjPOg; USERID=4339802; _utrace=7781a14c03acacba170d423ddb2a0c0e_2018-08-09; snsInfo[wx2a416286e96100ed]=%7B%22city%22%3A%22%E6%9D%AD%E5%B7%9E%22%2C%22country%22%3A%22%E4%B8%AD%E5%9B%BD%22%2C%22eleme_key%22%3A%22d7bd99eaf4bdaec6ba96220d99c6776e%22%2C%22headimgurl%22%3A%22http%3A%2F%2Fthirdwx.qlogo.cn%2Fmmopen%2Fvi_32%2FQ0j4TwGTfTJYk0mY3BuCX8sfBkQwxAqvdniawPnqJXNx2JuXNibMicvzdjz7BgfsibYcJgXjYLBLxgviakO9rBL2c0g%2F132%22%2C%22language%22%3A%22zh_CN%22%2C%22nickname%22%3A%22%E9%A2%9C%E8%89%B21934%22%2C%22openid%22%3A%22oEGLvjgfRqfHoYNHx0IxY1EwjLMg%22%2C%22privilege%22%3A%5B%5D%2C%22province%22%3A%22%E6%B5%99%E6%B1%9F%22%2C%22sex%22%3A1%2C%22unionid%22%3A%22o_PVDuIDTHd67WEarVE7vWW22XdM%22%2C%22name%22%3A%22%E9%A2%9C%E8%89%B21934%22%2C%22avatar%22%3A%22http%3A%2F%2Fthirdwx.qlogo.cn%2Fmmopen%2Fvi_32%2FQ0j4TwGTfTJYk0mY3BuCX8sfBkQwxAqvdniawPnqJXNx2JuXNibMicvzdjz7BgfsibYcJgXjYLBLxgviakO9rBL2c0g%2F132%22%7D; ubt_ssid=r1mrxh8blwtcd5gbl54qsacysewvl860_2018-08-09; perf_ssid=qek2wdc82t10tbh0p7ja67rna74bg6eg_2018-08-09"`) //saya//更改后
		// `_utrace=a9d347e5cb6b5712a4ccf2190b899823_2018-09-24; snsInfo[101204453]=%7B%22city%22%3A%22%22%2C%22constellation%22%3A%22%22%2C%22eleme_key%22%3A%226a8a0e0b0ce802715fabd62a9912a2d2%22%2C%22figureurl%22%3A%22http%3A%2F%2Fqzapp.qlogo.cn%2Fqzapp%2F101204453%2F243E6B480E2FBE5A20D0E6F963F187BE%2F30%22%2C%22figureurl_1%22%3A%22http%3A%2F%2Fqzapp.qlogo.cn%2Fqzapp%2F101204453%2F243E6B480E2FBE5A20D0E6F963F187BE%2F50%22%2C%22figureurl_2%22%3A%22http%3A%2F%2Fqzapp.qlogo.cn%2Fqzapp%2F101204453%2F243E6B480E2FBE5A20D0E6F963F187BE%2F100%22%2C%22figureurl_qq_1%22%3A%22http%3A%2F%2Fthirdqq.qlogo.cn%2Fqqapp%2F101204453%2F243E6B480E2FBE5A20D0E6F963F187BE%2F40%22%2C%22figureurl_qq_2%22%3A%22http%3A%2F%2Fthirdqq.qlogo.cn%2Fqqapp%2F101204453%2F243E6B480E2FBE5A20D0E6F963F187BE%2F100%22%2C%22gender%22%3A%22%E7%94%B7%22%2C%22is_lost%22%3A0%2C%22is_yellow_vip%22%3A%220%22%2C%22is_yellow_year_vip%22%3A%220%22%2C%22level%22%3A%220%22%2C%22msg%22%3A%22%22%2C%22nickname%22%3A%22yqr4%22%2C%22openid%22%3A%22243E6B480E2FBE5A20D0E6F963F187BE%22%2C%22province%22%3A%22%22%2C%22ret%22%3A0%2C%22vip%22%3A%220%22%2C%22year%22%3A%220%22%2C%22yellow_vip_level%22%3A%220%22%2C%22name%22%3A%22yqr4%22%2C%22avatar%22%3A%22http%3A%2F%2Fthirdqq.qlogo.cn%2Fqqapp%2F101204453%2F243E6B480E2FBE5A20D0E6F963F187BE%2F40%22%7D; ubt_ssid=un06vfudmtcfqcs02hupqua3ogwx1h9p_2018-09-24; perf_ssid=duvsqebcbmfiofplfmmzigivqpxdrmxr_2018-09-24;track_id=1538068201|419c749d933de0da803f0dc81c3e10837ddc8c983c159f112c|dc97786e3a95beb68de3000f197763bc;USERID=4339802;SID=ZK6G6VgV65Lo7YQCVUG1LhkvzjidXvaVeF8w`) //saya//实际上
		this.CookieInfo.Cookie)
	fmt.Println("luckydoorJson: "+this.UserName, luckydoorJson)
	//fmt.Println("this.CookieInfo.Cookie: "+this.UserName, this.CookieInfo.Cookie)
	myluckyMoneyResp := &luckyMoneyRespStruct{}
	err := json.Unmarshal([]byte(luckydoorJson), &myluckyMoneyResp)
	if err != nil {
		return nil, err
	}
	if myluckyMoneyResp.Name == "PHONE_IS_EMPTY" {
		this.ClearPhoneInfo()
		PublicUserList.OutputData()
	}
	if myluckyMoneyResp.Account != this.PhoneInfo.Phone {
		this.PhoneInfo.Phone = myluckyMoneyResp.Account
		PublicUserList.OutputData()
	}

	if len(myluckyMoneyResp.Promotion_records) == myELuckyMoney.lucky_number-1 {
		myELuckyMoney.EGonaLucky = true
	} else {
		myELuckyMoney.EGonaLucky = false
	}
	if len(myluckyMoneyResp.Promotion_records) >= myELuckyMoney.lucky_number {
		myELuckyMoney.isUsedLucky = true
	} else {
		myELuckyMoney.isUsedLucky = false
	}
	fmt.Println("先已经有", len(myluckyMoneyResp.Promotion_records), "人抢了")
	fmt.Println("第", myELuckyMoney.lucky_number, "个人是最大红包")
	fmt.Println(SplitLine)
	return myluckyMoneyResp, nil
	//todo 根据饿了么返回 修改他的手机号信息
}

type SendSMSRequestStruct struct {
	Mobile       string `json:"mobile,omitempty"`
	CaptchaValue string `json:"captcha_value,omitempty"`
	CaptchaHash  string `json:"captcha_hash,omitempty"`
}
type BindPhoneSendSMSRespStruct struct {
	ValidateToken string `json:"validate_token"`
	Message       string `json:"message"`
	Name          string `json:"name"`
}
type CapRespStruct struct {
	CaptchaHash  string `json:"captcha_hash"`
	CaptchaImage string `json:"captcha_image"`
}

func (this *UserStruct) BindPhoneSendSMS(mySendSMSRequest *SendSMSRequestStruct) (*BindPhoneSendSMSRespStruct, error) {
	this.CookieInfo.LastSMSTimesamp = time.Now().Format("2006-01-02 15:04:05")
	PublicUserList.OutputData()
	sendSMSRespJson, _, myCookies := UselessDownload.DownloadPOSTJsonCollectCookie(
		myEOrigin+`/restapi/eus/login/mobile_send_code`,
		string(mySendSMSRequest.converjson()),
		this.CookieInfo.Cookie)
	for index := range myCookies {
		this.ChangeCookieKeyValue(myCookies[index].Name, myCookies[index].Value)
	}
	fmt.Println("sendSMSRespJson:", sendSMSRespJson)
	myBindPhoneSendSMSResp := &BindPhoneSendSMSRespStruct{}
	err := json.Unmarshal([]byte(sendSMSRespJson), &myBindPhoneSendSMSResp)
	if err != nil {
		return nil, err
	}
	if len(myBindPhoneSendSMSResp.ValidateToken) != 0 {
		SMSUserMap[myBindPhoneSendSMSResp.ValidateToken] = this
		go DeleteMapAfter1Min(myBindPhoneSendSMSResp.ValidateToken)
	}
	switch myBindPhoneSendSMSResp.Name {
	case "VALIDATION_TOO_BUSY":
		{
			return myBindPhoneSendSMSResp, errors.New(myBindPhoneSendSMSResp.Message)
		}
	case "NEED_CAPTCHA":
		{
			return myBindPhoneSendSMSResp, nil
		}
	case "CAPTCHA_CODE_ERROR":
		{
			return myBindPhoneSendSMSResp, errors.New(myBindPhoneSendSMSResp.Message)
		}
	default:
		return myBindPhoneSendSMSResp, nil
	}
}

type BindPhoneCheckSMSRequestStuct struct {
	Mobile        string `json:"mobile"`
	ValidateCode  string `json:"validate_code"`
	ValidateToken string `json:"validate_token"`
}

type BindPhoneCheckSMSRespStuct struct {
	Message        string `json:"message"`
	Name           string `json:"name"`
	NeedBindMobile bool   `json:"need_bind_mobile"`
	UserID         int64  `json:"user_id"`
}

func CheckSMS(myVipCheckSMSRequest *VipCheckSMSRequestStruct) (*BindPhoneCheckSMSRespStuct, error) {
	mySMSUser, isOK := SMSUserMap[myVipCheckSMSRequest.ValidateToken]
	if !isOK {
		return nil, errors.New("ValidateToken:" + myVipCheckSMSRequest.ValidateToken + "已经超过两分钟了，不可用。")
	}
	myBindPhoneSendSMSResp := &BindPhoneSendSMSRespStruct{
		ValidateToken: myVipCheckSMSRequest.ValidateToken,
	}
	myBindPhoneCheckSMSResp, err := mySMSUser.BindPhoneCheckSMS(myVipCheckSMSRequest.Phone, myVipCheckSMSRequest.ValidateCode, myBindPhoneSendSMSResp)
	if err != nil {
		return nil, err
	}
	if myBindPhoneCheckSMSResp.UserID > 0 {
		mySMSUser.PhoneInfo.Phone = myVipCheckSMSRequest.Phone
		mySMSUser.PhoneInfo.PhoneOwner = myVipCheckSMSRequest.PhoneOwner
		PublicUserList.OutputData()
		return myBindPhoneCheckSMSResp, nil
	} else {
		return nil, errors.New(myBindPhoneCheckSMSResp.Message)
	}
}

type ChangePhoneRequestStruct struct {
	Sign  string `json:"sign"`
	Phone string `json:"phone"`
}

func (this *UserStruct) BindPhoneCheckSMS(myPhone string, myValidateCode string, myBindPhoneSendSMSResp *BindPhoneSendSMSRespStruct) (*BindPhoneCheckSMSRespStuct, error) {
	myBindPhoneCheckSMSRequest := &BindPhoneCheckSMSRequestStuct{
		Mobile:        myPhone,
		ValidateToken: myBindPhoneSendSMSResp.ValidateToken,
		ValidateCode:  myValidateCode,
	}
	bindPhoneCheckSMSJson, _, myCookies := UselessDownload.DownloadPOSTJsonCollectCookie(
		myEOrigin+`/restapi/eus/login/login_by_mobile`,
		string(myBindPhoneCheckSMSRequest.converjson()),
		this.CookieInfo.Cookie)
	for index := range myCookies {
		fmt.Println(myCookies[index].Name)
		fmt.Println(myCookies[index].Value)
		this.ChangeCookieKeyValue(myCookies[index].Name, myCookies[index].Value)
	}
	PublicUserList.OutputData()
	fmt.Println("bindPhoneCheckSMSJson:", bindPhoneCheckSMSJson)
	myBindPhoneCheckSMSResp := &BindPhoneCheckSMSRespStuct{}
	err := json.Unmarshal([]byte(bindPhoneCheckSMSJson), &myBindPhoneCheckSMSResp)
	if err != nil {
		return nil, err
	}
	if myBindPhoneCheckSMSResp.UserID == 0 {
		return myBindPhoneCheckSMSResp, nil
	}
	//this.ChangeUid(fmt.Sprintf("%d", myBindPhoneCheckSMSResp.UserID))
	myChangePhoneRequest := &ChangePhoneRequestStruct{
		Sign:  this.GetElemeKey(),
		Phone: myPhone,
	}
	myChangePhoneResp, _ := UselessDownload.Download_POST_Json(
		myEOrigin+`/restapi/marketing/hongbao/weixin/`+this.GetOpenID()+`/change`,
		string(myChangePhoneRequest.converjson()),
		"",
		this.CookieInfo.Cookie)
	if len(myChangePhoneResp) != 0 {
		rlog.V(1).Info("myChangePhoneResp:" + myChangePhoneResp)
	} else {

	}

	// myChangePhoneRequest := &ChangePhoneRequestStruct{
	// 	Sign:  "d7ad32f9273ccd9757794c11f4bbd84a", //todo
	// 	Phone: myPhone,
	// }
	// myChangePhoneResp, _ := UselessDownload.Download_POST_Json(
	// 	myEOrigin+`/restapi/marketing/hongbao/weixin/`+this.GetOpenID()+`/change`,
	// 	string(myChangePhoneRequest.converjson()),
	// 	"",
	// 	this.CookieInfo.Cookie)
	// if len(myChangePhoneResp) != 0 {
	// 	rlog.V(1).Info("myChangePhoneResp:" + myChangePhoneResp)
	// } else {

	// }

	return myBindPhoneCheckSMSResp, nil
}

func Register() (*UserStruct, error) {
	myUser := &UserStruct{}
	return myUser, nil
}

//tools----------------------------------------------------------------------
func (this *UserStruct) GetElemeKey() string {
	if len(this.CookieInfo.elemeKey) == 0 {
		cookieAfterUrldecode, isdecodeOk := url.QueryUnescape(this.CookieInfo.Cookie)
		if isdecodeOk != nil {
			return ""
		} else {
			this.CookieInfo.elemeKey = UselessHelper.GetJsonValue(cookieAfterUrldecode, "eleme_key")
		}
	}
	return this.CookieInfo.elemeKey
}
func (this *UserStruct) GetOpenID() string {
	if len(this.CookieInfo.openID) == 0 {
		cookieAfterUrldecode, isdecodeOk := url.QueryUnescape(this.CookieInfo.Cookie)
		if isdecodeOk != nil {
			return ""
		} else {
			this.CookieInfo.openID = UselessHelper.GetJsonValue(cookieAfterUrldecode, "openid")
		}
	}
	return this.CookieInfo.openID
}

func (this *UserStruct) HelpMeOpen(myELuckyMoney *ELuckyMoneyStruct) error {
	err := errors.New("人手不足")
	for index := range PublicUserList.UserList {
		if myELuckyMoney.EGonaLucky {
			myResp, err := this.OpenLuckyMoney(myELuckyMoney)
			if err != nil {
				return err
			}
			switch myResp.Ret_code {
			case 4:
				{
					err = nil
					log.Println("success")
				}

			case 2:
				err = errors.New("这个红包你之前已经抢过")
			default:
				{
					err = errors.New("未知错误")
				}
			}
			break
		}
		if PublicUserList.UserList[index].IsCanHelp() && PublicUserList.UserList[index].PhoneInfo.Phone != this.PhoneInfo.Phone {
			_, err := PublicUserList.UserList[index].OpenLuckyMoney(myELuckyMoney)
			if err != nil {
				return err
			}
			if myELuckyMoney.isUsedLucky {
				return errors.New("这个红包已经被使用过 没必要再抢了")
			}
			time.Sleep(time.Second * 5)
		} else {
			continue
		}
	}
	return err
}
func (this *UserStruct) IsCanHelp() bool {
	if len(this.PhoneInfo.Phone) == 0 {
		return false
	}
	if this.CookieInfo.OfferTimesEveryDay <= this.CookieInfo.OfferTimesToday {
		return false
	}
	return true
}
func MidNightTask() {

}
func (this *UserStruct) ChangeUid(newUid string) {
	myOldCookie := this.CookieInfo.Cookie
	myOldUid := UselessHelper.GetCookieValue(myOldCookie, "USERID")
	if len(myOldUid) < 10 && len(myOldUid) > 2 {
		myNewUid := strings.Replace(myOldCookie, myOldUid, newUid, -1)
		this.CookieInfo.Cookie = myNewUid
		return
	} else {
		this.CookieInfo.Cookie = this.CookieInfo.Cookie + ";USERID=" + newUid
	}
}

func (this *UserStruct) ChangeCookieKeyValue(key, newValue string) {
	myOldCookie := this.CookieInfo.Cookie
	myOldValue := UselessHelper.GetCookieValue(myOldCookie, key)
	if len(myOldValue) > 0 {
		myNewCookie := strings.Replace(myOldCookie, myOldValue, newValue, -1)
		this.CookieInfo.Cookie = myNewCookie
		return
	} else {
		this.CookieInfo.Cookie = this.CookieInfo.Cookie + "; " + key + "=" + newValue
	}
}
func (this *UserStruct) GetCookieTrackID() string {
	myOldCookie := this.CookieInfo.Cookie
	myOldValue := UselessHelper.GetCookieValue(myOldCookie, "track_id")
	return myOldValue
}
