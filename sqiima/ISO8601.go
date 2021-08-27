package sqiima

import (
	"strings"
	"time"

	"golang.org/x/xerrors"
)

// ISO8601 は ISO-8601（"2000-01-01T00:00:00+00:00"）フォーマットの日付オブジェクトを示します.
// Qiita API の日付は ATOM（ISO-8601）形式であるため、通常の Unmarshal/Marshal では変換に失敗するため.
type ISO8601 struct {
	time.Time
}

// Format は obj.Time を ISO8601 フォーマットの文字列で返します.
func (obj *ISO8601) Format() string {
	return obj.Time.Format("2006-01-02T15:04:05+09:00")
}

// Parse は inputTime を ISO8601 フォーマットとして解釈し time.Time 型で返します.
func (obj *ISO8601) Parse(inputTime string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02T15:04:05+09:00", inputTime, time.Local)
}

// MarshalJSON は json.Marshal() で struct から JSON に変換する際に呼び出されるメソッドで、
// time を ISO8601 形式に変換します.
func (obj *ISO8601) MarshalJSON() ([]byte, error) {
	return []byte(`"` + obj.Format() + `"`), nil
}

// UnmarshalJSON は json.Unmarshal() で JSON から struct に変換する際に呼び出されるメソッドで、
// ISO8601 形式の文字列を time.Time に変換します.
func (obj *ISO8601) UnmarshalJSON(bTime []byte) error {
	inputTime := string(bTime)
	if inputTime == "null" {
		return nil
	}

	// inputTime はダブルクォートされて届くので両端をトリム
	inputTime = strings.Trim(inputTime, "\"")

	pTime, err := obj.Parse(inputTime) // ISO8601 フォーマットから time.Time に変換
	if err != nil {
		return xerrors.Errorf("error parse time from ISO8601 format string: %v", err)
	}

	// フィールドにセット
	obj.Time = pTime

	return nil
}
