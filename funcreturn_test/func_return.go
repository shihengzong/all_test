package funcreturn_test

import "fmt"

func test() {
	var nums [2]int
	for i, v := range nums {
		fmt.Println("i==", i)
		if i == 0 {
			nums[i+1]++
		} else {
			fmt.Print(v)
		}
	}
}

//定义的type函数
type fun func(int)

func Tfunc(y int) {
	fmt.Println(3 + y)
}

func (f fun) Add(i int) {
	f(i)
}

//需要传递函数
func callback(i int) {
	fmt.Println("i am callBack")
	fmt.Println(i)
}

//main中调用的函数
func one(i int, f func(int)) {
	two(i, fun(f))
}

//one()中调用的函数
func two(i int, c Call) {
	c.call(i)
}

//fun实现的Cal接口的call()函数
func (f fun) call(i int) {
	f(i)
}

//接口
type Call interface {
	call(int)
}
