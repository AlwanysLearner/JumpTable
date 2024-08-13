package Jumptable

import (
	"fmt"
)

// heapify 是一个辅助函数，用于维护堆的性质
func heapify(arr []int, n, i int) {
	largest := i     // 初始化 largest 为根节点
	left := 2*i + 1  // 左子节点
	right := 2*i + 2 // 右子节点

	// 如果左子节点存在，且大于根节点
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// 如果右子节点存在，且大于目前最大的节点
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大值不是根节点
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // 交换
		heapify(arr, n, largest)                    // 递归地对受影响的子树进行 heapify
	}
}

// heapSort 主函数
func heapSort(arr []int) {
	n := len(arr)

	// 构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 一个个从堆顶取出元素，并对剩余元素进行 heapify
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // 交换
		heapify(arr, i, 0)              // 对减小后的堆进行 heapify
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Unsorted array:", arr)
	heapSort(arr)
	fmt.Println("Sorted array:", arr)
}
