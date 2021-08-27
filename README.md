[![Go Report Card](https://goreportcard.com/badge/github.com/Qithub-BOT/sqiima)](https://goreportcard.com/report/github.com/Qithub-BOT/sqiima)
[![Go Reference](https://pkg.go.dev/badge/github.com/Qithub-BOT/sqiima.svg)](https://pkg.go.dev/github.com/Qithub-BOT/sqiima)

# SQiiMA - JSON schema struct of Qiita API v2 for Go

`SQiiMA` は [Qiita API v2](https://qiita.com/api/v2/docs) を Go 言語（以下 Golang）で扱う際の JSON <ruby>スキーマ<rt>構造の定義</rt></ruby>を元に Golang の構造体（`struct type`）に起こしただけのシンプルなライブラリです。

Qiita API へのリクエストおよびレスポンスの JSON データを <ruby>`Marshal` / `Unmarshal`<rt>JSON ⇆ Go struct に変換</rt></ruby> する際に `import` して利用ください。

- [Qiita API v2 の JSON スキーマ定義](https://qiita.com/api/v2/schema) @ Qiita 公式
- [sqiima の構造体一覧](./sqiima/)

## `go get`

```shellsession
$ go get github.com/Qithub-BOT/sqiima
```

## Usage

```go
import "github.com/Qithub-BOT/sqiima/sqiima"
```

### シンプルな使い方

```go
// Qiita API にリクエストした結果を取得する
// (/api/v2/items エンドポイントは新着記事の JSON オブジェクトを配列で返します)
rawJSON := doSomethingToRequest("/api/v2/items")

// Go の初期化された空のオブジェクト作成
// Items は Item のスライス（可変長配列）です =`[]Item`
items := new(sqiima.Items)

// 空のオブジェクトに生 JSON を割り当て (JSON decode-ish)
if err := json.Unmarshal(rawJSON, item); err != nil {
	log.Fatalf("error to parse JSON: %v", err)
}

// 新着 Qiita 記事一覧のタイトル出力
for _, item := range *items {
	fmt.Println(item.Title)
}
```

### もう少し具体的な使い方

- 新着 Qiita 記事 20 件を取得して一覧を表示するサンプル
	- [./examples/request_items/](./examples/request_items/sample.go)
- 指定した記事 ID を取得して記事情報を表示するサンプル
	- [./examples/request_single_item/](./examples/request_single_item/sample.go)

## サポートしている JSON スキーマ対応表

JSON 側オブジェクト名 | Go 側構造体名（型名） | 説明
:--: | :--: | :--
"item" | [Item](./sqiima/Item.go) | ユーザからの投稿を表します。
"tag" | [Tag](./sqiima/Tag.go) | 記事に付けられた個々のタグを表します。
"like" | [Like](./sqiima/Like.go) | 記事につけられた「LGTM！」を表します。<strong>Qiita Team の LGTM API は2020年11月4日より廃止となりました。今後は絵文字リアクション API をご利用ください。</strong>
"user" | [User](./sqiima/User.go) | Qiita 上のユーザを表します。
"authenticated_user" | [AuthenticatedUser](./sqiima/AuthenticatedUser.go) | 現在のアクセストークンで認証中のユーザを表します。通常のユーザ情報よりも詳細な情報を含んでいます。
"access_token" | [AccessToken](./sqiima/AccessToken.go) | Qiita API v2 で認証・認可を行うためのアクセストークンを表します。
"comment" | [Comment](./sqiima/Comment.go) | 記事やプロジェクトに付けられたコメントを表します。プロジェクトへのコメントは Qiita Team でのみ有効です。

### Qiita Team でのみ有効な JSON

JSON 側オブジェクト名 | Go 側構造体名（型名） | 説明
:--: | :--: | :--
"team" | [Team](./sqiima/Team.go) | Qiita Team 上で所属しているチームを表します。
"group" | [Group](./sqiima/Group.go) | Qiita Team のグループを表します。
"reaction" | [Reaction](./sqiima/Reaction.go) | Qiita Team 上での絵文字リアクションを表します。
"template" | [Template](./sqiima/Template.go) | 記事のひな形に利用できるテンプレートを表します。
"project" | [Project](./sqiima/Project.go) | Qiita Team 上でのプロジェクトを表します。
"expanded_template" | [ExpandedTemplate](./sqiima/ExpandedTemplate.go) | テンプレートを保存する前に変数展開後の状態をプレビューできます。
"team_access_token" | [TeamAccessToken](./sqiima/TeamAccessToken.go) | Qiita API v2 で認証・認可を行うためのチーム別アクセストークンを表します。
"tagging" | [Tagging](./sqiima/Tagging.go) | 記事とタグとの関連を表します。
"remove_team_member" | [RemoveTeamMember](./sqiima/RemoveTeamMember.go) | 指定のユーザーをチームから離脱させます。自身とチームのオーナーはこの API では離脱させられません。
"team_invitation" | [TeamInvitation](./sqiima/TeamInvitation.go) | Qiita Teamでの招待中のメンバーを表します。管理者権限が必要です。

- [詳細なドキュメント](./sqiima/README.md)

## ライセンス

- MIT
- Copyright [The SQiiMA Contributors](https://github.com/Qithub-BOT/sqiima/graphs/contributors)
