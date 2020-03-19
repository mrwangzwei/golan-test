package utils

import (
	"strconv"
	"strings"
)

//字符串转slice，并且值为int
//dataType. 0：直接用string；10：int；64：int64
func StringToIntSlice(str string, spec string) (result []int, err error) {
	slice := strings.Split(str, spec)
	var res int
	for _, a := range slice {
		res, err = strconv.Atoi(a)
		if err != nil {
			return
		}
		result = append(result, res)
	}
	return
}

//获取数组重叠的部分
func GetIntSliceIn(arr1 []int, arr2 []int) (res []int) {
	var in bool
	for _, item := range arr1 {
		in = false
		for _, item2 := range arr2 {
			if item == item2 {
				in = true
			}
		}
		if in {
			res = append(res, item)
		}
	}
	return
}

//找出数组不同的部分
func GetIntSliceDiff(arr1 []int, arr2 []int) (res []int) {
	var in bool
	for _, item := range arr1 {
		in = false
		for _, item2 := range arr2 {
			if item == item2 {
				in = true
			}
		}
		if !in {
			res = append(res, item)
		}
	}
	return
}

//数组合并（不重复）
func IntSlicesMerge(arr1 []int, arr2 []int) (res []int) {
	var in bool
	res = arr1
	for _, item2 := range arr2 {
		in = false
		for _, item := range arr1 {
			if item2 == item {
				in = true
			}
		}
		if !in {
			res = append(res, item2)
		}
	}
	return
}
