/*
 * @Description: 数据结构 -- 单链表(Singly Linked List)
 * @Command: go test -v singly_linked_list_test.go
 * @Author: zongsh
 * @Date: 2019-08-27 11:14:37
 * @LastEditTime: 2019-08-27 15:34:52
 * @LastEditors: zongsh
 */
package linked_list_test

import (
	"errors"
	"fmt"
	"testing"
)

/*
 * @description:表里的元素
 */
type Elem struct {
	Name string
	Age  int
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
func (n *SinglyLinkedNode) Traverse() []Elem {
	var list []Elem
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
func (n *SinglyLinkedNode) Get(index int) Elem {
	p := n
	for i := 0; i < index; i++ {
		p = p.NextNode
		if nil == p {
			return Elem{}
		}
	}
	return p.Node
}

func TestSinglyLinkedList(t *testing.T) {
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
