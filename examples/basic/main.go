package main

import (
        "log"
        "bufio"
        "os"
        "fmt"

        "github.com/anderspitman/symbiote-go"
)

func main() {
        symbiote.Supervise()

        fmt.Println("Input text:")
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        err := scanner.Err()
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("You entered: %s\n", scanner.Text())
}
