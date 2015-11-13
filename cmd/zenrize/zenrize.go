package main

import (
	"bufio"
	"fmt"
	"github.com/sugyan/go-zenra"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result := zenra.Zenrize(scanner.Text())
		fmt.Println(result)
	}
}
