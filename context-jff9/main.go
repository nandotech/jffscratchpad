package main

import (
	"log"
	"fmt"
	"context"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	mySleepAndTalk(ctx, time.Second, "hello")

}

func mySleepAndTalk(c context.Context, sec time.Duration, s string) {
	select {
	case <-time.After(sec):
		fmt.Println(s)
	case<-c.Done():
		log.Print(c.Err())
	}
}
