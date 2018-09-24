package anchors

import (
	"fmt"
	"math/big"

	"context"

	"github.com/centrifuge/go-centrifuge/centrifuge/identity"

	"github.com/centrifuge/go-centrifuge/centrifuge/queue"
	"github.com/centrifuge/gocelery"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/go-errors/errors"
)

const (
	AnchorRepositoryConfirmationTaskName string = "AnchorRepositoryConfirmationTaskName"
	AnchorIDParam                        string = "AnchorIDParam"
	CentrifugeIDParam                    string = "CentrifugeIDParam"
	AddressParam                         string = "AddressParam"
)

type AnchorCommittedWatcher interface {
	WatchAnchorCommitted(opts *bind.WatchOpts, sink chan<- *EthereumAnchorRepositoryContractAnchorCommitted,
		from []common.Address, anchorID []*big.Int, centrifugeId []*big.Int) (event.Subscription, error)
}

// AnchoringConfirmationTask is a queued task to watch ID registration events on Ethereum using EthereumAnchoryRepositoryContract.
// To see how it gets registered see bootstrapper.go and to see how it gets used see setUpRegistrationEventListener method
type AnchoringConfirmationTask struct {
	// task parameters
	From         common.Address
	AnchorID     AnchorID
	CentrifugeID identity.CentID

	// state
	EthContextInitializer  func() (ctx context.Context, cancelFunc context.CancelFunc)
	AnchorRegisteredEvents chan *EthereumAnchorRepositoryContractAnchorCommitted
	EthContext             context.Context
	AnchorCommittedWatcher AnchorCommittedWatcher
}

func NewAnchoringConfirmationTask(
	anchorCommittedWatcher AnchorCommittedWatcher,
	ethContextInitializer func() (ctx context.Context, cancelFunc context.CancelFunc),
) *AnchoringConfirmationTask {
	return &AnchoringConfirmationTask{
		AnchorCommittedWatcher: anchorCommittedWatcher,
		EthContextInitializer:  ethContextInitializer,
	}
}

func (act *AnchoringConfirmationTask) Name() string {
	return AnchorRepositoryConfirmationTaskName
}

func (act *AnchoringConfirmationTask) Init() error {
	queue.Queue.Register(act.Name(), act)
	return nil
}

func (act *AnchoringConfirmationTask) Copy() (gocelery.CeleryTask, error) {
	return &AnchoringConfirmationTask{
		act.From,
		act.AnchorID,
		act.CentrifugeID,
		act.EthContextInitializer,
		act.AnchorRegisteredEvents,
		act.EthContext,
		act.AnchorCommittedWatcher,
	}, nil
}

// ParseKwargs - define a method to parse AnchorID, Address and RootHash
func (act *AnchoringConfirmationTask) ParseKwargs(kwargs map[string]interface{}) error {
	anchorID, ok := kwargs[AnchorIDParam]
	if !ok {
		return fmt.Errorf("undefined kwarg " + AnchorIDParam)
	}

	anchorIDBytes, err := getBytesAnchorID(anchorID)

	if err != nil {
		return fmt.Errorf("malformed kwarg [%s] because [%s]", AnchorIDParam, err.Error())
	}

	act.AnchorID = anchorIDBytes

	//parse the centrifuge id
	centrifugeId, ok := kwargs[CentrifugeIDParam]
	if !ok {
		return fmt.Errorf("undefined kwarg " + CentrifugeIDParam)
	}

	centrifugeIdBytes, err := getBytesCentrifugeID(centrifugeId)
	if err != nil {
		return fmt.Errorf("malformed kwarg [%s] because [%s]", CentrifugeIDParam, err.Error())
	}

	act.CentrifugeID = centrifugeIdBytes

	// parse the address
	address, ok := kwargs[AddressParam]
	if !ok {
		return fmt.Errorf("undefined kwarg " + AddressParam)
	}
	addressStr, ok := address.(string)
	if !ok {
		return fmt.Errorf("param is not hex string " + AddressParam)
	}
	addressTyped, err := getAddressFromHexString(addressStr)
	if err != nil {
		return fmt.Errorf("malformed kwarg [%s] because [%s]", AddressParam, err.Error())
	}
	act.From = addressTyped
	return nil
}

// RunTask calls listens to events from geth related to AnchoringConfirmationTask#AnchorID and records result.
func (act *AnchoringConfirmationTask) RunTask() (interface{}, error) {
	log.Infof("Waiting for confirmation for the anchorID [%x]", act.AnchorID)
	if act.EthContext == nil {
		act.EthContext, _ = act.EthContextInitializer()
	}
	watchOpts := &bind.WatchOpts{Context: act.EthContext}
	if act.AnchorRegisteredEvents == nil {
		act.AnchorRegisteredEvents = make(chan *EthereumAnchorRepositoryContractAnchorCommitted)
	}

	subscription, err := act.AnchorCommittedWatcher.WatchAnchorCommitted(watchOpts, act.AnchorRegisteredEvents,
		[]common.Address{act.From}, []*big.Int{act.AnchorID.toBigInt()},
		[]*big.Int{act.CentrifugeID.BigInt()})

	if err != nil {
		wError := errors.WrapPrefix(err, "Could not subscribe to event logs for anchor repository", 1)
		log.Errorf(wError.Error())
		return nil, wError
	}
	for {
		select {
		case err := <-subscription.Err():
			log.Errorf("Subscription error %s for anchor ID: %x\n", err.Error(), act.AnchorID)
			return nil, err
		case <-act.EthContext.Done():
			log.Errorf("Context closed before receiving AnchorRegistered event for anchor ID: %x\n", act.AnchorID)
			return nil, act.EthContext.Err()
		case res := <-act.AnchorRegisteredEvents:
			log.Infof("Received AnchorRegistered event from: %x, identifier: %x\n", res.From, res.AnchorId)
			subscription.Unsubscribe()
			return res, nil
		}
	}
}

func getBytesAnchorID(key interface{}) (AnchorID, error) {
	var fixed [AnchorIDLength]byte
	b, ok := key.([]interface{})
	if !ok {
		return fixed, errors.New("Could not parse interface to []byte")
	}
	// convert and copy b byte values
	for i, v := range b {
		fv := v.(float64)
		fixed[i] = byte(fv)
	}
	return fixed, nil
}

func getBytesCentrifugeID(key interface{}) (identity.CentID, error) {
	var fixed [identity.CentIDByteLength]byte
	b, ok := key.([]interface{})
	if !ok {
		return fixed, errors.New("Could not parse interface to []byte")
	}
	// convert and copy b byte values
	for i, v := range b {
		fv := v.(float64)
		fixed[i] = byte(fv)
	}
	return fixed, nil
}

func getAddressFromHexString(hex string) (common.Address, error) {
	return common.BytesToAddress(common.FromHex(hex)), nil
}