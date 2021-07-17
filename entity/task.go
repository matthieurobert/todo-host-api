package entity

import "github.com/todo-host/todo-host-api/config"

type Task struct {
	Id     int64
	UserId int64
	Title  string
	Body   string
}

func GetTasks(userId int64) ([]Task, error) {
	var tasks []Task

	err := config.POSTGRES.DB.Model(&tasks).Where("user_id = ?", userId).Select()

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTaskById(id int64) (*Task, error) {
	task := &Task{Id: id}
	err := config.POSTGRES.DB.Model(task).Where("id = ?", task.Id).Select()

	if err != nil {
		return nil, err
	}

	return task, err
}

func PostTask(task Task) (int64, error) {
	_, err := config.POSTGRES.DB.Model(&task).Insert()

	if err != nil {
		return 0, err
	}

	return task.Id, nil
}

func DeleteTaskById(id int64) (int64, error) {
	task := &Task{Id: id}

	_, err := config.POSTGRES.DB.Model(task).Where("id = ?", task.Id).Delete()

	if err != nil {
		return 0, err
	}

	return task.Id, nil
}

func UpdateTaskById(task Task) (int64, error) {
	_, err := config.POSTGRES.DB.Model(task).Update()

	if err != nil {
		return 0, err
	}

	return task.Id, nil
}
