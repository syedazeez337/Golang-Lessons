package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("file.go")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	fmt.Println("File contents: ")
	fmt.Println(string(content))

}

/*

	file, err := os.Create("file.go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fileContent := `
package main

func Add(x, y int) int {
	return x + y
}
`
	file.WriteString(fileContent)
	fmt.Println("File created succefully...")
	x := Add(3, 4)
	fmt.Println(x)
*/

// how to handle csv files
// how to get the data from the file
// how to use flags go tool compile -S main.go > main.s
// using flag
// how to parse a csv file
// structs, maps
