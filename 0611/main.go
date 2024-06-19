package main

func countBattleships(board [][]byte) int {
	n := len(board)
	m := len(board[0])

	ans := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'X' {
				ans++
				board[i][j] = '.'
				for ii := i+1; ii<n; ii++ {
					if board[ii][j] == 'X' {
						board[ii][j] = '.'
					} else {
						break
					}
				}
				for jj := j+1; jj<m; jj++ {
					if board[i][jj] == 'X' {
						board[i][jj] = '.'
					} else {
						break
					}
				}
			}
		}
	}

	return ans
}

func main() {

}