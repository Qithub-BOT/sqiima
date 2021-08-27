package sqiima

// Team は Qiita Team 上 で所属しているチームのオブジェクトを表します.
// Qiita Team でのみ有効です.
type Team struct {
	// ID : チームの一意な ID
	ID string `json:"id"`
	// Name : チームに設定されている名前を表します。
	Name string `json:"name"`
	// Active : チームが利用可能な状態かどうか
	Active bool `json:"active"`
}
