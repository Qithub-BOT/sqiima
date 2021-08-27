#shellcheck shell=sh
# =============================================================================
#  functions.sh のユニット・テスト
# =============================================================================

# -----------------------------------------------------------------------------
#  Sample Test
# -----------------------------------------------------------------------------

# parottry（サンプル関数のテスト例）
Describe 'parottry command'
    Include "${PATH_DIR_REQUEST_SRC}/functions.sh"

    It 'should print ok'
        # call は同じシェル内で関数を呼び出し（引数 'OK' 付）実行した結果を output にキャ
        # プチャします
        When call parottry 'ok'
        The output should equal 'ok'
    End

    It 'should print an error if arg is missing'
        # run はサブ・シェルで parottry を実行するので exit も検知できる
        When run parottry
        The stderr should equal 'argument missing'
        The status should be failure
    End
End

# -----------------------------------------------------------------------------
#  Function Tests（abc 順）
# -----------------------------------------------------------------------------
# ターゲットとなるファイルのインポート
Include "${PATH_DIR_REQUEST_SRC}/functions.sh"
# -----------------------------------------------------------------------------

# getCurlNoAuth 関数
Describe 'getCurlNoAuth functionality'
    It 'should return false if arg is missing'
        When run getCurlNoAuth
        The status should be failure
        The stderr should equal 'argument missing'
    End

    It 'should return true if args are correct'
        When run getCurlNoAuth '/api/v2/items'

        The status should be success
        The output should include "rendered_body"
        The output should include "coediting"
        The output should include "created_at"
    End
End

# checkRequirements 関数
Describe 'checkRequirements functionality'
    It 'should return true to proceed other tests'
        When run checkRequirements
        The status should be success
    End
End

# isCmdInstalled 関数
Describe 'isCmdInstalled functionality'
    It 'should return true since echo exits'
        When run isCmdInstalled "echo"
        The status should be success
    End

    It 'should return false since foobar does not exist'
        When run isCmdInstalled "foobar"
        The stderr should equal 'foobar is not installed'
        The status should be failure
    End

    It 'should print an error if arg is missing'
        When run isCmdInstalled
        The stderr should equal 'argument missing'
        The status should be failure
    End
End

# -----------------------------------------------------------------------------
#  Miscellaneous Tests（その他のテスト）
# -----------------------------------------------------------------------------

# API のホストが有効かチェック
Describe 'if URL_API_BASE is a valid host'
    It 'should return 200 HTTP status code in the header'
        When run curl -s -I "$URL_API_BASE"
        The output should include "HTTP/2 200"
        The output should include "domain=.qiita.com;"
    End
End
