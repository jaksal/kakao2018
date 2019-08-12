package kakao2018

import "fmt"

func printpuzzle(nb [][]byte) {
	for r := 0; r < len(nb); r++ {
		fmt.Println(string(nb[r]))
	}
	fmt.Println("==========")
}

func puzzle(m, n int, nb [][]byte) int {
	printpuzzle(nb)
	// for temp board.
	temp := make([][]byte, m)
	for i := 0; i < m; i++ {
		temp[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			if nb[i][j] == ' ' {
				temp[i][j] = ' '
			} else {
				temp[i][j] = 'O'
			}
		}
	}

	bCheck := false
	// check patten..
	for r := 0; r < m-1; r++ {
		row := nb[r]
		next := nb[r+1]
		for c := 0; c < n-1; c++ {
			if row[c] != ' ' &&
				row[c] == row[c+1] &&
				row[c] == next[c] &&
				row[c] == next[c+1] {
				temp[r][c] = 'X'
				temp[r][c+1] = 'X'
				temp[r+1][c] = 'X'
				temp[r+1][c+1] = 'X'

				bCheck = true
			}
		}
	}
	printpuzzle(temp)

	if !bCheck {
		return 0
	}

	// merge..
	result := 0

	for c := 0; c < n; c++ {
		for r := m - 1; r >= 0; r-- {
			if temp[r][c] == 'X' {
				bFind := false
				for t := r - 1; t >= 0; t-- {
					if temp[t][c] == 'O' {
						nb[r][c] = nb[t][c]
						nb[t][c] = ' '
						temp[r][c] = 'X'
						temp[t][c] = ' '

						bFind = true
						break
					}
				}
				if !bFind {
					nb[r][c] = ' '
				}

				//fmt.Println("[DBG]")
				//printpuzzle(temp)
				//printpuzzle(nb)

				result++
			}
		}
	}
	printpuzzle(nb)
	// printpuzzle(temp)

	fmt.Printf("change block=%d\n", result)

	return result
}

func k6th(m, n int, board []string) int {
	fmt.Printf("row=%d col=%d board=%+v\n", m, n, board)

	// for loop.
	nb := make([][]byte, m)
	for i := 0; i < m; i++ {
		nb[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			nb[i][j] = board[i][j]
		}
	}

	result := 0

	r := puzzle(m, n, nb)
	for r != 0 {
		result += r
		nr := puzzle(m, n, nb)

		if r == nr {
			break
		}
		r = nr
	}
	fmt.Printf("result=%d\n", result)

	return result
}
