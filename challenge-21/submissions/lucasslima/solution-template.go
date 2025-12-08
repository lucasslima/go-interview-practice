package main

import (
	"fmt"
)

func main() {
	// Example sorted array for testing
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	// Test binary search
	target := 7
	index := BinarySearch(arr, target)
	fmt.Printf("BinarySearch: %d found at index %d\n", target, index)

	// Test recursive binary search
	recursiveIndex := BinarySearchRecursive(arr, target, 0, len(arr)-1)
	fmt.Printf("BinarySearchRecursive: %d found at index %d\n", target, recursiveIndex)

	// Test find insert position
	insertTarget := 8
	insertPos := FindInsertPosition(arr, insertTarget)
	fmt.Printf("FindInsertPosition: %d should be inserted at index %d\n", insertTarget, insertPos)
}

// BinarySearch performs a standard binary search to find the target in the sorted array.
// Returns the index of the target if found, or -1 if not found.
func BinarySearch(arr []int, target int) int {
	// TODO: Implement this function
	left, right := 0, len(arr)-1
	for left <= right {
		current := left + (right-left)/2
		if arr[current] == target {
			return current
		}
		if arr[current] < target {
			left = current + 1
		} else { // arr[current] > target
			right = current - 1
		}
	}
	return -1
}

// BinarySearchRecursive performs binary search using recursion.
// Returns the index of the target if found, or -1 if not found.
func BinarySearchRecursive(arr []int, target int, left int, right int) int {
	if right < left {
		return -1
	}
	current := left + (right-left)/2
	if arr[current] < target {
		return BinarySearchRecursive(arr, target, current+1, right)
	} else if arr[current] > target {
		return BinarySearchRecursive(arr, target, left, current-1)
	} else {
		return current
	}
}

// FindInsertPosition returns the index where the target should be inserted
// to maintain the sorted order of the array.
func FindInsertPosition(arr []int, target int) int {
	// TODO: Implement this function
	if len(arr) == 0 {
		return 0
	}
	left, right := 0, len(arr)-1
	var current int
	for left <= right {
		current = left + (right-left)/2
		if target <= arr[left] {
			return left
		} else if target > arr[right] {
			return right + 1
		} else {
			if target > arr[current] {
				left = current + 1
			} else if target < arr[current] {
				right = current - 1
			} else {
				break
			}
		}
	}
	return current
}
