package main

import (
	"fmt"
)

func main() {
	//image

    var imageWidth float64 = 256
    var imageHeight float64 = 256

	//Render
	fmt.Printf("P3\n%v %v\n255\n", imageWidth, imageHeight)
	for j := 0; j < int(imageWidth); j++ {
		for i := 0; i < int(imageHeight); i++ {
            r := float64(i) / (imageWidth - 1)
            g := float64(j) / (imageHeight - 1)
            b := 0

            ir := int(255.999 * float64(r))
            ig := int(255.999 * float64(g))
            ib := int(255.999 * float64(b))

            fmt.Printf("%v %v %v\n", ir, ig, ib)

		}

	}
}
