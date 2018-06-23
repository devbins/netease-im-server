package nimserversdk

type Nim struct {
	APPKEY    string
	APPSECRET string
	User      User
	Friend    Friend
	Msg       Msg
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
	}

}
