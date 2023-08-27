// Author: Michael Bopp
// Filename: main.go
// Description: Go -> wasm file containing the logic for the Mandelbrot set fractal.
// Date Created: 8/26/23
// Last Edited: 8/26/23

package main

import (
	"math"
	"syscall/js"
)

const MAX_ITERATION = 80

// Define complex number structure in Go
type Complex struct {
	x, y float64
}

func mandelbrot(c Complex) (int, bool) {
	z := Complex{0, 0}
	n := 0
	var p Complex
	var d float64

	for d <= 2 && n < MAX_ITERATION {
		p = Complex{
			x: math.Pow(z.x, 2) - math.Pow(z.y, 2),
			y: 2 * z.x * z.y,
		}
		z = Complex{
			x: p.x + c.x,
			y: p.y + c.y,
		}
		d = math.Sqrt(math.Pow(z.x, 2) + math.Pow(z.y, 2))
		n += 1
	}

	return n, d <= 2
}

func draw(this js.Value, args []js.Value) interface{} {
	canvas := js.Global().Get("document").Call("getElementById", "myCanvas")
	ctx := canvas.Call("getContext", "2d")

	width := js.Global().Get("innerWidth").Int()
	height := js.Global().Get("innerHeight").Int()

	ctx.Get("canvas").Set("width", width)
	ctx.Get("canvas").Set("height", height)

	realSet := struct{ start, end float64 }{-2, 1}
	imaginarySet := struct{ start, end float64 }{-1, 1}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			c := Complex{
				x: realSet.start + (float64(i)/float64(width))*(realSet.end-realSet.start),
				y: imaginarySet.start + (float64(j)/float64(height))*(imaginarySet.end-imaginarySet.start),
			}

			_, isMandelbrotSet := mandelbrot(c)

			color := "#000" // default to black
			if !isMandelbrotSet {
				color = "#FFF" // white if not part of the Mandelbrot set
			}

			ctx.Set("fillStyle", color)
			ctx.Call("fillRect", i, j, 1, 1)
		}
	}

	return nil
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("draw", js.FuncOf(draw))
	<-c
}
