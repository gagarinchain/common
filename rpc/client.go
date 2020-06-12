package rpc

import (
	pb "github.com/gagarinchain/common/protobuff"
	"github.com/op/go-logging"
	"google.golang.org/grpc"
)

var log = logging.MustGetLogger("rpc")

type Client struct {
	conn *grpc.ClientConn
}

func (c *Client) Conn() *grpc.ClientConn {
	return c.conn
}

func NewClient(address string) *Client {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	//opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	return &Client{
		conn: conn,
	}
}

func (c *Client) Stop() {
	if err := c.conn.Close(); err != nil {
		log.Error(err)
	}
}

type OnReceiveProposalClient struct {
	*Client
	pbc pb.OnReceiveProposalClient
}

func (c *OnReceiveProposalClient) Pbc() pb.OnReceiveProposalClient {
	return c.pbc
}

func InitOnReceiveProposalClient(address string) *OnReceiveProposalClient {
	client := NewClient(address)
	return &OnReceiveProposalClient{
		Client: client,
		pbc:    pb.NewOnReceiveProposalClient(client.conn),
	}
}

type OnProposalClient struct {
	*Client
	pbc pb.OnProposalClient
}

func (c *OnProposalClient) Pbc() pb.OnProposalClient {
	return c.pbc
}

func InitOnProposalClient(address string) *OnProposalClient {
	client := NewClient(address)
	return &OnProposalClient{
		Client: client,
		pbc:    pb.NewOnProposalClient(client.conn),
	}
}

type OnVoteReceivedClient struct {
	*Client
	pbc pb.OnVoteReceivedClient
}

func (c *OnVoteReceivedClient) Pbc() pb.OnVoteReceivedClient {
	return c.pbc
}

func InitOnVoteReceivedClient(address string) *OnVoteReceivedClient {
	client := NewClient(address)
	return &OnVoteReceivedClient{
		Client: client,
		pbc:    pb.NewOnVoteReceivedClient(client.conn),
	}
}

type OnNewBlockCreatedClient struct {
	*Client
	pbc pb.OnNewBlockCreatedClient
}

func (c *OnNewBlockCreatedClient) Pbc() pb.OnNewBlockCreatedClient {
	return c.pbc
}

func InitOnNewBlockCreatedClient(address string) *OnNewBlockCreatedClient {
	client := NewClient(address)
	return &OnNewBlockCreatedClient{
		Client: client,
		pbc:    pb.NewOnNewBlockCreatedClient(client.conn),
	}
}

type OnBlockCommitClient struct {
	*Client
	pbc pb.OnBlockCommitClient
}

func (c *OnBlockCommitClient) Pbc() pb.OnBlockCommitClient {
	return c.pbc
}

func InitOnBlockCommitClient(address string) *OnBlockCommitClient {
	client := NewClient(address)
	return &OnBlockCommitClient{
		Client: client,
		pbc:    pb.NewOnBlockCommitClient(client.conn),
	}
}

type OnNextViewClient struct {
	*Client
	pbc pb.OnNextViewClient
}

func (c *OnNextViewClient) Pbc() pb.OnNextViewClient {
	return c.pbc
}

func InitOnNextViewClient(address string) *OnNextViewClient {
	client := NewClient(address)
	return &OnNextViewClient{
		Client: client,
		pbc:    pb.NewOnNextViewClient(client.conn),
	}
}

type OnNextEpochClient struct {
	*Client
	pbc pb.OnNextEpochClient
}

func (c *OnNextEpochClient) Pbc() pb.OnNextEpochClient {
	return c.pbc
}

func InitOnNextEpochClient(address string) *OnNextEpochClient {
	client := NewClient(address)
	return &OnNextEpochClient{
		Client: client,
		pbc:    pb.NewOnNextEpochClient(client.conn),
	}
}
