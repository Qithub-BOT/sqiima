package sqiima

// Template は記事のひな形に利用できるテンプレートのオブジェクトを表します.
// Note: Qiita Team でのみ有効です.
type Template struct {
	// Body : テンプレートの本文
	Body string `json:"body"`
	// ExpandedBody : 変数を展開した状態の本文
	ExpandedBody string `json:"expanded_body"`
	// ExpandedTitle : 変数を展開した状態のタイトル
	ExpandedTitle string `json:"expanded_title"`
	// Name : テンプレートを判別するための名前
	Name string `json:"name"`
	// Title : 生成される記事のタイトルの雛形
	Title string `json:"title"`
	// Tags : タグ一覧
	Tags []TemplateTagsItem `json:"tags"`
	// ExpandedTags : 変数を展開した状態のタグ一覧
	ExpandedTags []TemplateExpandedTagsItem `json:"expanded_tags"`
	// ID : テンプレートの一意な ID
	ID int64 `json:"id"`
}
