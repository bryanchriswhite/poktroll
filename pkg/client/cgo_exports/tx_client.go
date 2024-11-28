package main

// #include <client.h>
import "C"
import (
	"context"
	"fmt"

	bankv1beta1 "cosmossdk.io/api/cosmos/bank/v1beta1"
	"cosmossdk.io/depinject"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	cosmosproto "github.com/cosmos/gogoproto/proto"

	//"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/pokt-network/poktroll/api/poktroll/application"
	"github.com/pokt-network/poktroll/pkg/client"
	"github.com/pokt-network/poktroll/pkg/client/tx"
	apptypes "github.com/pokt-network/poktroll/x/application/types"
	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
)

var (
	interfaceRegistry = codectypes.NewInterfaceRegistry()
	cdc               = codec.NewProtoCodec(interfaceRegistry)
)

// TODO_IN_THIS_COMMIT: godoc...
// TODO_IN_THIS_COMMIT: add seperate constructor which supports options...
//
//export NewTxClient
func NewTxClient(depsRef C.GoRef, signingKeyName *C.char, cErr **C.char) C.GoRef {
	// TODO_CONSIDERATION: Could support a version of methods which receive a go context, created elsewhere..
	ctx := context.Background()

	deps, err := GetGoMem[depinject.Config](depsRef)
	if err != nil {
		*cErr = C.CString(err.Error())
		return CNilRef
	}

	//txCtxRef, err := NewTxContext(tcpURL, cErr)
	//if err != nil {
	//	*cErr = C.CString(err.Error())
	//	return 0
	//}
	//
	//txCtx, err := GetGoMem[client.TxContext](txCtxRef)
	//if err != nil {
	//	*cErr = C.CString(err.Error())
	//	return 0
	//}
	//
	//cfg := depinject.Supply(txCtx)

	keyOpt := tx.WithSigningKeyName(C.GoString(signingKeyName))

	txClient, err := tx.NewTxClient(ctx, deps, keyOpt)
	if err != nil {
		*cErr = C.CString(err.Error())
		return CNilRef
	}

	return C.GoRef(SetGoMem(txClient))
}

//export WithSigningKeyName
func WithSigningKeyName(keyName *C.char) C.GoRef {
	return C.GoRef(SetGoMem(tx.WithSigningKeyName(C.GoString(keyName))))
}

// TODO_IN_THIS_COMMIT: godoc...
// TODO_IMPROVE: support multiple msgs (if top-level JSON array).
// TODO_IMPROVE: support (in a seperate method) proto msg bytes.
//
//export SignAndBroadcast
func SignAndBroadcast(txClientRef C.GoRef, msgAnyJSON *C.char, cErr **C.char) C.GoRef {
	// TODO_CONSIDERATION: Could support a version of methods which receive a go context, created elsewhere..
	ctx := context.Background()

	txClient, err := GetGoMem[client.TxClient](txClientRef)
	if err != nil {
		*cErr = C.CString(err.Error())
		return CNilRef
	}

	msg, err := convertAnyMsgJSON(msgAnyJSON)
	if err != nil {
		*cErr = C.CString(err.Error())
		return CNilRef
	}

	eitherAsyncErr := txClient.SignAndBroadcast(ctx, msg)
	err, errCh := eitherAsyncErr.SyncOrAsyncError()
	if err != nil {
		*cErr = C.CString(err.Error())
		return CNilRef
	}

	return C.GoRef(SetGoMem(errCh))
}

// TODO_IN_THIS_COMMIT: move & godoc...
func convertAnyMsgJSON(msgAnyJSON *C.char) (cosmosproto.Message, error) {
	msgAnyJSONBz := []byte(C.GoString(msgAnyJSON))
	msgAny := new(anypb.Any)
	if err := cdc.UnmarshalJSON(msgAnyJSONBz, msgAny); err != nil {
		return nil, err
	}

	var resultMsg cosmosproto.Message
	switch msgAny.GetTypeUrl() {
	case "type.googleapis.com/cosmos.bank.v1beta1.MsgSend":
		apiMsg := new(bankv1beta1.MsgSend)
		if err := msgAny.UnmarshalTo(apiMsg); err != nil {
			return nil, err
		}

		msg := new(banktypes.MsgSend)

		// TODO_IN_THIS_COMMIT: automate the below via reflection...
		msg.FromAddress = apiMsg.GetFromAddress()
		msg.ToAddress = apiMsg.GetToAddress()

		coins := make(cosmostypes.Coins, len(apiMsg.GetAmount()))
		for i, apiCoin := range apiMsg.GetAmount() {
			coinAmt, ok := math.NewIntFromString(apiCoin.GetAmount())
			if !ok {
				return nil, fmt.Errorf("failed to parse coin amount %q", apiCoin.GetAmount())
			}

			coin := cosmostypes.Coin{
				Denom:  apiCoin.GetDenom(),
				Amount: coinAmt,
			}
			coins[i] = coin
		}
		msg.Amount = coins

		resultMsg = msg
	case "type.googleapis.com/poktroll.application.MsgStakeApplication":
		apiMsg := new(application.MsgStakeApplication)
		if err := msgAny.UnmarshalTo(apiMsg); err != nil {
			return nil, err
		}

		msg := new(apptypes.MsgStakeApplication)

		// TODO_IN_THIS_COMMIT: automate the below via reflection...
		msg.Address = apiMsg.GetAddress()

		stakeAmt, ok := math.NewIntFromString(apiMsg.GetStake().GetAmount())
		if !ok {
			return nil, fmt.Errorf("failed to parse stake amount %q", apiMsg.GetStake().GetAmount())
		}

		stake := &cosmostypes.Coin{
			Denom:  apiMsg.GetStake().GetDenom(),
			Amount: stakeAmt,
		}
		msg.Stake = stake

		services := make([]*sharedtypes.ApplicationServiceConfig, len(apiMsg.GetServices()))
		for i, apiService := range apiMsg.GetServices() {
			service := &sharedtypes.ApplicationServiceConfig{
				ServiceId: apiService.GetServiceId(),
			}
			services[i] = service
		}
		msg.Services = services

		resultMsg = msg
	default:
		panic("unsupported msg type")
	}

	return resultMsg, nil
}
