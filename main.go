package main

import "fmt"

//可変超引数で使用
func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	//明示的な変数の定義
	var i int = 100
	fmt.Println(i)

	var sta string = "Hello Go"
	fmt.Println(sta)

	//複数の変数に纏めて値を入れる

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 234
	s3 = "GO GO"
	fmt.Println(i3, s3)

	//暗黙的な変数の定義
	//変数名　:= 値
	i4 := 400
	fmt.Println(i4)

	i4 = 450

	//int型

	var inte int = 100

	fmt.Println(inte + 33)

	//スライス
	var sl []int
	fmt.Println(sl)

	//append 追加
	sl = append(sl, 22, 33, 44, 55)
	fmt.Println(sl)

	//make
	sl2 := make([]int, 20, 25)
	fmt.Println(sl2)

	//len・・レングス
	fmt.Println(len(sl2))
	//cap・・キャパシティ
	fmt.Println(cap(sl2))
	//copy
	sl3 := []int{1, 2, 3, 4, 5}
	sl4 := make([]int, 5, 10)
	fmt.Println(sl4)
	n := copy(sl4, sl3) //copy(コピー先, コピー元)
	fmt.Println(n, sl4) //nはコピーに成功した数を返す

	//for
	sl5 := []string{"A", "B", "C"}
	fmt.Println(sl5)

	for i, v := range sl5 {
		fmt.Println(i, v)
	}
	//可変超引数
	fmt.Println(Sum(1, 3, 5))
}
