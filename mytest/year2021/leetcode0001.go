package main

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	m, n, lands, dx, dy := len(grid), len(grid[0]), 0, [4]int{-1, 1, 0, 0}, [4]int{0, 0, -1, 1}
	var dfsFill func(i, j int) int
	dfsFill = func(i, j int) int {
		if grid[i][j] == '0' {
			return 0
		}
		grid[i][j] = '0'
		for k := 0; k < len(dx); k++ {
			x, y := i+dx[k], j+dy[k]
			if !(x < 0 || y < 0 || x >= m || y >= n) {
				dfsFill(x, y)
			}
		}
		return 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				lands += dfsFill(i, j)
			}
		}
	}
	return lands
}
func main() {

}
