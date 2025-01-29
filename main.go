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

	newLines := make([]string, len(lines))

	page := ""

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

				fmt.Printf("Length before change A: %d\n", len(newLines[index-(6*column)]))

				newLines[index-(6*column)] = newLines[index-(6*column)] + line

				fmt.Printf("Length after change A: %d\n", len(newLines[index-(6*column)]))

			} else {

				if column > 0 {

					fmt.Printf("Length before change B: %d\n", len(newLines[index-(6*column)]))

					newLines[index-(6*column)] = newLines[index-(6*column)] + line

					fmt.Printf("Length after change B: %d\n", len(newLines[index-(6*column)]))

					ocurrencyCounter += 1

				} else {

					newLines[index] = newLines[index] + line

				}
			}

		} else {

			if column > 0 {

				fmt.Printf("Length before change C: %s\n", newLines[index-(6*column)])

				newLines[index-(6*column)] = newLines[index-(6*column)] + line

				fmt.Printf(" DASA %s", line)

				fmt.Printf("Length after change C: %s\n", newLines[index-(6*column)])

			} else {

				newLines[index] = newLines[index] + line

			}
		}
	}

	a, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Erro ao criar o arquivo: %v", err)
	}
	defer a.Close()

	// Grava as strings concatenadas no arquivo de saída
	for _, line := range newLines {

		a.WriteString(line)

	}

}
