package main

import (
	"time"
	"log"
	"context"
	"math/rand"
	"sync"
)

type ctxXyzKeyType int
const ctxXyzKey ctxXyzKeyType = 0

func main() {
	log.SetFlags(log.Lmicroseconds)

	rand.Seed(time.Now().UnixNano())
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	const N = 5
	resultCh := make(chan int)

	var wg sync.WaitGroup

	for i := 0 ; i < N ; i++ {
		wg.Add(1)
		go func(ctx context.Context, index int) {
			defer wg.Done()
			result, err := doSomething(ctx)
			if err != nil {
				log.Println(index, "err:", err)
				return
			}
			select {
			case resultCh <- result:
				log.Println(index, "send result ok")
			case <-ctx.Done():
				log.Println(index, "err:", ctx.Err())
			}
		}(context.WithValue(ctx, ctxXyzKey, i), i)
	}

	result := <- resultCh
	cancel()
	wg.Wait()

	log.Println("result: ", result)
}

func doSomething(ctx context.Context) (int, error) {
	index := ctx.Value(ctxXyzKey)
	t := rand.Intn(10)+1
	d := time.Second * time.Duration(t)

	log.Printf("doSomething(index=%v) d=%v \n", index, d)

	select {
	case <-time.After(d):
		// done
	case <-ctx.Done():
		// cancel
		return 0, ctx.Err()
	}
	return t, nil
}
