package sqiima_test

import (
	"encoding/json"
	"testing"

	"github.com/Qithub-BOT/sqiima/sqiima"
	"github.com/stretchr/testify/assert"
)

func TestISO8601(t *testing.T) {
	inputJSON := []byte(`{"created_at":"2021-08-21T08:38:37+09:00"}`)

	type TempJSON struct {
		CreatedAt sqiima.ISO8601 `json:"created_at"`
	}

	item := new(TempJSON)

	// Unmarshal テスト
	{
		err := json.Unmarshal(inputJSON, item)
		assert.NoError(t, err)

		expect := "2021-08-21T08:38:37+09:00"
		actual := item.CreatedAt.Format()
		assert.Equal(t, expect, actual)
	}

	// Marshal テスト
	{
		outputJSON, err := json.Marshal(item)
		assert.NoError(t, err)

		expect := inputJSON
		actual := outputJSON
		assert.Equal(t, expect, actual)
		assert.Contains(t, string(outputJSON), "2021-08-21T08:38:37+09:00")
	}
}

func TestISO8601_null_value(t *testing.T) {
	inputJSON := []byte(`{"created_at":null}`)

	type TempJSON struct {
		CreatedAt sqiima.ISO8601 `json:"created_at"`
	}

	item := new(TempJSON)
	err := json.Unmarshal(inputJSON, item)

	assert.NoError(t, err)
	assert.Empty(t, item.CreatedAt, "'null' value should be empty")
}

func TestISO8601_parse_error(t *testing.T) {
	item := new(sqiima.ISO8601)
	err := item.UnmarshalJSON([]byte("hello"))

	assert.Error(t, err, "malformed tye should return error")
	assert.Contains(t, err.Error(), "error parse time from ISO8601 format string")
}
