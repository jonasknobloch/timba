package tim_test

import (
	"fmt"
	"github.com/google/uuid"
	"timba/tim"
	"time"
)

func Example() {
	backup := tim.NewBackup()

	task := tim.NewTask(uuid.New().String(), "foo", time.Now(), time.Now())

	task.AddRecord(tim.NewRecord(time.Now(), time.Now(), ""))
	task.AddRecord(tim.NewRecord(time.Now(), time.Now(), ""))
	task.AddRecord(tim.NewRecord(time.Now(), time.Now(), ""))

	group := tim.NewGroup(uuid.New().String(), "bar", time.Now(), time.Now())

	backup.AddGroup(group)
	backup.AddTask(task, group)

	fmt.Println(backup.ToJSON())
}
