package mock

import (
	"sync"

	"github.com/ElrondNetwork/elrond-go/crypto"
	"github.com/ElrondNetwork/elrond-go/vm"
)

// CryptoComponentsMock -
type CryptoComponentsMock struct {
	PubKey         crypto.PublicKey
	PrivKey        crypto.PrivateKey
	PubKeyString   string
	PrivKeyBytes   []byte
	PubKeyBytes    []byte
	BlockSig       crypto.SingleSigner
	TxSig          crypto.SingleSigner
	MultiSig       crypto.MultiSigner
	BlKeyGen       crypto.KeyGenerator
	TxKeyGen       crypto.KeyGenerator
	MsgSigVerifier vm.MessageSignVerifier
	mutMultiSig    sync.RWMutex
}

// PublicKey -
func (ccm *CryptoComponentsMock) PublicKey() crypto.PublicKey {
	return ccm.PubKey
}

// PrivateKey -
func (ccm *CryptoComponentsMock) PrivateKey() crypto.PrivateKey {
	return ccm.PrivKey
}

// PublicKeyString -
func (ccm *CryptoComponentsMock) PublicKeyString() string {
	return ccm.PubKeyString
}

// PublicKeyBytes -
func (ccm *CryptoComponentsMock) PublicKeyBytes() []byte {
	return ccm.PubKeyBytes
}

// PrivateKeyBytes -
func (ccm *CryptoComponentsMock) PrivateKeyBytes() []byte {
	return ccm.PrivKeyBytes
}

// BlockSigner -
func (ccm *CryptoComponentsMock) BlockSigner() crypto.SingleSigner {
	return ccm.BlockSig
}

// TxSingleSigner -
func (ccm *CryptoComponentsMock) TxSingleSigner() crypto.SingleSigner {
	return ccm.TxSig
}

// MultiSigner -
func (ccm *CryptoComponentsMock) MultiSigner() crypto.MultiSigner {
	ccm.mutMultiSig.RLock()
	defer ccm.mutMultiSig.RUnlock()

	return ccm.MultiSig
}

// SetMultiSigner -
func (ccm *CryptoComponentsMock) SetMultiSigner(ms crypto.MultiSigner) error {
	ccm.mutMultiSig.Lock()
	ccm.MultiSig = ms
	ccm.mutMultiSig.Unlock()

	return nil
}

// BlockSignKeyGen -
func (ccm *CryptoComponentsMock) BlockSignKeyGen() crypto.KeyGenerator {
	return ccm.BlKeyGen
}

// TxSignKeyGen -
func (ccm *CryptoComponentsMock) TxSignKeyGen() crypto.KeyGenerator {
	return ccm.TxKeyGen
}

// MessageSignVerifier -
func (ccm *CryptoComponentsMock) MessageSignVerifier() vm.MessageSignVerifier {
	return ccm.MsgSigVerifier
}

// IsInterfaceNil -
func (ccm *CryptoComponentsMock) IsInterfaceNil() bool {
	return ccm == nil
}
