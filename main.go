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

	s := rl.Vector2{X: 100, Y: 100}
	cp := rl.Vector2{X: 700, Y: 100}
	e := rl.Vector2{X: 700, Y: 350}

	var steps int32 = 10
	var line_thick float32 = 5.0

	pts_lb := linearBezier(s, e, steps)
	pts_qb := quadraticBezier(s, cp, e, steps)

	dotCol := rl.Red
	dotCol.A = 0xAA

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.LightGray)

		for i := 0; i < 10; i++ {
			rl.DrawLineEx(s, e, 3.0, rl.Black)
		}

		for _, p := range pts_lb {
			rl.DrawCircleV(p, line_thick, dotCol)
		}

		for i := 0; i < 10; i++ {
			rl.DrawLineEx(pts_qb[i], pts_qb[i+1], 3.0, rl.Blue)
		}

		for _, p := range pts_qb {
			rl.DrawCircleV(p, line_thick, dotCol)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func linearBezier(s rl.Vector2, e rl.Vector2, nSteps int32) []rl.Vector2 {
	// a straight line! basically learp

	pts := make([]rl.Vector2, nSteps)

	step := 1.0 / float32(nSteps)
	var i int32

	for i = 0; i < nSteps; i++ {

		a := float32(i) * step

		b := rl.Vector2{X: a, Y: a}
		c := rl.Vector2{X: 1.0 - a, Y: 1.0 - a}

		pts[i] = rl.Vector2Add(rl.Vector2Multiply(s, c), rl.Vector2Multiply(e, b))
	}

	return pts
}

func quadraticBezier(p0, p1, p2 rl.Vector2, nSteps int32) []rl.Vector2 {

	// a straight line! basically learp

	pts := make([]rl.Vector2, nSteps+1)

	step := 1.0 / float32(nSteps)
	var i int32

	for i = 0; i < nSteps; i++ {
		t := float32(i) * step

		x := p1.X + (1-t)*(1-t)*(p0.X-p1.X) + t*t*(p2.X-p1.X)
		y := p1.Y + (1-t)*(1-t)*(p0.Y-p1.Y) + t*t*(p2.Y-p1.Y)

		pts[i] = rl.Vector2{X: x, Y: y}

	}

	pts[nSteps] = p2

	return pts
}
