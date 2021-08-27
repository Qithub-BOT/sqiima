package sqiima

// Like は記事につけられた「LGTM！」のオブジェクトを表します.
// Note: Qiita Team の LGTM API は 2020 年 11 月 4 日より廃止となりました。今後は絵文字リアクション API をご利用ください.
type Like struct {
	// CreatedAt : データが作成された日時
	CreatedAt ISO8601 `json:"created_at"`
	// User : Qiita 上のユーザを表します。
	User LikeUser `json:"user"`
}
