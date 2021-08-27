#!/bin/sh
# =============================================================================
#  Qiita API v2 のレスポンス取得スクリプト (request.sh)
# =============================================================================
#  この shell スクリプトは Qiita API v2 のレスポンス（リクエスト結果）を cURL で取得して
#  ローカル（../../testdata/）に保存するだけのスクリプトです。
#
#  【注意事項】
#  「環境変数を設定」「Qiita のアクセストークン」と聞いてわからない場合は、危険なので利用しな
#  いでください。
#
#  【事前設定】
#  このスクリプトを利用するには、環境変数 QIITA_ACCESS_TOKEN に Qiita の個人用アクセス・
#  トークンを設定する必要があります。
#  Qiita の設定にある「アプリケーション」から「個人用アクセストークン」を全スコープ有効にして
#  発行してください。取得後は、登録を削除することをオススメします。
#
#  【実行方法】
#  $ ./request.sh
#
#  【PullRequest】
#  POSIX 互換でお願いします。
#  また、shfmt（構文解析）, shellcheck（静的解析）, shellspec（ユニットテスト）のテストを
#  パスしないとレビューされません。
#
#  【テストの実行方法】
#
#  - shfmt（構文解析）
#    リポジトリのルートから以下を実行してください
#      $ shfmt -f . | xargs shfmt -d -p -ci
#
#  【命名規約】
#  定数: 一度代入したら変更しない、なんちゃって定数的な変数は UPPER_SNAKE_CASE
#  変数: lower_snake_case
#  関数: このファイルでは定義せず ./request_src/functions.sh で定義してください

# -----------------------------------------------------------------------------
#  Setup
# -----------------------------------------------------------------------------

# shellcheck disable=SC3028
PATH_DIR_SCRIPT="$(cd "$(dirname "${BASH_SOURCE:-$0}")" && pwd)"
PATH_DIR_OUT="$(dirname "${PATH_DIR_SCRIPT}")"

# 関数ファイルと定数の読み込み
# shellcheck source=/workspaces/sqiima/testdata/tools/request_src/functions.sh
. "${PATH_DIR_SCRIPT}/request_src/functions.sh"

# 必須コマンドの確認
checkRequirements || {
    exit "$FAILURE"
}

# -----------------------------------------------------------------------------
#  認証不要のリクエスト
# -----------------------------------------------------------------------------

echo "$URL_API_BASE"
echo "$PATH_DIR_OUT"

# /api/v2/items のレスポンス
# -----------------------------------------------------------------------------
name_file_json='api.v2.items.response.json.tmp'
path_file_json="${PATH_DIR_OUT}/${name_file_json}"
echo "PATH:${path_file_json}"

if [ ! -e "$path_file_json" ]; then
    entrypoint='/api/v2/items' # 記事の一覧を作成日時の降順で返します
    result=$(getCurlNoAuth "$entrypoint" 1) && {
        echo "$result" >"$path_file_json"
    }
fi

# /api/v2/items/:item_ID のレスポンス
# -----------------------------------------------------------------------------
name_file_json='api.v2.items.item_id.response.json.tmp'
path_file_json="${PATH_DIR_OUT}/${name_file_json}"
echo "PATH:${path_file_json}"

if [ ! -e "$path_file_json" ]; then
    entrypoint='/api/v2/items/0ab42c53dcebc5a925f0' # 指定記事を返します
    result=$(getCurlNoAuth "$entrypoint" 1) && {
        echo "$result" >"$path_file_json"
    }
fi

# -----------------------------------------------------------------------------
echo 'done'
