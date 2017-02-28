package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

import _ "strconv"

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	alpha := make(map[string]int)

	for _, b := range input {
		norm := strings.ToLower(string(b))
		res, _ := regexp.MatchString("[a-z]", norm)

		if res {
			alpha[norm] = 1
		}
	}

	//fmt.Printf("%v\n", alpha)
	numAlpha := 0

	// loop over a to z
	for i := 97; i <= 122; i++ {
		char := fmt.Sprintf("%c", i)
		//fmt.Printf("checking for %s\n", char)

		// check if exists in map
		if _, present := alpha[char]; present {
			numAlpha++
		}
	}

	if numAlpha == 26 {
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}
