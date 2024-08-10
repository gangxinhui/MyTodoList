package dao

import (
	"MyTodoList/model"
	"MyTodoList/types"
	"context"
	"gorm.io/gorm"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewDBClient(ctx)}
}

// CreateTask 创建Task
func (s *TaskDao) CreateTask(task *model.Task) error {
	return s.Model(&model.Task{}).Create(&task).Error
}
func (s *TaskDao) ListTask(start, limit int, userId uint) (r []*model.Task, total int64, err error) {
	err = s.Model(&model.Task{}).Preload("User").Where("uid = ?", userId).
		Count(&total).
		Limit(limit).Offset((start - 1) * limit).
		Find(&r).Error

	return
}

// UpdateTask 修改
func (s *TaskDao) UpdateTask(uId uint, req *types.UpdateTaskReq) error {
	t := new(model.Task)
	err := s.Model(&model.Task{}).Where("id = ? AND uid=?", req.ID, uId).First(&t).Error
	if err != nil {
		return err
	}

	if req.Status != 0 {
		t.Status = req.Status
	}

	if req.Title != "" {
		t.Title = req.Title
	}

	if req.Content != "" {
		t.Content = req.Content
	}

	return s.Save(t).Error
}

// SearchTask 搜索
func (s *TaskDao) SearchTask(uId uint, info string) (tasks []*model.Task, err error) {
	err = s.Where("uid=?", uId).Preload("User").First(&tasks).Error
	if err != nil {
		return
	}

	err = s.Model(&model.Task{}).Where("title LIKE ? OR content LIKE ?",
		"%"+info+"%", "%"+info+"%").Find(&tasks).Error

	return
}

// Delete 删除
func (s *TaskDao) DeleteTask(uId, tId uint) error {
	r, err := s.FindTaskByIdAndUserID(uId, tId)
	if err != nil {
		return err
	}
	return s.Delete(&r).Error

}
