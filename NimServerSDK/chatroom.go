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
	ChatroomInfo `json:"chatroom"`
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
