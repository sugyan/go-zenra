package main

import (
	"bufio"
	"fmt"
	"github.com/sugyan/go-zenra"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result, err := zenra.Zenrize(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}
