package sqiima_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Qithub-BOT/sqiima/sqiima"
	"github.com/stretchr/testify/assert"
)

// ============================================================================
//  Tests （abc 順）
// ============================================================================

// ----------------------------------------------------------------------------
//  Item
// ----------------------------------------------------------------------------

func TestItem(t *testing.T) {
	// 空の記事オブジェクト作成
	item := new(sqiima.Item)
	// テストデータ（生 JSON）の読み込み
	jsonResponse := getTestData(t, "api.v2.items.item_id.response.json")
	// 生 JSON を記事オブジェクトに割り当て（JSON decode）
	err := json.Unmarshal(jsonResponse, item)
	failNowOnError(t, err, "failed to unmarshal test JSON data")
	// タイトルチェック
	{
		assert.Equal(t, item.Title, "【Golang】【wasm】テストの仕方と注意点 @ Go 1.16+")
	}
	// 作成日のチェック
	{
		// "2021-08-04T15:33:58+09:00" ISO8601 フォーマット
		expectTime := time.Date(2021, 8, 4, 15, 33, 58, 0, time.FixedZone("JST", 9*60*60))
		expect := expectTime.Local().String()
		actual := item.CreatedAt.Local().String()

		assert.NoErrorf(t, err, "failed to parse date")
		assert.Equal(t, expect, actual)
	}
	// タグ
	{
		assert.Equal(t, "Go", item.Tags[0].Name)
		assert.Equal(t, "test", item.Tags[1].Name)
		assert.Equal(t, "WebAssembly", item.Tags[2].Name)
		assert.Equal(t, "wasm", item.Tags[3].Name)
	}
	// 本文 Markdown
	{
		assert.Equal(t, "## WebAssembly 用のアプリでもテストがしたい", item.Body)
	}
}

// ----------------------------------------------------------------------------
//  Items
// ----------------------------------------------------------------------------

func TestItems(t *testing.T) {
	items := new(sqiima.Items)                                   // 空の記事オブジェクト作成
	jsonResponse := getTestData(t, "api.v2.items.response.json") // テストデータの読み込み

	// 記事オブジェクトに生 JSON を割り当て（JSON decode）
	err := json.Unmarshal(jsonResponse, items)
	failNowOnError(t, err, "failed to unmarshal test JSON data")

	assertType := sqiima.Item{}

	for _, item := range *items {
		assert.IsType(t, assertType, *item)
	}
}
