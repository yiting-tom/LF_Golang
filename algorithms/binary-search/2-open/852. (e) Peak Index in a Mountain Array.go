/* https://leetcode.com/problems/peak-index-in-a-mountain-array/
852. Peak Index in a Mountain Array

Let's call an array arr a mountain if the following properties hold:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... arr[i-1] < arr[i]
arr[i] > arr[i+1] > ... > arr[arr.length - 1]
Given an integer array arr that is guaranteed to be a mountain, return any i such that arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1].

*/

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)

	for left+1 < right {
		mid := left + (right-left)>>1

		if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return mid

		} else if arr[mid] > arr[mid+1] {
			right = mid

		} else {
			left = mid
		}
	}

	if arr[left] > arr[right] {
		return left
	} else {
		return right
	}
}