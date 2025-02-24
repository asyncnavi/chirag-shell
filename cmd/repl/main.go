package main

import (
	"bufio"
	"fmt"
	"github.com/asyncnavi/chirag-shell/internal/buitlins"
	"github.com/asyncnavi/chirag-shell/internal/executor"
	"os"
	"os/exec"
)

func init() {
	showDir := exec.Command("pwd")
	showDir.Stdout = os.Stdout

	err := showDir.Run()
	if err != nil {
		return
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {

		// Prints shell header ( with os name, current working directory)
		builtins.PrintShellHeader()

		fmt.Print("🪔 ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(os.Stderr, err)
		}

		if err := executor.ExecuteInput(input); err != nil {
			fmt.Println(os.Stderr, err)
		}
	}
}
