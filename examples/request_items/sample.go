/*
Package main は、sqiima パッケージの使い方のサンプルです.

新着記事 20 件（デフォルトの取得件数）を取得して、各々のタイトルを出力します.
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Qithub-BOT/sqiima/sqiima"
)

// OsExit は os.Exit のエイリアスで、テスト時にモック用に使われます.
var OsExit = os.Exit

func main() {
	// HTTP クライアントの作成
	client := new(http.Client)

	// HTTP リクエストの作成
	request, err := http.NewRequest("GET", "https://qiita.com/api/v2/items", nil)
	ExitOnError(err)
	request.Close = true // リクエスト完了直後に Close する

	// リクエストの実行（アクセストークンなしのリクエストは 60回/時 の制限があります）
	response, err := client.Do(request)
	ExitOnError(err)

	defer response.Body.Close()

	// レスポンスの読み取り
	rawJSON, err := io.ReadAll(response.Body)
	ExitOnError(err)

	// 空の記事オブジェクトを作成しポインタを取得
	items := new(sqiima.Items)

	// 記事オブジェクトに生 JSON を割り当て（JSON decode）
	err = json.Unmarshal(rawJSON, items)
	ExitOnError(err)

	// 記事の ID とタイトル出力
	for _, item := range *items {
		fmt.Println(item.ID, item.Title)
	}
}

// ExitOnError は err がエラーの場合、標準エラー出力にエラー内容を出力し、終了ステータス 1 で終了します.
func ExitOnError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)

		OsExit(1)
	}
}
