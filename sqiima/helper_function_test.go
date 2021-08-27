package sqiima_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
//  Variables and Constants for Test
// ----------------------------------------------------------------------------

const nameDirTestData = "testdata"

// ----------------------------------------------------------------------------
//  Helper Functions
// ----------------------------------------------------------------------------

// failNowOnError はエラーの場合に msg の内容をログに記録し、テストを即時終了します.
func failNowOnError(t *testing.T, err error, msg string) {
	t.Helper()

	if err != nil {
		t.Fatalf("%v\nError: %v", msg, err)
	}
}

// getPathDirTestData はリポジトリのテストデータまでのパスを返します.
func getPathDirTestData(t *testing.T) string {
	t.Helper()

	pathWd, err := os.Getwd()
	failNowOnError(t, err, "failed to get current work dir")

	// カレントにテストディレクトリがあるかチェック
	pathDirTestData := filepath.Join(pathWd, nameDirTestData)
	if isDir(t, pathDirTestData) {
		return pathDirTestData
	}

	// 1 階層上をチェック
	pathDirParent := filepath.Dir(pathWd)
	pathDirTestData = filepath.Join(pathDirParent, nameDirTestData)
	if isDir(t, pathDirTestData) {
		return pathDirTestData
	}

	// 2 階層上をチェック
	pathDirGrandParent := filepath.Dir(pathDirParent)
	pathDirTestData = filepath.Join(pathDirGrandParent, nameDirTestData)
	if isDir(t, pathDirTestData) {
		return pathDirTestData
	}

	t.Fatalf("failed to get data dir path. Path is not a dir: %v", pathDirTestData)

	return ""
}

// getTestData は nameFile の内容をテストディレクトリから読み取って返します.
func getTestData(t *testing.T, nameFile string) []byte {
	t.Helper()

	path := getPathDirTestData(t)
	pathFileData := filepath.Join(path, nameFile)

	byteData, err := os.ReadFile(pathFileData)
	failNowOnError(t, err, "failed to read test data from "+pathFileData)

	return byteData
}

// isDir はディレクトリが存在する場合に true を返します.
func isDir(t *testing.T, pathFile string) bool {
	t.Helper()

	if pathAbs, err := filepath.Abs(pathFile); err != nil {
		return false
	} else if fileInfo, err := os.Stat(pathAbs); os.IsNotExist(err) || !fileInfo.IsDir() {
		return false
	}

	return true
}

func Test_isDir(t *testing.T) {
	t.Helper()

	assert.False(t, isDir(t, "../reft"), "non-existing dir should be false")
	assert.False(t, isDir(t, "../reft/README.md"), "non-existing file should be false")

	assert.True(t, isDir(t, "../refs"), "existing dir should be true")
	assert.True(t, isDir(t, "../testdata"), "existing dir should be true")
	assert.False(t, isDir(t, "../refs/README.md"), "existing but file should be false")
}
