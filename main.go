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

	file, err := os.ReadFile("input.txt")

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

	highest := 0

	for _, line := range lines {
		if len(line) > highest {
			highest = len(line)
		}
	}

	for index, line := range lines {
		if len(line) < highest {
			difference := highest - len(line)
			line += strings.Repeat(" ", difference)
			lines[index] = line
		}

		fmt.Printf("Novo len da line: %d\n", len(line))
	}

	for index, line := range lines {

		if len(line) >= 128 && regex.MatchString(line[startIndexPage:endIndexPage]) {

			if page == "" {
				page = line[startIndexPage:endIndexPage]

			}

			if page[5:7] != line[startIndexPage+5:endIndexPage-3] {

				page = line[startIndexPage:endIndexPage]

				column += 1

				newLines[index-(120*column)] = newLines[index-(120*column)] + line

			} else {

				if column > 0 {

					newLines[index-(120*column)] = newLines[index-(120*column)] + line

					ocurrencyCounter += 1

				} else {

					newLines[index] = newLines[index] + line

				}
			}

		} else {

			if column > 0 {

				newLines[index-(120*column)] = newLines[index-(120*column)] + line

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

	for _, line := range newLines {
		// Remove todos os caracteres \r da linha antes de gravar
		cleanedLine := strings.Replace(line, "\r", "", -1)
		a.WriteString(cleanedLine + "\n") // Adiciona \n para nova linha
	}

}
