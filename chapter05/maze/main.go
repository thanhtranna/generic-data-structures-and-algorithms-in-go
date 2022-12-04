package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"chapter05/slicestack"
)

// Direction abstraction
type Direction int

const (
	N            int = 0
	NE               = 1
	E                = 2
	SE               = 3
	S                = 4
	SW               = 5
	W                = 6
	NW               = 7
	NotAvailable     = 8
)

func (d Direction) String() string {
	switch d {
	case 0:
		return "north"
	case NE:
		return "north-east"
	case E:
		return "east"
	case SE:
		return "south-east"
	case S:
		return "south"
	case SW:
		return "south-west"
	case W:
		return "west"
	case NW:
		return "north-west"
	case NotAvailable:
		return "not available"
	}
	return "unknown"
}

func (d Direction) PrintDirection() {
	fmt.Println("direction: ", d)
}

// ******************************

// Point abstraction
type Point struct {
	x, y int
}

func (p Point) Equals(other Point) bool {
	return p.x == other.x && p.y == other.y
}

func (p Point) PrintPoint() {
	fmt.Printf("<%d, %d>\n", p.x, p.y)
}

var None = Point{-1, -1}

// *********************************

// Path abstraction
type Path struct {
	point          Point
	move           Direction
	movesAvailable []Direction
}

func NewPath(point Point) Path {
	path := Path{point, Direction(NotAvailable), []Direction{}}
	path.move = NotAvailable
	// Initially all directions available
	path.movesAvailable = []Direction{0, NE, E, SE, S, SW, W, NW}
	return path
}

func (path *Path) RandomMove() Direction {
	// Sets value of move
	indicesAvailable := []int{}
	for index := 0; index < 8; index++ {
		if path.movesAvailable[index] != NotAvailable {
			indicesAvailable = append(indicesAvailable, index)
		}
	}
	count := len(indicesAvailable)
	if count > 0 {
		randomIndex := rand.Intn(count)
		path.move = path.movesAvailable[indicesAvailable[randomIndex]]
		path.movesAvailable[indicesAvailable[randomIndex]] = NotAvailable
		return path.move
	} else {
		return NotAvailable
	}
}

// ********************************

// MAZE abstraction
type Maze struct {
	rows, cols int
	start, end Point
	mazefile   string
	barriers   [][]bool
	current    Path
	moveCount  int
	pathStack  slicestack.Stack[Path]
	gameOver   bool
}

func NewMaze(rows int, cols int, start Point, end Point, mazefile string) (maze Maze) {
	maze.rows = rows
	maze.cols = cols
	maze.start = start
	maze.end = end

	// Initialize maze.barriers
	maze.barriers = make([][]bool, rows)
	for i := range maze.barriers {
		maze.barriers[i] = make([]bool, cols)
	}

	file, err := os.Open(mazefile)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var textlines []string
	for scanner.Scan() {
		textlines = append(textlines, scanner.Text())
	}
	defer file.Close()
	for row := 0; row < rows; row++ {
		line := textlines[row]
		for col := 0; col < cols; col++ {
			if string(line[col]) == "1" {
				maze.barriers[row][col] = true
			} else {
				maze.barriers[row][col] = false
			}
		}
	}
	maze.current = NewPath(start)
	maze.pathStack = slicestack.Stack[Path]{} // generic instance
	maze.pathStack.Push(maze.current)
	maze.barriers[start.x][start.y] = true
	return maze
}

func NewPosition(oldPosition Point, move Direction) Point {
	if move == Direction(N) {
		return Point{oldPosition.x, oldPosition.y - 1}
	} else if move == NE {
		return Point{oldPosition.x + 1, oldPosition.y - 1}
	} else if move == E {
		return Point{oldPosition.x + 1, oldPosition.y}
	} else if move == SE {
		return Point{oldPosition.x + 1, oldPosition.y + 1}
	} else if move == S {
		return Point{oldPosition.x, oldPosition.y + 1}
	} else if move == SW {
		return Point{oldPosition.x - 1, oldPosition.y + 1}
	} else if move == W {
		return Point{oldPosition.x - 1, oldPosition.y}
	} else {
		return Point{oldPosition.x - 1, oldPosition.y - 1}
	}
}

func (m *Maze) StepAhead() (Point, Point) {
	validMove := false
	backTrackPoint := None
	newPos := None
	for {
		if m.gameOver || validMove || m.pathStack.IsEmpty() {
			break
		}
		validMove = false
		m.current = m.pathStack.Pop()
		m.moveCount += 1
		nextMove := m.current.RandomMove()
		for {
			if validMove || nextMove == NotAvailable {
				break
			}
			newPos = NewPosition(m.current.point, m.current.move)
			if !m.barriers[newPos.y][newPos.x] {
				validMove = true
				if newPos.Equals(m.end) {
					for {
						if m.pathStack.IsEmpty() {
							break
						}
						m.pathStack.Pop()
					}
					m.gameOver = true
				}
				m.barriers[newPos.y][newPos.x] = true
				m.pathStack.Push(m.current)
				newPathObject := NewPath(newPos)
				m.pathStack.Push(newPathObject)
			} else {
				nextMove = m.current.RandomMove()
			}
		}
		if !validMove && !m.pathStack.IsEmpty() {
			fmt.Printf("\nBacktrack from %v to %v\n", m.current.point, m.pathStack.Top().point)
			backTrackPoint = m.pathStack.Top().point
		}
	}
	if m.pathStack.IsEmpty() {
		fmt.Println("No solution is possible")
		return None, None
	}
	return newPos, backTrackPoint
}

// *********************************************

func main() {
	// rand.Seed(time.Now().UnixNano())
	// myDirection := Direction(6)
	// myDirection.PrintDirection()
	// myPoint := Point{3, 4}
	// myPoint.PrintPoint()

	// result := myPoint.Equals(Point{3, 4})
	// fmt.Println("result", result)
	// myPath := NewPath(Point{3, 4})
	// randomMove := myPath.RandomMove()
	// fmt.Println(randomMove)
	// fmt.Println(myPath)

	rand.Seed(time.Now().UnixNano())
	start := Point{1, 1}
	end := Point{38, 38}
	maze := NewMaze(40, 40, start, end, "maze.txt")
	newPos, _ := maze.StepAhead()
	time.Sleep(1 * time.Second)
	if newPos != None {
		fmt.Println(newPos)
	}
	for {
		if newPos == None || newPos.Equals(end) {
			break
		}
		newPos, _ = maze.StepAhead()
		time.Sleep(100 * time.Millisecond)
		if newPos != None {
			fmt.Println(newPos)
		}
	}
	if newPos.Equals(end) {
		fmt.Println("SUCCESS!  Reached ", end)
	}
}

/* Output
direction:  west
<3, 4>
result true
south
{{3 4} 4 [0 1 2 3 8 5 6 7]}
*/
