package main

import (
	"fmt"
	"sync"
)

/*弄个单个的func man() {

	names := make(chan string,4)

		go func() {
			names <- "张三"
		}()
		go func(){
			names<-"李四"
		}()
		go func(){
			names<-"王五"
		}()
		go func(){
			names<- "赵六"
		}()
		for  {
			fmt.Println(<-names)
		}

}*/
func san(zs, ls chan bool) {
	defer wg.Done()
	defer close(ls)
	for i := 0; i < 10; i++ {
		<-zs
		fmt.Println("李四")
		ls <- true

	}
}
func si(ls, ww chan bool) {
	defer wg.Done()
	defer close(ww)
	for i := 0; i < 10; i++ {
		<-ls
		fmt.Println("王五")
		ww <- true

	}
}
func wu(ww, zl chan bool) {
	defer wg.Done()
	defer close(zl)
	for i := 0; i < 10; i++ {
		<-ww
		fmt.Println("赵六")
		zl <- true

	}
}
func liu(zl, zs chan bool) {
	defer wg.Done()
	defer close(zs)
	for i := 0; i < 10; i++ {
		<-zl
		fmt.Println("张三")
		zs <- true

	}
}

var wg sync.WaitGroup

func main() {
	zs := make(chan bool, 1)
	ls := make(chan bool, 1)
	ww := make(chan bool, 1)
	zl := make(chan bool, 1)
	zl <- true
	go san(zs, ls)
	go si(ls, ww)
	go wu(ww, zl)
	go liu(zl, zs)
	wg.Add(4)
	wg.Wait()
}
