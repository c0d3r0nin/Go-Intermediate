package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// TODO: the Shape interface (Area() int and Name() string), the three
// shape types, their String methods, and the sortable slice type.

type Shape interface {
	Area() int
	Name() string
	String() string
}

// Rectangle
type Rect struct{ W, H int }

func (r Rect) Area() int    { return r.W * r.H }
func (r Rect) Name() string { return "rect" }
func (r Rect) String() string {
	return fmt.Sprintf("rect %dx%d area=%d", r.W, r.H, r.Area())
}

// Square
type Square struct{ S int }

func (s Square) Area() int    { return s.S * s.S }
func (s Square) Name() string { return "square" }
func (s Square) String() string {
	return fmt.Sprintf("square %d area=%d", s.S, s.Area())
}

// Triangle
type Tri struct{ B, H int }

func (t Tri) Area() int    { return (t.B * t.H) / 2 }
func (t Tri) Name() string { return "tri" }
func (t Tri) String() string {
	return fmt.Sprintf("tri %dx%d area=%d", t.B, t.H, t.Area())
}

// num returns word i of a split line as an int (0 if the word is missing).
func num(f []string, i int) int {
	if i >= len(f) {
		return 0
	}
	n, _ := strconv.Atoi(f[i])
	return n
}

// keeps the sort import alive until you use it
//var _ = sort.Stable

// Sorting Shape slices
type Shapes []Shape

func (s Shapes) Len() int { return len(s) }

func (s Shapes) Less(i, j int) bool {
	if s[i].Area() != s[j].Area() {
		return s[i].Area() > s[j].Area()
	}
	return s[i].Name() < s[j].Name()
}

func (s Shapes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	var lines [][]string // each input line split into words: {"rect", "3", "4"}
	for sc.Scan() {
		f := strings.Fields(sc.Text())
		if len(f) > 0 {
			lines = append(lines, f)
		}
	}

	var shapes Shapes
	var unknown []string
	var rects, tris, squares int

	// TODO: turn each line into a Shape (or remember its unknown kind),
	// sort, print the shapes, the count line, then the unknown kinds.
	for _, f := range lines {
		switch f[0] {
		case "rect":
			shapes = append(shapes, Rect{W: num(f, 1), H: num(f, 2)})
			rects++

		case "square":
			shapes = append(shapes, Square{S: num(f, 1)})
			squares++

		case "tri":
			shapes = append(shapes, Tri{B: num(f, 1), H: num(f, 2)})
			tris++

		default:
			unknown = append(unknown, f[0])
		}
	}

	// Once all shapes have been stores into []Shape, sort the array.
	// Use sort.Stable

	sort.Stable(shapes)

	for _, shape := range shapes {
		fmt.Println(shape.String())
	}

	fmt.Printf("rects=%d squares=%d tris=%d\n", rects, squares, tris)

	for _, kind := range unknown {
		fmt.Println("unknown:", kind)
	}
}
