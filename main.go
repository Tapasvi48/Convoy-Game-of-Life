package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// window to show the game
type Window struct {
	width  int
	height int
}

// type state
type State int

const (
	ALIVE State = iota
	DEAD
)

func stateMap(st State) string {
	switch st {
	//
	case ALIVE:
		return "█"
	case DEAD:
		return " "
	default:
		return " "
	}
	//def
}

// game struct
type Game struct {
	rows  int
	cols  int
	state [][]State
}

func NewGame(rows, cols int) *Game {
	arr := make([][]State, rows)
	for i := 0; i < rows; i++ {
		arr[i] = make([]State, cols)
	}
	return &Game{
		rows:  rows,
		cols:  cols,
		state: arr,
	}
}

func (g *Game) RenderGrid(grid *tview.Table) {
	for x := 0; x < g.rows; x++ {
		for y := 0; y < g.cols; y++ {
			cellText := " " // Dead cell
			color := tcell.ColorGray

			if g.state[x][y] == ALIVE {
				cellText = "█" // Alive cell
				color = tcell.ColorGreen
			}

			grid.SetCell(x, y, tview.NewTableCell(cellText).
				SetAlign(tview.AlignCenter).
				SetTextColor(color))
		}
	}
}

func (g *Game) LoadRLE(rle string) {
	lines := strings.Split(rle, "\n")
	var rleData []string
	var width, height int

	for _, line := range lines {
		if strings.HasPrefix(line, "x =") {
			parts := strings.Split(line, ",")
			width, _ = strconv.Atoi(strings.Fields(parts[0])[2])
			height, _ = strconv.Atoi(strings.Fields(parts[1])[2])
		} else if !strings.HasPrefix(line, "#") && strings.TrimSpace(line) != "" {
			rleData = append(rleData, strings.TrimSpace(line))
		}
	}

	data := strings.Join(rleData, "")
	x, y := 0, 0
	count := 0

	for i := 0; i < len(data); i++ {
		char := string(data[i])

		if char >= "0" && char <= "9" {
			num, _ := strconv.Atoi(char)
			count = count*10 + num
		} else {
			if count == 0 {
				count = 1
			}

			switch char {
			case "b":
				for j := 0; j < count; j++ {
					if x < width && y < height {
						g.state[y][x] = DEAD
						x++
					}
				}
			case "o":
				for j := 0; j < count; j++ {
					if x < width && y < height {
						g.state[y][x] = ALIVE
						x++
					}
				}
			case "$":
				y++
				x = 0
			case "!":
				return
			}

			count = 0
		}
	}
}

func (g *Game) initGame(pattern string) {

	for x := 0; x < g.rows; x++ {
		for y := 0; y < g.cols; y++ {
			g.state[x][y] = DEAD
		}
	}

	switch pattern {
	case "random":
		for x := 0; x < g.rows; x++ {
			for y := 0; y < g.cols; y++ {
				randomValue := rand.Intn(2)
				g.state[x][y] = State(randomValue)
			}
		}
	case "glider":
		rle := `#N Gosper glider gun
#O Bill Gosper
#C A true period 30 glider gun.
#C The first known gun and the first known finite pattern with unbounded growth.
#C www.conwaylife.com/wiki/index.php?title=Gosper_glider_gun
x = 36, y = 9, rule = B3/S23
24bo11b$22bobo11b$12b2o6b2o12b2o$11bo3bo4b2o12b2o$2o8bo5bo3b2o14b$2o8b
o3bob2o4bobo11b$10bo5bo7bo11b$11bo3bo20b$12b2o!`

		g.LoadRLE(rle)

	}
}

func (g *Game) getNeighbourCount(x, y int) int {
	cnt := 0
	for delr := -1; delr <= 1; delr++ {
		for delc := -1; delc <= 1; delc++ {
			if delr == 0 && delc == 0 {
				continue
			}
			nx, ny := x+delr, y+delc
			if nx >= 0 && ny >= 0 && nx < g.rows && ny < g.cols && g.state[nx][ny] == ALIVE {
				cnt++
			}
		}
	}
	return cnt
}

func repeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

func (g *Game) Update() {
	//rules
	//Any live cell with two or three live neighbours lives on to the next generation.
	//Any live cell with more than three live neighbours dies, as if by overpopulation.
	//Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
	//<=3 >2 ->live   >3 ->dies //
	cpyState := make([][]State, g.rows)
	for i := range cpyState {
		cpyState[i] = make([]State, g.cols)
		copy(cpyState[i], g.state[i])
	}
	for x := 0; x < g.rows; x++ {
		for y := 0; y < g.cols; y++ {
			cnt := g.getNeighbourCount(x, y)
			if g.state[x][y] == ALIVE && (cnt == 2 || cnt == 3) {
				cpyState[x][y] = ALIVE
			} else if g.state[x][y] == DEAD && cnt == 3 {
				cpyState[x][y] = ALIVE
			} else {
				cpyState[x][y] = DEAD
			}
		}
	}
	g.state = cpyState

}
func main() {
	// rand.Seed(time.Now().UnixMicro())
	// game := NewGame(31, 45)
	// game.initGame("glider")

	// ticker := time.NewTicker(time.Millisecond * 100)
	// for {
	// 	<-ticker.C
	// 	game.Draw()
	// }

	rand.Seed(time.Now().UnixNano())
	game := NewGame(20, 60)
	game.initGame("glider")
	app := tview.NewApplication()
	grid := tview.NewTable().SetBorders(true).SetBordersColor(tcell.ColorGray)
	title := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("Conway's Game of Life").
		SetTextColor(tcell.ColorYellow)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(title, 1, 1, false).
		AddItem(grid, 0, 5, true)

	// since this work in routines only cause event handle seperate from ui
	go func() {
		ticker := time.NewTicker(time.Millisecond * 20)
		defer ticker.Stop()
		for range ticker.C {
			game.Update()
			app.QueueUpdateDraw(func() {
				// to avoid race around condition as the outside main thread is also using or setting it
				game.RenderGrid(grid)
			})
		}
	}()
	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}

}
