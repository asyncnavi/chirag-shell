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

	cmd := exec.Command(tokens[0], tokens[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
