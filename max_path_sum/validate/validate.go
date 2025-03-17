package validate

func IsValidTriangle(triangle [][]int) bool {
	for i := 0; i < len(triangle); i++ {
		if len(triangle[i]) != i+1 {
			return false
		}
	}
	return true
}
