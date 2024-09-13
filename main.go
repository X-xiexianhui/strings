package main

import "fmt"

func main() {
	ls := []string{"mno", "abc", "def", "ghi", "jkl", "pqr", "stu", "vwx", "she", "shells"}
	// LSDSort(ls, 3)
	Quick3String(ls)
	fmt.Printf("%v\n", ls)
}

func LSDSort(a []string, w int) {
	//通过前w个字符排序
	n := len(a)
	r := 256
	aux := make([]string, n)
	for d := w - 1; d >= 0; d-- {
		//根据第d个字符排序
		count := make([]int, r+1)
		for i := 0; i < n; i++ {
			count[int(a[i][d])+1]++
		}
		//将频率转换为索引
		for i := 0; i < r; i++ {
			count[i+1] += count[i]
		}
		//将元素分类
		for i := 0; i < n; i++ {
			aux[count[int(a[i][d])]] = a[i]
			count[int(a[i][d])]++
		}
		//回写
		for i := 0; i < n; i++ {
			a[i] = aux[i]
		}
	}
}

func insertSort(a []string, lo, hi, d int) {
	//从第d个字符串开始对a[lo:hi]排序
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && less(a[j], a[j-1], d); j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
func less(v, w string, d int) bool {
	return v[d:] < w[d:]
}

const (
	R = 256 //基数
	M = 15  //小数组的切换阈值
)

var aux []string

func charAt(s string, i int) int {
	if i < len(s) {
		return int(s[i])
	}
	return -1
}
func MSDSort(a []string) {
	n := len(a)
	aux = make([]string, n)
	msdSort(a, 0, n-1, 0)
}

func msdSort(a []string, lo, hi, d int) {
	//以第d个字符为关键字对a[lo:hi]排序
	if hi <= lo+M {
		insertSort(a, lo, hi, d)
		return
	}
	count := make([]int, R+2) //计算频率
	for i := lo; i <= hi; i++ {
		count[charAt(a[i], d)+2]++
	}
	for i := 0; i < R+1; i++ { //将频率转换为索引
		count[i+1] += count[i]
	}
	for i := lo; i <= hi; i++ { //将元素分类
		aux[count[charAt(a[i], d)+1]] = a[i]
		count[charAt(a[i], d)+1]++
	}
	for i := lo; i <= hi; i++ { //回写
		a[i] = aux[i-lo]
	}
	//递归的1️以每个字符为键进行排序
	for i := 0; i < R; i++ {
		msdSort(a, lo+count[i], lo+count[i+1]-1, d+1)
	}

}

func Quick3String(a []string) {
	quick3String(a, 0, len(a)-1, 0)
}
func quick3String(a []string, lo, hi, d int) {
	if hi <= lo {
		return
	}
	lt := lo
	gt := hi
	v := charAt(a[lo], d)
	i := lo + 1

	for i <= gt {
		t := charAt(a[i], d)
		if t < v {
			a[lt], a[i] = a[i], a[lt]
			lt++
			i++
		} else if t > v {
			a[gt], a[i] = a[i], a[gt]
			gt--
		} else {
			i++
		}
	}
	quick3String(a, lo, lt-1, d)
	if v >= 0 {
		quick3String(a, lt, gt, d+1)
	}
	quick3String(a, gt+1, hi, d)
}
