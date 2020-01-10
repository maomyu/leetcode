package main

//
func GetNum(data []int, d int) (sum int) {
	for i := 0; i < len(data)-2; i++ {
		for j := i + 1; j < len(data)-1; j++ {
			for z := j + 1; z < len(data); z++ {
				if data[z]-data[i] <= d {
					sum++
				}
			}
		}
	}
	return sum
}

func main() {
	// var n int
	// var d int
	// fmt.Scanf("%d %d\n", &n, &d)
	// data := make([]int, n)
	// for i := 0; i < len(data); i++ {
	// 	fmt.Scanf("%d", &data[i])
	// }
	// fmt.Println(GetNum(data, d))

}
