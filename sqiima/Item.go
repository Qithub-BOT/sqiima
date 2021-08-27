package sqiima

// Item はユーザの投稿のオブジェクトを表します.
type Item struct {
	// PageViewsCount : 閲覧数
	PageViewsCount *int64 `json:"page_views_count"`

	// CreatedAt : データが作成された日時
	CreatedAt ISO8601 `json:"created_at"`
	// UpdatedAt : データが最後に更新された日時
	UpdatedAt ISO8601 `json:"updated_at"`

	// Body : Markdown 形式の本文
	Body string `json:"body"`
	// ID : 記事の一意な ID
	ID string `json:"id"`
	// RenderedBody : HTML 形式の本文
	RenderedBody string `json:"rendered_body"`
	// Title : 記事のタイトル
	Title string `json:"title"`
	// URL : 記事の URL
	URL string `json:"url"`

	// Tags : 記事に付いたタグ一覧
	Tags []ItemTagsItem `json:"tags"`

	// TeamMembership : Qiita Team のチームメンバー情報を表します。
	TeamMembership ItemTeamMembership `json:"team_membership"`
	// Group : Qiita Team のグループを表します。
	Group Group `json:"group"`
	// User : Qiita 上のユーザを表します。
	User ItemUser `json:"user"`

	// CommentsCount : この記事へのコメントの数
	CommentsCount int64 `json:"comments_count"`
	// LikesCount : この記事への「LGTM！」の数（Qiita でのみ有効）
	LikesCount int64 `json:"likes_count"`
	// ReactionsCount : 絵文字リアクションの数（Qiita Team でのみ有効）
	ReactionsCount int64 `json:"reactions_count"`

	// Coediting : この記事が共同更新状態かどうか (Qiita Team でのみ有効)
	Coediting bool `json:"coediting"`
	// Private : 限定共有状態かどうかを表すフラグ (Qiita Team では無効)
	Private bool `json:"private"`
}
