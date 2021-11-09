package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	var step = 0

	for true {
		time.Sleep(time.Microsecond * 100)

		step++

		err := requestNoClose()
		if err != nil {
			fmt.Printf("[%d] requestNoClose failed: %s", step, err)
			continue
		}

		fmt.Printf("[%d] ok\n", step)
	}
}

func requestNoClose() error {
	go func() {
		http.ListenAndServe(":8081", nil)
	}()

	_, err := http.Get("https://www.baidu.com")
	if err != nil {
		return fmt.Errorf("http.Get failed: %w", err)
	}

	return nil
}