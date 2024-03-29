package crypto

import (
	"github.com/gagarinchain/common/eth/crypto/bls12_381"
	pb "github.com/gagarinchain/common/protobuff"
	"github.com/op/go-logging"
	"github.com/phoreproject/bls/g1pubs"
	"math/big"
)

var log = logging.MustGetLogger("crypto")

type PublicKey struct {
	v *g1pubs.PublicKey
}

func (key *PublicKey) V() *g1pubs.PublicKey {
	return key.v
}

func NewPublicKey(v *g1pubs.PublicKey) *PublicKey {
	return &PublicKey{v: v}
}

func (key *PublicKey) Bytes() []byte {
	b := key.v.Serialize()
	return b[:]
}

func FromBytes(b []byte) (key *PublicKey) {
	var bytes [48]byte
	copy(bytes[:], b[:48])

	publicKey, err := g1pubs.DeserializePublicKey(bytes)
	if err != nil {
		log.Error(err)
		return nil
	}

	return &PublicKey{
		v: publicKey,
	}
}
func PkFromBytes(b []byte) (key *PrivateKey) {
	var bytes [32]byte
	copy(bytes[:], b[:32])

	pk := g1pubs.DeserializeSecretKey(bytes)

	return &PrivateKey{
		v: pk,
	}
}

type PrivateKey struct {
	v *g1pubs.SecretKey
}

func (pk *PrivateKey) V() *g1pubs.SecretKey {
	return pk.v
}

func NewPrivateKey(v *g1pubs.SecretKey) *PrivateKey {
	return &PrivateKey{v: v}
}

func (pk *PrivateKey) PublicKey() *PublicKey {
	return &PublicKey{g1pubs.PrivToPub(pk.v)}
}

type Signature struct {
	pub  *g1pubs.PublicKey
	sign *g1pubs.Signature
}

func (s *Signature) IsEmpty() bool {
	return g1pubs.NewAggregatePubkey().Equals(*s.Pub())
}

func (s *Signature) Sign() *g1pubs.Signature {
	return s.sign
}

func (s *Signature) Pub() *g1pubs.PublicKey {
	return s.pub
}

func (s *Signature) ToProto() *pb.Signature {
	if s == nil {
		return nil
	}
	if s.IsEmpty() {
		return &pb.Signature{}
	}
	pkBytes := s.Pub().Serialize()
	signBytes := s.Sign().Serialize()
	return &pb.Signature{
		From:      pkBytes[:],
		Signature: signBytes[:],
	}
}

func (s *Signature) ToStorageProto() *pb.SignatureS {
	if s == nil {
		return nil
	}
	if s.IsEmpty() {
		return &pb.SignatureS{}
	}
	pkBytes := s.Pub().Serialize()
	signBytes := s.Sign().Serialize()
	return &pb.SignatureS{
		From:      pkBytes[:],
		Signature: signBytes[:],
	}
}

func SignatureFromProto(mes *pb.Signature) *Signature {
	return NewSignatureFromBytes(mes.From, mes.Signature)
}

func SignatureFromStorageProto(mes *pb.SignatureS) *Signature {
	return NewSignatureFromBytes(mes.From, mes.Signature)
}

func NewSignature(pk *g1pubs.PublicKey, sign *g1pubs.Signature) *Signature {
	return &Signature{pub: pk, sign: sign}
}

func NewSignatureFromBytes(pk []byte, sign []byte) *Signature {
	empty := EmptySignature()
	if len(pk) == 0 {
		return empty
	}
	pkBytes := bls12_381.ToBytes48(pk)
	key, e := g1pubs.DeserializePublicKey(pkBytes)
	if e != nil {
		log.Error(e)
		return empty
	}
	signBytes := bls12_381.ToBytes96(sign)
	signature, e := g1pubs.DeserializeSignature(signBytes)
	if e != nil {
		log.Error(e)
		return empty
	}

	return NewSignature(key, signature)
}

type SignatureAggregate struct {
	bitmap    *big.Int
	aggregate *g1pubs.Signature
}

func (sa *SignatureAggregate) Aggregate() *g1pubs.Signature {
	return sa.aggregate
}

func (sa *SignatureAggregate) N() int {
	n := 0
	for i := 0; i < sa.bitmap.BitLen(); i++ {
		if sa.bitmap.Bit(i) == 1 {
			n++
		}
	}

	return n
}

func (sa *SignatureAggregate) Bitmap() *big.Int {
	return sa.bitmap
}

func (sa *SignatureAggregate) IsValid(message []byte, committee []*PublicKey) (res bool) {
	pubs := sa.KeysBitMapped(committee)
	return VerifyAggregate(message, pubs, sa)
}

func (sa *SignatureAggregate) KeysBitMapped(committee []*PublicKey) []*PublicKey {
	var pubs []*PublicKey
	for i, p := range committee {
		if sa.Bitmap().Bit(i) == 1 {
			pubs = append(pubs, p)
		}
	}
	return pubs
}

func AggregateFromProto(mes *pb.SignatureAggregate) *SignatureAggregate {
	return NewAggregateFromBytes(mes.Bitmap, mes.Signature)
}
func AggregateFromStorage(mes *pb.SignatureAggregateS) *SignatureAggregate {
	return NewAggregateFromBytes(mes.Bitmap, mes.Signature)
}

func EmptySignature() *Signature {
	return &Signature{
		pub:  g1pubs.NewAggregatePubkey(),
		sign: g1pubs.NewAggregateSignature(),
	}
}

func EmptyAggregateSignatures() *SignatureAggregate {
	return &SignatureAggregate{
		bitmap:    big.NewInt(0),
		aggregate: g1pubs.NewAggregateSignature(),
	}
}

func NewAggregateFromBytes(bitmap []byte, sign []byte) *SignatureAggregate {
	b := big.NewInt(0).SetBytes(bitmap)

	signBytes := bls12_381.ToBytes96(sign)
	s, e := g1pubs.DeserializeSignature(signBytes)
	if e != nil {
		log.Error(e)
		return nil
	}

	return &SignatureAggregate{
		bitmap:    b,
		aggregate: s,
	}
}

func (sa *SignatureAggregate) ToProto() *pb.SignatureAggregate {
	sign := sa.aggregate.Serialize()
	return &pb.SignatureAggregate{
		Bitmap:    sa.bitmap.Bytes(),
		Signature: sign[:],
	}
}

func (sa *SignatureAggregate) ToStorageProto() *pb.SignatureAggregateS {
	sign := sa.aggregate.Serialize()
	return &pb.SignatureAggregateS{
		Bitmap:    sa.bitmap.Bytes(),
		Signature: sign[:],
	}
}
