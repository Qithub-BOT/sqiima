package sqiima

// ExpandedTemplate はテンプレートを保存する前に変数展開後の状態をプレビューするためのオブジェクトです.
// Qiita Team でのみ有効です.
type ExpandedTemplate struct {
	// ExpandedBody : 変数を展開した状態の本文
	ExpandedBody string `json:"expanded_body"`
	// ExpandedTitle : 変数を展開した状態のタイトル
	ExpandedTitle string `json:"expanded_title"`
	// ExpandedTags : 変数を展開した状態のタグ一覧
	ExpandedTags []ExpandedTemplateExpandedTagsItem `json:"expanded_tags"`
}
