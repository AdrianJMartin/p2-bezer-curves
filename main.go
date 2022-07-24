package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const SCREEN_WIDTH = 600
const SCREEN_HEIGHT = 600

func main() {

	// move this to create scene or transition scene later

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)

	const BEZIER_CURVE_STEPS = 10

	curveList := []drawable{
		CreateLinearBezierCurve(rl.Vector2{X: 100, Y: 100}, rl.Vector2{X: 200, Y: 200}, BEZIER_CURVE_STEPS),
		CreateQuadraticBezierCurve(rl.Vector2{X: 200, Y: 200}, rl.Vector2{X: 600, Y: 0}, rl.Vector2{X: 500, Y: 500}, BEZIER_CURVE_STEPS),
	}

	dotCol := rl.Red
	dotCol.A = 0xAA

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.LightGray)

		for _, i := range curveList {
			i.draw()
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

type drawable interface {
	draw()
}

type BezierCurveType int8

const (
	LINEAR    BezierCurveType = 1
	QUADRATIC BezierCurveType = 2
)

type BezierCurve struct {
	curveType      BezierCurveType
	nSteps         int
	pts            []rl.Vector2
	line_thickness float32
	line_color     rl.Color
	vertex_color   rl.Color
	vertex_size    float32
}

func (bc *BezierCurve) CurveDefaults() {

	bc.line_color = rl.Black
	bc.line_thickness = 3.0

	bc.vertex_color = rl.Red
	bc.vertex_size = 5.0
}

func CreateLinearBezierCurve(p0, p1 rl.Vector2, nSteps int) BezierCurve {
	// a straight line! basically learp

	pts := make([]rl.Vector2, nSteps+1)

	step := 1.0 / float32(nSteps)
	var i int

	for i = 0; i <= nSteps; i++ {

		a := float32(i) * step

		b := rl.Vector2{X: a, Y: a}
		c := rl.Vector2{X: 1.0 - a, Y: 1.0 - a}

		pts[i] = rl.Vector2Add(rl.Vector2Multiply(p0, c), rl.Vector2Multiply(p1, b))
	}

	var bc BezierCurve
	bc.CurveDefaults()
	bc.nSteps = nSteps
	bc.pts = pts
	bc.curveType = LINEAR

	return bc
}

func CreateQuadraticBezierCurve(p0, p1, p2 rl.Vector2, nSteps int) BezierCurve {
	// a straight line! basically learp
	pts := make([]rl.Vector2, nSteps+1)

	step := 1.0 / float32(nSteps)
	var i int

	for i = 0; i <= nSteps; i++ {

		t := float32(i) * step

		x := p1.X + (1-t)*(1-t)*(p0.X-p1.X) + t*t*(p2.X-p1.X)
		y := p1.Y + (1-t)*(1-t)*(p0.Y-p1.Y) + t*t*(p2.Y-p1.Y)

		pts[i] = rl.Vector2{X: x, Y: y}

	}

	var bc BezierCurve
	bc.CurveDefaults()
	bc.nSteps = nSteps
	bc.pts = pts
	bc.curveType = QUADRATIC

	return bc
}

func (c BezierCurve) draw() {

	for i := 0; i <= c.nSteps-1; i++ {

		fmt.Printf("%d , %d \n", i, i+1)

		rl.DrawLineEx(c.pts[i], c.pts[i+1], 3.0, rl.Black)
	}

	for _, p := range c.pts {
		rl.DrawCircleV(p, c.line_thickness, c.vertex_color)
	}

}

type CurveList []BezierCurve
