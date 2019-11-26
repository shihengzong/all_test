/*
 * @Description: please edit
 * @Command: please edit
 * @Author: zongsh
 * @Date: 2019-09-02 10:08:20
 * @LastEditTime: 2019-09-02 11:53:52
 * @LastEditors: zongsh
 */
package point_test

import (
	"fmt"
	"testing"

	"github.com/jinzhu/copier"
)

type Elem struct {
	Name string
	Age  int
}

func (t *Elem) ReEqual(e *Elem) {
	// t -> t   			外面的people地址 等于里面的t地址   people = 0xx10   t = 0xx10
	// t=e      t -> e		里面的t地址等于e的地址			                    t = xx121             所以t指向e, people不变
	t = e
}

func (t *Elem) ReParam(e *Elem) {
	t.Name = e.Name
}

func (t *Elem) ReCopy(e *Elem) {
	x := t
	x.Name = e.Name
}

func TestPointVar(t *testing.T) {
	var people *Elem
	e := &Elem{Name: "小明", Age: 18}
	people.ReEqual(e)              // 指针地址被该变
	fmt.Println("result:", people) // people == nil

	// -----------------------------------------------------------------------
	people2 := &Elem{Name: "小红", Age: 16}
	people2.ReEqual(e)               // 指针地址被该变
	fmt.Println("result2:", people2) // people不变: &{Name: "小红", Age: 16}

	// -----------------------------------------------------------------------
	people3 := &Elem{Name: "小蓝", Age: 20}
	people3.ReParam(e)               // 指针地址被不变
	fmt.Println("result3:", people3) // people的Name被修改从小蓝->小明   &{小明 20}

	// -----------------------------------------------------------------------
	people4 := &Elem{Name: "小黄", Age: 22}
	people4.ReCopy(e)                // 指针地址被不变
	fmt.Println("result4:", people4) // people的Name被修改从小黄->小明   &{小明 22}
}

func TestPointCopy(t *testing.T) {
	// 深拷贝和浅拷贝
	p1 := &Elem{
		Name: "Jane",
		Age:  18,
	}

	p2 := p1
	fmt.Printf("p1:%v\n", p1)
	fmt.Printf("p2:%v\n\n", p2)

	p2.Name = "zhangsan" // p1也被该变(浅拷贝)
	fmt.Printf("p1:%v\n", p1)
	fmt.Printf("p2:%v\n\n", p2)

	p3 := &Elem{}
	copier.Copy(p3, p1)
	fmt.Printf("p1:%v\n", p1)
	fmt.Printf("p2:%v\n\n", p3)
}

func TestPointEqual(t *testing.T) {
	people := &Elem{Name: "小明", Age: 18}
	people2 := &Elem{Name: "小红", Age: 16}
	people2 = people
	fmt.Printf("people:%v\n", people)
	fmt.Printf("people2:%v\n\n", people2)

	var people3 *Elem
	people3 = people
	fmt.Printf("people:%v\n", people)
	fmt.Printf("people3:%v\n\n", people3)
	fmt.Printf("%p", people) // %p 打印地址
}

func TestPointSilence(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s2[0] = 10
	fmt.Printf("s1:%p\n", s1)
	fmt.Printf("s2:%p\n", s2)
	fmt.Printf("s1:%v\n", s1)
	fmt.Printf("s2:%v\n\n", s2)

	s1 = append(s1, 11)
	fmt.Printf("s1:%p\n", s1)
	fmt.Printf("s2:%p\n", s2)
	fmt.Printf("s1:%v\n", s1)
	fmt.Printf("s2:%v\n\n", s2)
	// s1:0xc000074060
	// s2:0xc000074060
	// s1:[10 2 3 4 5]
	// s2:[10 2 3 4 5]

	// s1:0xc00007c280
	// s2:0xc000074060
	// s1:[10 2 3 4 5 11]
	// s2:[10 2 3 4 5]
}
