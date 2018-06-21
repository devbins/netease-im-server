package api

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	APPKEY                               = ""
	APPSECRET                            = ""
	NONCE                                = ""
	BASE_URL                             = "https://api.netease.im/nimserver/"
	ACTION_USER_CREATE                   = BASE_URL + "user/create.action"               // 创建云通信ID
	ACTION_USER_UPDATE                   = BASE_URL + "user/update.action"               // 更新云通信ID
	ACTION_USER_REFRESH_TOKEN            = BASE_URL + "user/refreshToken.action"         // 更新并获取token
	ACTION_USER_BLOCK                    = BASE_URL + "user/block.action"                // 封禁网易云通信ID
	ACTION_USER_UNBLOCK                  = BASE_URL + "user/unblock.action"              // 解禁网易云通信ID
	ACTION_USER_DONNOP_OPEN              = BASE_URL + "user/setDonnop.action"            // 设置桌面端在线时，移动端是否需要推送
	ACTION_USER_UPDATE_UINFO             = BASE_URL + "user/updateUinfo.action"          // 更新用户名片
	ACTION_USER_GET_UINFO                = BASE_URL + "user/getUinfos.action"            // 获取用户名片
	ACTION_USER_SET_SPECIAL_RELATION     = BASE_URL + "user/setSpecialRelation.action"   // 设置黑名单/静音
	ACTION_USER_LIST_BLACK_AND_MUTE_LIST = BASE_URL + "user/listBlackAndMuteList.action" // 查看指定用户的黑名单和静音列表

	ACTION_FRIEND_ADD    = BASE_URL + "friend/add.action"    // 加好友
	ACTION_FRIEND_UPDATE = BASE_URL + "friend/update.action" // 更新好友相关信息
	ACTION_FRIEND_DELETE = BASE_URL + "friend/delete.action" // 删除好友
	ACTION_FRIEND_GET    = BASE_URL + "friend/get.action"    // 获取好友关系
)

var client = http.Client{}

type BaseResp struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

type Info struct {
	Token  string `json:"token"`
	Accid  string `json:"accid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender int    `json:"gender"`
	Mobile string `json:"mobile"`
}

type TokenRespose struct {
	BaseResp
	Info `json:"info"`
}

type Uinfos struct {
	BaseResp
	Uinfos []Info `json:"uinfos"`
}

// Create ...
func Create(accid string) (*TokenRespose, error) {
	req, err := http.NewRequest("POST", ACTION_USER_CREATE, strings.NewReader("accid="+accid))
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
	req, err := http.NewRequest("POST", ACTION_USER_UPDATE, strings.NewReader("accid="+accid))
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
	req, err := http.NewRequest("POST", ACTION_USER_REFRESH_TOKEN, strings.NewReader(reqBody.Encode()))
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
	req, err := http.NewRequest("POST", ACTION_USER_BLOCK, strings.NewReader("accid="+accid))
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
	req, err := http.NewRequest("POST", ACTION_USER_UNBLOCK, strings.NewReader("accid="+accid))
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

// UpdateUinfo ...
func UpdateUinfo(params url.Values) (*BaseResp, error) {
	data, err := ResponseResult(ACTION_USER_UPDATE_UINFO, params)
	if err != nil {
		return nil, err
	}
	result := &BaseResp{}
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// GetUinfo ...
func GetUinfo(accids ...string) (*Uinfos, error) {
	data, err := ResponseResult(ACTION_USER_GET_UINFO, url.Values{"accids": accids})
	if err != nil {
		return nil, err
	}
	infos := &Uinfos{}
	err = json.Unmarshal(data, infos)
	if err != nil {
		return nil, err
	}
	return infos, nil

}

// SetDonnop ...
func SetDonnop(accid string, donnopOpen bool) (*BaseResp, error) {
	params := url.Values{"accid": {accid}, "donnopOpen": {strconv.FormatBool(donnopOpen)}}
	data, err := ResponseResult(ACTION_USER_DONNOP_OPEN, params)
	if err != nil {
		return nil, err
	}
	result := &BaseResp{}
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// ResponseResult ...
func ResponseResult(action string, params url.Values) ([]byte, error) {
	req, err := http.NewRequest("POST", action, strings.NewReader(params.Encode()))
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
	return resBody, nil

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
