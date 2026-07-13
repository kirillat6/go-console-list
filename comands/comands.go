package comands

import (
	"fmt"
	"homework/models"
	"time"

	"github.com/k0kubun/pp"
)
func CreateTask(title string,description string,isComplete bool,createdAt time.Time,tasks *[]models.Task,){
		if title != "" && description != "" {
			fmt.Println("Добавление задачи...")
			*tasks = append(*tasks, models.Task{
			Title: title,
			Description: description,
			IsComplete: isComplete,
			CreatedAt: createdAt,
			CompletedAt: nil,
			})
			fmt.Println("Задача добавлена")
		} else {
			pp.Println("Заполнены не все поля!")
		}
}

func GetAllTasks(tasks []models.Task) {
	fmt.Println()
	fmt.Println("Получение задач...")

	for k, v := range tasks {
		fmt.Println()
		pp.Println("Задача №", k + 1)
		fmt.Println("Название:", v.Title)		
		fmt.Println("Описание:", v.Description)
		fmt.Println("Выполнена:",v.IsComplete)
		fmt.Println("Дата создания:",v.CreatedAt)		
		fmt.Println("Дата выполнения:",v.CompletedAt)		
	}

}

func CompleteTask(tasks []models.Task, completedTask string) {
	for k, v := range tasks {
		if v.Title == completedTask {
			fmt.Println("Выполнение задачи...")
			tasks[k].IsComplete = true
			now := time.Now()
			tasks[k].CompletedAt = &now
			fmt.Println("Задача выполнена!")
		}
	}
}

func DeleteTask(tasks *[]models.Task, delTask string){
	if delTask != "" {
		for i, task := range *tasks {
			if task.Title == delTask {
				*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
				return
			}
		}
	} else {
		fmt.Println("Введите название задачи")
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