package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "regexp"

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	re, _ := regexp.Compile("[[:upper:]]")
	numUp := re.FindAllString(text, -1)
	fmt.Println(len(numUp) + 1)

}
