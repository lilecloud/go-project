package main

import (
	"fmt"
	"runtime"
	"testing"
	"unsafe"
)

func TestIf(t *testing.T) {
	if runtime.GOOS == "linux" {
		fmt.Println("linux")
	}
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
func TestStruct(t *testing.T) {
	type Empty struct{} // Empty是一个不包含任何字段的空结构体类型
	var s Empty
	// fmt.Println(s == nil)
	fmt.Println(unsafe.Sizeof(s))

}

func TestMap(t *testing.T) {
	var m map[string]string
	if m == nil {
		fmt.Println("m is nil")
	}

	m = map[string]string{}

	m1 := make(map[string]string, 8)
	m1["lil"] = "l"
	m1["ss"] = "2"

	for key, val := range m1 {
		fmt.Println("key=", key, " val=", val)
	}

	fmt.Println("m1[not]= ", m1["not"] == "")

}

func TestArray(t *testing.T) {
	var arr [5]int = [5]int{1, 2}
	fmt.Println(arr)

	arr1 := [...]int{1}
	fmt.Println(arr1)

	arr2 := [5]string{3: "hello", 4: "world"}
	fmt.Println(arr2)

	fmt.Println("arr2[3=]", arr2[3])

	arr3 := [...]struct {
		name string
		age  int
	}{
		{"liming", 10},
		{"xiaohua", 20},
	}

	fmt.Println(arr3)

	for index, val := range arr3 {
		fmt.Println("index=", index, " val=", val)
	}

	for i := 0; i < len(arr3); i++ {
		fmt.Println("i=", i, " val=", arr3[i])
	}

}

func TestSlice(t *testing.T) {
	arr := [...]int{1, 2, 3}
	fmt.Println(arr)

	arr1 := [2]int{1, 2}
	fmt.Println(arr1)

	var arr2 [5]int = [5]int{}
	fmt.Println(arr2)

	s1 := arr[1:1]

	fmt.Println(len(s1))

	s2 := make([]int, 2, 10)
	s2[0] = 22
	s2[1] = 11
	s2 = append(s2, 9)
	s2 = append(s2, 333)
	fmt.Printf("slice e : %v\n", s2)

	fmt.Printf("slice cap: %d\n", cap(s2))

	for index, val := range s2 {
		fmt.Println("index=", index, " val=", val)
	}
}
