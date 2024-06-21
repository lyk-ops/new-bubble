package models

import "new-bubble/dao"

type Todo struct {
	ID     int64  `json:"id"`
	Title  string `json:"title" gorm:"size:255;unique_index"`
	Status bool   `json:"status"`
}

/*
	Todo这个Model的增删改查操作都放在这里
*/
// CreateATodo 创建todo
func CreateATodo(todo *Todo) error {
	err := dao.DB.Create(&todo).Error
	return err
}

// GetAllTodo 获取所有todo
func GetAllTodo() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

// GetATodo 获取一个todo
func GetATodo(id string) (todo *Todo, err error) {
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}

// UpdateTodo
func UpdateTodo(todo *Todo) error {
	err := dao.DB.Save(&todo).Error
	return err
}

// DeleteTodo
func DeleteTodo(id string) error {
	err := dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return err
}
