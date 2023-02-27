package dao

import "gin-vue/model"

// Register 创建用户
func (m *manager) Register(user *model.User) {
	m.database.Create(user)
}

func (m *manager) GetUserByPhone(phone string) model.User {
	var user model.User
	m.database.Where("phone=?", phone).First(&user)
	return user
}

func (m *manager) GetUserByName(username string) model.User {
	var user model.User
	m.database.Where("username=?", username).First(&user)
	return user
}

func (m *manager) GetUserByID(id int) model.User {
	var user model.User
	m.database.Where("id=?", id).First(&user)
	return user
}
