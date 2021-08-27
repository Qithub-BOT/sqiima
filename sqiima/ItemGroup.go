package sqiima

// ItemGroup は Qiita Team のグループのオブジェクトを表します.
type ItemGroup struct {
	// CreatedAt : データが作成された日時
	CreatedAt ISO8601 `json:"created_at"`
	// UpdatedAt : データが最後に更新された日時
	UpdatedAt ISO8601 `json:"updated_at"`

	// Description : グループの詳細を表します。
	Description string `json:"description"`
	// URLName : グループのチーム上での一意な名前を表します。
	URLName string `json:"url_name"`
	// Name : グループに付けられた表示用の名前を表します。
	Name string `json:"name"`

	// ID : グループの一意な ID を表します。
	ID int64 `json:"id"`

	// Private : 非公開グループかどうかを表します。
	Private bool `json:"private"`
}
