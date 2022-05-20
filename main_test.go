package main

import (
	"encoding/hex"
	"fmt"
	"github.com/decred/base58"
	"testing"
)

func Test_Addr(t *testing.T) {
		addressBase58 := "Dcur2mcGjmENx4DhNqDctW5wJCVyT3Qeqkx"
		address := base58.Decode(addressBase58)
		address = address[2:]
		addressHex := hex.EncodeToString(address)
		fmt.Printf("addressHex:             %v\n", addressHex)
		fmt.Printf("addressHex:             %v\n", addressHex[:40])
		// OP_DUP OP_HASH160 OP_DATA_20 addressHex(RIPEMD-160)_remove_checksum OP_EQUALVERIFY OP_CHECKSIG
		fmt.Printf("addressHexScript: 76a914%v88ac\n", addressHex[:40])
}