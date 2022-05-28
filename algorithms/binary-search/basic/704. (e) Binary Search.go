/* https://leetcode.com/problems/binary-search/
704. (Easy) Binary Search

Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)>>1

		if nums[mid] == target {
			return mid

		} else if nums[mid] < target {
			left = mid + 1

		} else {
			right = mid - 1
		}
	}

	return -1
}