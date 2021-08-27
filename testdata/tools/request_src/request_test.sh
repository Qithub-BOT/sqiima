#shellcheck shell=sh

# Simple example of shellspec usage
Describe 'echo command'
    It 'should print ok'
        When call echo 'ok'
        The output should eq 'ok'
    End
End

# テスト中に使われるグローバル変数のテスト
Describe 'Global Variable'
    It 'check PATH_DIR_REQUEST_SRC is defined and is a valid path'
        The value "$PATH_DIR_REQUEST_SRC" should be defined
        The path "$PATH_DIR_REQUEST_SRC" should be exist
    End

    It 'check PATH_DIR_TOOLS is defined and is a valid path'
        The value "$PATH_DIR_TOOLS" should be defined
        The path "$PATH_DIR_TOOLS" should be exist
    End

    It 'check PATH_DIR_TESTDATA is defined and is a valid path'
        The value "$PATH_DIR_TESTDATA" should be defined
        The path "$PATH_DIR_TESTDATA" should be exist
    End

    It 'check PATH_DIR_WORK is defined and is avalid path'
        The value "$PATH_DIR_WORK" should be defined
        The path "$PATH_DIR_WORK" should be exist
    End

    It 'check PATH_DIR_REPO is defined and is a valid path'
        The value "$PATH_DIR_REPO" should be defined
        The path "${PATH_DIR_REPO}/.git" should be exist
    End
End
