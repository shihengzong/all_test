/*
 * @Description: 几种排序算法
 * @Command: please edit
 * @Author: zongsh
 * @Date: 2019-09-02 17:28:55
 * @LastEditTime: 2019-09-02 17:31:37
 * @LastEditors: zongsh
 */
package sort_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	num      = 100
	rangeNum = 100
)

//生成count个[start,end)结束的不重复的随机数
func GenerateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

// 冒泡排序
func maopao(buf []int) {
	times := 0
	flag := false
	for i := 0; i < len(buf)-1; i++ {
		for j := 1; j < len(buf)-i; j++ {
			if buf[j-1] > buf[j] {
				times++
				tmp := buf[j-1]
				buf[j-1] = buf[j]
				buf[j] = tmp
				flag = true
			}
		}
		fmt.Println("arr2=: ", buf)
		if !flag {
			break
		}
	}
	fmt.Println("maopao times: ", times)
	fmt.Println("arr=: ", buf)
}

// 选择排序 [3,1,2]
func xuanze(buf []int) {
	times := 0
	for i := 0; i < len(buf)-1; i++ {
		min := i
		for j := i; j < len(buf); j++ {
			times++
			if buf[min] > buf[j] {
				min = j
			}
		}
		if min != i {
			tmp := buf[i]
			buf[i] = buf[min]
			buf[min] = tmp
		}
		fmt.Println("min= ", min, "  arr=", buf)
	}
	fmt.Println("xuanze times: ", times)
}

// 插入排序
func charu(buf []int) {
	times := 0
	for i := 1; i < len(buf); i++ {
		for j := i; j > 0; j-- {
			if buf[j] < buf[j-1] {
				times++
				tmp := buf[j-1]
				buf[j-1] = buf[j]
				buf[j] = tmp
			} else {
				break
			}
		}
		fmt.Println("arr= ", buf)
	}
	fmt.Println("charu times: ", times)
}

// 希尔排序
func xier(buf []int) {
	times := 0
	tmp := 0
	length := len(buf)
	incre := length
	// fmt.Println("buf: ", buf)
	for {
		incre /= 2
		for k := 0; k < incre; k++ { //根据增量分为若干子序列
			for i := k + incre; i < length; i += incre {
				for j := i; j > k; j -= incre {
					fmt.Println("j: ", j, " data: ", buf[j], " j-incre: ", j-incre, " data: ", buf[j-incre])
					times++
					if buf[j] < buf[j-incre] {
						tmp = buf[j-incre]
						buf[j-incre] = buf[j]
						buf[j] = tmp
					} else {
						break
					}
				}
				fmt.Println("middle: ", buf)
			}
			fmt.Println("outer: ", buf)
		}
		fmt.Println("outer outer: ", buf, " incre: ", incre)

		if incre == 1 {
			break
		}
	}
	fmt.Println("after: ", buf)
	fmt.Println("xier times: ", times)
}

// 快速排序
func kuaisu(buf []int) {
	kuai(buf, 0, len(buf)-1)
}

func kuai(a []int, l, r int) {
	if l >= r {
		return
	}
	i, j, key := l, r, a[l] //选择第一个数为key
	for i < j {
		for i < j && a[j] > key { //从右向左找第一个小于key的值
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
		}

		for i < j && a[i] < key { //从左向右找第一个大于key的值
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
		}
	}
	//i == j
	a[i] = key
	kuai(a, l, i-1)
	kuai(a, i+1, r)
}

//归并排序
func guibing(buf []int) {
	tmp := make([]int, len(buf))
	merge_sort(buf, 0, len(buf)-1, tmp)
}

func merge_sort(a []int, first, last int, tmp []int) {
	if first < last {
		middle := (first + last) / 2
		merge_sort(a, first, middle, tmp)       //左半部分排好序
		merge_sort(a, middle+1, last, tmp)      //右半部分排好序
		mergeArray(a, first, middle, last, tmp) //合并左右部分
	}
}

func mergeArray(a []int, first, middle, end int, tmp []int) {
	// fmt.Printf("mergeArray a: %v, first: %v, middle: %v, end: %v, tmp: %v\n",
	//     a, first, middle, end, tmp)
	i, m, j, n, k := first, middle, middle+1, end, 0
	for i <= m && j <= n {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			k++
			i++
		} else {
			tmp[k] = a[j]
			k++
			j++
		}
	}
	for i <= m {
		tmp[k] = a[i]
		k++
		i++
	}
	for j <= n {
		tmp[k] = a[j]
		k++
		j++
	}

	for ii := 0; ii < k; ii++ {
		a[first+ii] = tmp[ii]
	}
	// fmt.Printf("sort: buf: %v\n", a)
}

// 堆排序
func duipai(buf []int) {
	temp, n := 0, len(buf)

	for i := (n - 1) / 2; i >= 0; i-- {
		MinHeapFixdown(buf, i, n)
	}

	for i := n - 1; i > 0; i-- {
		temp = buf[0]
		buf[0] = buf[i]
		buf[i] = temp
		MinHeapFixdown(buf, 0, i)
	}
}

func MinHeapFixdown(a []int, i, n int) {
	j, temp := 2*i+1, 0
	for j < n {
		if j+1 < n && a[j+1] < a[j] {
			j++
		}

		if a[i] <= a[j] {
			break
		}

		temp = a[i]
		a[i] = a[j]
		a[j] = temp

		i = j
		j = 2*i + 1
	}
}

func TestSort(t *testing.T) {
	// randSeed := rand.New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
	// var buf []int
	// for i := 0; i < num; i++ {
	// 	buf = append(buf, randSeed.Intn(rangeNum))
	// }

	buf := GenerateRandomNumber(1, 11, 10)
	buf = []int{1, 3, 9, 4, 5, 2}
	now := time.Now()
	//冒泡排序
	// maopao(buf)
	// 选择排序
	// xuanze(buf)
	// 插入排序
	// charu(buf)
	//希尔排序
	xier(buf)
	//快速排序
	// kuaisu(buf)
	// 归并排序
	// guibing(buf)
	// 堆排序
	//duipai(buf)
	// sort 排序
	// sort.Sort(IntSlice(buf))
	//sort 逆向排序
	//sort.Sort(sort.Reverse(sort.IntSlice(buf)))
	// fmt.Println(buf)

	//sort.string 排序
	//ss := []string{"surface", "ipad", "mac pro", "mac air", "think pad", "idea pad"}
	//sort.Strings(ss)
	//fmt.Println(ss)

	fmt.Println(time.Since(now))
}
