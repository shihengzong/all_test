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
)

type Elem struct {
	Name string
	Age  int
}

func (t *Elem) ReEqual(e *Elem) {
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
