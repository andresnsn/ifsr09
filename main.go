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

	newFileIndex := 0

	for index, line := range lines {

		if len(line) >= 128 && regex.MatchString(line[118:128]) {

			if page == "" {

				page = line[118:128]

			}

			if page[5:7] != line[122:123] {

				page = line[118:128]

			} else {
				newLines[index] += line
			}

		} else {
			newLines[index] += line
		}
	}

}
