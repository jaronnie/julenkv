package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jaronnie/julenkv/client/go-julenkv"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	conn, err := julenkv.Connect("0.0.0.0:5200")
	if err != nil {
		log.Fatal(err)
	}
	for {
		printHeader()
		input, _ := reader.ReadString('\n')
		cmd := formatInput(input)
		reply, err := conn.Do(cmd...)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(reply)
		}
	}
}

func printHeader() {
	fmt.Print("julenkv-cli> ")
}

func formatInput(cmd string) []string {
	cmd = strings.TrimRight(cmd, "\n")
	formatInputFields := strings.Fields(cmd)
	return formatInputFields
}
