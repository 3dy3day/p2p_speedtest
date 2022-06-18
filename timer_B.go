package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	now := time.Now()
	result_rec, err_rec := exec.Command("bash", "-c", "./ftrans recv -p test1234 --signaling ws://YOURIP/ws").Output()
	if err_rec != nil {
		fmt.Println(string(result_rec))
	} else {
		fmt.Println(err_rec)
	}

	result, err := exec.Command("bash", "-c", "./ftrans send ../test.jpg -p test1234 --sig ws://YOURIP/ws").Output()
	if err != nil {
		fmt.Println(string(result))
	} else {
		fmt.Println(err)
	}

	fmt.Printf("Time elapsed: %vms\n", time.Since(now).Milliseconds())
}
