//
//  Created by 摸金校尉 on 17/11/20.
//  Copyright (c) 2018年 摸金校尉. All rights reserved.
//
package UselessHelper

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"

	//"hash/crc32"
	"os"
	"path/filepath"

	//"github.com/rlds/rlog"
	//"strconv"
	"time"

	"regexp"
	"strings"
)

func GetCookieValue(myCookie string, key string) string {
	defer RecoverMedicine("GetCookieValue")
	keyReg := regexp.MustCompile(key + `\s*=\s*([^;]+)[;]*`)
	QRURLArr := keyReg.FindAllStringSubmatch(myCookie, 1)
	if len(QRURLArr) == 0 {
		return ""
	}
	if len(QRURLArr[0]) < 2 {
		return ""
	}
	return QRURLArr[0][1]
}
func GetObjectValue(myJson string, key string) string {
	defer RecoverMedicine("getJsonValue")
	keyReg := regexp.MustCompile(key + `\s*:\s*"([^"]+)"`)
	QRURLArr := keyReg.FindAllStringSubmatch(myJson, 1)
	if len(QRURLArr) == 0 {
		return ""
	}
	if len(QRURLArr[0]) < 2 {
		return ""
	}
	return QRURLArr[0][1]
}
func GetJsonValue(myJson string, key string) string {
	defer RecoverMedicine("getJsonValue")
	keyReg := regexp.MustCompile(key + `"\s*:\s*"([^"]+)"`)
	QRURLArr := keyReg.FindAllStringSubmatch(myJson, 1)
	if len(QRURLArr) == 0 {
		return ""
	}
	if len(QRURLArr[0]) < 2 {
		return ""
	}
	return QRURLArr[0][1]
}
func TrimCannotbeseen(src string) (afterTrim string) {
	defer RecoverMedicine("TrimCannotbeseen")
	afterTrim = strings.TrimFunc(src, func(w rune) bool {
		if w < 32 {
			return true
		}
		if w == '\n' {
			return true
		}
		if w == '\t' {
			return true
		}
		if w == '\r' {
			return true
		}
		if w == ' ' {
			return true
		}
		return false
	})
	return
}
func RecoverMedicine(funcname string) {
	if err := recover(); err != nil {
		//rlog.V(1).Info(funcname+" panic:",err)
	}
}

func NowTime_s() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetAllFileData(filepath string) []byte {
	f, err := os.Open(filepath)
	if err != nil {
		//rlog.Info("GetAllFileDataErr_1:" + err.Error())
		return nil
	}
	var n int64

	if fi, err2 := f.Stat(); err2 == nil {
		if size := fi.Size(); size < 1e9 {
			n = size
		}
	} else {
		return nil
	}
	buf := bytes.NewBuffer(make([]byte, 0, n+bytes.MinRead))
	defer buf.Reset()
	_, err = buf.ReadFrom(f)
	f.Close()
	if err != nil {
		//rlog.Info("GetAllFileDataErr_2:" + err.Error())
		return nil
	}
	return buf.Bytes()
}

//存储并替换文件内容
func SaveReplaceFile(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	f.Write(data)
	f.Close()
	return nil
}

//判断文件是否存在
func IsFile(filepath string) error {
	_, err := os.Stat(filepath)
	return err
}

func DelFile(path string) error {
	return os.Remove(path)
}

func MkAlldir(dir string) bool {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		//rlog.V(1).Info("建立文件夹出现错误[" + err.Error() + "]")
		return false
	}
	return true
}

//删除文件夹
func DelDir(path string) error {
	return os.RemoveAll(path)
}
func GetMd5Str(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func GetConfPath(ipath string) (path string) {
	//file, _ := exec.LookPath(os.Args[0])
	path, _ = filepath.Abs(ipath)
	// fmt.Println(path)
	// fmt.Println(path + "/../" + ipath)
	// cpath, _ = filepath.Abs(path + "/../" + ipath)
	return
}
