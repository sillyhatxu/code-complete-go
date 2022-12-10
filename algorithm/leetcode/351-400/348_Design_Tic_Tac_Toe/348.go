package _48_Design_Tic_Tac_Toe

type TicTacToe struct {
	rows, cols                   []int
	diagonal, antiDiagonal, size int
}

func Constructor(n int) TicTacToe {
	return TicTacToe{
		rows:         make([]int, n, n),
		cols:         make([]int, n, n),
		diagonal:     0,
		antiDiagonal: 0,
		size:         n,
	}
}

func (this *TicTacToe) Move(row int, col int, player int) int {
	sign := -1
	if player == 1 {
		sign = 1
	}
	this.rows[row] += sign
	this.cols[col] += sign
	if row == col {
		//diagonal [0,0];[1,1];[2,2]
		this.diagonal += sign
	}
	if row+col+1 == this.size {
		//anti diagonal [0,2];[1,1];[2,0]
		this.antiDiagonal += sign
	}
	if this.abs(this.rows[row]) == this.size ||
		this.abs(this.cols[col]) == this.size ||
		this.abs(this.diagonal) == this.size ||
		this.abs(this.antiDiagonal) == this.size {
		return player
	}
	return 0
}

func (this *TicTacToe) abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
