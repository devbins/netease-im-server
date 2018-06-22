package nimserversdk

type Nim struct {
	APPKEY    string
	APPSECRET string
	NONCE     string
	User      User
	Friend    Friend
}

// New ...
func NewNim(appkey, appsecret, nonce string) *Nim {
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
	}

}
