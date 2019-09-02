/*
 * @Description: please edi
 * @Command: please edit
 * @Author: zongsh
 * @Date: 2019-08-28 16:55:42
 * @LastEditTime: 2019-08-30 17:59:27
 * @LastEditors: zongsh
 */
package main

import (
	"all_test/linked_list"
	"fmt"
	"time"
)

func main() {
	// 测试链表
	linked_list.AllTest()
	// ChanTest()
}

func ChanTest() {

	c := make(chan int, 10)
	go func(ct chan int) {
	For:
		for {
			select {
			case data, ok := <-ct: // channel 被关闭时ok会变成false,在关闭过程中一直可读不可写
				if !ok {
					fmt.Println("chan closed")
					break For
				}
				fmt.Println("receive from channel:", ok, data)
			}
			time.Sleep(10 * time.Microsecond) // 模拟关闭所用时间
		}
	}(c)
	c <- 1
	close(c)
	time.Sleep(5)
}
