package symbiote

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Supervise() {
	if len(os.Args) >= 2 && os.Args[1] == "-symbiote-supervised" {
		os.Args = append(os.Args[:1], os.Args[2:]...)
		return
	}

	for {
		fmt.Fprintf(os.Stderr, "symbiote supervisor - Starting process\n")
		args := append([]string{"-symbiote-supervised"}, os.Args[1:]...)
		cmd := exec.Command(os.Args[0], args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		fmt.Fprintf(os.Stderr, "symbiote supervisor - Process exited\n")

		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}

	os.Exit(0)
}
