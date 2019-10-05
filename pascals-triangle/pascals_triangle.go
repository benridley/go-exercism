package pascal

// Triangle computes pascal's triangle up to a given row
func Triangle(n int) (tri [][]int) {
	for i := 0; i < n; i++ {
		row := make([]int, i+1)
		for j := 0; j < len(row); j++ {
			if j == 0 || j == len(row)-1 {
				row[j] = 1
			} else {
				row[j] = tri[i-1][j] + tri[i-1][j-1]
			}
		}
		tri = append(tri, row)
	}
	return tri
}
