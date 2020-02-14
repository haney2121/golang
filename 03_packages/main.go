package main

import (
	"fmt"
	"math"

	"github.com/haney2121/crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(3.4))
	fmt.Println(math.Ceil(3.4))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("Justin"))
}
