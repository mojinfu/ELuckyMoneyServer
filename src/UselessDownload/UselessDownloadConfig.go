//  Created by 摸金校尉 on 17/11/20.
package UselessDownload

import (
	"time"
	"net/http"
	"net/url"
	"math/rand"
	"errors"
	"net/http/cookiejar"
)

var MyJar *cookiejar.Jar
var myProxy *url.URL
var FiddlerProxy *url.URL
//var MyClient *http.Client
var MyClient_WithoutProxy *http.Client
var transport_WithoutProxy http.Transport
var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
var myTimeout = time.Duration(30000* time.Millisecond)
type proxyStruct struct {
	Status string `json:"status"`
	Code int `json:"code"`
	Proxy []string `json:"data"`
}
var proxyServerERR =errors.New("proxyServerERR")
var jsonunmarshalERR =errors.New("jsonunmarshalERR")
var fiddlerProxyAdress string = "http://127.0.0.1:8888"