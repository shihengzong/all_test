/*
 * @Description: 数据结构 -- 双向链表(Double Linked List)
 * @Command: go test -v double_linked_list.go
 * @Author: zongsh
 * @Date: 2019-08-27 11:14:37
 * @LastEditTime: 2019-08-28 16:34:20
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
type DouElem struct {
	Name string
	Age  int
}

/*
 * @description:表元素集合
 * @param   Node       DouElem   当前节点信息
 * @param   NextDate   *DoubleLinkedNode   下一个节点的指针
 * @param   PreNode   *DoubleLinkedNode    前一个节点的指针
 */
type DoubleLinkedNode struct {
	PreNode  *DoubleLinkedNode
	Node     DouElem
	NextNode *DoubleLinkedNode
}

/*
 * @description:初始化一个头节点
 */
func InitDoubleNode() *DoubleLinkedNode {
	return &DoubleLinkedNode{
		PreNode:  nil,
		Node:     DouElem{Name: "头节点", Age: 0},
		NextNode: nil,
	}
}

/*
 * @description:在末尾新增一个节点
 */
func (n *DoubleLinkedNode) AddNode(e DouElem) {
	p := n
	pw := n
	for nil != p {
		pw = p         // 当前节点
		p = p.NextNode // 下一个节点
	}
	pw.NextNode = &DoubleLinkedNode{PreNode: pw, Node: e}
}

/*
 * @description:遍历表
 */
func (n *DoubleLinkedNode) Traverse() []DouElem {
	var list []DouElem
	p := n
	for nil != p {
		p = p.NextNode
		if nil != p {
			list = append(list, p.Node)
		}
	}
	return list
}

/*
 * @description:在固定位置插入节点，复杂度为o(n)
 */
func (n *DoubleLinkedNode) Insert(index int, e DouElem) error {
	p := n
	j := 1
	for nil != p && j < index {
		p = p.NextNode
		j++
	}

	if nil == p || j > index {
		return errors.New(fmt.Sprintf("please check iindex:%d", index))
	}

	insertLinked := &DoubleLinkedNode{PreNode: p, Node: e, NextNode: p.NextNode}
	p.NextNode = insertLinked
	return nil
}

/*
 * @description:在固定位置删除节点，复杂度为o(n)
 */
func (n *DoubleLinkedNode) Delete(index int) error {
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
func (n *DoubleLinkedNode) Get(index int) DouElem {
	p := n
	for i := 0; i < index; i++ {
		p = p.NextNode
		if nil == p {
			return DouElem{}
		}
	}
	return p.Node
}
