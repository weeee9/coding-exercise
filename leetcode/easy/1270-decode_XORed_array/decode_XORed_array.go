package decode

func decode(encoded []int, first int) []int {
    result := make([]int, len(encoded)+1)
    result[0] = first

    for i := 1; i< len(result); i++ {
        result[i] = result[i-1] ^ encoded[i-1]
    }

    return result
}