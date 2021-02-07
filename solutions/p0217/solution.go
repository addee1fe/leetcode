package solution

func containsDuplicate(nums []int) bool {
    m := make(map[int]bool,0)
    for _, v := range nums{
        if _, ok := m[v]; ok{
            return true
        }
        m[v] = true
    }
    return false
}
