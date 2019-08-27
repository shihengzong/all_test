/*
 * @Description: 数据结构 -- 线性表(linear list)
 * @Command: go test -v list_test.go
 * @Author: zongsh
 * @Date: 2019-08-26 15:36:40
 * @LastEditTime: 2019-08-27 15:31:07
 * @LastEditors: zongsh
 * @detail:一个线性表是n个具有相同特性的数据元素的有限序列;
 		   数据元素之间的关系是一对一的关系，即除了第一个和最后一个数据元素之外，其它数据元素都是首尾相接的;
*/

package list_test

import (
	"errors"
	"fmt"
	"sort"
	"testing"
)

/*
 * @description:线性表里的元素
 */
type Elem struct {
	Name string
	Age  int
}

/*
 * @description:线性表元素集合
 * @param   List     []Elem  元素数组
 * @param   MaxLen   int  最大长度
 * @param   CurrLen  int  当前长度
 */
type LinearList struct {
	List    []Elem
	MaxLen  int
	CurrLen int
}

/*
 * @description:初始化一个空列表
 */
func InitList(maxLen int) *LinearList {
	return &LinearList{
		List:    make([]Elem, maxLen),
		MaxLen:  maxLen,
		CurrLen: 0,
	}
}

/*
 * @description:清空一个列表
 */
func (list *LinearList) ClearList() *LinearList {
	list = InitList(list.MaxLen)
	return list
}

/*
 * @description:判断列表是否已满
 */
func (list *LinearList) IsFull() bool {
	return list.MaxLen <= list.CurrLen
}

/*
 * @description:判断列表是否为空
 */
func (list *LinearList) IsEmpty() bool {
	return list.CurrLen <= 0
}

/*
 * @description:获取列表长度
 */
func (list *LinearList) ListLen() int {
	return list.CurrLen
}

/*
 * @description:在列表的i位置上添加一个元素
 * @param:  index int   插入位置
 * @param:  e 	 Elem  插入元素
 */
func (list *LinearList) Add(index int, e Elem) error {
	// 判断是否越界
	if index < 1 || index > list.MaxLen+1 {
		return errors.New("下标越界!")
	}

	// 判断列表是否已满
	if isFull := list.IsFull(); isFull {
		return errors.New("列表已满!")
	}

	// 进行添加
	for i := list.CurrLen; i >= index; i-- {
		list.List[i] = list.List[i-1]
	}
	list.List[index-1] = e
	list.CurrLen++
	return nil
}

/*
 * @description:删除列表里指定元素
 * @param:  index int   删除位置
 * @return:
 */
func (list *LinearList) Del(index int) error {
	// 判断是否越界
	if index < 1 || index > list.CurrLen {
		return errors.New("下标越界!")
	}

	// 判断列表是否为空
	if isEmpty := list.IsEmpty(); isEmpty {
		return errors.New("列表为空!")
	}

	// 进行删除
	for i := index; i <= list.CurrLen; i++ {
		list.List[i-1] = list.List[i]
	}
	list.List[list.CurrLen-1] = Elem{}
	list.CurrLen--
	return nil
}

/*
 * @description:追加一个元素
 * @param:  e 	 Elem  插入元素
 */
func (list *LinearList) Append(e Elem) error {
	// 判断列表是否已满
	if isFull := list.IsFull(); isFull {
		return errors.New("列表已满!")
	}

	// 进行追加
	list.List[list.CurrLen] = e
	list.CurrLen++
	return nil
}

/*
 * @description:获取index位置的元素
 * @param:  index int   获取位置
 */
func (list *LinearList) Get(index int) Elem {
	// 判断是否越界
	if index < 1 || index > list.CurrLen {
		return Elem{}
	}
	return list.List[index-1]
}

/*
 * @description: 遍历所有元素
 */
func (list *LinearList) Traverse() string {
	end := list.CurrLen
	if list.CurrLen == 0 {
		end = 1
	}
	return fmt.Sprintf("%v", list.List[0:end])
}

type Elems []Elem

func (s Elems) Len() int      { return len(s) }
func (s Elems) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// 按名字排序
type SortByName struct {
	Elems
}

func (m SortByName) Less(i, j int) bool {
	return m.Elems[i].Name > m.Elems[j].Name
}

// 按年龄排序
type SortByAge struct {
	Elems
}

func (m SortByAge) Less(i, j int) bool {
	return m.Elems[i].Age > m.Elems[j].Age
}

/*
 * @description: 给表元素排序
 */
func (list *LinearList) Sort() {
	sort.Sort(SortByName{list.List})
}

func TestLinearList(t *testing.T) {
	list := InitList(100)

	// 添加元素
	list.Add(1, Elem{Name: "1小明", Age: 30})
	list.Add(2, Elem{Name: "3小红", Age: 80})
	list.Add(3, Elem{Name: "2小蓝", Age: 18})
	list.Add(1, Elem{Name: "4小绿", Age: 14})
	list.Append(Elem{Name: "5小花", Age: 60})

	// 通过表name字段排序
	list.Sort()

	// 删除列表第五个元素
	list.Del(5)

	// 获取列表第四个元素
	fmt.Println(list.Get(4))

	// 遍历列表有效元素
	fmt.Println(list.Traverse())

	// a := &Elem{Name: "lisi", Age: 27}

	// b := a

	// b = &Elem{Name: "zhaoliu", Age: 30}
	// fmt.Println(a)
	// fmt.Println(b)
}
