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

		err := requestWithClose()
		if err != nil {
			fmt.Printf("[%d] requestWithClose failed: %s", step, err)
			continue
		}

		fmt.Printf("[%d] ok\n", step)
	}
}

func requestWithClose() error {
	go func() {
		http.ListenAndServe(":8081", nil)
	}()

	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		return fmt.Errorf("http.Get failed: %w", err)
	}
	defer resp.Body.Close()

	return nil
}
