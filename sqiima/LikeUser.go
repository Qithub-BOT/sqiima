package sqiima

// LikeUser は Qiita 上のユーザのオブジェクトを表します.
type LikeUser struct {
	// Description : 自己紹介文
	Description *string `json:"description"`
	// FacebookID : Facebook ID
	FacebookID *string `json:"facebook_id"`
	// GithubLoginName : GitHub ID
	GithubLoginName *string `json:"github_login_name"`
	// LinkedinID : LinkedIn ID
	LinkedinID *string `json:"linkedin_id"`
	// Location : 居住地
	Location *string `json:"location"`
	// Name : 設定している名前
	Name *string `json:"name"`
	// Organization : 所属している組織
	Organization *string `json:"organization"`
	// TwitterScreenName : Twitterのスクリーンネーム
	TwitterScreenName *string `json:"twitter_screen_name"`
	// WebsiteURL : 設定しているWebサイトのURL
	WebsiteURL *string `json:"website_url"`

	// ID : ユーザID
	ID string `json:"id"`
	// ProfileImageURL : 設定しているプロフィール画像のURL
	ProfileImageURL string `json:"profile_image_url"`

	// ItemsCount : このユーザが qiita.com 上で公開している記事の数 (Qiita Teamでの記事数は含まれません)
	ItemsCount int64 `json:"items_count"`
	// PermanentID : ユーザごとに割り当てられる整数のID
	PermanentID int64 `json:"permanent_id"`
	// FolloweesCount : このユーザがフォローしているユーザの数
	FolloweesCount int64 `json:"followees_count"`
	// FollowersCount : このユーザをフォローしているユーザの数
	FollowersCount int64 `json:"followers_count"`

	// TeamOnly : Qiita Team専用モードに設定されているかどうか
	TeamOnly bool `json:"team_only"`
}
