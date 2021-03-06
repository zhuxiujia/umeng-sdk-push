package umeng_sdk_push

type PayloadBodyAndroid struct {
	Ticker      string `json:"ticker"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	PlayVibrate bool   `json:"play_vibrate"`
	PlayLights  bool   `json:"play_lights"`
	PlaySound   bool   `json:"play_sound"`
	AfterOpen   string `json:"after_open"`
}

type PayloadAndroid struct {
	DisplayType string                 `json:"display_type"`
	Body        *PayloadBodyAndroid    `json:"body"`
	Extra       map[string]interface{} `json:"extra"`
}

type NotificationAndroid struct {
	UmengNotification
}

func (this *NotificationAndroid) setConfig(secret string) {
	this.host = HOST
	this.appMasterSecret = secret
	this.postPath = POST_PATH
	this.uploadPath = UPLOAD_PATH
}
