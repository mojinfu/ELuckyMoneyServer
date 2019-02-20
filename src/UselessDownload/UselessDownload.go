//  Created by 摸金校尉 on 17/11/20.
package UselessDownload

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	//"rlog"
	"net"
	"net/url"
	"strings"
	// /"net/http/cookiejar"
	//"golang.org/x/net/publicsuffix"
)

func init() {
	var err error
	//MyJar, err = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if nil != err {
		fmt.Println("cooke jar init err")
	}
	FiddlerProxy, err = url.Parse(fiddlerProxyAdress)
	if nil != err {
		fmt.Println(err.Error())
	}
	var transport_WithoutProxy = http.Transport{
		TLSHandshakeTimeout: myTimeout,
		Dial:                dialTimeout,
		DisableKeepAlives:   true,
		//Proxy:               http.ProxyURL(FiddlerProxy), //saya
	}
	MyClient_WithoutProxy = &http.Client{
		Transport: &transport_WithoutProxy,
		//Jar :MyJar,
		Timeout: myTimeout,
	}
}
func InitSelf() {
	//nice to old
}
func dialTimeout(network, addr string) (net.Conn, error) {
	deadline := time.Now().Add(myTimeout)
	c, err := net.DialTimeout(network, addr, myTimeout)
	if err != nil {
		return nil, err
	}
	c.SetDeadline(deadline)
	return c, nil
}

func Download_POST(myUrl string, vv url.Values) (string, bool) {
	httpreq, errr := http.NewRequest("POST", myUrl, strings.NewReader(vv.Encode()))
	if errr != nil {
		//rlog.V(1).Info("Download_POST  NewRequest err"+errr.Error() )
		return "", false
	}
	resp, errr := MyClient_WithoutProxy.Do(httpreq)
	if resp != nil {
		defer resp.Body.Close()
	}
	if errr != nil {
		//rlog.V(1).Info("Download_POST_Welsh http post err"+errr.Error())
		return "", false
	}
	body, errr := ioutil.ReadAll(resp.Body)
	if errr != nil {
		//rlog.V(1).Info("Download_POST_Welsh ioutil err"+errr.Error())
		return "", false
	}
	return string(body), true
}

func Download_GET(myUrl string) (string, bool) {
	httpreq, errr := http.NewRequest("GET", myUrl, nil)
	if errr != nil {
		//rlog.V(1).Info("Download_GET  NewRequest err"+errr.Error() )
		return "", false
	}
	httpreq.Header.Add("Connection", "close")
	resp, errr := MyClient_WithoutProxy.Do(httpreq)
	if resp != nil {
		defer resp.Body.Close()
	}
	if errr != nil {
		//rlog.V(1).Info("Download_GET http get err"+errr.Error())
		return "", false
	}
	body, errr := ioutil.ReadAll(resp.Body)
	if errr != nil {
		return "", false
	}
	return string(body), true
}

func Download_GET302Location(myUrl string) (string, bool) {
	httpreq, errr := http.NewRequest("GET", myUrl, nil)
	if errr != nil {
		//rlog.V(1).Info("Download_GETLocation  NewRequest err"+errr.Error() )
		return "", false
	}
	httpreq.Header.Add("Connection", "keep-alive")
	resp, errr := MyClient_WithoutProxy.Do(httpreq)
	if resp != nil {
		defer resp.Body.Close()
	}
	if errr != nil {
		//rlog.V(1).Info("Download_GET http get err"+errr.Error())
		return "", false
	}
	if resp.StatusCode == 302 {
		LocationUrl, errr := resp.Location()
		if errr == nil {
			return LocationUrl.Path, true
		}
	} else {
		//rlog.V(1).Info("resp.StatusCode:",resp.StatusCode)
	}
	return "", false
}
func Download_PUT_Json(myUrl string, myJosnBody string, myCookie string) (string, bool) {
	httpreq, errr := http.NewRequest("PUT", myUrl, strings.NewReader(myJosnBody))
	if errr != nil {
		//rlog.V(1).Info("Download_Put_Json  NewRequest err"+errr.Error() )
		return "", false
	}
	httpreq.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 7.0; SM-G9200 Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/53.0.2785.49 Mobile MQQBrowser/6.2 TBS/043632 Safari/537.36 MicroMessenger/6.5.7.1041 NetType/WIFI Language/zh_CN")
	httpreq.Header.Add("Referer", "https://h5.ele.me/hongbao/?from=singlemessage&isappinstalled=0")
	httpreq.Header.Add("Cookie", myCookie)

	resp, errr := MyClient_WithoutProxy.Do(httpreq)
	if resp != nil {
		defer resp.Body.Close()
	}
	if errr != nil {
		//rlog.V(1).Info("Download_Put_Json http post err"+errr.Error())
		return "", false
	}
	body, errr := ioutil.ReadAll(resp.Body)
	if errr != nil {
		//rlog.V(1).Info("Download_Put_Json ioutil err"+errr.Error())
		return "", false
	}
	return string(body), true
}

func Download_POST_Json(myUrl string, myJosnBody string, myeosid string, myCookie string) (string, bool) {
	httpreq, errr := http.NewRequest("POST", myUrl, strings.NewReader(myJosnBody))
	if errr != nil {
		//rlog.V(1).Info("Download_POST  NewRequest err"+errr.Error() )
		return "", false
	}
	httpreq.Header.Add("Content-Type", "text/plain;charset=UTF-8")
	httpreq.Header.Add("Accept", "*/*")
	httpreq.Header.Add("Accept-Encoding", "")
	httpreq.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 7.0; SM-G9200 Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/53.0.2785.49 Mobile MQQBrowser/6.2 TBS/043632 Safari/537.36 MicroMessenger/6.5.7.1041 NetType/WIFI Language/zh_CN")
	httpreq.Header.Add("Referer", "https://h5.ele.me/hongbao/?from=singlemessage&isappinstalled=0")
	httpreq.Header.Add("X-Shard", "eosid="+myeosid)
	httpreq.Header.Add("Cookie", myCookie)

	resp, errr := MyClient_WithoutProxy.Do(httpreq)
	if resp != nil {
		defer resp.Body.Close()
	}
	if errr != nil {
		//rlog.V(1).Info("Download_POST_Welsh http post err"+errr.Error())
		return "", false
	}
	body, errr := ioutil.ReadAll(resp.Body)
	if errr != nil {
		//rlog.V(1).Info("Download_POST_Welsh ioutil err"+errr.Error())
		return "", false
	}
	//resp.Cookies()
	return string(body), true
}

func DownloadPOSTJsonCollectCookie(myUrl string, myJosnBody string, myCookie string) (string, bool, []*http.Cookie) {
	httpreq, errr := http.NewRequest("POST", myUrl, strings.NewReader(myJosnBody))
	if errr != nil {
		//rlog.V(1).Info("Download_POST  NewRequest err"+errr.Error() )
		return "", false, nil
	}
	httpreq.Header.Add("Content-Type", "text/plain;charset=UTF-8")
	httpreq.Header.Add("Accept", "*/*")
	httpreq.Header.Add("Accept-Encoding", "")
	httpreq.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 7.0; SM-G9200 Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/53.0.2785.49 Mobile MQQBrowser/6.2 TBS/043632 Safari/537.36 MicroMessenger/6.5.7.1041 NetType/WIFI Language/zh_CN")
	httpreq.Header.Add("Referer", "https://h5.ele.me/hongbao/?from=singlemessage&isappinstalled=0")
	httpreq.Header.Add("Cookie", myCookie)

	resp, errr := MyClient_WithoutProxy.Do(httpreq)
	if resp != nil {
		defer resp.Body.Close()
	}
	if errr != nil {
		//rlog.V(1).Info("Download_POST_Welsh http post err"+errr.Error())
		return "", false, nil
	}
	body, errr := ioutil.ReadAll(resp.Body)
	if errr != nil {
		//rlog.V(1).Info("Download_POST_Welsh ioutil err"+errr.Error())
		return "", false, nil
	}
	//resp.Cookies()
	return string(body), true, resp.Cookies()
}
