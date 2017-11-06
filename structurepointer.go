package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64 // x, y float64
	y float64
}

func distance(p, q *Point) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	var p *Point = &Point{}       // p := &Point{}
	var q *Point = &Point{10, 10} // q := &Point
	var r *Point = new(Point)     // r := new(Point)
	r.x, r.y = 100, 100
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(p.x)
	fmt.Println(p.y)
	fmt.Println(q.x)
	fmt.Println(q.y)
	fmt.Println(r.x)
	fmt.Println(r.y)
	fmt.Println(distance(p, q))
	fmt.Println(distance(p, r))
	fmt.Println(distance(q, r))
}
