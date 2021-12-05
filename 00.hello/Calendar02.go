package main

//func main() {
//	var years [200]int
//	// 给year数组赋值 1900:2100
//	for index := range years {
//		for i := 1900; i < 2100; i++ {
//			years[index] = i
//		}
//	}
//	//var(
//	//	bigMonth = [7]int{1,3,5,7,8,10,12}
//	//	smallMonth = [4]int{4,6,9,11}
//	//	febMonth = [1]int{2}
//	//)
//	// 大小月声名
//	var (
//		bigMonth   [12][7][31]int
//		smallMonth [12][7][30]int
//	)
//	// 闰年平年二月声明
//	for _, year := range years {
//		if year%4 == 0 {
//			var leapFebMonth [12][7][29]int
//		} else {
//			var commonFebMonth [12][7][28]int
//		}
//	}
//	for index, month := range bigMonth {
//		switch index {
//		case 1, 3, 5, 7, 8, 10, 12:
//			for week := bigMonth[index] {
//				fmt.Printf("%d ", week)
//			}
//		}
//	}
//
//}
