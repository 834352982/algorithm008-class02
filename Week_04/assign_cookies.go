func findContentChildren(g []int, s []int) int {
    if len(g) == 0 || len(s) == 0 {
        return 0
    }

    var count int

    sort.Ints(g)
    sort.Ints(s)

    for i, j := 0, 0; i < len(g) && j < len(s); {
        if g[i] <= s[j] {
            count ++
            i ++
        }
        // 看下一个饼干是否满足胃口
        j ++
    }

    return count
}