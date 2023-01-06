package types

import (
	"errors"
	"fmt"
	"github.com/232425wxy/meta--/crypto"
	"github.com/232425wxy/meta--/crypto/bls12"
	"github.com/232425wxy/meta--/crypto/sha256"
	"github.com/232425wxy/meta--/proto/pbtypes"
	"time"
)

/*⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓*/

// 区块

type SimpleBlock struct {
	Hash []byte `json:"hash"`
}

func (sb *SimpleBlock) ToProto() *pbtypes.SimpleBlock {
	if sb == nil {
		return nil
	}
	return &pbtypes.SimpleBlock{Hash: sb.Hash}
}

func SimpleBlockFromProto(pb *pbtypes.SimpleBlock) *SimpleBlock {
	return &SimpleBlock{Hash: pb.Hash}
}

type Block struct {
	LastBlock SimpleBlock `json:"last_block"` // 上个区块的哈希值
	Header    Header      `json:"header"`
	Data      Data        `json:"data"`
	Decision  Decision    `json:"decision"` // 人们对当前区块的投票决定
}

func (b *Block) ValidateBasic() error {
	if b == nil {
		return errors.New("nil block")
	}
	if err := b.Data.ValidateBasic(); err != nil {
		return err
	}
	return nil
}

func (b *Block) Hash() []byte {
	h := sha256.New()
	h.Write(b.LastBlock.Hash)
	h.Write([]byte(fmt.Sprintf("%d", b.Header.Height)))
	h.Write([]byte(b.Header.Timestamp.String()))
	h.Write([]byte(b.Header.Proposer))
	h.Write(b.Data.RootHash)
	b.Header.Hash = h.Sum(nil)
	return b.Header.Hash
}

func (b *Block) ToProto() *pbtypes.Block {
	// 不包括对当前区块的投票决定
	if b == nil {
		return nil
	}
	pb := &pbtypes.Block{
		LastBlock: b.LastBlock.ToProto(),
		Header:    b.Header.ToProto(),
		Data:      b.Data.ToProto(),
	}
	return pb
}

/*⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓*/

// 区块头

type Header struct {
	Hash      []byte    `json:"hash"` // 当前区块哈希
	Height    int64     `json:"height"`
	Timestamp time.Time `json:"timestamp"`
	Proposer  crypto.ID `json:"proposer"`
}

func (h *Header) ToProto() *pbtypes.Header {
	if h == nil {
		return nil
	}
	return &pbtypes.Header{
		Hash:      h.Hash,
		Height:    h.Height,
		Timestamp: h.Timestamp,
		Proposer:  string(h.Proposer),
	}
}

/*⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓*/

// 区块体

type Data struct {
	RootHash []byte `json:"root_hash"`
	Txs      Txs    `json:"txs"`
}

func (d *Data) ToProto() *pbtypes.Data {
	if d == nil {
		return nil
	}
	_txs := make([][]byte, len(d.Txs))
	for i, tx := range d.Txs {
		copy(_txs[i], tx)
	}
	return &pbtypes.Data{
		RootHash: d.RootHash,
		Txs:      _txs,
	}
}

// ValidateBasic ♏ | 作者 ⇨ 吴翔宇 | (｡･∀･)ﾉﾞ嗨
//
// ValidateBasic 方法验证区块体部分的交易数据大小不能超过1MB。
func (d *Data) ValidateBasic() error {
	size := 0
	for _, tx := range d.Txs {
		size += len(tx)
	}
	if size > 1024*1024 {
		return fmt.Errorf("exceed data limit: %d > %d", size, 1024*1024)
	}
	return nil
}

/*⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓*/

// 共识投票

type Decision struct {
	Signature *bls12.AggregateSignature
}
