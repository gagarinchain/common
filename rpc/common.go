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

func (c *CommonClient) GetView(ctx context.Context) {
	timer := time.NewTimer(time.Second)
	for {
		select {
		case <-timer.C:
			if view, err := c.pbc.GetCurrentView(ctx, &pb.GetCurrentViewRequest{}); err != nil {
				log.Error(err)
			} else {
				log.Info(view)
			}
			timer = time.NewTimer(time.Second)
		case <-ctx.Done():
			c.Stop()
		}

	}

}
