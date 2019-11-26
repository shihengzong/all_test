/*
 * @Description: os  -- 文件的创建. 删除. 读取. 和写入
 * @Command: please edit
 * @Author: zongsh
 * @Date: 2019-09-02 17:14:24
 * @LastEditTime: 2019-09-03 17:05:25
 * @LastEditors: zongsh
 */
package os_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// /usr/local/bin:/usr/bin:/bin:/usr/local/games:/usr/games:/sbin:/usr/sbin:/usr/lib/go/bin:/usr/local/python3/bin:/home/zongsh-deppin/pip/pip.ini

//创建一个文件夹
func CreatFile() {
	f, err := os.Create("hello_world.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = f.Write([]byte("Hello World !\n"))
	_, err = f.Write([]byte("How are you ?\n"))
	CheckErr(err)
}

// 打开并读取文件
func OpenFiles() {
	fileName := "hello_world.txt"
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	cByte, err := ioutil.ReadAll(f)
	CheckErr(err)
	fmt.Println(string(cByte))
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func FileOSTest() {
	var wireteString = "I'm fine. Thank you. And you ?"
	var filename = "./hello_world.txt"
	var f *os.File
	var err1 error
	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	CheckErr(err1)
	n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
	CheckErr(err1)
	fmt.Printf("写入 %d 个字节\n", n)
}

func TestOs(t *testing.T) {
	//CreatFile()
	//TestOs()
	//Parse()
	//CreatFile()
	//FileOSTest()
	OpenFiles()
}
