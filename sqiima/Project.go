package sqiima

// Project は Qiita Team 上でのプロジェクトを表します.
// Qiita Team でのみ有効です.
type Project struct {
	// CreatedAt : データが作成された日時
	CreatedAt ISO8601 `json:"created_at"`
	// UpdatedAt : データが最後に更新された日時
	UpdatedAt ISO8601 `json:"updated_at"`
	// Body : Markdown形式の本文
	Body string `json:"body"`
	// Name : プロジェクト名
	Name string `json:"name"`
	// RenderedBody : HTML形式の本文
	RenderedBody string `json:"rendered_body"`
	// ID : プロジェクトのチーム上での一意なID
	ID int64 `json:"id"`
	// ReactionsCount : 絵文字リアクション数
	ReactionsCount int64 `json:"reactions_count"`
	// Archived : このプロジェクトが進行中かどうか
	Archived bool `json:"archived"`
}
