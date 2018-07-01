package nimserversdk

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Nim struct {
	APPKEY    string
	APPSECRET string
	User      User
	Friend    Friend
	Msg       Msg
	Team      Team
}

// New ...
func NewNim(appkey, appsecret string) *Nim {
	return &Nim{
		APPKEY:    appkey,
		APPSECRET: appsecret,
		User: User{
			APPKEY:    appkey,
			APPSECRET: appsecret,
		},
		Friend: Friend{
			APPKEY:    appkey,
			APPSECRET: appsecret,
		},
		Msg: Msg{
			APPKEY:    appkey,
			APPSECRET: appsecret,
		},
		Team: Team{
			APPKEY:    appkey,
			APPSECRET: appsecret,
		},
	}

}

// ResponseResult ...
func ResponseResult(appkey string, appSecret string, action string, params url.Values) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, action, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	fillHeader(req, appkey, appSecret)
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
func fillHeader(req *http.Request, appkey, appsecret string) {
	curTime := strconv.Itoa(int(time.Now().Unix()))
	nonce := strconv.Itoa(rand.Int())
	req.Header.Add("AppKey", appkey)
	req.Header.Add("Nonce", nonce)
	req.Header.Add("CurTime", curTime)
	req.Header.Add("CheckSum", getCheckSUm(appsecret, nonce, curTime))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}
