/* https://leetcode.com/problems/search-insert-position/
35. (Easy) Search Insert Position

Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.
*/

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)>>1

		if nums[mid] == target {
			return mid

		} else if nums[mid] < target {
			left = mid + 1

		} else {
			right = mid
		}
	}
	return right
}