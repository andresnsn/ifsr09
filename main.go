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

	var cut []string

	for i := 0; i < len(lines); i++ {

		if len(cut) > 0 {

		}

		line := lines[i]

		if len(line) >= 128 {

			if regex.MatchString(line[117:128]) {

				if page == "" {

					page = line[117:128]

				}

				if page[5:7] != line[117:128] {

					page = line[117:128]

					cut = lines[i:]

					i = 0

				}

			}
		}
	}
}
