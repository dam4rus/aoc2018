package day6

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/adam-lavrik/go-imath/ix"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type Point struct {
	X, Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func ParsePoints(line []string) ([]Point, error) {
	var points []Point
	for _, line := range line {
		point, err := ParsePoint(line)
		if err != nil {
			return nil, err
		}
		points = append(points, point)
	}
	return points, nil
}

func ParsePoint(line string) (Point, error) {
	pointRegex := regexp.MustCompile("^([0-9]+), ([0-9]+)$")
	if matches := pointRegex.FindStringSubmatch(line); matches != nil {
		x, err := strconv.Atoi(matches[1])
		if err != nil {
			return Point{}, fmt.Errorf("failed to parse X coordinate of point: %w", err)
		}
		y, err := strconv.Atoi(matches[2])
		if err != nil {
			return Point{}, fmt.Errorf("failed to parse Y coordinate of point: %w", err)
		}
		return Point{
			X: x,
			Y: y,
		}, nil
	}
	return Point{}, fmt.Errorf("invalid point: %s", line)
}

func (p Point) manhattanDistance(otherPoint Point) int {
	return ix.Abs(p.X-otherPoint.X) + ix.Abs(p.Y-otherPoint.Y)
}

func (p Point) findNearestPointOnGrid(grid Grid) *Point {
	distanceCountMap := make(map[int]int)
	var nearestPoint *Point
	var nearestDistance int
	for _, point := range grid.Points {
		if point == p {
			return &point
		}
		point := point
		distance := p.manhattanDistance(point)
		distanceCountMap[distance] += 1
		if nearestPoint == nil || distance < nearestDistance {
			nearestPoint = &point
			nearestDistance = distance
		}
	}
	if distanceCountMap[nearestDistance] > 1 {
		return nil
	}
	return nearestPoint
}

type Grid struct {
	Points []Point
	width  int
	height int
}

func NewGrid(points []Point) *Grid {
	var width, height int
	for _, point := range points {
		width = ix.Max(point.X, width)
		height = ix.Max(point.Y, height)
	}
	return &Grid{
		Points: points,
		width:  width + 1,
		height: height + 1,
	}
}

func (g Grid) CalculateAreas() map[Point]int {
	areas := make(map[Point]int)
	for x := 0; x <= g.width; x++ {
		for y := 0; y <= g.height; y++ {
			point := NewPoint(x, y)
			nearestPoint := point.findNearestPointOnGrid(g)
			if nearestPoint == nil {
				continue
			}
			if areas[*nearestPoint] == -1 {
				continue
			}
			if g.isPointAtBorder(point) {
				areas[*nearestPoint] = -1
			} else {
				areas[*nearestPoint] += 1
			}
		}
	}
	return areas
}

func (g Grid) CalculateSafeAreaSize(totalDistanceThreshold int) (safeAreaSize int) {
	for x := 0; x <= g.width; x++ {
		for y := 0; y <= g.height; y++ {
			currentPoint := NewPoint(x, y)
			var totalDistance int
			for _, point := range g.Points {
				totalDistance += point.manhattanDistance(currentPoint)
			}
			if totalDistance < totalDistanceThreshold {
				safeAreaSize += 1
			}
		}
	}
	return
}

func (g Grid) FindLargestArea() int {
	areas := g.CalculateAreas()
	return slices.Max(maps.Values(areas))
}

func (g Grid) isPointAtBorder(point Point) bool {
	return point.X == 0 || point.X == g.width || point.Y == 0 || point.Y == g.height
}
