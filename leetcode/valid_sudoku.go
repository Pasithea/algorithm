// https://leetcode.com/problems/valid-sudoku/


func isValidSudoku(board [][]byte) bool {
	length := len(board)
	for r := 0; r < length; r++ {
		if hasRepeat(board[r]) {
			return false
		}
	}
	for c := 0; c < length; c++ {
		s := make([]byte, length)
		for r := 0; r < length; r++ {
			s[r] = board[r][c]
		}
		if hasRepeat(s) {
			return false
		}
		s = make([]byte, length)
	}
	sep := length / 3
	for i := 0; i < sep; i++ {
		for j := 0; j < sep; j++ {
			s := []byte{}
			for ii := i*sep; ii < (i+1)*sep; ii++ {
				for jj := j*sep; jj < (j+1)*sep; jj++ {
					s = append(s, board[ii][jj])
				}
			}
			if hasRepeat(s) {
				return false
			}
		}
	}
	return true
}

func hasRepeat(s []byte) bool {
	x := make(map[byte]bool)
	for _, e := range s {
		if e == '.' {
			continue
		}
		if _, ok := x[e]; ok {
			return true
		}
		x[e] = true
	}
	return false
}
