package main

import (
	"bufio"
	"fmt"
	"os"
)

var todo = []string{}

func newTodo() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("New todo: ")
		input, _ := reader.ReadString('\n')
		if input == "\n" {
			break
		}
		todo = append(todo, input)
	}
}

func viewTodo() {
	for i := 0; i < len(todo); i++ {
		fmt.Print(i+1, ". ", todo[i])
	}
}

func getInt() int {
	var input int
	fmt.Scan(&input)
	return input
}

func main() {
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Print("Welcome to Todo\n[0] Add a new todo\n[1] View todos\n[2] Delete todo\n")

		switch getInt() {
		case 0:
			newTodo()

		case 1:
			viewTodo()
			fmt.Scanln()

		case 2:
			viewTodo()
			fmt.Print("\nEnter number of Todo to remove: ")
			num := getInt() - 1
			todo[num] = "\x1b[9m" + todo[num] + "\x1b[29m"

		default:
			break
		}
	}
}
