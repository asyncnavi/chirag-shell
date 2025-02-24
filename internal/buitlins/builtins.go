package builtins

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/fatih/color"
)

func PrintShellHeader() {
	osName := runtime.GOOS

	pwd, err := os.Getwd()
	if err != nil {
		pwd = "unknown"
	}

	currentTime := time.Now().Format("02 Jan 15:04:05")
	fmt.Printf("[%s] %s (%s) \n", color.RedString(osName), color.GreenString(pwd), currentTime)
}

func PrintOsName() error {
	cmd := exec.Command("uname")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
