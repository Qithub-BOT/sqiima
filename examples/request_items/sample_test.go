package main

import (
	"errors"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	t.Log("Test of main at: examples/request_items/sample_test.go:" + t.Name())

	out := capturer.CaptureStderr(func() {
		main()
	})

	assert.Empty(t, out, "STDERR should be empty when success. Got: %v", out)
}

func TestExitOnError(t *testing.T) {
	// OsExit のモック前にバックアップとリストア
	oldOsExit := OsExit
	defer func() {
		OsExit = oldOsExit
	}()

	capturedCode := 0 // Exit ステータスをキャプチャする変数（この値が 1 にならないといけない）

	// OsExit をモック
	//os.Exit せずに、受け取ったステータスを変数に代入してキャプチャするだけ
	OsExit = func(code int) {
		capturedCode = code
	}

	// ダミー・エラーの作成
	msgError := "this is a dummy error"
	dummyError := errors.New(msgError)

	// 実行（標準エラー出力をキャプチャ）
	out := capturer.CaptureStderr(func() {
		ExitOnError(dummyError)
	})

	// 診断（標準エラー出力されたかチェック）
	expectMsg := msgError
	actualMsg := out
	assert.Contains(t, actualMsg, expectMsg, "it should print to STDERR on error")

	// 診断（終了ステータスが 1 だったかチェック）
	expectCode := 1
	actualCode := capturedCode
	assert.Equal(t, expectCode, actualCode, "on error it should exit with code 1")
}
