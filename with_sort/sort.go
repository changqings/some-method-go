package with_sort

import "sort"

// 遍历ss,从index=1开始，取出值tmpValue与beforeIndex的值进行比较
// 如果tmpValue比beforeIndex的值小，则把beforeIndex+1的值替换成ss[beforeIndex]
// 然后beforeIndex向前移动了一位，并继续比较ss[beforeIndex]与tmpValue的大小
// 如果ss[beforeIndex]比tmpValue大，则把beforeIndex+1的值替换成ss[beforeIndex]
// #"此时beforceIndex已经是左移过的值"
// 直到发现beforeIndex的值不比tmpValue大时,把beforeIndex+1的值替换成tmpValue
func ByInsert(ss []int) []int {

	for index := 1; index < len(ss); index++ {

		tmpValue := ss[index]
		beforeIndex := index - 1

		for beforeIndex >= 0 && ss[beforeIndex] > tmpValue {
			ss[beforeIndex+1] = ss[beforeIndex]
			beforeIndex--
		}

		ss[beforeIndex+1] = tmpValue

	}
	return ss
}

func ByGMSort(ss []int) []int {
	sort.Ints(ss)
	return ss
}
