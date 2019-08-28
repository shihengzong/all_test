/*
 * @Description: 测试文件
 * @Command: go run main .go
 * @Author: zongsh
 * @Date: 2019-08-28 16:29:24
 * @LastEditTime: 2019-08-28 17:23:11
 * @LastEditors: zongsh
 */
package linked_list

import (
	"fmt"
)

// 测试单链表
func SinglyLinkedListTest() {
	// 初始化头结点
	list := InitNode()

	// 添加节点
	list.AddNode(Elem{Name: "1张三", Age: 10})
	list.AddNode(Elem{Name: "2李四", Age: 20})
	list.AddNode(Elem{Name: "3王五", Age: 30})
	list.AddNode(Elem{Name: "4赵柳", Age: 40})
	list.AddNode(Elem{Name: "5小明", Age: 50})
	// 遍历列表
	fmt.Println(list.Traverse())

	// 删除第二个节点
	list.Delete(2)
	// 遍历列表
	fmt.Println(list.Traverse())

	// 在第二节点处插入 "小红"
	list.Insert(2, Elem{Name: "6小红", Age: 18})
	// 遍历列表
	fmt.Println(list.Traverse())

	// 获取第二节点
	fmt.Println(list.Get(2))
}

// 测试双链表
func DoubleLinkedListTest() {
	// 初始化头结点
	list := InitDoubleNode()

	// 添加节点
	list.AddNode(DouElem{Name: "1张三", Age: 10})
	list.AddNode(DouElem{Name: "2李四", Age: 20})
	list.AddNode(DouElem{Name: "3王五", Age: 30})
	list.AddNode(DouElem{Name: "4赵柳", Age: 40})
	list.AddNode(DouElem{Name: "5小明", Age: 50})
	// 遍历列表
	fmt.Println(list.Traverse())

	// 删除第二个节点
	list.Delete(2)
	// 遍历列表
	fmt.Println(list.Traverse())

	// 在第二节点处插入 "小红"
	list.Insert(2, DouElem{Name: "6小红", Age: 18})
	// 遍历列表
	fmt.Println(list.Traverse())

	// 获取第二节点
	fmt.Println(list.Get(2))
}

// 测试循环链表
func CycleLinkedListTest() {
	// 初始化头结点
	list := InitCycleNode()

	// 添加节点
	list.AddNode(CycleElem{Name: "1张三", Age: 10})
	list.AddNode(CycleElem{Name: "2李四", Age: 20})
	list.AddNode(CycleElem{Name: "3王五", Age: 30})
	list.AddNode(CycleElem{Name: "4赵柳", Age: 40})
	list.AddNode(CycleElem{Name: "5小明", Age: 50})
	// 遍历列表
	fmt.Println(list.Traverse())

	// 删除第二个节点
	list.Delete(2)
	// 遍历列表
	fmt.Println(list.Traverse())

	// 在第二节点处插入 "小红"
	list.Insert(2, CycleElem{Name: "6小红", Age: 18})
	// 遍历列表
	fmt.Println(list.Traverse())

	// 获取第二节点
	fmt.Println(list.Get(2))
}

// 测试栈
func StackTest() {
	// 初始化头结点
	stack := InitStack()

	// 添加节点
	stack.Push(Elem{Name: "1张三", Age: 10})
	stack.Push(Elem{Name: "2王五", Age: 20})
	stack.Push(Elem{Name: "3赵六", Age: 30})
	stack.Push(Elem{Name: "4小明", Age: 40})
	// 遍历链表
	fmt.Println(stack.List.Traverse())

	// 取出栈顶节点
	fmt.Println("已取出节点:", stack.Pop())
	// 遍历链表
	fmt.Println(stack.List.Traverse())

	// 获取栈顶节点
	fmt.Println(stack.Peek())

	// 获取栈的长度
	fmt.Println(stack.Size())
}

// 测试队列
func QueueTest() {
	// 初始化头结点
	queue := InitQueue()

	// 添加节点
	queue.JoinQueue(Elem{Name: "1张三", Age: 10})
	queue.JoinQueue(Elem{Name: "2王五", Age: 20})
	queue.JoinQueue(Elem{Name: "3赵六", Age: 30})
	queue.JoinQueue(Elem{Name: "4小明", Age: 40})
	// 遍历链表
	fmt.Println(queue.List.Traverse())

	// 取出队头节点
	fmt.Println("已取出节点:", queue.Get())
	// 遍历链表
	fmt.Println(queue.List.Traverse())

	// 获取队头节点
	fmt.Println(queue.Peek())

	// 获取队列的长度
	fmt.Println(queue.Size())
}

func AllTest() {
	// // 测试单链表
	// SinglyLinkedListTest()

	// // 测试双链表
	// DoubleLinkedListTest()

	// // 测试循环链表
	// CycleLinkedListTest()

	// // 测试栈
	// StackTest()

	// 测试队列
	QueueTest()
}
