package main

func rotate(matrix [][]int)  {
	length:=len(matrix)
	for i := 0; i < length/2; i++ {
		tmp:=matrix[i]
		matrix[i] = matrix[length - i -1]
		matrix[length -1 -i] = tmp
	}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			tmp:=matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
}

func main() {
	matrix := [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}
	rotate(matrix)
}
