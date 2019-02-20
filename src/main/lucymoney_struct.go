package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/rlds/rlog"
)

type luckyMoneyRespStruct struct {
	Message           string                    `json:"message"`
	Name              string                    `json:"name"`
	Promotion_records []Promotion_recordsStruct `json:"promotion_records"`
	Account           string                    `json:"account"`
	Ret_code          int
}
type luckyMoneyRequestStruct struct {
	Device_id       string `json:"device_id"`
	Group_sn        string `json:"group_sn"`
	Hardware_id     string `json:"hardware_id"`
	Method          string `json:"method"`
	Phone           string `json:"phone"`
	Platform        int    `json:"platform"`
	Sign            string `json:"sign"`
	Track_id        string `json:"track_id"`
	Unionid         string `json:"unionid"`
	Weixin_avatar   string `json:"weixin_avatar"`
	Weixin_username string `json:"weixin_username"`
}

//luckyMoneyRequestStruct
// {
//     "method": "phone",
//     "group_sn": "2a0b7a45559d6415",
//     "sign": "d7bd99eaf4bdaec6ba96220d99c6776e",
//     "phone": "",
//     "device_id": "",
//     "hardware_id": "",
//     "platform": 0,
//     "track_id": "1536291426|69395bec3a8e693fe23b6564c998fa34785f8a8f38b7dc5335|fe2368f9da9d8aefbce7ddc40a2a1cbe",
//     "weixin_avatar": "http://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTJYk0mY3BuCX8sfBkQwxAqvdniawPnqJXNx2JuXNibMicvzdjz7BgfsibYcJgXjYLBLxgviakO9rBL2c0g/132",
//     "weixin_username": "颜色1934",
//     "unionid": "o_PVDuIDTHd67WEarVE7vWW22XdM",
//     "latitude": 29.3484490185984,
//     "longitude": 119.7525971711162
// }

type ELuckyMoneyStruct struct {
	LuckyBoyPhone     string `json:"phone"`
	EUrl              string `json:"e_url"`
	HungryManName     string `json:"name"`
	registerSuccess   bool
	luckyMoneySuccess bool

	lucky_number int
	platform     string
	sn           string
	EGonaLucky   bool
	isUsedLucky  bool
}

type HungryManStruct struct {
	WechatCookie string `json:"WechatCookie"`
	Phone        string `json:"Phone"`
	openid       string
	eleme_key    string
	realHungry   bool
	nickname     string
}

const myEOrigin string = "https://h5.ele.me"

type registerStruct struct {
	Sign  string `json:"sign"`
	Phone string `json:"phone"`
}

type luckydoorReturnJsonStruct struct {
	Promotion_records []Promotion_recordsStruct `json:"promotion_records"`
	Account           string                    `json:"account"`
	Ret_code          int                       `json:"ret_code"`
}

type Promotion_recordsStruct struct {
	Amount       float32 `json:"amount"`
	Created_at   int     `json:"created_at"`
	Is_lucky     bool    `json:"is_lucky"`
	Sns_avatar   string  `json:"sns_avatar"`
	Sns_username string  `json:"sns_username"`
}

const alreadyCome int = 2
const successEdoor4 int = 4
const successEdoor3 int = 3
const failureEdoor int = 1

const fivemost int = 5

func (this *luckyMoneyRequestStruct) converjson() []byte {
	body, err := json.Marshal(this)
	if err != nil {
		rlog.V(1).Info("func converjson error:" + err.Error())
		return nil
	}
	return body
}
func UrlSTD(myUrl string) string {
	myUrl = strings.Replace(myUrl, "#", "?", 1)
	// if strings.Count(myUrl, "&") == strings.Count(myUrl, "=") {
	// 	myUrl = strings.Replace(myUrl, "&", "?", 1)
	// }
	return myUrl
}
func NewELuckyMoney(myUrl string) (*ELuckyMoneyStruct, error) {
	myELuckyMoney := &ELuckyMoneyStruct{
		EUrl: myUrl,
	}
	myELuckyMoney.EUrl = UrlSTD(myELuckyMoney.EUrl)
	fmt.Println("myELuckyMoney.EUrl:", myELuckyMoney.EUrl)
	myEUrl, idParseOK := url.Parse(myELuckyMoney.EUrl)
	if idParseOK != nil {
		return nil, errors.New("idParseOK:" + idParseOK.Error())
	}
	myELuckyMoney.sn = myEUrl.Query().Get("sn")
	fmt.Println("lucky_number:", myEUrl.Query())
	lucky_number, err := strconv.Atoi(myEUrl.Query().Get("lucky_number"))
	if err != nil {
		return nil, errors.New("Atoi:" + err.Error())
	}
	myELuckyMoney.lucky_number = lucky_number
	return myELuckyMoney, nil
}
