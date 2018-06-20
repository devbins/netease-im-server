package api

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	APPKEY               = ""
	APPSECRET            = ""
	NONCE                = ""
	BASE_URL             = "https://api.netease.im/nimserver/"
	ACTION_CREATE        = BASE_URL + "user/create.action"       // 创建云通信ID
	ACTION_UPDATE        = BASE_URL + "user/update.action"       // 更新云通信ID
	ACTION_REFRESH_TOKEN = BASE_URL + "user/refreshToken.action" // 更新并获取token
	ACTION_BLOCK         = BASE_URL + "user/block.action"        // 封禁网易云通信ID
	ACTION_UNBLOCK       = BASE_URL + "user/unblock.action"      // 解禁网易云通信ID
)

var client = http.Client{}

type BaseResp struct {
	Code int `json:"code"`
}

type Info struct {
	Token string `json:"token"`
	Accid string `json:"accid"`
	Name  string `json:"name"`
}

type TokenRes struct {
	BaseResp
	Info `json:"info"`
}

// create ...
func create(accid, name string) {
	req, err := http.NewRequest("POST", ACTION_CREATE, strings.NewReader("accid="+accid+"&name="+name))
	if err != nil {
		fmt.Println(err)
	}
	fillHeader(req)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resBody)

}

// fillHeader ...
func fillHeader(req *http.Request) {
	curTime := strconv.Itoa(int(time.Now().Unix()))
	req.Header.Add("AppKey", APPKEY)
	req.Header.Add("Nonce", NONCE)
	req.Header.Add("CurTime", curTime)
	req.Header.Add("CheckSum", getCheckSUm(APPSECRET, NONCE, curTime))
	req.Header.Add("Cotent-Type", "application/x-www-form-urlencoded")
}

// RefreshToken ...
func RefreshToken(accid string) (*TokenRes, error) {
	req, err := http.NewRequest("POST", ACTION_REFRESH_TOKEN, strings.NewReader("accid="+accid))
	if err != nil {
		return nil, err
	}

	fillHeader(req)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	token := &TokenRes{}
	err = json.Unmarshal(resBody, token)
	if err != nil {
		return nil, err
	}

	return token, nil

}

// getCheckSUm ...
func getCheckSUm(appSecret, nonce, curTime string) string {
	sum := appSecret + nonce + curTime
	h := sha1.New()
	h.Write([]byte(sum))
	resutlt := h.Sum(nil)
	return string(resutlt)
}
