package sqiima

// Reaction は Qiita Team 上での絵文字リアクションを表します.
// Qiita Team でのみ有効です.
type Reaction struct {
	// CreatedAt : データが作成された日時
	CreatedAt ISO8601 `json:"created_at"`
	// ImageURL : 絵文字画像の URL
	ImageURL string `json:"image_url"`
	// Name : 絵文字の識別子
	Name string `json:"name"`
	// User : Qiita 上のユーザを表します。
	User ReactionUser `json:"user"`
}
