package entity

import "github.com/todo-host/todo-host-api/config"

type Task struct {
	Id     int64
	UserId int64
	Title  string
	Body   string
}

func GetTasks(UserId int64) ([]Task, error) {
	var tasks []Task
	var userTasks []Task

	err := config.POSTGRES.DB.Model(&tasks).Select()

	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		err := config.POSTGRES.DB.Model(task).Where("userId = ?", task.UserId).Select()

		if err != nil {
			return nil, err
		}

		userTasks = append(userTasks, task)
	}

	return userTasks, nil
}

func GetTaskById(userId int64, id int64) (*Task, error) {
	task := &Task{Id: id}
	err := config.POSTGRES.DB.Model(task).Where("if = ?", task.Id).Select()

	if err != nil {
		return nil, err
	}

	if task.UserId != userId {
		return nil, nil
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
