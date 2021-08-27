package sqiima

// TemplateTagsItem はテンプレートのタグのオブジェクトを表します.
// Qiita Team でのみ有効です.
type TemplateTagsItem struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions,omitempty"`
}
