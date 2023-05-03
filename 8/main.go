package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Exec(line string) {
	arguments := strings.Split(line, " ")
	switch arguments[0] {
	case "cd":
		err := os.Chdir(arguments[1])
		if err != nil {
			log.Println(err)
		}
	case "pwd":
		res, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(res)
	case "echo":
		fmt.Println(arguments[1])
	case "kill":
		pid, err := strconv.Atoi(arguments[1])
		if err != nil {
			log.Println(err)
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			log.Println(err)
		}
		proc.Kill()
	case "ps":
		procs, err := exec.Command("ps", arguments[1]).Output()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(procs))
	case "exec":
		cmd := exec.Command(arguments[1])
		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		Exec(scanner.Text())
	}
}
