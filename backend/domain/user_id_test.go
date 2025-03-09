package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserId(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		value := "d04a8be984d6d2318e7268bb1b79cfa07890"
		user, _ := NewUserId(value)
		assert.Equal(t, user, UserId{value: value})
	})

	t.Run("ユーザーIDが存在しない", func(t *testing.T) {
		value := ""
		expected := "ユーザーIDの生成に失敗しました"
		_, err := NewUserId(value)
		assert.EqualError(t, err, expected)
	})

	t.Run("最大文字列超過(37文字)", func(t *testing.T) {
		value := "d04a8be984d6d2318e7268bb1b79cfa0789071"
		expected := "ユーザーIDの生成に失敗しました"
		_, err := NewUserId(value)
		assert.EqualError(t, err, expected)
	})
}

func TestUserIdValue(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		value := "value test"
		userId, _ := NewUserId(value)
		assert.Equal(t, userId.Value(), value)
	})
}
