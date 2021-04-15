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

	//map
	//明示的なmapの宣言
	var m = map[string]int{"A": 1, "B": 2}
	fmt.Println(m)

	//暗黙的なmapの宣言
	m2 := map[int]string{
		1: "test",
		2: "develop",
	}
	fmt.Println(m2)

	//make関数でmapの作成
	m3 := make(map[int]string)
	fmt.Println(m3)

	//mapに値の追加
	m3[1] = "JAPAN" //map[key] = 追加したい値
	m3[2] = "USA"
	m3[3] = "CHINA"
	fmt.Println(m3)
	//値の取り出し ※mapにない値も取得できてしまうのでエラーハンドリングが必要
	fmt.Println(m3[2])
	//errorハンドリング
	v, ok := m3[3] //値と取得出来たかをboolで取得することができる
	if !ok {       //okがboolで返ってくるので判別できる
		fmt.Println("Not Value")
	}
	fmt.Println(v, ok)

	//delete
	delete(m3, 3)
	fmt.Println(m3)
	//len関数
	fmt.Println(len(m3))
	//map for
	m4 := map[string]int{"Apple": 100, "Banana": 200}
	for k, v := range m4 {
		fmt.Println(k, v)
	}

	//channel
	//複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造
	var ch1 chan int

	//受信専用
	//var ch2 <-chan int

	//送信専用
	//var ch3 chan<- int

	ch1 = make(chan int)     //初期化と書き込み可能な状態にする
	ch2 := make(chan int, 5) //make関数で作成もできる 5はcap

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	//channnelにデータを送信
	ch2 <- 1 //1と言うデータをch2に送信
	ch2 <- 2
	ch2 <- 3
	fmt.Println(len(ch2))

	//channelからデータの受信
	i6 := <-ch2 //送信(データを挿入した物から順番に取得できる)
	fmt.Println(i6)
	fmt.Println("len", len(ch2))
	//channnelを閉じる
	close(ch1)
	//channnel for
	ch3 := make(chan int, 3)
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	close(ch3) //channelをforで展開する場合は、closeしてから行う
	for i := range ch3 {
		fmt.Println(i)
	}

	//select
	ch4 := make(chan int, 2)
	ch5 := make(chan string, 2)

	ch4 <- 5
	ch5 <- "test"
	ch4 <- 2
	ch5 <- "develop"

	select { //selectの処理はchannelの処理しか書けない。また処理はランダムに実行される
	case v1 := <-ch4:
		fmt.Println(v1 + 2000)
	case v2 := <-ch5:
		fmt.Println(v2 + "branch")
	default:
	}
	//pointer type

	//ポインタとは型とアドレスを組み合わせたデータ型
	//参照型のデータ型はポインタを使わなくても良い(スライス、マップなどは参照渡しできるため)

	var n3 int = 100
	fmt.Println(n3)
	fmt.Println(&n3) //n3のメモリ内のアドレスを表示

	var p *int = &n3 //pはn3のアドレスを参照する
	fmt.Println(p)   //pを出力するとアドレスが表示される
	fmt.Println(*p)  //値を表示させたい場合は*をつけると表示される
	*p = 255
	fmt.Println(n3) //　*pに値を代入するとn3の値も更新される
}
