package sqiima

// Tag は記事に付けられた個々のタグのオブジェクトを表します.
type Tag struct {
	// IconURL : このタグに設定されたアイコン画像の URL
	IconURL *string `json:"icon_url"`
	// ID : タグを特定するための一意な名前
	ID string `json:"id"`
	// FollowersCount : このタグをフォローしているユーザの数
	FollowersCount int64 `json:"followers_count"`
	// ItemsCount : このタグが付けられた記事の数
	ItemsCount int64 `json:"items_count"`
}
