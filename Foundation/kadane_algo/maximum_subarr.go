package kadane_algo

// implementation of kadane's algorithm
func maxSubArray(arr []int) int {
	sum, maxi := 0, arr[0]

	for i := 0; i < len(arr); i++ {
		sum = max(sum+arr[i], arr[i])
		maxi = max(maxi, sum)
	}
	return maxi
}
