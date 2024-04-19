package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "foo", "bar")
	userID := 10
	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result: ", val)
	fmt.Println("time taken: ", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	val := ctx.Value("foo")
	fmt.Println("value of foo: ", val)
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	
	respch := make(chan Response)
	
	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()
		respch <- Response{
			value: val,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("Fetching data from third party took too long")
		case resp := <-respch:
			return resp.value, resp.err
		}
	}
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 150)

	return 25, nil
}
