package gobasic

import (
	"fmt"
	"time"
)

func ShowConcurrency() {
	fmt.Println(" ---------- gorutine ---------- ")
	// noRutine()
	// goRutine()
	goroutines2()
	fmt.Println(" ---------- channel ---------- ")
	// goChannels()

}

// example goroutines 1
func doSometing(index int) {
	time.Sleep(time.Second * 1)
	fmt.Println("do something: ", index)
}

func noRutine() {
	start := time.Now()
	for i := 0; i < 5; i++ {
		doSometing(i)
	}
	time.Sleep(2 * time.Second)
	fmt.Println(time.Since(start).Seconds())
}

func goRutine() {
	start := time.Now()
	for i := 0; i < 5; i++ {
		go doSometing(i)
	}
	go fmt.Println("go rutine: ")
	time.Sleep(2 * time.Second)
	fmt.Println(time.Since(start).Seconds())
}

// example goroutines 2
func goroutines2() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println(i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

// example go channels
func goChannels() {
	messages := make(chan string)

	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Print("msg: ", msg)
}
