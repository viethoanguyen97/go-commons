package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ans := make([][]uint8, dy)

	for i := range ans {
		ans[i] = make([]uint8, dx)
		for j := range ans[i] {
			p := (i + j) % 3
			switch p {
			case 0:
				ans[i][j] = uint8((i + j) / 2)
			case 1:
				ans[i][j] = uint8(i * j)
			default:
				ans[i][j] = uint8(i ^ j)
			}
			//ans[i][j] = uint8(((j*i)/2) % 255) //Really cool photo!!!
		}
	}

	return ans
}

func main() {
	pic.Show(Pic)
}
