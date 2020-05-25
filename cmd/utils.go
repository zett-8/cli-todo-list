package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type IntList []int

func (il IntList) contains(n int) bool {
	for _, v := range il {
		if v == n {
			return true
		}
	}
	return false
}

type Todo struct {
	id   int
	done bool
	todo string
	tags []string
}

func (t Todo) containsTag(tag string) bool {
	if string(tag[0]) != "@" {
		tag = "@" + tag
	}

	for _, v := range t.tags {
		if v == tag {
			return true
		}
	}

	return false
}

func getFilePath() string {
	path, err := os.Executable()
	if err != nil {
		ErrorMsg(err)
	}

	return filepath.Join(path + "_data.txt")
}

func ErrorMsg(err error) {
	if os.Getenv("APP_ENV") == "dev" {
		fmt.Println(err)
	} else {
		fmt.Println("oops, something went wrong")
	}
}

func ReadData() ([]Todo, error) {

	f, err := os.OpenFile(getFilePath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		ErrorMsg(err)
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadFile(getFilePath())
	if err != nil {
		ErrorMsg(err)
		return nil, err
	}

	var list []string
	if len(b) > 0 {
		list = strings.Split(string(b), "\n")
	}

	var todos []Todo
	for i, v := range list {
		sp := strings.Split(v, "_*_*_*_")
		done := false
		todo := sp[0]
		tags := strings.Split(sp[1], " ")

		if len(todo) > 2 && todo[:2] == "!!" {
			done = true
			todo = todo[2:]
		}

		todos = append(todos, Todo{i + 1, done, todo, tags})
	}

	return todos, nil
}

func WriteData(todos []Todo) error {
	var strList []string

	for _, v := range todos {
		todo := v.todo
		if v.done {
			todo = "!!" + todo
		}
		tags := strings.Join(v.tags, " ")

		strList = append(strList, todo+"_*_*_*_"+tags)
	}

	mydata := []byte(strings.Join(strList, "\n"))

	err := ioutil.WriteFile(getFilePath(), mydata, 0777)
	if err != nil {
		return err
	}

	ConsoleList(todos, "", false)

	return nil
}

func ConsoleList(todos []Todo, tag string, without bool) {
	if tag != "" {
		var filtered []Todo
		for _, todo := range todos {
			if todo.containsTag(tag) {
				filtered = append(filtered, todo)
			}
		}

		todos = filtered
	}

	if without {
		var filtered []Todo
		for _, todo := range todos {
			if !todo.done {
				filtered = append(filtered, todo)
			}
		}

		todos = filtered
	}

	for i, t := range todos {
		bold := color.New(color.Bold).SprintFunc()
		faint := color.New(color.Faint).SprintFunc()

		numberColumn := bold("[" + strconv.Itoa(i+1) + "]")
		todo := t.todo
		if t.done {
			todo = faint(t.todo)
		}
		tags := color.MagentaString(strings.Join(t.tags, " "))

		fmt.Printf("%s %s   %s\n", numberColumn, todo, tags)
	}
}

func Reset() {
	err := os.Remove(getFilePath())
	if err != nil {
		fmt.Println(err)
		return
	}
}
