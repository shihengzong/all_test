/*
 * @Description: 测试文件
 * @Command: go run main .go
 * @Author: zongsh
 * @Date: 2019-08-28 16:29:24
 * @LastEditTime: 2019-09-02 09:56:18
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

// 测试树
func TreeQueueTest() {
	// 先生成一个树
	tree := Elem{
		Name: "root节点",
		Age:  7,
	}

	// 左子树
	tree.LeftChild = &Elem{
		Name: "左1节点",
		Age:  4,
	}

	tree.LeftChild.LeftChild = &Elem{
		Name: "左1节点 - 左2节点",
		Age:  2,
	}

	tree.LeftChild.RightChild = &Elem{
		Name: "左1节点 - 右2节点",
		Age:  6,
	}

	tree.LeftChild.LeftChild.LeftChild = &Elem{
		Name: "左1节点 - 左2节点 - 左3节点",
		Age:  1,
	}

	tree.LeftChild.LeftChild.RightChild = &Elem{
		Name: "左1节点 - 左2节点 - 右3节点",
		Age:  3,
	}

	tree.LeftChild.RightChild.LeftChild = &Elem{
		Name: "左1节点 - 右2节点 - 左3节点",
		Age:  5,
	}

	// 右子树
	tree.RightChild = &Elem{
		Name: "右1节点",
		Age:  9,
	}
	tree.RightChild.LeftChild = &Elem{
		Name: "右1节点 - 左2节点",
		Age:  8,
	}

	tree.RightChild.RightChild = &Elem{
		Name: "右1节点 - 右2节点",
		Age:  10,
	}

	list := InitTreeStack()
	// 进行前序遍历 (7,4,2,1,3,6,5,9,8,10)
	// list.PreTraverse(tree)
	// 中序
	list.InTravesal(&tree)
	// 后序
	// list.PostTravesal(&tree)

	// 进行层次遍历(7,4,9,2,6,8,10,1,3,5)
	// list := InitTreeQueue()
	// list.LevelTraverse(tree)

}

// 测试排序二叉树
func TreeSortTest() {
	var es []Elem
	es = append(es,
		Elem{
			Age: 5,
		},
		Elem{
			Age: 1,
		},
		Elem{
			Age: 4,
		},
		Elem{
			Age: 2,
		},
		Elem{
			Age: 3,
		},
		Elem{
			Age: 6,
		},
		Elem{
			Age: 7,
		},
		Elem{
			Age: 8,
		},
		Elem{
			Age: 9,
		},
	)
	count := len(es)
	e := Elem{}
	// t := &Elem{Name: "树信息", Age: 0}
	var t *Elem
	for i := 0; i < count; i++ {
		e = es[i]
		t.TreeSort(e)
	}

	// 采用中序遍历
	list := InitTreeStack()
	list.InTravesal(t)
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
	// QueueTest()

	// 测试树
	// TreeQueueTest()

	// 测试排序二叉树
	TreeSortTest()
}
