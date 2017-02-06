func findMaxConsecutiveOnes(nums []int) int {
    maxLen := 0
    sum := 0
    for _, value := range nums {
        sum = (sum + value) * value
        if sum > maxLen {
            maxLen = sum
        }
    }
    return maxLen
}