/*
 * @Description: 数据结构 -- 栈 (先进后出 , 用单链表实现)
 * @Command: please edit
 * @Author: zongsh
 * @Date: 2019-08-27 17:38:51
 * @LastEditTime: 2019-08-28 17:18:17
 * @LastEditors: zongsh
 */

package linked_list

// 栈信息
type Stack struct {
	List *SinglyLinkedNode
}

// Init 初始化栈
func InitStack() *Stack {
	return &Stack{List: InitNode()}
}

// Push 压入栈,后来的放在最上面
func (s *Stack) Push(e Elem) error {
	return s.List.Insert(1, e)
}

// Pop 压出栈
func (s *Stack) Pop() *Elem {
	node := s.List.Get(1)
	if node != nil {
		s.List.Delete(1)
		return node
	}
	return nil
}

// Peek 查看栈顶元素
func (s *Stack) Peek() *Elem {
	node := s.List.Get(1)
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
