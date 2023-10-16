package color

import (
	"fmt"

	"github.com/AxterDoesCode/goRTWeekend/pkg/Vec3"
)

type Color = Vec3.Vec3

func WriteColor(pixelColor Color) {
	ir := int(255.999 * pixelColor.X())
	ig := int(255.999 * pixelColor.Y())
	ib := int(255.999 * pixelColor.Z())
	fmt.Printf("%v %v %v\n", ir, ig, ib)
}
