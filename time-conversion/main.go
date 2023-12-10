package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
	"strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */


// Complete the timeConversion function below.
func timeConversion(s string) string {

	time := strings.Split(s, ":")
	// If the last two characters are PM, add 12 to the hour
	if strings.Contains(time[2], "PM") {
		// If the hour is 12, do not add 12
		if time[0] != "12" {
			hour, _ := strconv.Atoi(time[0])
			time[0] = strconv.Itoa(hour + 12)
		}
	} else {
		// If the hour is 12, set it to 00
		if time[0] == "12" {
			time[0] = "00"
		}
	}
	
	// Remove the AM/PM from the seconds
	time[2] = time[2][0:2]
	
	// Return the time in military format
	return strings.Join(time, ":")
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

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
