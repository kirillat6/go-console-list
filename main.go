package main

import (
	"time"

	"github.com/k0kubun/pp"
)

type Task struct {
	Title string
	Description string
	IsComplete bool
	CreatedAt time.Time
	CompletedAt *time.Time
}

func main(){
	tasks := make([]Task, 0)

	CreateTask("Погулять с собакой", "Описание", false, time.Now(), &tasks)
	CreateTask("Погулять с котом", "Описание", false, time.Now(), &tasks)
	CreateTask("Погулять с собакой", "Описание", false, time.Now(), &tasks)

	CompleteTask(tasks, "Погулять с собакой")

	DeleteTask(&tasks, "Погулять с котом")

	GetAllTasks(tasks)

}

func CreateTask(title string,description string,isComplete bool,createdAt time.Time,tasks *[]Task,){
		if title != "" && description != "" {
			*tasks = append(*tasks, Task{
			Title: title,
			Description: description,
			IsComplete: isComplete,
			CreatedAt: createdAt,
			CompletedAt: nil,
			})
		} else {
			pp.Println("Заполнены не все поля!")
		}
}

func GetAllTasks(tasks []Task) {
	pp.Println()
	pp.Println("Получение задач...")

	for k, v := range tasks {
		pp.Println()
		pp.Println("Задача №", k + 1)
		pp.Println("Название:", v.Title)		
		pp.Println("Описание:", v.Description)
		pp.Println("Выполнена:",v.IsComplete)
		pp.Println("Дата создания:",v.CreatedAt)		
		pp.Println("Дата выполнения:",v.CompletedAt)		
	}

}

func CompleteTask(tasks []Task, completedTask string) {
	for k, v := range tasks {
		if v.Title == completedTask {
			pp.Println("Выполнение задачи...")
			tasks[k].IsComplete = true
			now := time.Now()
			tasks[k].CompletedAt = &now
		}
	}
}

func DeleteTask(tasks *[]Task, delTask string){
	for i, task := range *tasks {
			if task.Title == delTask {
				// Изменяем слайс по указателю
				*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
				return
			}
		}
}

func Help() {
	pp.Println()
	pp.Println("Список команд:")
	pp.Println("help — эта команда позволяет узнать доступные команды и их формат")
	pp.Println("add {заголовок задачи из одного слова} {текст задачи из одного или" + 
	"нескольких слов} — эта команда позволяет добавлять новые задачи в список задач")
	pp.Println("list — эта команда позволяет получить полный список всех задач")
	pp.Println("del {заголовок существующей задачи} — эта команда позволяет" + 
	"удалить задачу по её заголовку")
	pp.Println("done {заголовок существующей задачи} — эта команда позволяет" + 
	"отменить задачу как выполненную")
	pp.Println("events — эта команда позволяет получить список всех событий")
	pp.Println("exit — эта команда позволяет завершить выполнение программы")

} 