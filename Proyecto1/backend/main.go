package main

import (
	cpucontroller "backend/controllers/cpu.controller"
	ramcontroller "backend/controllers/ram.controller"
	"fmt"
	"time"
)

func main() {
	// rutina para cpu
	go func() {
		for {
			cpuInfo, err := cpucontroller.Get()
			if err != nil {
				fmt.Println(err.Error())
			}
			if err := cpucontroller.Post(*cpuInfo); err != nil {
				fmt.Println(err.Error())
			}
			time.Sleep(2500 * time.Millisecond)
		}
	}()
	// rutina para ram
	go func() {
		for {
			ramInfo, err := ramcontroller.Get()
			if err != nil {
				fmt.Println(err.Error())
			}
			ramInfo.Total = (ramInfo.Total / 1024) / 1024
			ramInfo.Used = (ramInfo.Used / 1024) / 1024
			ramInfo.Free = (ramInfo.Free / 1024) / 1024
			if err := ramcontroller.Post(*ramInfo); err != nil {
				fmt.Println(err.Error())
			}
			time.Sleep(2500 * time.Millisecond)
		}
	}()
	server := NewServer(":8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
