/*
 * @Description: please edit
 * @Command: please edit
 * @Author: zongsh
 * @Date: 2019-09-02 10:08:20
 * @LastEditTime: 2019-09-02 10:20:57
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

func TestPointVar(t *testing.T) {
	var people *Elem
	e := &Elem{Name: "小明", Age: 18}
	people.ReEqual(e)
	fmt.Println("result:", people) // people == nil

	// -----------------------------------------------------------------------
	people2 := &Elem{}
	people2.ReEqual(e)
	fmt.Println("result2:", people2) // people == nil
}
