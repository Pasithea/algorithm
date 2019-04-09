// https://leetcode.com/problems/rotate-image/


func rotate(matrix [][]int)  {
	length := len(matrix)
	i, j := 0, length-1
	for i < j {
		matrix[i], matrix[j] = matrix[j], matrix[i]
		i++
		j--
	}
	fmt.Println(matrix)
	for i := 0; i < length; i++ {
		for j := i+1; j < length; j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}
}
