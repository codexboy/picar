package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	//"time"
	"electrical"
	"net/http"
	"infrared"
	"time"
)

const (
	inf_pin uint8 = 4
	lele_pin1 uint8 = 17
	lele_pin2 uint8 = 22
	rele_pin1 uint8 = 23
	rele_pin2 uint8 = 24
	// pin_echo = 21
	//pin_trig = 22
)

func clo(){
	fmt.Println("defer close")
	rpio.Close()
}

var left *electrical.Elect
var right *electrical.Elect
var inf *infrared.Infrared
var flag bool = false

func leftTrun(){
	stopTrun()
	right.Forward()

	flag = true
	for flag {
		if inf.Check() {
			stopTrun()
		}else {
			right.Forward()
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func rightTrun(){
	stopTrun()
	left.Forward()

	flag = true
	for flag {
		if inf.Check() {
			stopTrun()
		}else {
			left.Forward()
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func stopTrun(){
	flag = false
	left.Stop()
	right.Stop()
}

func forwardTrun(){
	stopTrun()
	left.Forward()
	right.Forward()

	flag = true
	for flag {
		if inf.Check() {
			stopTrun()
		}else {
			left.Forward()
			right.Forward()
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func backTruc(){
	stopTrun()
	left.Backup()
	right.Backup()
}

func forwardHandle(w http.ResponseWriter,r *http.Request){
	fmt.Println("forward")
	forwardTrun()
}


func backHandle(w http.ResponseWriter,r *http.Request){
	fmt.Println("back")
	backTruc()
}


func stopHandle(w http.ResponseWriter,r *http.Request){
	fmt.Println("stop")
	stopTrun()
}


func leftHandle(w http.ResponseWriter,r *http.Request){
	fmt.Println("left")
	leftTrun()
}

func rightHandle(w http.ResponseWriter,r *http.Request){
	fmt.Println("right")
	rightTrun()
}



func main() {
	err := rpio.Open()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer clo()

	inf = infrared.NewInfrared(inf_pin)

	left = electrical.NewElect(lele_pin1,lele_pin2)
	right = electrical.NewElect(rele_pin1,rele_pin2)

	stopTrun()

	//for {
	//	time.Sleep(time.Millisecond * 200)
	//	if inf.Check() {
	//		left.Stop()
	//		right.Stop()
	//	} else {
	//		left.Forward()
	//		right.Forward()
	//	}
	//}




	http.HandleFunc("/forward",forwardHandle)
	http.HandleFunc("/backup",backHandle)
	http.HandleFunc("/stop",stopHandle)
	http.HandleFunc("/left",leftHandle)
	http.HandleFunc("/right",rightHandle)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err = http.ListenAndServe(":8888",nil)
	if err != nil {
		fmt.Println("http started")
	}
}
