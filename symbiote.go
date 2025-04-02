package reanimator

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Supervise() {

	foundIdx := -1
	for i, arg := range os.Args {
		if arg == "-reanimator-supervised" {
			foundIdx = i
		}
	}

	if foundIdx < 0 {
		return
	}

	args := append(os.Args[:foundIdx], os.Args[foundIdx+1:]...)

	for {
		fmt.Fprintf(os.Stderr, "reanimator supervisor - Starting process\n")
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		fmt.Fprintf(os.Stderr, "reanimator supervisor - Process exited\n")

		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}

	os.Exit(0)
}
