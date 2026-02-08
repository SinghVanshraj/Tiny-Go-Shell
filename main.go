package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/user"
	"os/exec"
	"strings"
)

var commandsHistory[]string

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		hostname,err := os.Hostname()
		u, err:= user.Current()
		getDir, err := os.Getwd()
		fmt.Printf("%s@%s: %s$ ", u.Username,hostname,getDir)
		input, err := reader.ReadString('\n')
		commandsHistory = append(commandsHistory, input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr,err)
		}

	}
}

var ErrNoPath = errors.New("path required")

func listHistory() {
	for i := 0; i < len(commandsHistory); i++ {
		fmt.Println(commandsHistory[i])
	}
}

func execInput(input string) error {

	input = strings.TrimSuffix(input, "\n");

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return ErrNoPath
		}

		return os.Chdir(args[1])
		
	case "history":
		listHistory()
		return nil
	
	case "exit" :
		os.Exit(0)

	}

	cmd := exec.Command(args[0],args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}