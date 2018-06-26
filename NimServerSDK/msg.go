package nimserversdk

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
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
	res, err := ResponseResult(msg.APPKEY, msg.APPSECRET, ACTION_MSG_SNED_MSG, url.Values{"from": {from}, "ope": {strconv.Itoa(ope)}, "to": {to}, "type": {strconv.Itoa(msgType)}, "body": {body}, "antispam": {strconv.FormatBool(antispam)}, "antispamCustom": {antispamCustom}, "option": {option}, "pushcontent": {pushcontent}, "payload": {payload}, "ext": {ext}, "forcepushlist": {forcepushlist}, "forcepushcontent": {forcepushcontent}, "bid": {bid}, "useYidun": {strconv.Itoa(useYidun)}, "markRead": {strconv.Itoa(markRead)}})
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
func (msg *Msg) SendBatchMsg(fromAccid string, toAccids string, msgType int, body string, option string, pushcontent string, payload string, ext string, bid string, useYidun int) (*SendResult, error) {
	res, err := ResponseResult(msg.APPKEY, msg.APPSECRET, ACTION_MSG_SNED_BATCH_MSG, url.Values{"fromAccid": {fromAccid}, "toAccids": {toAccids}, "type": {strconv.Itoa(msgType)}, "body": {body}, "option": {option}, "pushcontent": {pushcontent}, "payload": {payload}, "ext": {ext}, "bid": {bid}, "useYidun": {strconv.Itoa(useYidun)}})
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
	res, err := ResponseResult(msg.APPKEY, msg.APPSECRET, ACTION_MSG_SEND_ATTACH_MSG, url.Values{"from": {from}, "msgtype": {strconv.Itoa(msgType)}, "to": {to}, "attach": {attach}, "pushcontent": {pushcontent}, "payload": {payload}, "sound": {sound}, "save": {strconv.Itoa(save)}, "option": {option}})
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
func (msg *Msg) SendBatchAttachMsg(fromAccid string, toAccids string, attach string, pushcontent string, payload string, sound string, save int, option string) (*SendResult, error) {
	res, err := ResponseResult(msg.APPKEY, msg.APPKEY, ACTION_MSG_SEND_BATCH_ATTACH_MSG, url.Values{"fromAccid": {fromAccid}, "toAccids": {toAccids}, "attach": {attach}, "pushcontent": {pushcontent}, "payload": {payload}, "sound": {sound}, "save": {strconv.Itoa(save)}, "option": {option}})
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
func (msg *Msg) Upload(content []byte, fileType string, isHttps bool) (*UploadResult, error) {
	res, err := ResponseResult(msg.APPKEY, msg.APPSECRET, ACTION_MSG_UPLOAD, url.Values{"content": {base64.StdEncoding.EncodeToString(content)}, "type": {fileType}, "ishttps": {strconv.FormatBool(isHttps)}})
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
func (msg *Msg) UploadByMultiPart(content []byte, fileType string, isHttps bool) (*UploadResult, error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	err := w.WriteField("type", fileType)
	if err != nil {
		return nil, err
	}
	err = w.WriteField("ishttps", strconv.FormatBool(isHttps))
	if err != nil {
		return nil, err
	}
	fw, err := w.CreateFormField("content")
	if err != nil {
		return nil, err
	}
	_, err = fw.Write(content)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", ACTION_MSG_UPLOAD_MULTIPART, buf)
	if err != nil {
		return nil, err
	}
	fillHeader(req, msg.APPKEY, msg.APPSECRET)
	req.Header.Set("Content-Type", w.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	uploadResult := &UploadResult{}
	err = json.Unmarshal(resBody, uploadResult)
	if err != nil {
		return nil, err
	}
	return uploadResult, nil

}

// MsgRecall ...
func (msg *Msg) Recall(deleteMsgid string, timetag string, msgType int, from, to, msgDesc, ignoreTime string) (*BaseResp, error) {
	res, err := ResponseResult(msg.APPKEY, msg.APPSECRET, ACTION_MSG_RECALL, url.Values{"deleteMsgid": {deleteMsgid}, "timetag": {timetag}, "type": {strconv.Itoa(msgType)}, "from": {from}, "to": {to}, "msg": {msgDesc}, "ignoreTime": {ignoreTime}})
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
	res, err := ResponseResult(msg.APPKEY, msg.APPSECRET, ACTION_MSG_BROADCAST, url.Values{"body": {body}, "from": {from}, "isOffline": {strconv.FormatBool(isOffline)}, "ttl": {strconv.Itoa(ttl)}, "targetOs": {targetOs}})
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
