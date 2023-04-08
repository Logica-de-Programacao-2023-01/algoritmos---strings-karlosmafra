package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma sequência de números: ")
	scanner.Scan()
	str := scanner.Text()

	int, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		fmt.Println("A string não contém apenas números.")
	} else {
		fmt.Println("A string contém apenas números.")
		fmt.Println(int)
	}
}
