package kata

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 已知：一个字符串in
// 未知：一个字符串out
// 条件：
// （1）in是以空格分割的整数
// （2）out是以空格分割的最大值和最小值
// （3）所有的整数为Int32类型
// （4）in至少包含1个数字
// （5）out必须包含两个由空格分割的数字，且最大值在前面
func HighAndLow(in string) string {
	numerals := strings.Split(in, " ")
	max := math.MinInt32
	min := math.MaxInt32
	for _, str := range numerals {
		num, _ := strconv.Atoi(str)
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	var appender strings.Builder
	appender.WriteString(fmt.Sprint(max))
	appender.WriteString(" ")
	appender.WriteString(fmt.Sprint(min))
	return appender.String()
}
