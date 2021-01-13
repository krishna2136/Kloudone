package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("^[A-Z]{1}[a-z]{4}[0-9]{2}")
	fmt.Println(re.FindString("kloustest"))
	fmt.Println(re.FindString("kloudtest1"))
	fmt.Println(re.FindString("kloudtest2"))
	fmt.Println(re.FindString("kloudtest3"))
}
