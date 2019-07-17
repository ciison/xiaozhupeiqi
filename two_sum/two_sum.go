package main

import "fmt"

//解决思路
// 1. 先将切片排序
// 2. 利用夹逼的方法, 已经排序好的切片的前后两个输的和 和定值比较
// 3. 如果和比定值小, 则表明左边的数小, 将左下标＋ 1 继续 步骤2
// 4. 如果和比定值大, 则表明右边的数字大, 将右下标 -1 继续步骤 2
// 5. 如果坐下标和右下标 相等, 则表明没有找到这两个数字

// 快排的实现

func quickSort(data []int, left, right int) {
	// 如果左下标大于右下标, 表明递归结束
	if left >= right {
		return
	}

	var (
		i, j   = left, right
		haviot = data[left]
	)

	for i < j {
		for j > i && haviot <= data[j] {
			j--
		}

		for i < j && haviot > data[i] {
			i++
		}

		if i != j {
			data[i], data[j] = data[j], data[i]
		}

	}

	// 交换基准数字
	data[i], data[left] = data[left], data[i]

	// 左边
	quickSort(data, 0, i)
	// 右边
	quickSort(data, i+1, right)
}

// 查找两个数的和为定值的两个数
func findMix(data []int, mix int) (left, right int) {

	for i, j := 0, len(data)-1; i < j; {
		sum := data[i] + data[j]

		if sum == mix {
			return data[i], data[j]
		}

		// 如果连个数的和比定值小, 表明左边的值太小, 坐下标增大
		// 否则右下标减小
		if sum < mix {
			i++
		} else {
			j--
		}
	}

	// 表示没有找到
	return 0, 0
}

func main() {

	data := []int{9, 6, 5, 3, 7, 8, 1, 14, 11, 99, 101}
	quickSort(data, 0, len(data)-1)
	fmt.Println(data)

	fmt.Println(findMix(data, 15))
	fmt.Println(findMix(data, 100))
}
