/*
 * @Description: 数据结构 -- 树
 * @Command: please edit
 * @Author: zongsh
 * @Date: 2019-08-29 10:58:34
 * @LastEditTime: 2019-09-02 15:37:51
 * @LastEditors: zongsh
 * @Detail: 1. 叶子节点(没有子节点的节点)   2.树的度(当前节点有几个子节点 度就是几)   3.树的深度(有几层深度就是几)
 			4.排序二叉树:　 1)、若他的左子树不为空，则左子树上所有节点的值均小于它的根节点的值。
						　　2)、若它的右子树不为空，则右子树上所有节点的值均大于它的根节点的值。
						　　3)、它的左、右子树也分别为排序二叉树。
						　　4)、二叉树节点的值不允许重复
*/
package linked_list

import "fmt"

/**
 * @-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:
 * 前序遍历(用栈实现,先遍历root节点,再遍历左子树,最后遍历右子树)
 */
// 树的栈
type TreeStack struct {
	List *Stack
}

// 初始化一个树的栈
func InitTreeStack() *TreeStack {
	return &TreeStack{
		List: InitStack(),
	}
}

// 前序遍历(用栈实现,先遍历root节点,再遍历左子树,最后遍历右子树)
func (t *TreeStack) PreTraverse(node Elem) []Elem {
	var list []Elem
	// 放入root节点
	t.List.Push(node)
	// 获取队里元素个数
	count := 1
	for count > 0 {
		// 取出栈顶元素
		e := t.List.Pop()
		fmt.Println(e)

		list = append(list, *e)

		// 先判断右边子节点是否为空
		if nil != e.RightChild {
			t.List.Push(*e.RightChild)
		}
		// 再判断左边子节点是否为空
		if nil != e.LeftChild {
			t.List.Push(*e.LeftChild)
		}
		count, _ = t.List.List.Traverse()
	}
	fmt.Println(list)
	return list
}

/**
 * @-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:
 * 中序遍历(用栈实现,排序二叉树遍历)
 */
func (s *TreeStack) InTravesal(root *Elem) {
	if root == nil {
		return
	}
	cur := root

	for {
		for cur != nil {
			s.List.Push(*cur)
			cur = cur.LeftChild
		}

		count, _ := s.List.List.Traverse()
		if count < 1 {
			break
		}

		cur = s.List.Pop()
		fmt.Println(cur)
		cur = cur.RightChild
	}
}

/**
 * @-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:
 * 后序遍历(用栈实现)
 */
func (s *TreeStack) PostTravesal(root *Elem) {
	num := 0
	if root == nil {
		return
	}

	out := InitTreeStack()
	s.List.Push(*root)
	count := 1
	for count > 0 {
		count, _ = s.List.List.Traverse()

		if cur := s.List.Pop(); cur != nil {
			// fmt.Println(cur)
			out.List.Push(*cur)

			if cur.LeftChild != nil {
				s.List.Push(*cur.LeftChild)
			}

			if cur.RightChild != nil {
				s.List.Push(*cur.RightChild)
			}
		}
	}

	count2 := 1
	for count2 > 0 {
		cur := out.List.Pop()
		fmt.Println(cur)
		num++
		if num > 100 {
			break
		}
		count2, _ = out.List.List.Traverse()
	}
}

/**
 * @-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:
 * 层次遍历(用队列实现)
 */
// 树的队列
type TreeQueue struct {
	List *Queue
}

// 初始化一个树队列
func InitTreeQueue() *TreeQueue {
	return &TreeQueue{
		List: InitQueue(),
	}
}

// 层次遍历(用栈实现,先遍历root节点,再遍历左子树,最后遍历右子树)
func (t *TreeQueue) LevelTraverse(node Elem) []*Elem {
	var list []*Elem
	// 放入root节点
	t.List.JoinQueue(node)
	// 获取队里元素个数
	count, _ := t.List.List.Traverse()
	for count >= 1 {
		// 取出队头元素
		e := t.List.Get()
		fmt.Println(e)

		list = append(list, e)

		// 先判断左边子节点是否为空
		if nil != e.LeftChild {
			t.List.JoinQueue(*e.LeftChild)
		}
		// 再判断右边子节点是否为空
		if nil != e.RightChild {
			t.List.JoinQueue(*e.RightChild)
		}
	}
	fmt.Println(list)
	return list
}

/**
 * @-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:
 * 排序二叉树
 */

func (t *Elem) TreeSort(e Elem) *Elem {
	defer fmt.Println(t)
	if t == nil {
		t = &Elem{Age: e.Age}
		fmt.Println(t)
		return t
	}
	// 小的放左边
	if t.Age > e.Age {
		t.LeftChild = t.LeftChild.TreeSort(e)
	} else {
		t.RightChild = t.RightChild.TreeSort(e)
	}
	return t
}

/**
 * @-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
 * @description: 数据结构 -- 红黑树
 * @ 平衡二叉树是一个排序二叉树
 * @ 红黑是是一个平衡二叉树
 * @ 平衡二叉树的左子树和右子树深度相差不能大于 1
 */
