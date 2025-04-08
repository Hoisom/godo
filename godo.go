package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type todo struct {
	Desc string `json:desc`
	Done bool   `json:done`
}

var board = []todo{}

const filename = "todos.json"

func loadTodos() {
	data, err := os.ReadFile(filename)
	if err != nil {
	} else {
		err = json.Unmarshal(data, &board)
	}
}

func saveTodos() error {
	data, err := json.MarshalIndent(board, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func newTodo() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("New todo: ")
		input, _ := reader.ReadString('\n')
		if input == "\n" {
			break
		}
		board = append(board, todo{input[:len(input)-1], false})
	}
}

func viewTodo() {
	for i := 0; i < len(board); i++ {
		if board[i].Done {
			fmt.Print(i+1, ". ", "\x1b[9m"+board[i].Desc+"\x1b[29m", "\n")
			continue
		}

		fmt.Print(i+1, ". ", board[i].Desc, "\n")
	}
}

func getInt() int {
	var input int
	fmt.Scan(&input)
	return input
}

func main() {
	loadTodos()
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Print("Welcome to Godo!\n[0] Add a new todo\n[1] View todos\n[2] Delete todo\n")

		switch getInt() {
		case 0:
			newTodo()
			saveTodos()

		case 1:
			viewTodo()
			saveTodos()
			fmt.Scanln()

		case 2:
			viewTodo()
			fmt.Print("\nEnter number of Todo to remove: ")
			num := getInt() - 1
			board[num].Done = true

		default:
			break
		}
	}
}
