// Surface computes an SVG rendering of a 3-D surface function.
package ch3

import (
	"fmt"
	"io/ioutil"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func Surface() {
	var surfaceStr string
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	surfaceStr += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) {
				fmt.Println("ax | ay is nan.")
				fmt.Println(">>>>>>>>>>>>>>>>>>> a stop >>>>>>>>>>>>>>>>>")
			}
			if math.IsNaN(bx) || math.IsNaN(by) {
				fmt.Println("bx | dy is nan.")
				fmt.Println(">>>>>>>>>>>>>>>>>>> a stop >>>>>>>>>>>>>>>>>")
			}
			if math.IsNaN(cx) || math.IsNaN(cy) {
				fmt.Println("cx | cy is nan.")
				fmt.Println(">>>>>>>>>>>>>>>>>>> a stop >>>>>>>>>>>>>>>>>")
			}
			if math.IsNaN(dx) || math.IsNaN(dy) {
				fmt.Println("dx | dy is nan.")
				fmt.Println(">>>>>>>>>>>>>>>>>>> a stop >>>>>>>>>>>>>>>>>")
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			surfaceStr += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
	surfaceStr += fmt.Sprintf("</svg>")

	content := []byte(surfaceStr)
	err := ioutil.WriteFile(
		"/Users/qiujiayu/GolandProjects/awesomeProject/utils/surface.html",
		content,
		0644)
	if err != nil {
		fmt.Println(">>>>>>>>>>>>>>>>>>> error start >>>>>>>>>>>>>>>>>")
		fmt.Println(err)
		fmt.Println(">>>>>>>>>>>>>>>>>>>  error end  >>>>>>>>>>>>>>>>>")
	}
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := eggbox(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}


func eggbox(x, y float64) float64 { //鸡蛋盒
	r := 0.2 * (math.Cos(x) + math.Cos(y))
	return r
}

func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	r := y*y/a2 - x*x/b2
	return r
}