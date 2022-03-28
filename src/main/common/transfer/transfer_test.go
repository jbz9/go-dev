package transfer

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	var a []int
	var b = []int{1, 1, 2}
	a = append(a, 1)       // 追加1个元素
	a = append(a, 1, 2, 3) // 追加多个元素, 手写解包方式
	a = append(a, b...)    // 追加一个切片, 切片需要解包
	defer println(1)
	defer println(2)
}

// 值传递 指针传递 引用传递
func TestB(t *testing.T) {
	var a = 1
	fmt.Printf("%#v\n", &a)
	updateDataA(a)
	fmt.Printf("%#v\n%d", &a, a)

}

//值传递
func updateDataA(b int) int {
	//实参 a在传递给函数 updateData 的形参 b 后，在 函数 的内部
	//b 会被当作局部变量在栈上分配空间，并且完全拷贝 a 的值
	//a、b拥有完全不同的内存地址， 说明他们虽然值相同（b拷贝的a，值肯定一样），
	//但是分别在内存中不同的地方，也因此在函数 vFoo 内部如果改变 b 的值，a 是不会受到影响的。
	fmt.Printf("%#v\n", &b)
	return b + 1
}

//指针传递
func updateDataB(b int) int {
	//实参 a在传递给函数 updateData 的形参 b 后，在 函数 的内部
	//b 会被当作局部变量在栈上分配空间，并且完全拷贝 a 的值
	//a、b拥有完全不同的内存地址， 说明他们虽然值相同（b拷贝的a，值肯定一样），
	//但是分别在内存中不同的地方，也因此在函数 vFoo 内部如果改变 b 的值，a 是不会受到影响的。
	fmt.Printf("%#v\n", &b)
	return b + 1
}
