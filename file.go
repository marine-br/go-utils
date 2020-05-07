package utils

import (
	"bufio"
	"log"
	"os"
)

type processText func(string)

//ReadFileLineByLine Lê linha a linha de um arquivo de texto e envia para uma função de callback tratar
func ReadFileLineByLine(filename string, callback processText) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		callback(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
