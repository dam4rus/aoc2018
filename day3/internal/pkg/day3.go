package day3

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

type Rect struct {
	Position      Point
	Width, Height int
}

type Claim struct {
	Id   int
	Rect Rect
}

func (rect Rect) Left() int {
	return rect.Position.X
}

func (rect Rect) Right() int {
	return rect.Position.X + rect.Width
}

func (rect Rect) Top() int {
	return rect.Position.Y
}

func (rect Rect) Bottom() int {
	return rect.Position.Y + rect.Height
}

func (rect Rect) Points() []Point {
	var points []Point
	for x := rect.Left(); x < rect.Right(); x += 1 {
		for y := rect.Top(); y < rect.Bottom(); y += 1 {
			points = append(points, Point{X: x, Y: y})
		}
	}
	return points
}

func (rect Rect) Intersects(otherRect Rect) bool {
	return rect.horizontallyIntersects(otherRect) && rect.verticallyIntersects(otherRect)
}

func (rect Rect) Intersection(otherRect Rect) *Rect {
	if !rect.Intersects(otherRect) {
		return nil
	}
	var left int
	if rect.Left() > otherRect.Left() {
		left = rect.Left()
	} else {
		left = otherRect.Left()
	}
	var top int
	if rect.Top() > otherRect.Top() {
		top = rect.Top()
	} else {
		top = otherRect.Top()
	}
	var right int
	if rect.Right() < otherRect.Right() {
		right = rect.Right()
	} else {
		right = otherRect.Right()
	}
	var bottom int
	if rect.Bottom() < otherRect.Bottom() {
		bottom = rect.Bottom()
	} else {
		bottom = otherRect.Bottom()
	}
	return &Rect{
		Position: Point{
			X: left,
			Y: top,
		},
		Width:  right - left,
		Height: bottom - top,
	}
}

func (rect Rect) horizontallyIntersects(otherRect Rect) bool {
	return (otherRect.Left() < rect.Right()) && (otherRect.Right() > rect.Left())
}

func (rect Rect) verticallyIntersects(otherRect Rect) bool {
	return (otherRect.Top() < rect.Bottom()) && (otherRect.Bottom() > rect.Top())
}

func ParseClaim(line string) (*Claim, error) {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text()[1:])
	if err != nil {
		return nil, err
	}
	scanner.Scan()
	scanner.Scan()
	positionWord := scanner.Text()
	position := strings.Split(positionWord, ",")
	left, err := strconv.Atoi(position[0])
	if err != nil {
		return nil, err
	}
	top, err := strconv.Atoi(position[1][:len(position[1])-1])
	if err != nil {
		return nil, err
	}
	scanner.Scan()
	sizeWord := scanner.Text()
	size := strings.Split(sizeWord, "x")
	width, err := strconv.Atoi(size[0])
	if err != nil {
		return nil, err
	}
	height, err := strconv.Atoi(size[1])
	if err != nil {
		return nil, err
	}
	return &Claim{
		Id: id,
		Rect: Rect{
			Position: Point{X: left, Y: top},
			Width:    width,
			Height:   height,
		},
	}, nil
}

func ParseInput(lines []string) ([]Claim, error) {
	var claims []Claim
	for _, line := range lines {
		claim, err := ParseClaim(line)
		if err != nil {
			return []Claim{}, err
		}
		claims = append(claims, *claim)
	}
	return claims, nil
}

func OverlappingPoints(claims []Claim) []Point {
	overlappingPointMap := make(map[Point]struct{})
	for i, claim := range claims {
		for _, otherClaim := range claims[i+1:] {
			intersection := claim.Rect.Intersection(otherClaim.Rect)
			if intersection != nil {
				// fmt.Println(*intersection)
				for _, point := range intersection.Points() {
					overlappingPointMap[point] = struct{}{}
				}
			}
		}
	}
	var overlappingPoints []Point
	for key, _ := range overlappingPointMap {
		overlappingPoints = append(overlappingPoints, key)
	}
	return overlappingPoints
}

func NonOverlapping(claims []Claim) (int, error) {
	nonOverlappingClaims := make(map[Claim]struct{})
	for _, claim := range claims {
		nonOverlappingClaims[claim] = struct{}{}
	}
	for i, claim := range claims {
		for _, otherClaim := range claims[i+1:] {
			intersection := claim.Rect.Intersection(otherClaim.Rect)
			if intersection != nil {
				delete(nonOverlappingClaims, claim)
				delete(nonOverlappingClaims, otherClaim)
			}
		}
	}
	if len(nonOverlappingClaims) > 1 {
		return 0, errors.New("found more non overlapping claims")
	}
	for key, _ := range nonOverlappingClaims {
		return key.Id, nil
	}
	return 0, errors.New("not found non overlapping claims")
}
