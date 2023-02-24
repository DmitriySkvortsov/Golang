package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func countCall(f func(string)) func(string) {
	cnt := 0
	funcname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return func(s string) {
		cnt++
		fmt.Printf("Функция %s вызвана %d раз\n", funcname, cnt)
		f(s)
	}
}

func metricTimeCall(f func(string)) func(string) {
	return func(s string) {
		start := time.Now()
		f(s)
		fmt.Println("Execution time: ", time.Now().Sub(start))
	}
}

func myprint(s string) {
	fmt.Println(s)
}

func main() {

	countedPrint := countCall(myprint)
	countedPrint("Hello world")
	countedPrint("Hi")

	countAndMetricPrint := metricTimeCall(countedPrint)
	countAndMetricPrint("Hello")
	countAndMetricPrint("World")

}
