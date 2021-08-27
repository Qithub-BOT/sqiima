package sqiima

// TeamAccessToken は Qiita API v2 で認証・認可を行うためのチーム別アクセストークンのオブジェクトを表します.
// Qiita Team でのみ有効です.
type TeamAccessToken struct {
	// ClientID : 登録されたAPIクライアントを特定するための ID
	ClientID string `json:"client_id"`
	// Token : アクセストークンを表現する文字列
	Token string `json:"token"`
	// Scopes : アクセストークンに許された操作の一覧
	Scopes []string `json:"scopes"`
}
