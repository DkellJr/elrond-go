package factory

import (
	"fmt"

	"github.com/ElrondNetwork/elrond-go/config"
	"github.com/ElrondNetwork/elrond-go/core/check"
	"github.com/ElrondNetwork/elrond-go/data/state"
	factoryState "github.com/ElrondNetwork/elrond-go/data/state/factory"
	"github.com/ElrondNetwork/elrond-go/data/trie/factory"
	"github.com/ElrondNetwork/elrond-go/sharding"
)

//TODO: merge this with data components

// StateComponentsFactoryArgs holds the arguments needed for creating a state components factory
type StateComponentsFactoryArgs struct {
	Config           config.Config
	GenesisConfig    *sharding.Genesis
	ShardCoordinator sharding.Coordinator
	Core             CoreComponentsHolder
	Tries            *TriesComponents
}

type stateComponentsFactory struct {
	config           config.Config
	genesisConfig    *sharding.Genesis
	shardCoordinator sharding.Coordinator
	core             CoreComponentsHolder
	tries            *TriesComponents
}

// NewStateComponentsFactory will return a new instance of stateComponentsFactory
func NewStateComponentsFactory(args StateComponentsFactoryArgs) (*stateComponentsFactory, error) {
	if args.GenesisConfig == nil {
		return nil, ErrNilGenesisConfiguration
	}
	if args.Core == nil {
		return nil, ErrNilCoreComponents
	}
	if check.IfNil(args.Core.PathHandler()) {
		return nil, ErrNilPathManager
	}
	if args.Tries == nil {
		return nil, ErrNilTriesComponents
	}
	if check.IfNil(args.ShardCoordinator) {
		return nil, ErrNilShardCoordinator
	}

	return &stateComponentsFactory{
		config:           args.Config,
		genesisConfig:    args.GenesisConfig,
		core:             args.Core,
		tries:            args.Tries,
		shardCoordinator: args.ShardCoordinator,
	}, nil
}

// Create creates the state components
func (scf *stateComponentsFactory) Create() (*StateComponents, error) {
	processPubkeyConverter, err := factoryState.NewPubkeyConverter(scf.config.AddressPubkeyConverter)
	if err != nil {
		return nil, fmt.Errorf("%w for ProcessPubkeyConverter: %s", ErrPubKeyConverterCreation, err.Error())
	}

	validatorPubkeyConverter, err := factoryState.NewPubkeyConverter(scf.config.ValidatorPubkeyConverter)
	if err != nil {
		return nil, fmt.Errorf("%w for ValidatorPubkeyConverter: %s", ErrPubKeyConverterCreation, err.Error())
	}

	accountFactory := factoryState.NewAccountCreator()
	merkleTrie := scf.tries.TriesContainer.Get([]byte(factory.UserAccountTrie))
	accountsAdapter, err := state.NewAccountsDB(merkleTrie, scf.core.Hasher(), scf.core.InternalMarshalizer(), accountFactory)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrAccountsAdapterCreation, err.Error())
	}

	inBalanceForShard, err := scf.genesisConfig.InitialNodesBalances(scf.shardCoordinator)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInitialBalancesCreation, err.Error())
	}

	accountFactory = factoryState.NewPeerAccountCreator()
	merkleTrie = scf.tries.TriesContainer.Get([]byte(factory.PeerAccountTrie))
	peerAdapter, err := state.NewPeerAccountsDB(merkleTrie, scf.core.Hasher(), scf.core.InternalMarshalizer(), accountFactory)
	if err != nil {
		return nil, err
	}

	return &StateComponents{
		PeerAccounts:             peerAdapter,
		AddressPubkeyConverter:   processPubkeyConverter,
		ValidatorPubkeyConverter: validatorPubkeyConverter,
		AccountsAdapter:          accountsAdapter,
		InBalanceForShard:        inBalanceForShard,
	}, nil
}
