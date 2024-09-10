package wordfunctions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var firstLetter string
var secondLetter string

func wordStorer() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input the first letter for the word you want to generate here: ")
	firstLetter, _ := reader.ReadString('\n')
	firstLetter = strings.TrimSpace(firstLetter)

	fmt.Println("The first letter will be ", firstLetter)

	fmt.Print("Input the second letter for the word you want to generate here: ")
	secondLetter, _ := reader.ReadString('\n')
	secondLetter = strings.TrimSpace(secondLetter)

	fmt.Println("The second letter will be ", secondLetter)
}
