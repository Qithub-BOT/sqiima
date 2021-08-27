package sqiima

// ItemTagsItem はタグのオブジェクトを示します.
type ItemTagsItem struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions,omitempty"`
}
