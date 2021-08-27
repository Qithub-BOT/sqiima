package sqiima_test

import (
	"os"
	"path/filepath"
	"testing"
)

// ----------------------------------------------------------------------------
//  Variables and Constants for Test
// ----------------------------------------------------------------------------

const nameDirTestData = "testdata"

var pathDirTestData = ""

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

	if pathDirTestData != "" {
		return pathDirTestData
	}

	pathWd, err := os.Getwd()
	failNowOnError(t, err, "failed to get current work dir")

	pathDirTestData = filepath.Join(filepath.Dir(pathWd), nameDirTestData)

	if !isDir(t, pathDirTestData) {
		t.Fatalf("failed to get data dir path. Path is not a dir: %v", pathDirTestData)
	}

	return pathDirTestData
}

// getTestData は nameFile の内容をテストディレクトリから読み取って返します.
func getTestData(t *testing.T, nameFile string) []byte {
	t.Helper()

	pathFileData := filepath.Join(getPathDirTestData(t), nameFile)

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
