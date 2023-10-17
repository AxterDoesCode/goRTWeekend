package main

import (
	"fmt"
	"log"

	"github.com/AxterDoesCode/goRTWeekend/pkg/Vec3"
)

func main() {
	//image
	var imageWidth float64 = 256
	var imageHeight float64 = 256

	//Render
	fmt.Printf("P3\n%v %v\n255\n", imageWidth, imageHeight)
	for j := 0; j < int(imageWidth); j++ {
		log.Printf("\nScanlines remaning: %v", int(imageHeight)-j)
		for i := 0; i < int(imageHeight); i++ {
			r := float64(i) / (imageWidth - 1)
			g := float64(j) / (imageHeight - 1)
			b := float64(160)
			pixel_color := Vec3.NewVec3WithValues(r, g, b)
			Vec3.WriteColor(pixel_color)
		}

	}
}
