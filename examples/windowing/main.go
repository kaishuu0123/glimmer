package main

import (
	"fmt"

	glimmer "github.com/kaishuu0123/glimmer"
)

func main() {
	glimmer.InitDisplayLoop("famigo", 256*2+40, 240*2+40, 256, 240, func(sharedState *glimmer.WindowState) {
		fmt.Println("HELLO")
	})
}
