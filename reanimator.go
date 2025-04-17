package reanimator

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Supervise() {

	args := []string{}

	found := false
	for _, arg := range os.Args {
		if arg == "-reanimator-supervised" {
			found = true
			continue
		}
		args = append(args, arg)
	}

	// If the flag was found, it was stripped above so we can pass the remaining args to the program.
	// If it wasn't found, add it and run the program supervised
	if found {
		os.Args = args
		return
	} else {
		args = append(os.Args, "-reanimator-supervised")
	}

	for {
		fmt.Fprintf(os.Stderr, "\nreanimator supervisor - Starting process\n")
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		fmt.Fprintf(os.Stderr, "\nreanimator supervisor - Process exited")

		if err == nil {
			fmt.Fprintf(os.Stderr, " with code 0. Exiting...\n")
			break
		} else if exitError, ok := err.(*exec.ExitError); ok {

			code := exitError.ExitCode()

			fmt.Fprintf(os.Stderr, " with code %d", code)

			if code == 1 {
				fmt.Fprintf(os.Stderr, ". Exiting...\n")
				break
			} else {
				fmt.Fprintf(os.Stderr, ". Restarting...\n")
			}
		} else {
			fmt.Println(err)
		}

		time.Sleep(3 * time.Second)
	}

	os.Exit(0)
}
