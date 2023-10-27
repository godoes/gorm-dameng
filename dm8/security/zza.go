/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */

package security

import (
	"crypto/rand"
	"errors"
	"io"
	"math/big"
)

type dhGroup struct {
	p *big.Int
	g *big.Int
}

func newDhGroup(prime, generator *big.Int) *dhGroup {
	return &dhGroup{
		p: prime,
		g: generator,
	}
}

func (dg *dhGroup) P() *big.Int {
	p := new(big.Int)
	p.Set(dg.p)
	return p
}

func (dg *dhGroup) G() *big.Int {
	g := new(big.Int)
	g.Set(dg.g)
	return g
}

// GeneratePrivateKey 生成本地公私钥
func (dg *dhGroup) GeneratePrivateKey(randReader io.Reader) (key *DhKey, err error) {
	if randReader == nil {
		randReader = rand.Reader
	}
	// 0 < x < p
	x, err := rand.Int(randReader, dg.p)
	if err != nil {
		return
	}
	zero := big.NewInt(0)
	for x.Cmp(zero) == 0 {
		x, err = rand.Int(randReader, dg.p)
		if err != nil {
			return
		}
	}
	key = new(DhKey)
	key.x = x

	// y = g ^ x mod p
	key.y = new(big.Int).Exp(dg.g, x, dg.p)
	key.group = dg
	return
}

func (dg *dhGroup) ComputeKey(publicKey *DhKey, privateKey *DhKey) (kye *DhKey, err error) {
	if dg.p == nil {
		err = errors.New("DH: invalid group")
		return
	}
	if publicKey.y == nil {
		err = errors.New("DH: invalid public key")
		return
	}
	if publicKey.y.Sign() <= 0 || publicKey.y.Cmp(dg.p) >= 0 {
		err = errors.New("DH parameter out of bounds")
		return
	}
	if privateKey.x == nil {
		err = errors.New("DH: invalid private key")
		return
	}
	k := new(big.Int).Exp(publicKey.y, privateKey.x, dg.p)
	key := new(DhKey)
	key.y = k
	key.group = dg
	return
}
