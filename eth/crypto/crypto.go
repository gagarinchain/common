// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package crypto

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"github.com/gagarinchain/common/eth"
	"github.com/gagarinchain/common/eth/common"
	"github.com/gagarinchain/common/eth/crypto/bls12_381"
	"github.com/phoreproject/bls/g1pubs"
	"golang.org/x/crypto/sha3"
	"math/big"
)

var (
	secp256k1N, _  = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	secp256k1halfN = new(big.Int).Div(secp256k1N, big.NewInt(2))
)

var errInvalidPubkey = errors.New("invalid secp256k1 public key")

// Keccak256 calculates and returns the Keccak256 hash of the input data.
func Keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

// Keccak256Hash calculates and returns the Keccak256 hash of the input data,
// converting it to an internal Hash data structure.
func Keccak256Hash(data ...[]byte) (h common.Hash) {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	d.Sum(h[:0])
	return h
}

// Keccak512 calculates and returns the Keccak512 hash of the input data.
func Keccak512(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak512()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

//FOLLOWING 2 FUNCTIONS ARE USED FOR CONTRACTS IN ETH AND WE POSSIBLY WILL NEVER NEED THEM @dabasov

// CreateAddress creates an ethereum address given the bytes and the nonce
//func CreateAddress(b common.Address, nonce uint64) common.Address {
//	data, _ := EncodeToBytes([]interface{}{b, nonce})
//	return common.BytesToAddress(Keccak256(data)[12:])
//}

// CreateAddress2 creates an ethereum address given the address bytes, initial
// contract code hash and a salt.
//func CreateAddress2(b common.Address, salt [32]byte, inithash []byte) common.Address {
//	return common.BytesToAddress(Keccak256([]byte{0xff}, b.Bytes(), salt[:], inithash)[12:])
//}

func GenerateKey() (*PrivateKey, error) {
	secretKey, e := g1pubs.RandKey(rand.Reader)
	if e != nil {
		return nil, e
	}
	return NewPrivateKey(secretKey), nil
}

// ValidateSignatureValues verifies whether the signature values are valid with
// the given chain rules. The v value is assumed to be either 0 or 1.
func ValidateSignatureValues(v byte, r, s *big.Int, homestead bool) bool {
	if r.Cmp(eth.Big1) < 0 || s.Cmp(eth.Big1) < 0 {
		return false
	}
	// reject upper range of s values (ECDSA malleability)
	// see discussion in secp256k1/libsecp256k1/include/secp256k1.h
	if homestead && s.Cmp(secp256k1halfN) > 0 {
		return false
	}
	// Frontier: allow s to be in full N range
	return r.Cmp(secp256k1N) < 0 && s.Cmp(secp256k1N) < 0 && (v == 0 || v == 1)
}

func PubkeyToAddress(pub *PublicKey) common.Address {
	pubBytes := pub.v.Serialize()
	return common.BytesToAddress(pubBytes[:])
}

func zeroBytes(bytes []byte) {
	for i := range bytes {
		bytes[i] = 0
	}
}

func Sign(msg []byte, pk *PrivateKey) *Signature {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, 0)
	sig := g1pubs.SignWithDomain(bls12_381.ToBytes32(msg), pk.v, bls12_381.ToBytes8(b))
	return NewSignature(pk.PublicKey().v, sig)
}

func Verify(msg []byte, s *Signature) (res bool) {
	defer func() {
		if r := recover(); r != nil {
			log.Error("Recovered in f", r)
			res = false
		}
	}()

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, 0)
	return g1pubs.VerifyWithDomain(bls12_381.ToBytes32(msg), s.pub, s.sign, bls12_381.ToBytes8(b))
}

func VerifyAggregate(msg []byte, pubs []*PublicKey, s *SignatureAggregate) (res bool) {
	defer func() {
		if r := recover(); r != nil {
			log.Error("Recovered in f", r)
			res = false
		}
	}()

	if s == nil {
		return false
	}
	var gpubs []*g1pubs.PublicKey
	for _, pub := range pubs {
		gpubs = append(gpubs, pub.V())
	}
	aggregatePublicKey := g1pubs.AggregatePublicKeys(gpubs)

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, 0)

	return g1pubs.VerifyWithDomain(bls12_381.ToBytes32(msg), aggregatePublicKey, s.aggregate, bls12_381.ToBytes8(b))
}

func AggregateSignatures(bitmap *big.Int, signs []*Signature) *SignatureAggregate {
	var ss []*g1pubs.Signature
	for _, v := range signs {
		if v == nil {
			continue
		}
		ss = append(ss, v.sign)
	}
	g1signs := g1pubs.AggregateSignatures(ss)

	return &SignatureAggregate{
		bitmap:    bitmap,
		aggregate: g1signs,
	}
}
