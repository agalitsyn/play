// Level: medium
//
// Source: https://leetcode.com/problems/number-of-islands/
//
package main

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0x31 {
				count++
				find(grid, i, j)
			}
		}
	}
	return count
}

func find(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == 0x30 {
		return
	}

	grid[i][j] = 0x30

	find(grid, i+1, j)
	find(grid, i-1, j)
	find(grid, i, j+1)
	find(grid, i, j-1)
}
