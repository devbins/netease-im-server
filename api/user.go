package api

import (
	"crypto/sha1"
	"fmt"
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

// refreshToken ...
func refreshToken(accid string) {
	req, err := http.NewRequest("POST", ACTION_REFRESH_TOKEN, strings.NewReader("accid="+accid))
	if err != nil {
		fmt.Println(err)
	}
	curTime := strconv.Itoa(int(time.Now().Unix()))
	req.Header.Add("AppKey", APPKEY)
	req.Header.Add("Nonce", NONCE)
	req.Header.Add("CurTime", curTime)
	req.Header.Add("CheckSum", getCheckSUm(APPSECRET, NONCE, curTime))
	req.Header.Add("Cotent-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

}

// getCheckSUm ...
func getCheckSUm(appSecret, nonce, curTime string) string {
	sum := appSecret + nonce + curTime
	h := sha1.New()
	h.Write([]byte(sum))
	resutlt := h.Sum(nil)
	return string(resutlt)
}
