package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const SCREEN_WIDTH = 800
const SCREEN_HEIGHT = 450

func main() {

	// move this to create scene or transition scene later

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)

	s := Point{X: 100, Y: 100}
	e := Point{X: 700, Y: 350}

	pts := linearBezier(s, e, 10)

	dotCol := rl.Red
	dotCol.A = 0xAA

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.LightGray)

		for i := 0; i < 10-1; i++ {
			rl.DrawLine(pts[i].X, pts[i].Y, pts[i+1].X, pts[i+1].Y, rl.Black)
			rl.DrawLineEx(
				rl.Vector2{X: float32(pts[i].X), Y: float32(pts[i].Y)},
				rl.Vector2{X: float32(pts[i+1].X), Y: float32(pts[i+1].Y)},
				3.0,
				rl.Black,
			)
		}

		for _, p := range pts {
			rl.DrawCircleV(rl.Vector2{X: float32(p.X), Y: float32(p.Y)}, 5.0, dotCol)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

type Point struct {
	X int32
	Y int32
}

type PointF struct {
	X float64
	Y float64
}

func (pf *PointF) FromPoint(p Point) {
	pf.X = float64(p.X)
	pf.Y = float64(p.Y)
}

func (p1 PointF) sub(p2 PointF) PointF {
	return PointF{X: p1.X - p2.X, Y: p1.Y - p2.Y}
}

func (p1 PointF) mul(v float64) PointF {
	return PointF{X: p1.X * v, Y: p1.Y * v}
}

func (p1 PointF) add(p2 PointF) PointF {
	return PointF{X: p1.X + p2.Y, Y: p1.Y + p2.Y}
}

func linearBezier(s Point, e Point, nSteps int) []Point {
	// a straight line! basically learp

	pts := make([]Point, nSteps)

	var sp PointF
	sp.FromPoint(s)

	var ep PointF
	ep.FromPoint(e)

	d := ep.sub(sp)

	step := 1.0 / float64(nSteps)

	for i := 0; i < nSteps; i++ {
		pf := d.mul(float64(i) * step).add(sp)
		pts[i] = Point{X: int32(pf.X), Y: int32(pf.Y)}
	}

	return pts
}
