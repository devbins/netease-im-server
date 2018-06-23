package nimserversdk

import (
	"encoding/json"
	"net/url"
)

type Msg struct {
	APPKEY    string
	APPSECRET string
}

type Data struct {
	MsgId    string `json:"msgid"`
	Antispam bool   `json:"antispam"`
}

type SendResult struct {
	BaseResp
	Data Data `json:"data"`
}

type SendBatchResutl struct {
	BaseResp
	UnRegister []string `json:"unregister"`
}

type UploadResult struct {
	BaseResp
	Url string `json:"url"`
}

type Broadcast struct {
	BaseResp
	BroadcastMsg `json:"msg"`
}
type BroadcastMsg struct {
	ExpireTime  int64    `json:"expireTime"`
	Body        string   `json:"body"`
	CreateTime  int64    `json:"createTime"`
	IsOffline   bool     `json:"isOffline"`
	BroadcastId int64    `json:"broadcastId"`
	TargetOs    []string `json:"targetOs"`
}

// SendMsg ...
func (msg *Msg) SendMsg(from string, ope int, to string, msgType int, body string, antispam bool, antispamCustom string, option string, pushcontent string, payload string, ext string, forcepushlist string, forcepushcontent string, forcepushall string, bid string, useYidun int, markRead int) (*SendResult, error) {
	res, err := ResponseResult(msg.APPKEY, msg.APPSECRET, ACTION_MSG_SNED_MSG, url.Values{})
	if err != nil {
		return nil, err
	}
	sendResult := &SendResult{}
	err = json.Unmarshal(res, sendResult)
	if err != nil {
		return nil, err
	}
	return sendResult, nil

}

// SendBatchMsg ...
func (msg *Msg) SendBatchMsg(fromAccid string, toAccid string, msgType int, body string, option string, pushcontent string, payload string, ext string, bid string, useYidun int) (*SendResult, error) {
	res, err := ResponseResult(msg.APPKEY, msg.APPSECRET, ACTION_MSG_SNED_BATCH_MSG, url.Values{})
	if err != nil {
		return nil, err
	}
	batchResult := &SendResult{}
	err = json.Unmarshal(res, batchResult)
	if err != nil {
		return nil, err
	}
	return batchResult, nil

}

// SendAttachMsg ...
func (msg *Msg) SendAttachMsg(from string, msgType int, to string, attach string, pushcontent string, payload string, sound string, save int, option string) (*BaseResp, error) {
	res, err := ResponseResult(msg.APPKEY, msg.APPSECRET, ACTION_MSG_SEND_ATTACH_MSG, url.Values{})
	if err != nil {
		return nil, err
	}
	resMsg := &BaseResp{}
	err = json.Unmarshal(res, resMsg)
	if err != nil {
		return nil, err
	}
	return resMsg, nil
}

// SendBatchAttachMsg ...
func (msg *Msg) SendBatchAttachMsg(fromAccid string, toAccids []string, attach string, pushcontent string, payload string, sound string, save int, option string) (*SendResult, error) {
	res, err := ResponseResult(msg.APPKEY, msg.APPKEY, ACTION_MSG_SEND_BATCH_ATTACH_MSG, url.Values{})
	if err != nil {
		return nil, err
	}
	sendResult := &SendResult{}
	err = json.Unmarshal(res, sendResult)
	if err != nil {
		return nil, err
	}
	return sendResult, nil

}

// MsgUpload ...
func (msg *Msg) MsgUpload(content string, fileType string, isHttps bool) (*UploadResult, error) {
	res, err := ResponseResult(msg.APPKEY, msg.APPSECRET, ACTION_MSG_UPLOAD, url.Values{})
	if err != nil {
		return nil, err
	}
	uploadResult := &UploadResult{}
	err = json.Unmarshal(res, uploadResult)
	if err != nil {
		return nil, err
	}
	return uploadResult, nil

}

// MsgUpload ...
func (msg *Msg) MsgUploadByMultiPart(content string, fileType string, isHttps bool) (*UploadResult, error) {
	res, err := ResponseResult(msg.APPKEY, msg.APPSECRET, ACTION_MSG_UPLOAD, url.Values{})
	if err != nil {
		return nil, err
	}
	uploadResult := &UploadResult{}
	err = json.Unmarshal(res, uploadResult)
	if err != nil {
		return nil, err
	}
	return uploadResult, nil

}

// MsgRecall ...
func (msg *Msg) MsgRecall(deleteMsgid string, timetag string, msgType int, from, to, msgDesc, ignoreTime string) (*BaseResp, error) {
	res, err := ResponseResult(msg.APPKEY, msg.APPSECRET, ACTION_MSG_RECALL, url.Values{})
	if err != nil {
		return nil, err
	}
	resResult := &BaseResp{}
	err = json.Unmarshal(res, resResult)
	if err != nil {
		return nil, err
	}
	return resResult, nil

}

// Broadcast ...
func (msg *Msg) Broadcast(body string, from string, isOffline bool, ttl int, targetOs string) (*Broadcast, error) {
	res, err := ResponseResult(msg.APPKEY, msg.APPSECRET, ACTION_MSG_BROADCAST, url.Values{})
	if err != nil {
		return nil, err
	}
	broadcastResult := &Broadcast{}
	err = json.Unmarshal(res, broadcastResult)
	if err != nil {
		return nil, err
	}
	return broadcastResult, nil

}
