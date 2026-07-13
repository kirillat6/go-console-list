package main

import (
	"bufio"
	"fmt"
	"homework/comands"
	"homework/models"
	"os"
	"strings"
	"time"

	"github.com/k0kubun/pp"
)


func main(){
	tasks := make([]models.Task, 0)
	comands.Help()
	fmt.Println("")
	
	for {
	fmt.Print("Введите команду: ")
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); !ok {
		pp.Println("Ошибка ввода")
		return
	}
	text := scanner.Text()
	fields := strings.Fields(text)
		switch fields[0] {
			case "help":
				comands.Help()
			case "add":
				str := ""
				for i := 2; i < len(fields); i++ {
					str += (fields[i] + " ")
				}
				comands.CreateTask(fields[1], str, false, time.Now(), &tasks)
			case "list":
				comands.GetAllTasks(tasks)
			case "del":
				comands.DeleteTask(&tasks, fields[1])
			case "done":
				comands.CompleteTask(tasks, fields[1])
			case "exit":
				return
		}
	}
}

