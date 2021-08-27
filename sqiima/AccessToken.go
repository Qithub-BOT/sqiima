package sqiima

// AccessToken は Qiita API v2 で認証・認可を行うためのアクセストークンのオブジェクトを表します.
type AccessToken struct {
	// ClientID : 登録された API クライアントを特定するための ID
	ClientID string `json:"client_id"`
	// Token : アクセストークンを表現する文字列
	Token string `json:"token"`

	// Scopes : アクセストークンに許された操作の一覧
	Scopes []string `json:"scopes"`
}
