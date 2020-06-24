package rpc

import (
	"context"
	pb "github.com/gagarinchain/common/protobuff"
	"time"
)

type CommonClient struct {
	*Client
	pbc pb.CommonServiceClient
}

func (c *CommonClient) Pbc() pb.CommonServiceClient {
	return c.pbc
}

func InitCommonClient(address string) *CommonClient {
	client := NewClient(address)
	return &CommonClient{
		Client: client,
		pbc:    pb.NewCommonServiceClient(client.conn),
	}
}

func (c *CommonClient) PollView(ctx context.Context) chan int32 {
	timer := time.NewTimer(10 * time.Microsecond)
	res := make(chan int32)
	currentView := int32(0)
	for {
		select {
		case <-timer.C:
			if view, err := c.pbc.GetCurrentView(ctx, &pb.GetCurrentViewRequest{}); err != nil {
				log.Error(err)
			} else {
				if currentView != view.View {
					res <- view.View
				}
			}
			timer = time.NewTimer(10 * time.Microsecond)
		case <-ctx.Done():
			close(res)
		}

	}

}
