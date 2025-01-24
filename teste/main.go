package main

import "fmt"

func main() {

	teste := make([]string, 10)

	line := "anyInfo"

	index := 9

	column := 2

	teste[index-(2*column)] += line

	fmt.Printf("Resultado da expressão: %d", index-(2*column))

	for i, value := range teste {
		fmt.Printf("Index: %d\n", i)
		fmt.Printf("Value: %s\n\n", value)
	}

	fmt.Print(teste)

	testando := []string{"info1", "info2", "info3", "info4", "info5"}

	novo := []string{"new1", "new2", "new3", "new4", "new5", "new6"}

	for index, value := range testando {

		novo[index+1] += value
	}

	fmt.Println(novo)

}
