package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	file, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Erro ao ler o arquivo: ", err)
		return
	}

	regex := regexp.MustCompile(`^PAGE=\d{2},\d{2}$`)

	lines := strings.Split(string(file), "\n")

	page := ""

	newLines := make([]string, 60)

	startIndexPage := 118
	endIndexPage := 128

	column := 1

	for index, line := range lines {

		if len(line) >= 128 && regex.MatchString(line[startIndexPage:endIndexPage]) {

			if page == "" {

				page = line[startIndexPage:endIndexPage]

			}

			if page[5:7] != line[startIndexPage+5:endIndexPage-3] {

				page = line[startIndexPage:endIndexPage]

				column += 1

			} else {

				if column > 1 {

					newLines[index-(60*column)] += line

				} else {
					fmt.Printf("Índice: %d, Linha: %s\n", index, line)
					newLines[index] += line

				}
			}

		} else {

			if column > 1 {

				newLines[index-(60*column)] += line

			} else {

				newLines[index] += line

			}
		}
	}

}
