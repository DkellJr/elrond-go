package mock

import (
	"github.com/ElrondNetwork/elrond-go/core"
	"github.com/ElrondNetwork/elrond-go/data/state"
)

type oneShardCoordinatorMock struct {
	noShards        uint32
	ComputeIdCalled func(state.AddressContainer) uint32
}

func NewOneShardCoordinatorMock() *oneShardCoordinatorMock {
	return &oneShardCoordinatorMock{noShards: 1}
}

func (scm *oneShardCoordinatorMock) NumberOfShards() uint32 {
	return scm.noShards
}

func (scm *oneShardCoordinatorMock) ComputeId(address state.AddressContainer) uint32 {
	if scm.ComputeIdCalled != nil {
		return scm.ComputeIdCalled(address)
	}

	return uint32(0)
}

func (scm *oneShardCoordinatorMock) SelfId() uint32 {
	return 0
}

func (scm *oneShardCoordinatorMock) SetSelfId(_ uint32) error {
	return nil
}

func (scm *oneShardCoordinatorMock) SameShard(_, _ state.AddressContainer) bool {
	return true
}

func (scm *oneShardCoordinatorMock) CommunicationIdentifier(destShardID uint32) string {
	if destShardID == core.MetachainShardId {
		return "_0_META"
	}

	return "_0"
}

// IsInterfaceNil returns true if there is no value under the interface
func (scm *oneShardCoordinatorMock) IsInterfaceNil() bool {
	return scm == nil
}