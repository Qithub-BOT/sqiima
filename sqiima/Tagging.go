package sqiima

// Tagging は記事とタグとの関連を表します.
type Tagging struct {
	// Name : タグを特定するための一意な名前
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}
