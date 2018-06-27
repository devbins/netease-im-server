package nimserversdk

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Chatroom struct {
	APPKEY    string
	APPSECRET string
}

type ChatroomResult struct {
	BaseResp
	ChatroomInfo ChatroomInfo `json:"chatroom"`
}

type ChatroomInfo struct {
	Roomid       int    `json:"roomid"`
	Valid        bool   `json:"valid"`
	Announcement string `json:"announcement"`
	Name         string `json:"name"`
	Broadcasturl string `json:"broadcasturl"`
	Ext          string `json:"ext"`
	Creator      string `json:"creator"`
}

type ChatroomDetailInfoResult struct {
	BaseResp
	ChatroomDetailInfo ChatroomDetailInfo `json:"chatroom"`
}

type ChatroomDetailInfo struct {
	Roomid          int    `json:"roomid"`
	Valid           bool   `json:"valid"`
	Muted           bool   `json:"muted"`
	Announcement    string `json:"announcement"`
	Name            string `json:"name"`
	Broadcasturl    string `json:"broadcasturl"`
	Onlineusercount int    `json:"onlineusercount"`
	Ext             string `json:"ext"`
	Creator         string `json:"creator"`
	Queuelevel      int    `json:"queuelevel"`
}

// Create ...
func (chatroom *Chatroom) Create(creator string, name string, announcement string, broadcasturl string, ext string, queuelevel int) (*ChatroomResult, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_CREATE, url.Values{"creator": {creator}, "name": {name}, "announcement": {announcement}, "broadcasturl": {broadcasturl}, "ext": {ext}, "queuelevel": {strconv.Itoa(queuelevel)}})
	if err != nil {
		return nil, err
	}
	chatroomResult := &ChatroomResult{}
	err = json.Unmarshal(res, chatroomResult)
	if err != nil {
		return nil, err
	}
	return chatroomResult, nil

}

// Get ...
func (chatroom *Chatroom) Get(romid int64, needOnlineUserCount bool) (*ChatroomDetailInfoResult, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_GET, url.Values{"roomid": {strconv.FormatInt(romid, 10)}, "needOnlineUserCount": {strconv.FormatBool(needOnlineUserCount)}})
	if err != nil {
		return nil, err
	}
	chatroomDetailInfoResult := &ChatroomDetailInfoResult{}
	err = json.Unmarshal(res, chatroomDetailInfoResult)
	if err != nil {
		return nil, err
	}
	return chatroomDetailInfoResult, nil

}
