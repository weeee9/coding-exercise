package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	amOrPm := s[8:]

	hourPart := s[:2]

	switch amOrPm {
	case "AM":
		if hourPart == "12" {
			return fmt.Sprintf("00%s", s[2:8])
		}
		return s[:8]

	case "PM":
		if hourPart == "12" {
			return s[:8]
		}

		tensDigit := hourPart[0]
		tensDigit += 1
		unitsDigit := hourPart[1]
		unitsDigit += 2

		return fmt.Sprintf("%c%c%s", tensDigit, unitsDigit, s[2:8])
	}

	return ""
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
