package handler

import (
	"context"
	"coding-challenge/logic"
	"time"
	s "github.com/Route42/ct-lib/service"
)

func Handler(service *s.Service, ctx context.Context, cancel context.CancelFunc) {
	//launches 5 goroutines
	for i := 0; i < 5; i++ {
		go logic.CreateAndSubmitResults(ctx, service)
	}

	select {
		// if it's done we cancel the execution of the program
	case <- ctx.Done():
		cancel()
		// if more than 30 seconds have passed we again cancel the execution
	case <- time.After(time.Second * 30):
		cancel()
	}
}