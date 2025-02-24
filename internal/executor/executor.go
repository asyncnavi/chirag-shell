package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func isExitRequest(tokens []string) bool {
	return (len(tokens) == 1) && tokens[0] == "exit"
}

func ExitShell() {
	fmt.Println("Thank you for using ChiragSH 🙏")
	os.Exit(1)
}

func ExecuteInput(input string) error {

	tokens := strings.Fields(input)

	if len(tokens) == 0 {
		return nil
	}

	if isExitRequest(tokens) {
		ExitShell()
	}

	if tokens[0] == "cd" {
		if len(tokens) < 2 {
			return fmt.Errorf("cd: missing argumnent")
		}
		if err := os.Chdir(tokens[1]); err != nil {
			return fmt.Errorf("cd %v", err)
		}
		return nil
	}

	cmd := exec.Command(tokens[0], tokens[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
