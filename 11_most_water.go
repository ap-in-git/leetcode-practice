package main

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxArea := 0
	for l < r {
		area := 0
		if height[l] <= height[r] {
			area = (r - l) * height[l]
			l++
		} else {
			area = (r - l) * height[r]
			r--
		}
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
