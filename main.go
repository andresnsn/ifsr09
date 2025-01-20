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

	var newLines []string

	for index, line := range lines {

		if len(line) >= 128 {

			if regex.MatchString(line[118:128]) {

				fmt.Printf("Posição do match: %d\n", index)

				if page == "" {

					page = line[118:128]

				}

				if page[5:7] != line[122:123] {

					page = line[118:128]

				} else {
					newLines = append(newLines, line)
				}

			}
		} else {
			newLines = append(newLines, line)
		}
	}

}
