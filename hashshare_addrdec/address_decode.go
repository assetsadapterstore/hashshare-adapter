package hashshare_addrdec

import (
	"github.com/blocktree/go-owcdrivers/addressEncoder"
)

var (
	alphabet = addressEncoder.ZECAlphabet
)

var (

	HSS_mainnetAddressP2PKH         = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x28}, Suffix: nil}
	HSS_testnetAddressP2PKH         = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x8b}, Suffix: nil}
	HSS_mainnetPrivateWIFCompressed = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "", HashLen: 32, Prefix: []byte{0xd4}, Suffix: []byte{0x01}}
	HSS_testnetPrivateWIFCompressed = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "", HashLen: 32, Prefix: []byte{0xEF}, Suffix: []byte{0x01}}
	HSS_mainnetAddressP2SH          = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x0d}, Suffix: nil}
	HSS_testnetAddressP2SH          = addressEncoder.AddressType{EncodeType: "base58", Alphabet: alphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x13}, Suffix: nil}

	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	IsTestNet bool
}

//AddressDecode 地址解析
func (dec *AddressDecoderV2) AddressDecode(addr string, opts ...interface{}) ([]byte, error) {

	cfg := HSS_mainnetAddressP2PKH
	if dec.IsTestNet {
		cfg = HSS_testnetAddressP2PKH
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	return addressEncoder.AddressDecode(addr, cfg)
}

//AddressEncode 地址编码
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {

	cfg := HSS_mainnetAddressP2PKH
	if dec.IsTestNet {
		cfg = HSS_testnetAddressP2PKH
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	address := addressEncoder.AddressEncode(hash, cfg)
	return address, nil
}
