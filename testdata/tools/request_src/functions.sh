#!/bin/sh
# shellcheck shell=sh disable=SC2034
# =============================================================================
#  関数一覧
# =============================================================================
# - 関数名は lowerCamelCase()
# - POSIX 互換であること
# - 期待する出力やエラーのテスト（functions_test.sh）を作成してから実装すること

# ----------------------------------------------------------------------------
#  定数
# ----------------------------------------------------------------------------
SUCCESS=0
FAILURE=1
TRUE=0
FALSE=1
URL_API_BASE='https://qiita.com' # 最後にスラッシュは付けない

# ----------------------------------------------------------------------------
#  関数
# ----------------------------------------------------------------------------

# checkRequirements は API のリクエストに最低限必要なコマンドがインストールされていた場合
# 0 (true) を返します。
checkRequirements() {
    isCmdInstalled 'curl' || {
        echo >&2 'cURL must be installed to request API'

        return $FALSE
    }

    # cURL が URL エンコード（--data-urlencode オプション）に対応したバージョンか確認
    curl --help all 2>&1 | grep data-urlencode 2>/dev/null 1>/dev/null || {
        echo >&2 '--data-urlencode option missing in curl command. Use latest'

        return $FALSE
    }

    isCmdInstalled 'jq' || {
        echo >&2 'jq must be installed to parse received JSON data'

        return $FALSE
    }
}

# curlNoAuth は第 1 引数のエントリーポイントに GET で HTTP リクエストした際のレスポンスの
# JSON を標準出力に出力します。JSON は整形済みです。
# この関数はアクセストークンを必要としない Qiita APIにのみ有効です。
getCurlNoAuth() {
    if [ $# -lt 1 ]; then
        echo >&2 'argument missing'

        return $FALSE
    fi

    url_request="${URL_API_BASE:?'URL_API_BASE undefined'}${1}"

    response=$(curl -sS -f -H "Content-Type: application/json" "$url_request" 2>&1) || {
        echo >&2 "$response"

        return $FALSE
    }

    echo "$response" | jq -e .

    return "$?"
}

# isCmdInstalled は第 1 引数のコマンドが存在する場合に 0 を返します。
isCmdInstalled() {
    if [ $# -lt 1 ]; then
        echo >&2 'argument missing'

        return $FALSE
    fi

    cmd="$1"

    if ! type "$cmd" 2>/dev/null 1>/dev/null; then
        echo >&2 "${cmd} is not installed"

        return $FALSE
    fi
}

# parotty は第 1 引数を echo するだけのサンプル関数です。
parottry() {
    if [ $# -lt 1 ]; then
        echo >&2 'argument missing'
        exit $FAILURE
    fi

    echo "$@"
}
