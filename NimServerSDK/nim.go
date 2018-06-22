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
		NONCE:     nonce,
		User: User{
			APPKEY:    appkey,
			APPSECRET: appsecret,
			NONCE:     nonce,
		},
		Friend: Friend{
			APPKEY:    appkey,
			APPSECRET: appsecret,
			NONCE:     nonce,
		},
	}

}
