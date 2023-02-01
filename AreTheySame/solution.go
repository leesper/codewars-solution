package kata

import (
	"reflect"
	"sort"
)

// 输入：两个整数数组array1和array2
// 输出：布尔值
// 条件：
// （1）如果array2中每个元素都是array1中元素的平方，返回true否则false
// （2）array1或array2可能为空数组
// （3）array1或array2可能为空指针，此时应返回false
func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil {
		return false
	}

	a1 := array1[:]
	a2 := array2[:]

	for i, n := range a1 {
		a1[i] = n * n
	}

	sort.Ints(a1)
	sort.Ints(a2)

	return reflect.DeepEqual(a1, a2)
}
