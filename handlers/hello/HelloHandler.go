package hello

import "github.com/gin-gonic/gin"

func Hello(context *gin.Context) {
	context.JSON(200, gin.H{
		"code": 20,
		"msg":  "hello,world",
	})
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	first := arr[0]
	low := make([]int, 0, 0)
	high := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	mid = append(mid, first)
	for i := 0; i < len(arr); i++ {
		if arr[i] < first {
			low = append(low, arr[i])
		} else if arr[i] > first {
			high = append(high, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, high = QuickSort(low), QuickSort(high)
	myarr := append(append(low, mid...), high...)
	return myarr
}

//func HeapSortMax(arr []int, length int) []int {
//	if length <= 1 {
//		return arr
//	}
//	depth := length/2 - 1
//	for i := depth; i >= 0; i-- {
//		topmax := i
//		leftChild := 2*i + 1
//		rightChild := 2*i + 2
//
//	}
//}
