package entity

import (
	"os"
	"testing"

	"github.com/todo-host/todo-host-api/config"
)

func TestMain(m *testing.M) {
	config.InitConfig()
	CreateSchema(config.POSTGRES.DB)

	exitCode := m.Run()

	config.POSTGRES.DB.Close()

	os.Exit(exitCode)
}

func TestPostTask(t *testing.T) {
	task := Task{
		Title:  "test",
		Body:   "This is testing",
		UserId: 1,
	}

	_, err := PostTask(&task)

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetTasks(t *testing.T) {
	task1 := Task{
		Title:  "test1",
		Body:   "This is testing",
		UserId: 1,
	}

	PostTask(&task1)

	task2 := Task{
		Title:  "test2",
		Body:   "This is testing",
		UserId: 1,
	}

	PostTask(&task2)

	tasksfound, err := GetTasks(1)

	if err != nil {
		t.Errorf(err.Error())
	}

	if len(tasksfound) == 0 {
		t.Errorf("Couldn't find tasks")
	}
}

func TestGetTaskById(t *testing.T) {
	task := Task{
		Title:  "test",
		Body:   "This is testing",
		UserId: 1,
	}

	id, _ := PostTask(&task)

	taskFound, err := GetTaskById(id)

	if err != nil {
		t.Errorf(err.Error())
	}

	if id != taskFound.Id {
		t.Errorf("Couldn't find task by id")
	}
}
