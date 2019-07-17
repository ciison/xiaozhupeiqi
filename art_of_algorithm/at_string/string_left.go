package main

import "fmt"

var (
	data = []byte{'H', 'O', 'W', 'A', 'R', 'E', 'Y', 'O', 'U'}
)

//
func reverseString(data []byte, n int) {

	m := n % len(data)

	// 这里开始做转换
	reverseStringN(data, 0, m)
	reverseStringN(data, m+1, len(data)-1)
	reverseStringN(data, 0, len(data)-1)

}

// 这里是调换 字符的位置
func reverseStringN(data []byte, left, right int) {

	for left < right {
		data[left], data[right] = data[right], data[left]
		left++
		right--
	}
}
func main() {

	fmt.Println(string(data))
	// 这里表示字符做移动三位
	reverseString(data, 2)
	fmt.Println(string(data))
}
