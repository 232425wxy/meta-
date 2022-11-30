package bls12

import (
	"math/big"
)

/*⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓*/

// 项目级全局常量

const (
	// PrivateKeyFileType ♏ |作者：吴翔宇| 🍁 |日期：2022/11/30|
	//
	// PrivateKeyFileType PEM格式的私钥。
	PrivateKeyFileType = "BLS12-381 PRIVATE KEY"

	// PublicKeyFileType ♏ |作者：吴翔宇| 🍁 |日期：2022/11/30|
	//
	// PublicKeyFileType PEM格式的公钥。
	PublicKeyFileType = "BLS12-381 PUBLIC KEY"
)

/*⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓⛓*/

// 包级全局变量

var domain = []byte("BLS12-381-SIG:REDACTABLE-BLOCKCHAIN")

// curveOrder ♏ | 作者 ⇨ 吴翔宇 | (｡･∀･)ﾉﾞ嗨
//
//	---------------------------------------------------------
//
// curveOrder 椭圆曲线G1的阶。
var curveOrder, _ = new(big.Int).SetString("73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001", 16)
