package main

import (
	"fmt"
	"log"

	"github.com/AxterDoesCode/goRTWeekend/pkg/Color"
	"github.com/AxterDoesCode/goRTWeekend/pkg/Vec3"
)


func main() {
	//image
    var imageWidth float64 = 256
    var imageHeight float64 = 256

	//Render
	fmt.Printf("P3\n%v %v\n255\n", imageWidth, imageHeight)
	for j := 0; j < int(imageWidth); j++ {
        log.Printf("\nScanlines remaning: %v", int(imageHeight) - j)
		for i := 0; i < int(imageHeight); i++ {
            r := float64(i) / (imageWidth - 1)
            g := float64(j) / (imageHeight - 1)
            b := float64(0)
            pixel_color := color.Color(Vec3.NewVec3WithValues(r,g,b))
            color.WriteColor(pixel_color)
		}

	}
}
