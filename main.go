package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// 0 a 29
func main() {

	file, err := os.ReadFile("input.txt")

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

			fmt.Printf("Linha que eu dei match: %d\n", index)

			if page == "" {
				page = line[startIndexPage:endIndexPage]

			}

			if page[5:7] != line[startIndexPage+5:endIndexPage-3] {

				fmt.Printf("Entrou, pois: %s | %s\n", page[5:7], line[startIndexPage+5:endIndexPage-3])

				page = line[startIndexPage:endIndexPage]

				fmt.Printf("Número antigo da coluna: %d\n", column)

				column += 1

				fmt.Printf("Número novo da coluna: %d\n", column)

				newLines[index-(121*column)] += line

				fmt.Printf("Resultado aqui: %d\n", index-(61*column))

				fmt.Printf("Index atual: %d\n", index)

			} else {

				if column > 0 {

					newLines[index-(121*column)] += line

					fmt.Printf("Ocorrência X na linha: %d\n", index)

					ocurrencyCounter += 1

				} else {

					newLines[index] += line

				}
			}

		} else {

			if column > 0 {

				newLines[index-(121*column)] += line

				fmt.Printf("Ocorrência Z na linha: %d\n", index)

				ocurrencyCounter += 1

			} else {

				newLines[index] += line

			}
		}
	}

	fmt.Printf("Ocurrency Counter: %d", ocurrencyCounter)

	os.WriteFile("output.txt", []byte(strings.Join(newLines, "\n")), 0644)

}
