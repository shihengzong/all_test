/*
 * @Description: 数据结构 -- 循环链表(Cycle Linked List)
 * @Command: go test -v cycle_linked_list.go
 * @Author: zongsh
 * @Date: 2019-08-27 15:58:53
 * @LastEditTime: 2019-08-28 16:34:43
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
type CycleElem struct {
	Name string
	Age  int
}

/*
 * @description:表元素集合
 * @param   Node       CycleLinkedNode   当前节点信息
 * @param   NextDate   *CycleElem   下一个节点的指针
 * @param   PreNode   *CycleLinkedNode    前一个节点的指针
 */
type CycleLinkedNode struct {
	PreNode  *CycleLinkedNode
	Node     CycleElem
	NextNode *CycleLinkedNode
}

/*
 * @description:初始化一个头节点
 */
func InitCycleNode() *CycleLinkedNode {
	return &CycleLinkedNode{
		PreNode:  nil,
		Node:     CycleElem{Name: "头节点", Age: 0},
		NextNode: nil,
	}
}

/*
 * @description:在末尾新增一个节点
 */
func (n *CycleLinkedNode) AddNode(e CycleElem) {
	p := n
	pw := n
	for nil != p {
		pw = p         // 当前节点
		p = p.NextNode // 下一个节点
		if pw.NextNode == n {
			break
		}
	}

	// 循环链表 (尾节点的nextNode指向头节点)
	pw.NextNode = &CycleLinkedNode{PreNode: pw, Node: e, NextNode: n}
	n.PreNode = pw.NextNode
}

/*
 * @description:遍历表
 */
func (n *CycleLinkedNode) Traverse() []CycleElem {
	var list []CycleElem
	p := n
	pw := n
	for nil != p {
		pw = p
		p = p.NextNode
		if nil != p {
			list = append(list, pw.Node)
		}
		if pw.NextNode == n {
			break
		}
	}
	return list
}

/*
 * @description:在固定位置插入节点，复杂度为o(n)
 */
func (n *CycleLinkedNode) Insert(index int, e CycleElem) error {
	p := n
	j := 1
	for nil != p && j < index {
		p = p.NextNode
		j++
	}

	if nil == p || j > index {
		return errors.New(fmt.Sprintf("please check iindex:%d", index))
	}

	insertLinked := &CycleLinkedNode{PreNode: p, Node: e, NextNode: p.NextNode}
	// 循环链表
	if nil == p.NextNode {
		insertLinked = &CycleLinkedNode{PreNode: p, Node: e, NextNode: n}
		n.PreNode = insertLinked
	}
	p.NextNode = insertLinked
	return nil
}

/*
 * @description:在固定位置删除节点，复杂度为o(n)
 */
func (n *CycleLinkedNode) Delete(index int) error {
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
func (n *CycleLinkedNode) Get(index int) CycleElem {
	p := n
	for i := 0; i < index; i++ {
		p = p.NextNode
		if nil == p {
			return CycleElem{}
		}
	}
	return p.Node
}
