package common

import (
	"github.com/gagarinchain/common/eth/crypto"
	pb "github.com/gagarinchain/common/protobuff"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"testing"
)

func TestEmptyBlockSize(t *testing.T) {

	h := &pb.BlockHeader{
		Hash:       crypto.Keccak256Hash([]byte("hash")).Bytes(),
		ParentHash: crypto.Keccak256Hash([]byte("parent")).Bytes(),
		QcHash:     crypto.Keccak256Hash([]byte("qc")).Bytes(),
		DataHash:   nil,
		TxHash:     nil,
		StateHash:  nil,
		Height:     134123123,
		Timestamp:  0,
	}
	any1, _ := ptypes.MarshalAny(h)

	sync := &pb.SynchronizePayload{
		Height:    134123123,
		Cert:      nil,
		Signature: nil,
	}

	any2, _ := ptypes.MarshalAny(sync)

	m1, _ := proto.Marshal(&pb.Message{
		Type:    pb.Message_SYNCHRONIZE,
		Payload: any1,
	})
	m2, _ := proto.Marshal(&pb.Message{
		Type:    pb.Message_SYNCHRONIZE,
		Payload: any2,
	})

	println(len(m1))
	println(len(m2))
}
