/*
 * Author: pan <panxuzhengxuxin@gmail.com>
 *
 * go语言实现的快速排序算法
 *
 */
package main

import (
	"fmt"
)

//实现排序
func Qsort(source []int, left, right int) {
	keyElement := source[left]
	i := left
	j := right
	for {
		for j = right; j > i; j = j - 1 {
			if source[j] < keyElement {
				source[i] = source[j]
				break
			}
		}
		for i = left; i < j; i = i + 1 {
			if source[i] > keyElement {
				source[j] = source[i]
				break
			}
		}

		if i >= j {
			break
		}
	}
	source[i] = keyElement
	//对关键字左边排序
	if left < i {
		Qsort(source, left, i-1)
	}
	//对关键字右边排序
	if right > i {
		Qsort(source, i+1, right)
	}

}

func main() {
	source := []int{5, 3, 7, 2, 6, 1, 4, 9, 8}
	Qsort(source, 0, len(source)-1)
	for i := 0; i < len(source); i++ {
		fmt.Println(source[i])
	}
}
