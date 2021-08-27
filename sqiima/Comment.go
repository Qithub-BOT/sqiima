package sqiima

// Comment は 記事やプロジェクトに付けられたコメントのオブジェクトを表します.
// プロジェクトへのコメントは Qiita Team でのみ有効です.
type Comment struct {
	// CreatedAt : データが作成された日時
	CreatedAt ISO8601 `json:"created_at"`
	// UpdatedAt : データが最後に更新された日時
	UpdatedAt ISO8601 `json:"updated_at"`

	// Body : コメントの内容を表す Markdown 形式の文字列
	Body string `json:"body"`
	// ID : コメントの一意な ID
	ID string `json:"id"`
	// RenderedBody : コメントの内容を表す HTML 形式の文字列
	RenderedBody string `json:"rendered_body"`

	// User : Qiita 上のユーザを表します。
	User CommentUser `json:"user"`
}
