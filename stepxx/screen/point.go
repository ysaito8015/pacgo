package screen

import "errors"

type Point struct {
	row, col int
}

func (p Point) Up() (Point, error) {
	p.row--
	if p.row < 0 {
		p.row = len(maze) - 1
	}
	if !IsLegal(p) {
		return Point{}, errors.New("invalid position")
	}
	return p, nil
}

func (p Point) Down() (Point, error) {
	p.row++
	if p.row == len(maze) {
		p.row = 0
	}
	if !IsLegal(p) {
		return Point{}, errors.New("invalid position")
	}
	return p, nil
}

func (p Point) Left() (Point, error) {
	p.col--
	if p.col < 0 {
		p.col = len(maze[0]) - 1
	}
	if !IsLegal(p) {
		return Point{}, errors.New("invalid position")
	}
	return p, nil
}

func (p Point) Right() (Point, error) {
	p.col++
	if p.col == len(maze[0]) {
		p.col = 0
	}
	if !IsLegal(p) {
		return Point{}, errors.New("invalid position")
	}
	return p, nil
}

func IsLegal(pos Point) bool {
	if maze[pos.row][pos.col] == '#' {
		return false
	}
	return true
}
