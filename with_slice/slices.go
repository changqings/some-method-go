package with_slice

// 使用map的key惟一性进行删除
func RemoveDuplicats(input []int) []int {

	encountered := map[int]bool{}
	result := []int{}

	for _, v := range input {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result

}

// 删除元素，使用append()实现, append()会重新分配内存
// 每删除一个元素，后面元素会进行多次移动
func RemoveValue(ss []int, value int) []int {

	for i, j := range ss {
		if j == value {
			ss = append(ss[:i], ss[i+1:]...)
		}
	}
	return ss
}

// 删除元素，每次只移动一个元素
func RemoveValuePerf(ss []int, value int) []int {

	j := 0
	// 取切片的每一个值，与value比较大小
	// 如果值不等于value,写回原切片，从j=0开始
	// 如果等于value，则不操作
	// 最后，对切片进行截取
	for _, v := range ss {
		if v != value {
			ss[j] = v
			j++
		}
	}

	return ss[:j]
}

// 删除元素，使用指定索引(0 <= index <= length)
// 使用append(), 这里append()已经考虑了边界情况
// 即删除最后一个元素时，ss[:index+1]不会panic
func RemoveIndex(ss []int, index int) []int {

	if index < 0 || index >= len(ss) {
		return ss
	}

	tmp := append(ss[:index], ss[index+1:]...)
	return tmp

}

// 删除元素，使用指定索引(0 <= index <= length)
// 在原切片操作
func RemoveIndexPerf(ss []int, index int) []int {
	if index < 0 || index >= len(ss) {
		return ss
	}

	j := 0
	for i := range ss {
		if i < index {
			j++
		} else if i >= index && i < len(ss)-1 {
			ss[j] = ss[i+1]
			j++
		}
	}

	return ss[:j]

}
