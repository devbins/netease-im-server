package nimserversdk

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
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

	ACTION_MSG_SNED_MSG              = BASE_URL + "msg/sendMsg.action"            // 发送普通消息
	ACTION_MSG_SNED_BATCH_MSG        = BASE_URL + "msg/sendBatchMsg.action"       // 批量发送点对点普通消息
	ACTION_MSG_SEND_ATTACH_MSG       = BASE_URL + "msg/sendAttachMsg.action"      // 发送自定义系统通知
	ACTION_MSG_SEND_BATCH_ATTACH_MSG = BASE_URL + "msg/sendBatchAttachMsg.action" // 批量发送点对点自定义系统通知
	ACTION_MSG_UPLOAD                = BASE_URL + "msg/upload.action"             // 文件上传
	ACTION_MSG_UPLOAD_MULTIPART      = BASE_URL + "msg/fileUpload.action"         // 文件上传 multipart 方式
	ACTION_MSG_RECALL                = BASE_URL + "msg/recall.action"             // 消息撤回
	ACTION_MSG_BROADCAST             = BASE_URL + "broadcastMsg.action"

	ACTION_TEAM_CREATE             = BASE_URL + "team/create.action"          // 创建群
	ACTION_TEAM_ADD                = BASE_URL + "team/add.action"             // 拉人入群
	ACTION_TEAM_KICK               = BASE_URL + "team/kick.action"            // 踢人出群
	ACTION_TEAM_REMOVE             = BASE_URL + "team/remove.action"          // 解散群
	ACTION_TEAM_UPDATE             = BASE_URL + "team/update.action"          // 编辑群资料
	ACTION_TEAM_QUERY              = BASE_URL + "team/query.action"           // 群信息与成员列表查询
	ACTION_TEAM_QUERY_DETAIL       = BASE_URL + "team/queryDetail.action"     // 获取群组详细信息
	ACTION_TEAM_GET_MARK_READ_INFO = BASE_URL + "team/getMarkReadInfo.action" // 获取群组已读消息的已读详情信息
	ACTION_TEAM_CHANGE_OWNER       = BASE_URL + "team/changeOwner.action"     // 移交群主
	ACTION_TEAM_ADD_MANAGER        = BASE_URL + "team/addManager.action"      // 任命管理员
	ACTION_TEAM_REMOVE_MANAGER     = BASE_URL + "team/removeManager.action"   // 移除管理员
	ACTION_TEAM_JOIN_TEAM          = BASE_URL + "team/joinTeams.action"       // 获取某用户所加入的群信息
	ACTION_TEAM_UPDATE_TEAM_NICK   = BASE_URL + "team/updateTeamNick.action"  // 修改群昵称
	ACTION_TEAM_MUTE               = BASE_URL + "team/muteTeam.action"        // 修改消息提醒开关
	ACTION_TEAM_MUTE_LIST          = BASE_URL + "team/muteTlist.action"       // 禁言群成员
	ACTION_TEAM_LEAVE              = BASE_URL + "team/leave.action"           // 主动退群
	ACTION_TEAM_MUTE_ALL           = BASE_URL + "team/muteTlistAll.action"    // 将群组整体禁言
	ACTION_TEAM_LIST_TEAM_MUTE     = BASE_URL + "team/listTeamMute.action"    // 获取群组禁言列表

	ACTION_CHATROOM_CREATE              = BASE_URL + "chatroom/create.action"           // 创建聊天室
	ACTION_CHATROOM_GET                 = BASE_URL + "chatroom/get.action"              // 查询聊天室信息
	ACTION_CAHTROOM_GET_BATCH           = BASE_URL + "chatroom/getBatch.action"         // 批量查询聊天室信息
	ACTION_CHATROOM_UPDATE              = BASE_URL + "chatroom/update.action"           // 更新聊天室信息
	ACTION_CHATROOM_TOGGLE_CLOSE_STAT   = BASE_URL + "chatroom/toggleCloseStat.action"  // 修改聊天室开/关闭状态
	ACTION_CHATROOM_SET_MEMBER_ROLE     = BASE_URL + "chatroom/setMemberRole.action"    // 设置聊天室内用户角色
	ACTION_CHATROOM_REQUEST_ADDR        = BASE_URL + "chatroom/requestAddr.action"      // 请求聊天室地址
	ACTION_CHATROOM_SEND_MSG            = BASE_URL + "chatroom/sendMsg.action"          // 发送聊天室消息
	ACTION_CHATROOM_ADD_ROBOT           = BASE_URL + "chatroom/addRobot.action"         // 往聊天室内添加机器人
	ACTION_CHATROOM_REMOVE_ROBOT        = BASE_URL + "chatroom/removeRobot.action"      // 从聊天室内删除机器人
	ACTION_CHATROOM_TEMPORARY_MUTE      = BASE_URL + "chatroom/temporaryMute.action"    // 设置临时禁言状态
	ACTION_CHATROOM_QUEUE_OFFER         = BASE_URL + "chatroom/queueOffer.action"       // 往聊天室有序队列中新加或更新元素
	ACTION_CHATROOM_QUEUE_POLL          = BASE_URL + "chatroom/queuePoll.action"        // 从队列中取出元素
	ACTION_CHATROOM_QUEUE_LIST          = BASE_URL + "chatroom/queueList.action"        // 排序；列出队列中所有元素
	ACTION_CHATROOM_QUEUE_DROP          = BASE_URL + "chatroom/queueDrop.action"        // 删除清理整个队列
	ACTION_CAHTROOM_QUEUE_INIT          = BASE_URL + "chatroom/queueInit.action"        // 初始化队列
	ACTION_CHATROOM_MUTE_ROOM           = BASE_URL + "chatroom/muteRoom.action"         // 将聊天室整体禁言
	ACTION_CHATROOM_TOPN                = BASE_URL + "chatroom/topn.action"             // 查询聊天室统计指标TopN
	ACTION_CHATROOM_MEMBERS_BY_PAGE     = BASE_URL + "chatroom/membersByPage.action"    // 分页获取成员列表
	ACTION_CHATROOM_QUERY_MEMBERS       = BASE_URL + "chatroom/queryMembers.action"     // 批量获取在线成员信息
	ACTION_CHATROOM_UPDATE_MY_ROOM_ROLE = BASE_URL + "chatroom/updateMyRoomRole.action" // 变更聊天室内的角色信息

)

type User struct {
	APPKEY    string
	APPSECRET string
}

var client = http.Client{}

type BaseResp struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

type Info struct {
	Accid  string `json:"accid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender int    `json:"gender"`
	Mobile string `json:"mobile"`
}

type TokenInfo struct {
	Token string `json:"token"`
	Accid string `json:"accid"`
}

type TokenRespose struct {
	BaseResp
	TokenInfo TokenInfo `json:"info"`
}

type Uinfos struct {
	BaseResp
	Uinfos []Info `json:"uinfos"`
}

// Create ...
func (this *User) Create(accid string) (*TokenRespose, error) {
	req, err := http.NewRequest("POST", ACTION_USER_CREATE, strings.NewReader("accid="+accid))
	if err != nil {
		return nil, err
	}
	fillHeader(req, this.APPKEY, this.APPSECRET)
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
func (this *User) Update(accid, props, token string) (*BaseResp, error) {
	req, err := http.NewRequest("POST", ACTION_USER_UPDATE, strings.NewReader(url.Values{"accid": {accid}, "props": {props}, "token": {token}}.Encode()))
	if err != nil {
		return nil, err
	}

	fillHeader(req, this.APPKEY, this.APPSECRET)

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
func (this *User) RefreshToken(accid string) (*TokenRespose, error) {
	resBody, err := ResponseResult(this.APPKEY, this.APPSECRET, ACTION_USER_REFRESH_TOKEN, url.Values{"accid": {accid}})
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
func (this *User) Block(accid string) (*BaseResp, error) {
	req, err := http.NewRequest("POST", ACTION_USER_BLOCK, strings.NewReader("accid="+accid))
	if err != nil {
		return nil, err
	}
	fillHeader(req, this.APPKEY, this.APPSECRET)
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
func (this *User) UnBlock(accid string) (*BaseResp, error) {
	req, err := http.NewRequest("POST", ACTION_USER_UNBLOCK, strings.NewReader("accid="+accid))
	if err != nil {
		return nil, err
	}
	fillHeader(req, this.APPKEY, this.APPSECRET)
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
func (this *User) UpdateUinfo(gender int, accid, name, icon, sign, email, birth, mobile, ex string) (*BaseResp, error) {
	params := url.Values{"accid": {accid}, "name": {name}, "icon": {icon}, "sign": {sign}, "email": {email}, "birth": {birth}, "mobile": {mobile}, "ex": {ex}, "gender": {strconv.Itoa(gender)}}
	data, err := ResponseResult(this.APPKEY, this.APPSECRET, ACTION_USER_UPDATE_UINFO, params)
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
func (this *User) GetUinfo(accids []string) (*Uinfos, error) {
	ids, err := json.Marshal(accids)
	if err != nil {
		return nil, err
	}
	data, err := ResponseResult(this.APPKEY, this.APPSECRET, ACTION_USER_GET_UINFO, url.Values{"accids": {string(ids)}})
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
func (this *User) SetDonnop(accid string, donnopOpen bool) (*BaseResp, error) {
	params := url.Values{"accid": {accid}, "donnopOpen": {strconv.FormatBool(donnopOpen)}}
	data, err := ResponseResult(this.APPKEY, this.APPSECRET, ACTION_USER_DONNOP_OPEN, params)
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
func ResponseResult(appkey string, appSecret string, action string, params url.Values) ([]byte, error) {
	req, err := http.NewRequest("POST", action, strings.NewReader(params.Encode()))
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
