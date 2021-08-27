# shellcheck shell=sh disable=SC2034
# =============================================================================
#  spec_helper.sh は各テストの実行前に呼び出されるスクリプトです
# =============================================================================
#  See: https://github.com/shellspec/shellspec#spec_helper

# テスト中に利用可能なグローバル変数
PATH_DIR_REQUEST_SRC="${SHELLSPEC_SPECDIR:?'SHELLSPEC_SPECDIR not set'}"
PATH_DIR_TOOLS="$(cd "$(dirname "$PATH_DIR_REQUEST_SRC")" && pwd)"
PATH_DIR_TESTDATA="$(cd "$(dirname "$PATH_DIR_TOOLS")" && pwd)"
PATH_DIR_REPO="$(cd "$(dirname "$PATH_DIR_TESTDATA")" && pwd)"
PATH_DIR_WORK="${SHELLSPEC_TMPDIR:?'SHELLSPEC_TMPDIR not set'}"

# Avoid too many requests
sleep 1