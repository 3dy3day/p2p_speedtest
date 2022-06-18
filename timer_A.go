package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	now := time.Now()
	result, err := exec.Command("bash", "-c", "./ftrans send ./test.jpg -p test1234 --sig ws://18.183.47.243:8081/ws").Output()
	if err != nil {
		fmt.Println(string(result))
	} else {
		fmt.Println(err)
	}

	result_rec, err_rec := exec.Command("bash", "-c", "./ftrans recv -p test1234 --signaling ws://18.183.47.243:8081/ws").Output()
	if err_rec != nil {
		fmt.Println(string(result_rec))
	} else {
		fmt.Println(err_rec)
	}

	fmt.Printf("Time elapsed: %vms\n", time.Since(now).Milliseconds())
}
