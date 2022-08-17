package main

func main() {

}

func trap(height []int) int {
	l, r := 0, len(height)-1
	leftMax, rightMax := height[l], height[r]
	trap := 0

	for l < r {
		if leftMax < rightMax {
			l += 1
			leftMax = max1(leftMax, height[l])
			trap += leftMax - height[l]
		} else {
			r -= 1
			rightMax = max(rightMax, height[r])
			trap += rightMax - height[r]
		}
	}

	return trap
}

func max1(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

