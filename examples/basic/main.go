package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/anderspitman/reanimator-go"
)

func main() {
	reanimator.Supervise()

	fmt.Println("Input text:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You entered: %s\n", scanner.Text())
}
