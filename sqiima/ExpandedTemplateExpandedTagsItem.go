package sqiima

// ExpandedTemplateExpandedTagsItem は拡張テンプレートに付いたタグのオブジェクトです.
// Qiita Team でのみ有効です.
type ExpandedTemplateExpandedTagsItem struct {
	// Name : タグの名前
	Name string `json:"name"`
	// Versions: タグが付けられたバージョンの一覧
	Versions []string `json:"versions,omitempty"`
}
