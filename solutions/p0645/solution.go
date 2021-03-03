package solution 

func findErrorNums(nums []int) []int {
    m := make(map[int]int, len(nums))
    dup, miss := -1, 1
    for _, v := range nums{
        m[v]++
    }
    for i := 1; i <= len(nums); i++{
        if _, ok := m[i]; ok {
            if m[i] == 2{
                dup = i
            }
        } else{
            miss = i
        }
    }
    return []int{dup, miss}
}
