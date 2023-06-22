package main

import "fmt"

func main() {
	var str string
	fmt.Print("Digite uma palavra: ")
	fmt.Scan(&str)

	var unicas string

	for car := range str {
		var cont int
		for c := range str {
			if str[car] == str[c] {
				cont++
			}
		}
		if cont == 1 {
			unicas += string(str[car])
		}
	}

	fmt.Println(unicas)
}
