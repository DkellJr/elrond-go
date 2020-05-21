package mock

import (
	"sync"

	"github.com/ElrondNetwork/elrond-go/core"
	"github.com/ElrondNetwork/elrond-go/data/state"
	"github.com/ElrondNetwork/elrond-go/data/typeConverters"
	"github.com/ElrondNetwork/elrond-go/hashing"
	"github.com/ElrondNetwork/elrond-go/marshal"
	"github.com/ElrondNetwork/elrond-go/storage"
)

// CoreComponentsMock -
type CoreComponentsMock struct {
	IntMarsh            marshal.Marshalizer
	TxMarsh             marshal.Marshalizer
	VmMarsh             marshal.Marshalizer
	Hash                hashing.Hasher
	UInt64ByteSliceConv typeConverters.Uint64ByteSliceConverter
	AddrPubKeyConv      state.PubkeyConverter
	ValPubKeyConv       state.PubkeyConverter
	StatusHdl           core.AppStatusHandler
	mutStatus           sync.RWMutex
	PathHdl             storage.PathManagerHandler
	ChainIdCalled       func() string
	mutIntMarshalizer   sync.RWMutex
}

// InternalMarshalizer -
func (ccm *CoreComponentsMock) InternalMarshalizer() marshal.Marshalizer {
	ccm.mutIntMarshalizer.RLock()
	defer ccm.mutIntMarshalizer.RUnlock()

	return ccm.IntMarsh
}

// SetInternalMarshalizer -
func (ccm *CoreComponentsMock) SetInternalMarshalizer(m marshal.Marshalizer) error {
	ccm.mutIntMarshalizer.Lock()
	ccm.IntMarsh = m
	ccm.mutIntMarshalizer.Unlock()

	return nil
}

// TxMarshalizer -
func (ccm *CoreComponentsMock) TxMarshalizer() marshal.Marshalizer {
	return ccm.TxMarsh
}

// VmMarshalizer -
func (ccm *CoreComponentsMock) VmMarshalizer() marshal.Marshalizer {
	return ccm.VmMarsh
}

// Hasher -
func (ccm *CoreComponentsMock) Hasher() hashing.Hasher {
	return ccm.Hash
}

// Uint64ByteSliceConverter -
func (ccm *CoreComponentsMock) Uint64ByteSliceConverter() typeConverters.Uint64ByteSliceConverter {
	return ccm.UInt64ByteSliceConv
}

// AddressPubKeyConverter -
func (ccm *CoreComponentsMock) AddressPubKeyConverter() state.PubkeyConverter {
	return ccm.AddrPubKeyConv
}

// ValidatorPubKeyConverter -
func (ccm *CoreComponentsMock) ValidatorPubKeyConverter() state.PubkeyConverter {
	return ccm.ValPubKeyConv
}

// StatusHandler -
func (ccm *CoreComponentsMock) StatusHandler() core.AppStatusHandler {
	ccm.mutStatus.RLock()
	defer ccm.mutStatus.RUnlock()

	return ccm.StatusHdl
}

// SetStatusHandler -
func (ccm *CoreComponentsMock) SetStatusHandler(statusHandler core.AppStatusHandler) error {
	ccm.mutStatus.Lock()
	ccm.StatusHdl = statusHandler
	ccm.mutStatus.Unlock()

	return nil
}

// PathHandler -
func (ccm *CoreComponentsMock) PathHandler() storage.PathManagerHandler {
	return ccm.PathHdl
}

// ChainID -
func (ccm *CoreComponentsMock) ChainID() string {
	if ccm.ChainIdCalled != nil {
		return ccm.ChainIdCalled()
	}
	return "undefined"
}

// IsInterfaceNil -
func (ccm *CoreComponentsMock) IsInterfaceNil() bool {
	return ccm == nil
}
