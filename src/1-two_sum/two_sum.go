func twoSum(nums []int, target int) []int {
    table := make(map[int]int)
    for i,v := range nums {
        complement := target - v
        if index, ok := table[complement]; ok {
            return []int{index, i}
        } else {
            table[v] = i
        }
    }
    return []int{-1, -1}
}
