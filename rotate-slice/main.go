//  Go code that rotates given slice by given number of steps to the right or to the left.

// Input: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] and k = 3 (rotate by 3 steps to the right)
// Output: [7, 8, 9, 10, 1, 2, 3, 4, 5, 6]

// Input: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] and k = -3 (rotate by 3 steps to the left)
// Output: [4, 5, 6, 7, 8, 9, 10, 1, 2, 3]

package main

import "fmt"

func main() {

	// Create a slice of integers to store the given numbers
	var arr []int

	// Get user input for how many numbers to enter
	fmt.Println("Enter the size: ")
	var n int
	fmt.Scanf("%d", &n)

	// Get user input for the numbers
	fmt.Println("Enter the numbers: ")
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanf("%d", &num)
		arr = append(arr, num)
	}

	// Get user input for the number of rotations
	fmt.Println("Enter the number of rotations: ")
	var k int
	fmt.Scanf("%d", &k)

	// Get the rotation way
	fmt.Println("Enter the rotation way: ")
	var way string
	fmt.Scanf("%s", &way)

	// call the rotate function to rotate the slice of integers
	rotatedArr := rotate(arr, k, way)

	// Print the rotated slice of integers
	fmt.Println(rotatedArr)
}

func rotate(arr []int, stepNum int, direction string) []int {

	// Rotate the slice of integers by k positions if direction is right
	if direction == "right" {
		for i := 0; i < stepNum; i++ {
			arr = append(arr[len(arr)-1:], arr[:len(arr)-1]...)
		}
	}

	// Rotate the slice of integers by k positions if direction is left
	if direction == "left" {
		for i := 0; i < stepNum; i++ {
			arr = append(arr[len(arr)-stepNum:], arr[:len(arr)-stepNum]...)
		}
	}

	return arr
}
