package main

// https://leetcode.cn/problems/number-of-valid-move-combinations-on-chessboard/?envType=daily-question&envId=2024-12-04
const SIZE = 8

type Move struct {
	x, y, dx, dy, step int
}

type Dir struct {
	dx, dy int
}

func countCombinations(pieces []string, positions [][]int) int {
	N := len(pieces)
	rookDirs := []Dir{
		Dir{-1, 0}, Dir{0, -1}, Dir{1, 0}, Dir{0, 1},
	}
	bishopDirs := []Dir{
		Dir{-1, -1}, Dir{1, -1}, Dir{-1, 1}, Dir{1, 1},
	}
	pieceDirsMap := map[string][]Dir{
		"rook":   rookDirs,
		"bishop": bishopDirs,
		"queen":  append(rookDirs, bishopDirs...),
	}

	ans := 0
	allMoves := [][]Move{}
	for i := 0; i < N; i++ {
		dirs := pieceDirsMap[pieces[i]]
		position := positions[i]
		row, col := position[0]-1, position[1]-1
		allMoves = append(allMoves, genMoves(row, col, dirs))
	}
	path := make([]Move, N)

	var dfs func(int)
	dfs = func(i int) {
		if i == N {
			ans++
			return
		}
		for _, move1 := range allMoves[i] {
			valid := true
			for _, move2 := range path[:i] {
				if !isValid(&move1, &move2) {
					valid = false
					break
				}
			}
			if valid {
				path[i] = move1
				dfs(i + 1)
			}
		}
	}

	dfs(0)
	return ans
}

func genMoves(x0 int, y0 int, dirs []Dir) []Move {
	moves := []Move{Move{x0, y0, 0, 0, 0}}
	for _, dir := range dirs {
		x, y := x0+dir.dx, y0+dir.dy
		step := 1
		for 0 <= x && x < SIZE && 0 <= y && y < SIZE {
			moves = append(moves, Move{x0, y0, dir.dx, dir.dy, step})
			x += dir.dx
			y += dir.dy
			step++
		}
	}
	return moves
}

func isValid(move1 *Move, move2 *Move) bool {
	x1, y1 := move1.x, move1.y
	x2, y2 := move2.x, move2.y
	for i := 0; i < max(move1.step, move2.step); i++ {
		if i < move1.step {
			x1 += move1.dx
			y1 += move1.dy
		}
		if i < move2.step {
			x2 += move2.dx
			y2 += move2.dy
		}
		if x1 == x2 && y1 == y2 {
			return false
		}
	}
	return true
}

func main() {

}
