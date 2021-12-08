package logic

import (
	"context"
	
	s "github.com/Route42/ct-lib/service"
)

func CreateAndSubmitResults(ctx context.Context, service *s.Service) {
	for {
		select {
		case <-ctx.Done():
			return
		default: 
			result, err := service.Work()
		
			if err == nil && result != nil && result.IsValid() {
				service.SubmitResult(result)
			}
		}
	}
}