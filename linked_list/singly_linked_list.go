/*
 * @Description: 数据结构 -- 单链表(Singly Linked List)
 * @Command: go test -v singly_linked_list.go
 * @Author: zongsh
 * @Date: 2019-08-27 11:14:37
 * @LastEditTime: 2019-08-29 14:28:31
 * @LastEditors: zongsh
 */
package linked_list

import (
	"errors"
	"fmt"
)

/*
 * @description:表里的元素
 */
type Elem struct {
	Name       string
	Age        int
	LeftChild  *Elem // 左子节点(用作树)
	RightChild *Elem // 右子节点(用作树)
}

/*
 * @description:表元素集合
 * @param   Node       Elem   当前节点信息
 * @param   NextDate   *SinglyLinkedNode   下一个节点的指针
 */
type SinglyLinkedNode struct {
	Node     Elem
	NextNode *SinglyLinkedNode
}

/*
 * @description:初始化一个头节点
 */
func InitNode() *SinglyLinkedNode {
	return &SinglyLinkedNode{
		Node:     Elem{Name: "头节点", Age: 0},
		NextNode: nil,
	}
}

/*
 * @description:在末尾新增一个节点
 */
func (n *SinglyLinkedNode) AddNode(e Elem) {
	p := n
	pw := n
	for nil != p {
		pw = p         // 当前节点
		p = p.NextNode // 下一个节点
	}
	pw.NextNode = &SinglyLinkedNode{Node: e}
}

/*
 * @description:遍历表
 */
func (n *SinglyLinkedNode) Traverse() (int, []Elem) {
	var list []Elem
	count := 0
	p := n
	for nil != p {
		p = p.NextNode
		if nil != p {
			list = append(list, p.Node)
			count++
		}
	}
	return count, list
}

/*
 * @description:在固定位置插入节点，复杂度为o(n)
 */
func (n *SinglyLinkedNode) Insert(index int, e Elem) error {
	p := n
	j := 1
	for nil != p && j < index {
		p = p.NextNode
		j++
	}

	if nil == p || j > index {
		return errors.New(fmt.Sprintf("please check iindex:%d", index))
	}

	insertLinked := &SinglyLinkedNode{Node: e, NextNode: p.NextNode}
	p.NextNode = insertLinked
	return nil
}

/*
 * @description:在固定位置删除节点，复杂度为o(n)
 */
func (n *SinglyLinkedNode) Delete(index int) error {
	p := n
	j := 1
	for nil != p && j < index {
		p = p.NextNode
		j++
	}

	if nil == p || j > index {
		return errors.New(fmt.Sprintf("please check iindex:%d", index))
	}

	p.NextNode = p.NextNode.NextNode
	return nil
}

/*
 * @description:获取固定位置节点，复杂度为o(n)
 */
func (n *SinglyLinkedNode) Get(index int) *Elem {
	p := n
	for i := 0; i < index; i++ {
		p = p.NextNode
		if nil == p {
			return nil
		}
	}
	return &p.Node
}
