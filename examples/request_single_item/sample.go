/*
Package main は、sqiima パッケージの使い方のサンプルです.

指定された Qiita の記事 ID を取得して記事の諸情報を出力します.
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/Qithub-BOT/sqiima/sqiima"
)

// OsExit は os.Exit のエイリアスで、テスト時にモック用に使われます.
var OsExit = os.Exit

func main() {
	// リクエスト情報
	baseURL := "https://qiita.com/"  // API のホスト
	entrypoint := "/api/v2/items/"   // API のエントリーポイント
	itemID := "84f19475647baf0ebe1f" // Qiita の記事 ID（「再帰とは」@ https://qiita.com/c_r_5/items/84f19475647baf0ebe1f）

	// リクエストの作成
	reqURL := GenReqQuery(baseURL, entrypoint, itemID)
	method := "GET"
	request, err := http.NewRequest(method, reqURL, nil)
	ExitOnError(err)
	request.Close = true // リクエスト完了直後に Close する

	// HTTP クライアントの作成
	client := new(http.Client)

	// リクエストの実行（アクセストークンなしのリクエストは 60回/時 の制限があります）
	response, err := client.Do(request)
	ExitOnError(err)

	defer response.Body.Close()

	// レスポンスの読み取り
	rawJSON, err := io.ReadAll(response.Body)
	ExitOnError(err)

	// 空の記事オブジェクト作成
	item := new(sqiima.Item)

	// 記事オブジェクトに生 JSON を割り当て（JSON decode）
	err = json.Unmarshal(rawJSON, item)
	ExitOnError(err)

	// 割り当て内容の出力（一部）
	fmt.Println("Item ID:", item.ID)                    // 記事の一意な ID
	fmt.Println("User ID:", item.User.ID)               // 記事の一意な ID
	fmt.Println("URL:", item.URL)                       // 記事の URL
	fmt.Println("Title:", item.Title)                   // 記事のタイトル
	fmt.Println("Body:", item.Body)                     // Markdown 形式の本文
	fmt.Println("RenderedBody:", item.RenderedBody)     // HTML 形式の本文
	fmt.Println("CommentsCount:", item.CommentsCount)   // この記事へのコメントの数
	fmt.Println("CreatedAt:", item.CreatedAt)           // データが作成された日時
	fmt.Println("UpdatedAt:", item.UpdatedAt)           // データが最後に更新された日時
	fmt.Println("Tags:", item.Tags)                     // 記事に付いたタグ一覧
	fmt.Println("User:", item.User)                     // Qiita 上のユーザ情報
	fmt.Println("LikesCount:", item.LikesCount)         // この記事への「LGTM！」の数
	fmt.Println("TeamMembership:", item.TeamMembership) // Qiita Team のチームメンバー情報
	fmt.Println("Group:", item.Group)                   // Qiita Team のグループ
	fmt.Println("PageViewsCount:", item.PageViewsCount) // 閲覧数（要アクセストークン）
}

// GenReqQuery はリクエスト先の URL を生成します.
func GenReqQuery(baseURL, entrypoint, itemID string) string {
	reqURL, err := url.Parse(baseURL + entrypoint + itemID)
	ExitOnError(err)

	return reqURL.String()
}

// ExitOnError は err がエラーの場合、標準エラー出力に内容を出力し終了ステータス 1 で終了します.
func ExitOnError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)

		OsExit(1)
	}
}
