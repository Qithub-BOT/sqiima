package sqiima

// TeamInvitation は Qiita Team での招待中のメンバーのオブジェクトを表します.
// Note: Qiita Team でのみ有効です。管理者権限が必要です.
type TeamInvitation struct {
	// Email : 招待中のメンバーの email アドレスです。
	Email string `json:"email"`
	// URL : 招待用URLです。有効期限は 1 日です。
	URL string `json:"url"`
}
