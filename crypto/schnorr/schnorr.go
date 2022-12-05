package schnorr

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"github.com/232425wxy/meta--/crypto/hash/sha256"
	"math/big"
)

/*⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓*/

// API 项目级全局函数

func GenerateKey() (*PublicKey, *PrivateKey) {
	var err error
	privateKey := new(PrivateKey)
	publicKey := new(PublicKey)

	privateKey.key = randGen(p)
	if err != nil {
		panic(fmt.Sprintf("Schnorr: failed to generate private key: %q", err))
	}
	publicKey.key = new(big.Int).Exp(g, privateKey.key, p)
	return publicKey, privateKey
}

/*⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓*/

// 定义公私钥

// PublicKey ♏ | 作者 ⇨ 吴翔宇 | (｡･∀･)ﾉﾞ嗨
//
//	---------------------------------------------------------
//
// PublicKey
type PublicKey struct {
	key *big.Int
}

// PrivateKey ♏ | 作者 ⇨ 吴翔宇 | (｡･∀･)ﾉﾞ嗨
//
//	---------------------------------------------------------
//
// PrivateKey
type PrivateKey struct {
	key *big.Int
}

// Signature ♏ | 作者 ⇨ 吴翔宇 | (｡･∀･)ﾉﾞ嗨
//
//	---------------------------------------------------------
//
// Signature
type Signature struct {
	sig     *big.Int
	sum     []byte
	message []byte
}

// Sign ♏ | 作者 ⇨ 吴翔宇 | (｡･∀･)ﾉﾞ嗨
//
//	---------------------------------------------------------
//
// Sign 签名算法，给定的输入有：私钥sk，消息msg，哈希函数H。现在依次执行以下步骤来计算签名：
//  1. 选取随机值k mod p，计算：K=g**k；
//  2. 计算哈希值：sum=H(K,msg)
//  3. 用私钥计算签名：sig=k-sk*sum
//  4. 组装签名：<sig, sum, msg>
func (key *PrivateKey) Sign(message []byte) (*Signature, error) {
	k := randGen(p)
	K := new(big.Int).Exp(g, k, p)
	sum := hash(K.Bytes(), message)
	e := new(big.Int).SetBytes(sum)
	xe := new(big.Int).Mul(e, key.key)
	sig := e.Sub(k, xe)
	return &Signature{sig: sig, sum: sum, message: message}, nil
}

// Verify ♏ | 作者 ⇨ 吴翔宇 | (｡･∀･)ﾉﾞ嗨
//
//	---------------------------------------------------------
//
// Verify 验签算法，给定的输入有：公钥pk，签名<sig, sum, msg>。现在依次执行以下步骤来验证签名：
//  1. 解析签名，计算：x=pk**sum
//  2. 计算：y=g**sig
//  3. 计算：z=x*y，然后求哈希值sum'=H(z,msg)，比较sum'和sum一不一样。
func (key *PublicKey) Verify(signature *Signature) bool {
	sum_ := new(big.Int).SetBytes(signature.sum)
	g_sig := new(big.Int).Exp(g, signature.sig, p)
	ye := new(big.Int).Exp(key.key, sum_, p)
	rv := new(big.Int).Mul(g_sig, ye)
	rv.Mod(rv, p)
	ha := hash(rv.Bytes(), signature.message)
	return bytes.Equal(ha, signature.sum)
}

/*⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓*/

// 定义群的生成元和阶：群的阶是256比特长

var p, _ = new(big.Int).SetString("c0a0f171e583d4efb262c1783a6ed6d995fc1d4eea476149cf8ea40078d27ad7", 16)

var q = new(big.Int).Div(new(big.Int).Sub(p, new(big.Int).SetInt64(1)), new(big.Int).SetInt64(2))

var g, _ = new(big.Int).SetString("3e446bf8e43afebc6b49bc7220d19a415c9bda8cbbcc25189c67d27b33df73cf", 16)

/*⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓*/

// 定义不可导出的工具函数

func randGen(upper *big.Int) *big.Int {
	randomBig, err := rand.Int(rand.Reader, upper)
	if err != nil {
		panic(err)
	}
	return randomBig
}

func hash(bzs ...[]byte) []byte {
	h := sha256.New()
	for _, bz := range bzs {
		h.Write(bz)
	}
	return h.Sum(nil)
}
