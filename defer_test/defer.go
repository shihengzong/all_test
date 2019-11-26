/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-18 11:49:07
 * @LastEditTime: 2019-09-18 11:50:21
 * @LastEditors: Please set LastEditors
 */
package defer_test

import "fmt"

func DeferTset() int {
	t := 1
	defer func() {
		t += 2
	}()
	fmt.Printf("[费用减免] %s 是否结清: %t \n", "123132", false)
	fmt.Printf("[费用减免]第 %d 未结清: %s \n", 2, "weihuan")
	return t
}
