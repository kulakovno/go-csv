package main

import (
	"bufio"
	"strings"
	"encoding/csv"
	"io/ioutil"
	"os"
	"fmt"
	"io"
)
const RESULT_FILE_NAME = "result.xlsx"

func main() {

	os.Remove(RESULT_FILE_NAME)
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		fileName := f.Name()
		if strings.Contains(fileName, ".csv") {
			lines := parseFile(fileName)
			addLinesToFile(lines)
			fmt.Printf("Добавлен " + fileName)
		}

	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Нажмите любую клавишу для завершения")
	reader.ReadRune()

	os.Exit(0)
	// Load a CSV file.

}

func parseFile(file string) [][]string {
	var result [][]string
	f, _ := os.Open(file)
	slicedName := strings.Split(file, "_")
	slicedExt := strings.Split(slicedName[len(slicedName)-1], ".")
	DIST := slicedExt[0]
	fmt.Println(DIST)
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	i := 0
	for {
		record, err := r.Read()

		// Stop at EOF.
		if err == io.EOF {
			break
		}

		//skip first 3 lines
		i++
		if i < 4 {
			continue
		}

		resultRecord := strings.Split(record[0], ";")
		resultRecord[0] = DIST
		//resultRecord[len(resultRecord)-1] = "\n"
		//resultRecord = append(resultRecord, "\n")
		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.
		//fmt.Println(strings.Join(resultRecord, ";"))
		resultString := strings.Join(resultRecord, ";")
		resultSlice := make([]string, 1)
		resultSlice[0] = resultString
		result = append(result, resultSlice)
		//fmt.Println(len(record))
		//for value := range record {
		//	fmt.Printf("  %v\n", record[value])
		//}
	}
	return result
}

func addLinesToFile(lines [][]string) {
	resultFile, _ := os.OpenFile(RESULT_FILE_NAME, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	writer := csv.NewWriter(resultFile)
	writer.WriteAll(lines)
}
