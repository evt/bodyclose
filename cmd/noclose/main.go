package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		http.ListenAndServe(":8082", nil)
	}()

	var step = 0

	for {
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
	_, err := http.Get("https://www.baidu.com")
	if err != nil {
		return fmt.Errorf("http.Get failed: %w", err)
	}

	return nil
}
