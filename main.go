package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// 0 a 29
func main() {

	file, err := os.ReadFile("other.txt")

	if err != nil {
		fmt.Println("Erro ao ler o arquivo: ", err)
		return
	}

	regex := regexp.MustCompile(`^PAGE=\d{2},\d{2}$`)

	lines := strings.Split(string(file), "\n")

	page := ""

	newLines := make([]string, 300)

	startIndexPage := 118
	endIndexPage := 128

	ocurrencyCounter := 0

	column := 0

	for index, line := range lines {

		if len(line) >= 128 && regex.MatchString(line[startIndexPage:endIndexPage]) {

			if page == "" {
				page = line[startIndexPage:endIndexPage]

			}

			if page[5:7] != line[startIndexPage+5:endIndexPage-3] {

				page = line[startIndexPage:endIndexPage]

				column += 1

				newLines[index-(6*column)] += line

				// fmt.Printf("newLines antes: %s\n", newLines[index-(6*column)])

				// fmt.Printf("Index A: %d\n", index)

				// fmt.Printf("Resultado aqui A: %d\n", index-(6*column))

				if index-(6*column) == 0 {
					// fmt.Printf("Linha atual: %s\n", line)
					// fmt.Printf("Concatenação: %s\n", newLines[0])
				}
				fmt.Printf("Concatenação: %s\n", newLines[0])
				fmt.Printf("Expressão: %d\n", index-(6*column))

			} else {

				if column > 0 {

					newLines[index-(6*column)] += line

					// fmt.Printf("Index B: %d\n", index)

					// fmt.Printf("Resultado aqui B: %d\n", index-(6*column))

					ocurrencyCounter += 1

				} else {

					newLines[index] += line

				}
			}

		} else {

			if column > 0 {

				newLines[index-(6*column)] += line

				// fmt.Printf("Index C: %d\n", index)

				// fmt.Printf("Resultado aqui C: %d\n", index-(6*column))

				ocurrencyCounter += 1

			} else {

				newLines[index] += line

			}
		}
	}

	// Criar o arquivo
	a, _ := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Erro ao criar o arquivo: %v", err)
	}
	defer a.Close()

	// Escrever as strings no arquivo
	for _, line := range newLines {
		_, err := a.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("Erro ao escrever no arquivo: %v", err)
		}
	}

}
