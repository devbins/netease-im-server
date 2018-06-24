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
