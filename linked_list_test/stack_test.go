/*
 * @Description: 数据结构 -- 栈 (先进后出 , 用单链表实现)
 * @Command: please edit
 * @Author: zongsh
 * @Date: 2019-08-27 17:38:51
 * @LastEditTime: 2019-08-27 17:57:39
 * @LastEditors: zongsh
 */

package linked_list_test

// 栈信息
type Stack struct {
	List *SinglyLinkedNode
}

// Init 初始化栈
func InitStack() *Stack {
	return &Stack{List: InitNode()}
}

// Push 压入栈
func (s *Stack) Push(e Elem) error {
	return s.List.Insert(0, e)
}

// Pop 压出栈
func (s *Stack) Pop() interface{} {
	node := s.List.Get(0)
	if node != nil {
		s.List.Delete(0)
		return node
	}
	return nil
}

// Peek 查看栈顶元素
func (s *Stack) Peek() interface{} {
	node := s.List.Get(0)
	if node != nil {
		return node
	}
	return nil
}

// Size 获取栈的长度
func (s *Stack) Size() int {
	count, _ := s.List.Traverse()
	return count
}
