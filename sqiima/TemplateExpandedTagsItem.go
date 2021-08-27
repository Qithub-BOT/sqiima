package sqiima

// TemplateExpandedTagsItem はテンプレートの拡張タグのオブジェクトを表します.
// Qiita Team でのみ有効です.
type TemplateExpandedTagsItem struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions,omitempty"`
}
