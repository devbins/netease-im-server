package nimserversdk

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type Team struct {
	APPKEY    string
	APPSECRET string
}

type CreateResult struct {
	BaseResp
	Tid    string `json:"tid"`
	Faccid Exceed `json:"faccid"`
}

type Exceed struct {
	Accid []string `json:"accid"`
	Msg   string   `json:"msg"`
}

type AddResult struct {
	BaseResp
	Faccid Exceed `json:"faccid"`
}

type QueryResult struct {
	BaseResp
	Tinfos []Tinfo `json:"tinfos"`
}

type ReadInfoResult struct {
	BaseResp
	Data ReadInfo `json:"data"`
}

type ReadInfo struct {
	ReadSize     int      `json:"readSize"`
	UnreadSize   int      `json:"unreadSize"`
	Readaccids   []string `json:"readAccids"`
	UnreadAccids []string `json:"unreadAccids"`
}

type Tinfo struct {
	Tname        string   `json:"tname"`
	Announcement string   `json:"announcement"`
	Owner        string   `json:"owner"`
	Maxusers     int      `json:"maxusers"`
	Joinmode     int      `json:"joinmode"`
	Tid          int      `json:"tid"`
	Intro        string   `json:"intro"`
	Size         int      `json:"size"`
	Custom       string   `json:"custom"`
	ClientCustom string   `json:"clientCustom"`
	Mute         bool     `json:"mute"`
	Createtime   int64    `json:"createtime"`
	Updatetime   int64    `json:"updatetime"`
	Admins       []string `json:"admins"`
	Members      []string `json:"members"`
}

type TeamMember struct {
	Createtime int64  `json:"createtime"`
	Updatetime int64  `json:"updatetime"`
	Nick       string `json:"nick"`
	Accid      string `json:"accid"`
	Mute       bool   `json:"mute"`
	Custom     string `json:"custom"`
}

type TDetailInfo struct {
	Icon         string       `json:"icon"`
	Announcement string       `json:"announcement"`
	Uptinfomode  int          `json:"uptinfomode"`
	Maxusers     int          `json:"maxusers"`
	Intro        string       `json:"intro"`
	Upcustommode int          `json:"upcustommode"`
	Tname        string       `json:"tname"`
	Beinvitemode int          `json:"beinvitemode"`
	Joinmode     int          `json:"joinmode"`
	Tid          int          `json:"tid"`
	Invitemode   int          `json:"invitemode"`
	Mute         bool         `json:"mute"`
	Custom       string       `json:"custom"`
	ClientCustom string       `json:"clientCustom"`
	Createtime   int64        `json:"createtime"`
	Updatetime   int64        `json:"updatetime"`
	Owner        TeamMember   `json:"owner"`
	Admins       []TeamMember `json:"admins"`
	Members      []TeamMember `json:"members"`
}

type QueryDetailResult struct {
	BaseResp
	TDetailInfo `json:"tinfo"`
}

type JoinTeamResult struct {
	BaseResp
	Count int        `json:"count"`
	Infos []TeamInfo `json:"infos"`
}

type TeamInfo struct {
	Owner    string `json:"owner"`
	Tname    string `json:"tname"`
	MaxUsers int    `json:"maxusers"`
	Tid      int    `json:"tid"`
	Size     int    `json:"size"`
	Custom   string `json:"custom"`
}

// Create ...
func (team *Team) Create(tname string, owner string, members string, announcement string, intro string, msg string, magree int, joinmode int, custom string, icon string, beinvitemode int, invitemode int, uptinfomode int, upcustommode int) (*CreateResult, error) {

	res, err := ResponseResult(team.APPKEY, team.APPSECRET, ACTION_TEAM_CREATE, url.Values{"tname": {tname}, "owner": {owner}, "members": {members}, "announcement": {announcement}, "intro": {intro}, "msg": {msg}, "magree": {strconv.Itoa(magree)}, "joinmode": {strconv.Itoa(joinmode)}, "custom": {custom}, "icon": {icon}, "beinvitemode": {strconv.Itoa(beinvitemode)}, "invitemode": {strconv.Itoa(invitemode)}, "uptinfomode": {strconv.Itoa(uptinfomode)}, "upcustommode": {strconv.Itoa(upcustommode)}})
	if err != nil {
		return nil, err
	}
	result := &CreateResult{}
	err = json.Unmarshal(res, result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// Add ...
func (team *Team) Add(tid string, owner string, members string, magree int, msg string, attach string) (*AddResult, error) {
	res, err := ResponseResult(team.APPKEY, team.APPSECRET, ACTION_TEAM_ADD, url.Values{"tid": {tid}, "owner": {owner}, "members": {members}, "magree": {strconv.Itoa(magree)}, "msg": {msg}, "attach": {attach}})
	if err != nil {
		return nil, err
	}
	addRusult := &AddResult{}
	err = json.Unmarshal(res, addRusult)
	if err != nil {
		return nil, err
	}
	return addRusult, err

}

// Kick ...
func (team *Team) Kick(tid, owner, member, members, attach string) (*BaseResp, error) {
	res, err := ResponseResult(team.APPKEY, team.APPSECRET, ACTION_TEAM_KICK, url.Values{"tid": {tid}, "owner": {owner}, "member": {member}, "members": {members}, "attach": {attach}})
	if err != nil {
		return nil, err
	}
	result := &BaseResp{}
	err = json.Unmarshal(res, result)

	if err != nil {
		return nil, err
	}
	return result, nil

}

// Remove ...
func (team *Team) Remove(tid, owner string) (*BaseResp, error) {
	res, err := ResponseResult(team.APPKEY, team.APPSECRET, ACTION_TEAM_REMOVE, url.Values{"tid": {tid}, "owner": {owner}})
	if err != nil {
		return nil, err
	}
	result := &BaseResp{}
	err = json.Unmarshal(res, result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// Update ...
func (team *Team) Update(tid string, tname string, owner string, announcement string, intro string, joinmode int, custom string, icon string, beinvitemode, invitemode, uptinfomode, upcustommode int) (*BaseResp, error) {
	res, err := ResponseResult(team.APPKEY, team.APPSECRET, ACTION_TEAM_UPDATE, url.Values{"tid": {tid}, "tname": {tname}, "owner": {owner}, "announcement": {announcement}, "intro": {intro}, "joinmode": {strconv.Itoa(joinmode)}, "custom": {custom}, "icon": {icon}, "beinvitemode": {strconv.Itoa(beinvitemode)}, "invitemode": {strconv.Itoa(invitemode)}, "uptinfomode": {strconv.Itoa(uptinfomode)}, "upcustommode": {strconv.Itoa(upcustommode)}})
	if err != nil {
		return nil, err
	}
	result := &BaseResp{}
	err = json.Unmarshal(res, result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// Query ...
func (team *Team) Query(tids string, ope int) (*QueryResult, error) {
	res, err := ResponseResult(team.APPKEY, team.APPSECRET, ACTION_TEAM_QUERY, url.Values{"tids": {tids}, "ope": {strconv.Itoa(ope)}})
	if err != nil {
		return nil, err
	}
	queryResult := &QueryResult{}
	err = json.Unmarshal(res, queryResult)
	if err != nil {
		return nil, err
	}
	return queryResult, nil

}

// QueryDetail ...
func (team *Team) QueryDetail(tid int64) (*TDetailInfo, error) {
	res, err := ResponseResult(team.APPKEY, team.APPSECRET, ACTION_TEAM_QUERY_DETAIL, url.Values{"tid": {strconv.FormatInt(tid, 10)}})
	if err != nil {
		return nil, err
	}
	tinfo := &TDetailInfo{}
	err = json.Unmarshal(res, tinfo)
	if err != nil {
		return nil, err
	}
	return tinfo, err

}

// GetMarkReadInfo ...
func (team *Team) GetMarkReadInfo(tid int64, msgId int64, fromAccid string, snapshot bool) (*ReadInfoResult, error) {
	res, err := ResponseResult(team.APPKEY, team.APPSECRET, ACTION_TEAM_GET_MARK_READ_INFO, url.Values{"tid": {strconv.FormatInt(tid, 10)}, "msgid": {strconv.FormatInt(msgId, 10)}, "fromAccid": {fromAccid}, "snapshot": {strconv.FormatBool(snapshot)}})
	if err != nil {
		return nil, err
	}
	readInfoResult := &ReadInfoResult{}
	err = json.Unmarshal(res, readInfoResult)
	if err != nil {
		return nil, err
	}
	return readInfoResult, nil

}

// ChangetOwner ...
func (team *Team) ChangetOwner(tid string, owner string, newowner string, leave int) (*BaseResp, error) {
	res, err := ResponseResult(team.APPKEY, team.APPSECRET, ACTION_TEAM_CHANGE_OWNER, url.Values{"tid": {tid}, "owner": {owner}, "newowner": {newowner}, "leave": {strconv.Itoa(leave)}})
	if err != nil {
		return nil, err
	}
	result := &BaseResp{}
	err = json.Unmarshal(res, result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// AddManager ...
func (team *Team) AddManager(tid, owner, members string) (*BaseResp, error) {
	res, err := ResponseResult(team.APPKEY, team.APPSECRET, ACTION_TEAM_ADD_MANAGER, url.Values{"tid": {tid}, "owner": {owner}, "members": {members}})
	if err != nil {
		return nil, err
	}
	result := &BaseResp{}
	err = json.Unmarshal(res, result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// RemoveManager ...
func (team *Team) RemoveManager(tid, owner, members string) (*BaseResp, error) {
	res, err := ResponseResult(team.APPKEY, team.APPSECRET, ACTION_TEAM_REMOVE_MANAGER, url.Values{"tid": {tid}, "owner": {owner}, "members": {members}})
	if err != nil {
		return nil, err
	}
	result := &BaseResp{}
	err = json.Unmarshal(res, result)
	if err != nil {
		return nil, err
	}
	return result, nil

}

// JoinTeams ...
func (team *Team) JoinTeams(accid string) (*TeamInfo, error) {
	res, err := ResponseResult(team.APPKEY, team.APPSECRET, ACTION_TEAM_JOIN_TEAM, url.Values{"accid": {accid}})
	if err != nil {
		return nil, err
	}
	teamInfo := &TeamInfo{}
	err = json.Unmarshal(res, teamInfo)
	if err != nil {
		return nil, err
	}
	return teamInfo, nil

}
