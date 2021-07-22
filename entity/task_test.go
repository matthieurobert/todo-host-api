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

	_, err := PostTask(task)

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestGetTaskById(t *testing.T) {
	task := Task{
		Title:  "test",
		Body:   "This is testing",
		UserId: 1,
	}

	id, _ := PostTask(task)

	taskFound, err := GetTaskById(id)

	if err != nil {
		t.Errorf(err.Error())
	}

	if id != taskFound.Id {
		t.Errorf("Couldn't find task by id")
	}
}
