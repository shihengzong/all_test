/*
 * @Description: please edi
 * @Command: please edit
 * @Author: zongsh
 * @Date: 2019-08-28 16:55:42
 * @LastEditTime: 2019-09-18 11:51:57
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"all_test/excel_test"
	"fmt"
)

func main() {
	// 测试链表
	// linked_list.AllTest()
	// 测试defer
	// fmt.Println(defer_test.DeferTset())
	excel_test.NewExcel()
	fmt.Println(excel_test.ReadExcel())
}
