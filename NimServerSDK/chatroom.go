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

type BatchChatroomResult struct {
	BaseResp
	NoExistRooms []int64              `json:"noExistRooms"`
	FailRooms    []int64              `json:"failRooms"`
	SuccRooms    []ChatroomDetailInfo `json:"succRooms"`
}

type ToggleCloseStatResutl struct {
	BaseResp
	ChatroomInfo ChatroomInfo `json:"desc"`
}

type SetMember struct {
	Roomid int64  `json:"roomid"`
	Level  int    `json:"level"`
	Accid  string `json:"accid"`
	Type   string `json:"type"`
}

type SetMemberResult struct {
	SetMember SetMember `json:"desc"`
	BaseResp
}

type AddrResult struct {
	Addr []string `json:"addr"`
	BaseResp
}

type ChatroomSendMsgResult struct {
	BaseResp
	Desc SendMsgResult `json:"desc"`
}

type SendMsgResult struct {
	Time             string `json:"time"`
	FromAvator       string `json:"fromAvator"`
	Msgid_client     string `json:"msgid_client"`
	FromClientType   string `json:"fromClientType"`
	Attach           string `json:"attach"`
	RoomId           string `json:"roomId"`
	FromAccount      string `json:"fromAccount"`
	FromNick         string `json:"fromNick"`
	Type             string `json:"type"`
	Ext              string `json:"ext"`
	HighPriorityFlag int    `json:"highPriorityFlag"`
}

type RobotDesc struct {
	FailAccids    string `json:"failAccids"`
	SuccessAccids string `json:"successAccids"`
	OldAccids     string `json:"oldAccids"`
}

type RobotResult struct {
	BaseResp
	RobotDesc `json:"desc"`
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

// GetBatch ...
func (chatroom *Chatroom) GetBatch(roomids string, needOnlineUserCount bool) (*BatchChatroomResult, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_GET_BATCH, url.Values{"roomids": {roomids}, "needOnlineUserCount": {strconv.FormatBool(needOnlineUserCount)}})
	if err != nil {
		return nil, err
	}
	batchChatroomResult := &BatchChatroomResult{}
	err = json.Unmarshal(res, batchChatroomResult)
	if err != nil {
		return nil, err
	}
	return batchChatroomResult, nil

}

// Update ...
func (chatroom *Chatroom) Update(roomid int64, name string, announcement string, broadcasturl string, ext string, needNotify bool, notifyExt string, queuelevel int) (*ChatroomResult, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_UPDATE, url.Values{"roomid": {strconv.FormatInt(roomid, 10)}, "name": {name}, "announcement": {announcement}, "broadcasturl": {broadcasturl}, "ext": {ext}, "needNotify": {strconv.FormatBool(needNotify)}, "notifyExt": {notifyExt}, "queuelevel": {strconv.Itoa(queuelevel)}})
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

// name ...
func (chatroom *Chatroom) ToggleCloseStat(roomid int64, operator string, valid bool) (*ToggleCloseStatResutl, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_TOGGLE_CLOSE_STAT, url.Values{"roomid": {strconv.FormatInt(roomid, 10)}, "operator": {operator}, "valid": {strconv.FormatBool(valid)}})
	if err != nil {
		return nil, err
	}
	toggleCloseStat := &ToggleCloseStatResutl{}
	err = json.Unmarshal(res, toggleCloseStat)
	if err != nil {
		return nil, err
	}
	return toggleCloseStat, nil

}

// SetMemberRole ...
func (chatroom *Chatroom) SetMemberRole(roomid int64, operator string, target string, opt int, optvalue bool, notifyExt string) (*SetMemberResult, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_SET_MEMBER_ROLE, url.Values{"roomid": {strconv.FormatInt(roomid, 10)}, "operator": {operator}, "target": {target}, "opt": {strconv.Itoa(opt)}, "optvalue": {strconv.FormatBool(optvalue)}, "notifyExt": {notifyExt}})
	if err != nil {
		return nil, err
	}
	setMemberResult := &SetMemberResult{}
	err = json.Unmarshal(res, setMemberResult)
	if err != nil {
		return nil, err
	}
	return setMemberResult, nil

}

// name ...
func (chatroom *Chatroom) RequestAddr(roomid int64, accid string, clienttype int) (*AddrResult, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_REQUEST_ADDR, url.Values{"roomid": {strconv.FormatInt(roomid, 10)}, "accid": {accid}, "clienttype": {strconv.Itoa(clienttype)}})
	if err != nil {
		return nil, err
	}
	addrResult := &AddrResult{}
	err = json.Unmarshal(res, addrResult)
	if err != nil {
		return nil, err
	}
	return addrResult, nil

}

// SendMsg ...
func (chatroom *Chatroom) SendMsg(roomid int64, msgId string, fromAccid string, msgType int, resendFlag int, attach string, ext string, antispam string, antispamCustom string, skipHistory int, bid string, highPriority bool, useYidun int, needHighPriorityMsgResend bool) (*ChatroomSendMsgResult, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_SEND_MSG, url.Values{"roomid": {strconv.FormatInt(roomid, 10)}, "msgId": {msgId}, "fromAccid": {fromAccid}, "msgType": {strconv.Itoa(msgType)}, "attach": {attach}, "ext": {ext}, "antispam": {antispam}, "antispamCustom": {antispamCustom}, "skipHistory": {strconv.Itoa(skipHistory)}, "bid": {bid}, "highPriority": {strconv.FormatBool(highPriority)}, "useYidun": {strconv.Itoa(useYidun)}, "needHighPriorityMsgResend": {strconv.FormatBool(needHighPriorityMsgResend)}})
	if err != nil {
		return nil, err
	}
	sendResult := &ChatroomSendMsgResult{}
	err = json.Unmarshal(res, sendResult)
	if err != nil {
		return nil, err
	}
	return sendResult, nil

}

// AddRobot ...
func (chatroom *Chatroom) AddRobot(roomid int64, accids string, roleExt string, notifyExt string) (*RobotResult, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_ADD_ROBOT, url.Values{"roomid": {strconv.FormatInt(roomid, 10)}, "accids": {accids}, "roleExt": {roleExt}, "notifyExt": {notifyExt}})
	if err != nil {
		return nil, err
	}
	robotReuslt := &RobotResult{}
	err = json.Unmarshal(res, robotReuslt)
	if err != nil {
		return nil, err
	}
	return robotReuslt, nil

}

// RemoveRobot ...
func (chatroom *Chatroom) RemoveRobot(roomid int64, accids string) (*RobotResult, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_REMOVE_ROBOT, url.Values{"roomid": {strconv.FormatInt(roomid, 10)}, "accids": {accids}})
	if err != nil {
		return nil, err
	}
	robotResult := &RobotResult{}
	err = json.Unmarshal(res, robotResult)
	if err != nil {
		return nil, err
	}
	return robotResult, nil

}

// QueueOffer ...
func (chatroom *Chatroom) QueueOffer(roomid int64, key, value, operator, transient string) (*BaseResp, error) {
	res, err := ResponseResult(chatroom.APPKEY, chatroom.APPSECRET, ACTION_CHATROOM_QUEUE_OFFER, url.Values{"roomid": {strconv.FormatInt(roomid, 10)}, "key": {key}, "value": {value}, "operator": {operator}, "transient": {transient}})
	if err != nil {
		return nil, err
	}
	baseResp := &BaseResp{}
	err = json.Unmarshal(res, baseResp)
	if err != nil {
		return nil, err
	}
	return baseResp, nil

}
