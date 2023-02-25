package dao

import "gin-vue/model"

// Register 创建用户
func (m *manager) Register(user *model.User) {
	m.database.Create(user)
}
