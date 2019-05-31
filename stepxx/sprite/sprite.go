package sprite

// Sprite is any game character
type Sprite interface {
	Move()
	Pos() (int, int)
	Img() string
}

func makeMove(oldPos Point, dir string) Point {
	var fn func() (Point, error)
	switch dir {
	case "UP":
		fn = oldPos.Up
	case "DOWN":
		fn = oldPos.Down
	case "LEFT":
		fn = oldPos.Left
	case "RIGHT":
		fn = oldPos.Right
	default:
		return oldPos
	}

	pos, err := fn()
	if err != nil {
		return oldPos
	}

	return pos
}
