package main

import (
	"fmt"
	"math/rand"
)

type board struct {
	values [][]int
}

func collapseNums(nums []int) []int {
	// base case, only one number left
	if len(nums) == 1 {
		return nums
	} else if nums[0] == nums[1] {
		nums[0] = 0
		a := append(nums[1:2], nums[2:]...)
		return collapseNums(a)
	} else {
		return append(collapseNums(nums[0:1]), collapseNums(nums[1:])...)
	}
}

func collapseNums2(nums [][]int) [][]int {
	// base case, only one number left
	if len(nums) == 1 {
		return nums
	} else if nums[0][0] == nums[1][0] {
		//copy to new var
		a := nums
		// add value to inner slice

		a[0] = append(a[0], a[0][0])
		// cocatonate into new slice, removing the second entry
		//var b [][]int
		b := append(a[0:1], a[2:]...)

		return collapseNums2(b)
	} else {
		return append(collapseNums2(nums[0:1]), collapseNums2(nums[1:])...)
	}
}

func collapseNums3(nums []int, final *[]int) []int {
	// base case, only one number left
	if len(nums) == 1 {
		return nums
	} else if nums[0] == nums[1] {
		*final = append((*final), 0)
		a := append(nums[1:2], nums[2:]...)
		return collapseNums3(a, final)
	} else {
		return append(collapseNums3(nums[0:1], final), collapseNums3(nums[1:], final)...)
	}
}

func collapseNums4(nums [][]int, output *[]int) [][]int {
	// base case, only one number left
	if len(nums) == 1 {
		return nums
	} else if nums[0][0] == nums[1][0] {
		//add a zero to final
		*output = append((*output), 0)

		//copy to new var
		a := nums
		// add value to inner slice
		a[1] = append(a[1], a[0][0])
		// cocatonate into new slice, removing the second entry
		//var b [][]int
		b := append(a[1:2], a[2:]...)

		return collapseNums4(b, output)
	} else {
		return append(collapseNums4(nums[0:1], output), collapseNums4(nums[1:], output)...)
	}
}

func main() {
	// make an empty matrix
	board := make([][]int, 4)
	for i := 0; i < 4; i++ {
		board[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			board[i][j] = rand.Intn(10)
		}
		//fmt.Println(board[i])
	}

	a := make([]int, 0)
	// b := []int{2, 3, 2, 3}
	// a = append(a, collapseNums3(b, &a)...)
	// //fmt.Println(collapseNums3(b, &a))
	// fmt.Println(a)
	// //fmt.Println(collapseNums(b))

	m := make([][]int, 4)
	for i := 0; i < len(m); i++ {
		//m[i][0] = 0
		m[i] = make([]int, 1)
		//m[i][0] = 2
	}

	m[0][0] = 2
	m[1][0] = 2
	m[2][0] = 3
	m[3][0] = 3

	final := make([]int, 0)

	j := collapseNums4(m, &a)
	fmt.Println(j)

	for _, v := range j {
		fmt.Println(v)
		total := 0
		for _, v1 := range v {
			total += v1
		}
		final = append(final, total)

	}

	fmt.Println(final)
	a = append(a, final...)
	fmt.Println(a)

	//matrix := [4][4]int{}
	//fmt.Println(board)

}
