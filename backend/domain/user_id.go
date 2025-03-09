package domain

import "errors"

const USER_ID_MAX_LENGTH = 36
const USER_ID_ERROR_TEXT = "ユーザーIDの生成に失敗しました"

type UserId struct {
	value string
}

func NewUserId(value string) (UserId, error) {
	if value == "" {
		return UserId{}, errors.New(USER_ID_ERROR_TEXT)
	}
	if len(value) > USER_ID_MAX_LENGTH {
		return UserId{}, errors.New(USER_ID_ERROR_TEXT)
	}
	return UserId{value: value}, nil
}

func (ui UserId) Value() string {
	return ui.value
}
