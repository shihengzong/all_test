/*
 * @Description: please edit
 * @Command: please edit
 * @Author: zongsh
 * @Date: 2019-09-02 11:19:26
 * @LastEditTime: 2019-09-02 11:20:18
 * @LastEditors: zongsh
 */
package channel_test

import (
	"fmt"
	"testing"
	"time"
)

func TestChanClose(t *testing.T) {
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
