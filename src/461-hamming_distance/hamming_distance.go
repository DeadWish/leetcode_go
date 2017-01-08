func hammingDistance(x int, y int) int {
    i, a := 0, x^y
    for a != 0 {
        a &= a - 1
        i++
    }
    return i
}