package main

import (
	"ELuckyMoneyServer/src/UselessHelper"
	"fmt"
	"net/http"
	"testing"
)

func TestNewELuckyMoney(t *testing.T) {
	t.Error("AC")
	A, B := NewELuckyMoney("https://h5.ele.me/hongbao/#hardware_id=&is_lucky_group=True&lucky_number=6&track_id=&platform=0&sn=2a08bc52482cf499&theme_id=3049&device_id=&refer_user_id=4339802")
	fmt.Println("A:", A)
	fmt.Println("B:", B)
}

func TestGetElemeKey(t *testing.T) {
	t.Error("AC")
	myUser := getUser()
	A := myUser.GetElemeKey()
	fmt.Println("A:", A)
}

func TestOpenLuckyMoney(t *testing.T) {
	t.Error("AC")
	PublicUserList.LoadData()

	myUser := PublicUserList.GetUserByUserName("yqr4")
	fmt.Println(myUser.GetElemeKey())
	myELuckyMoney, _ := NewELuckyMoney("https://h5.ele.me/hongbao/#hardware_id=&is_lucky_group=True&lucky_number=6&track_id=&platform=0&sn=2a08bc52482cf499&theme_id=3049&device_id=&refer_user_id=4339802")
	A, B := myUser.OpenLuckyMoney(myELuckyMoney)
	fmt.Println("A:", A)
	fmt.Println("A:", A.Message)
	fmt.Println("B:", B)
}

func TestBindPhoneSendSMS(t *testing.T) {
	t.Error("AC")
	PublicUserList.LoadData()
	myUser := PublicUserList.GetEmptyPhoneUser()
	SMSUserMap = make(map[string]*UserStruct)
	A, B := myUser.BindPhoneSendSMS(
		&SendSMSRequestStruct{
			Mobile: "15967184143"})
	fmt.Println("A:", A)
	fmt.Println("B:", B)
}

func TestCheckSMS(t *testing.T) {
	t.Error("AC")
	PublicUserList.LoadData()
	myUser := PublicUserList.GetEmptyPhoneUser()
	SMSUserMap = make(map[string]*UserStruct)
	SMSUserMap["30385d29334089037e2e4194d44a67eac2d649cb0d5164a4dc716ea282a6ca47"] = myUser
	myVipCheckSMSRequest := &VipCheckSMSRequestStruct{
		Phone:         "15967184143",
		ValidateToken: "331586",
		ValidateCode:  "30385d29334089037e2e4194d44a67eac2d649cb0d5164a4dc716ea282a6ca47",
		PhoneOwner:    "mojinfu self",
	}
	CheckSMS(
		myVipCheckSMSRequest,
	)
	fmt.Printf("myUser:%+v\n", myUser)
}
func TestOutputData(t *testing.T) {
	t.Error("AC")
	PublicUserList.LoadData()
	PublicUserList.OutputData()
	// myUser := PublicUserList.GetEmptyPhoneUser()
	// fmt.Println(myUser.CookieInfo.Cookie)
}
func TestGetLastUserListFilePath(t *testing.T) {
	t.Error("AC")
	fmt.Println(GetLastUserListFilePath("./"))
}
func TestGetCookieValue(t *testing.T) {
	t.Error("AC")
	fmt.Println(UselessHelper.GetCookieValue("track_id=1536291426|69395bec3a8e693fe23b6564c998fa34785f8a8f38b7dc5335|fe2368f9da9d8aefbce7ddc40a2a1cbe; SID=fmO8EPeH1aVXdQswrLaySC9bcUpOqzx6mshg; USERID=4339802; _utrace=7781a14c03acacba170d423ddb2a0c0e_2018-08-09; snsInfo[wx2a416286e96100ed]=%7B%22city%22%3A%22%E6%9D%AD%E5%B7%9E%22%2C%22country%22%3A%22%E4%B8%AD%E5%9B%BD%22%2C%22eleme_key%22%3A%22d7bd99eaf4bdaec6ba96220d99c6776e%22%2C%22headimgurl%22%3A%22http%3A%2F%2Fthirdwx.qlogo.cn%2Fmmopen%2Fvi_32%2FQ0j4TwGTfTJYk0mY3BuCX8sfBkQwxAqvdniawPnqJXNx2JuXNibMicvzdjz7BgfsibYcJgXjYLBLxgviakO9rBL2c0g%2F132%22%2C%22language%22%3A%22zh_CN%22%2C%22nickname%22%3A%22%E9%A2%9C%E8%89%B21934%22%2C%22openid%22%3A%22oEGLvjgfRqfHoYNHx0IxY1EwjLMg%22%2C%22privilege%22%3A%5B%5D%2C%22province%22%3A%22%E6%B5%99%E6%B1%9F%22%2C%22sex%22%3A1%2C%22unionid%22%3A%22o_PVDuIDTHd67WEarVE7vWW22XdM%22%2C%22name%22%3A%22%E9%A2%9C%E8%89%B21934%22%2C%22avatar%22%3A%22http%3A%2F%2Fthirdwx.qlogo.cn%2Fmmopen%2Fvi_32%2FQ0j4TwGTfTJYk0mY3BuCX8sfBkQwxAqvdniawPnqJXNx2JuXNibMicvzdjz7BgfsibYcJgXjYLBLxgviakO9rBL2c0g%2F132%22%7D; ubt_ssid=r1mrxh8blwtcd5gbl54qsacysewvl860_2018-08-09; perf_ssid=qek2wdc82t10tbh0p7ja67rna74bg6eg_2018-08-09", "USERID"))
}
func TestGetPictureType(t *testing.T) {
	t.Error("AC")
	fmt.Println(getPictureType("data:image/jpeg;base64"))
}
func TestChangeCookieKeyValue(t *testing.T) {
	t.Error("AC")
	myUser := &UserStruct{
		CookieInfo: CookieInfoStruct{
			Cookie: ";USERID=2114457242",
		},
	}
	myCookies := []*http.Cookie{
		&http.Cookie{Name: "A", Value: "A"},
		&http.Cookie{Name: "B", Value: "B"},
	}
	for index := range myCookies {
		fmt.Println(myCookies[index].Name)
		fmt.Println(myCookies[index].Value)
		myUser.ChangeCookieKeyValue(myCookies[index].Name, myCookies[index].Value)
	}
	fmt.Println(myUser.CookieInfo.Cookie)
}
