package dto

import "gin-vue/model"

type UserDto struct {
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:  user.Username,
		Phone: user.Phone,
	}
}
