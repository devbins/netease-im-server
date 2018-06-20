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
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

type Info struct {
	Token string `json:"token"`
	Accid string `json:"accid"`
	Name  string `json:"name"`
}

type TokenRespose struct {
	BaseResp
	Info `json:"info"`
}

// Create ...
func Create(accid string) (*TokenRespose, error) {
	req, err := http.NewRequest("POST", ACTION_CREATE, strings.NewReader("accid="+accid))
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
	token := &TokenRespose{}
	err = json.Unmarshal(resBody, token)
	if err != nil {
		return nil, err
	}
	return token, nil

}

// Update ...
func Update(accid string) (*BaseResp, error) {
	req, err := http.NewRequest("POST", ACTION_UPDATE, strings.NewReader("accid="+accid))
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

	resCode := &BaseResp{}

	err = json.Unmarshal(resBody, resCode)
	if err != nil {
		return nil, err
	}
	return resCode, nil

}

// RefreshToken ...
func RefreshToken(accid string) (*TokenRespose, error) {
	reqBody := url.Values{"accid": {accid}}
	req, err := http.NewRequest("POST", ACTION_REFRESH_TOKEN, strings.NewReader(reqBody.Encode()))
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
	token := &TokenRespose{}
	err = json.Unmarshal(resBody, token)
	if err != nil {
		return nil, err
	}

	return token, nil

}

// Block ...
func Block(accid string) (*BaseResp, error) {
	req, err := http.NewRequest("POST", ACTION_BLOCK, strings.NewReader("accid="+accid))
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
	data := &BaseResp{}
	err = json.Unmarshal(resBody, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UnBlock ...
func UnBlock(accid string) (*BaseResp, error) {
	req, err := http.NewRequest("POST", ACTION_UNBLOCK, strings.NewReader("accid="+accid))
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
	data := &BaseResp{}
	err = json.Unmarshal(resBody, data)
	if err != nil {
		return nil, err
	}
	return data, nil

}

// getCheckSUm ...
func getCheckSUm(appSecret, nonce, curTime string) string {
	sum := appSecret + nonce + curTime
	h := sha1.New()
	h.Write([]byte(sum))
	result := h.Sum(nil)
	return fmt.Sprintf("%x", result)
}

// fillHeader ...
func fillHeader(req *http.Request) {
	curTime := strconv.Itoa(int(time.Now().Unix()))
	req.Header.Add("AppKey", APPKEY)
	req.Header.Add("Nonce", NONCE)
	req.Header.Add("CurTime", curTime)
	req.Header.Add("CheckSum", getCheckSUm(APPSECRET, NONCE, curTime))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}
