package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma sequência numérica (sem espaço ou vírgula): ")
	scanner.Scan()
	str := scanner.Text()

	cres := true

	ints, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		fmt.Println("A string não é uma sequência de números inteiros.")
	} else {
		for i := 0; i < len(str)-1; i++ {
			if str[i+1] < str[i] {
				cres = false
			}
		}
		if cres == true {
			fmt.Println("A string é uma sequência crescente.")
			fmt.Println(ints)
		} else {
			fmt.Println("A string não é uma sequência crescente.")
			fmt.Println(ints)
		}
	}
}
