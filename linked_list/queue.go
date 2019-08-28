/*
 * @Description: 数据结构 -- 队列 (先进先出 , 用单链表实现)
 * @Command: please edit
 * @Author: zongsh
 * @Date: 2019-08-28 16:16:10
 * @LastEditTime: 2019-08-28 17:19:15
 * @LastEditors: zongsh
 */

package linked_list

// 队列信息
type Queue struct {
	List *SinglyLinkedNode
}

// 初始化队列
func InitQueue() *Queue {
	return &Queue{List: InitNode()}
}

// 加入队列
func (q *Queue) JoinQueue(e Elem) {
	// 排队是后来的排在队尾,AddNode()方法
	q.List.AddNode(e)
}

// 取出队列元素,应取出先来的,即get(1)
func (q *Queue) Get() *Elem {
	node := q.List.Get(1)
	if node != nil {
		q.List.Delete(1)
		return node
	}
	return nil
}

// Peek 查看队头元素
func (q *Queue) Peek() *Elem {
	node := q.List.Get(1)
	if node != nil {
		return node
	}
	return nil
}

// Size 获取栈的长度
func (q *Queue) Size() int {
	count, _ := q.List.Traverse()
	return count
}
