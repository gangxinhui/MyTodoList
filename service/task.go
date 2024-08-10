package service

import (
	"MyTodoList/dao"
	"MyTodoList/model"
	"MyTodoList/pkg/ctl"
	"MyTodoList/pkg/util"
	"MyTodoList/types"
	"context"
	"sync"
	"time"
)

// 单例模式
var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

type TaskSrv struct {
}

func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

func (s *TaskSrv) CreateTask(ctx context.Context, req *types.CreateTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	user, err := dao.NewUserDao(ctx).FindUserByUserId(u.Id)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	task := &model.Task{
		User:      *user,
		Uid:       user.ID,
		Title:     req.Title,
		Content:   req.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
	}
	//创建新的task数据库
	err = dao.NewTaskDao(ctx).CreateTask(task)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}
func (s *TaskSrv) ListTask(ctx context.Context, req *types.TaskList) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	tasks, total, err := dao.NewTaskDao(ctx).ListTask(req.Start, req.Limit, u.Id)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	taskRespList := make([]*types.TaskResp, 0)
	for _, v := range tasks {
		taskRespList = append(taskRespList, &types.TaskResp{
			ID:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			View:      v.View(),
			Status:    v.Status,
			CreatedAt: v.CreatedAt.Unix(),
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	return ctl.RespList(taskRespList, total), nil
}
func (s *TaskSrv) ShowTask(ctx context.Context, req *types.ShowTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	task, err := dao.NewTaskDao(ctx).FindTaskByIdAndUserID(u.Id, req.Id)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	respTask := &types.TaskResp{
		ID:        task.ID,
		Title:     task.Title,
		Content:   task.Content,
		View:      task.View(),
		Status:    task.Status,
		CreatedAt: task.CreatedAt.Unix(),
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
	task.AddView() // 增加点击数
	return ctl.RespSuccessWithData(respTask), nil

}
func (s *TaskSrv) UpdateTask(ctx context.Context, req *types.UpdateTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	err = dao.NewTaskDao(ctx).UpdateTask(u.Id, req)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}
func (s *TaskSrv) SearchTask(ctx context.Context, req *types.SearchTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	tasks, err := dao.NewTaskDao(ctx).SearchTask(u.Id, req.Info)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	taskRespList := make([]*types.TaskResp, 0)
	for _, v := range tasks {
		taskRespList = append(taskRespList, &types.TaskResp{
			ID:        v.ID,
			Title:     v.Title,
			Content:   v.Content,
			Status:    v.Status,
			View:      v.View(),
			CreatedAt: v.CreatedAt.Unix(),
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	return ctl.RespList(taskRespList, int64(len(taskRespList))), nil
}
func (s *TaskSrv) DeleteTask(ctx context.Context, req *types.DeleteTaskReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	err = dao.NewTaskDao(ctx).DeleteTask(u.Id, req.Id)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}
