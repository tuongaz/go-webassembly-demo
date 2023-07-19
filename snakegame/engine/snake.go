package engine

import (
	"snakegame/coord"
)

const (
	snakeInitialLength = 3
)

type Direction string

const (
	UP    Direction = "up"
	RIGHT Direction = "right"
	DOWN  Direction = "down"
	LEFT  Direction = "left"
)

func newSnake(head coord.Coord) *Snake {
	body := make([]coord.Coord, snakeInitialLength)
	for i := 0; i < snakeInitialLength; i++ {
		body[snakeInitialLength-i-1] = coord.Coord{X: head.X, Y: head.Y + i}
	}

	return &Snake{
		Direction: UP,
		Body:      body,
	}
}

type Snake struct {
	Direction Direction     `json:"direction"`
	Body      []coord.Coord `json:"body"`
	Points    int           `json:"points"`
}

func (s *Snake) changeDirection(direction Direction) {
	switch direction {
	case UP:
		if s.Direction == DOWN {
			return
		}
	case RIGHT:
		if s.Direction == LEFT {
			return
		}
	case DOWN:
		if s.Direction == UP {
			return
		}
	case LEFT:
		if s.Direction == RIGHT {
			return
		}
	}
	s.Direction = direction
}

func (s *Snake) addPoints(p int) {
	s.Points += p
}

func (s *Snake) head() coord.Coord {
	head := s.Body[len(s.Body)-1]
	return coord.Coord{X: head.X, Y: head.Y}
}

func (s *Snake) grow() {
	head := s.head()

	switch s.Direction {
	case UP:
		head.Y--
	case RIGHT:
		head.X++
	case DOWN:
		head.Y++
	case LEFT:
		head.X--
	}

	s.Body = append(s.Body, head)
}

func (s *Snake) move() error {
	head := s.head()

	switch s.Direction {
	case UP:
		head.Y--
	case RIGHT:
		head.X++
	case DOWN:
		head.Y++
	case LEFT:
		head.X--
	}

	if s.isOnBody(head) {
		return s.die("snake hit itself")
	}

	s.Body = append(s.Body[1:], head)

	return nil
}

func (s *Snake) die(msg string) error {
	return &DieError{
		msg: msg,
	}
}

func (s *Snake) isOnBody(head coord.Coord) bool {
	for _, b := range s.Body {
		if b.X == head.X && b.Y == head.Y {
			return true
		}
	}

	return false
}
