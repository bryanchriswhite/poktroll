package main

/*
#include <client.h>
#include <stdlib.h>
typedef void* (*callback_type)(void*, char**);
//void bridge_callback(callback_fn *cb, void* data, char** err);
static void bridge_callback(callback_fn *cb, void* data, char** err) {
    if (cb) {
        cb(data, err);
    }
}
*/
import "C"
import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	bankv1beta1 "cosmossdk.io/api/cosmos/bank/v1beta1"
	"cosmossdk.io/depinject"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	cosmosproto "github.com/cosmos/gogoproto/proto"
	"github.com/gogo/protobuf/proto"
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
func NewTxClient(depsRef C.go_ref, signingKeyName *C.char, cErr **C.char) C.go_ref {
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

	return C.go_ref(SetGoMem(txClient))
}

//export WithSigningKeyName
func WithSigningKeyName(keyName *C.char) C.go_ref {
	return C.go_ref(SetGoMem(tx.WithSigningKeyName(C.GoString(keyName))))
}

// TODO_IN_THIS_COMMIT: godoc...
// TODO_IMPROVE: support multiple msgs (if top-level JSON array).
//
//export TxClient_SignAndBroadcastAny
func TxClient_SignAndBroadcastAny(txClientRef C.go_ref, msgAnyJSON *C.char, cErr **C.char) C.go_ref {
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

	return C.go_ref(SetGoMem(errCh))
}

//export TxClient_SignAndBroadcast
func TxClient_SignAndBroadcast(
	txClientRef C.go_ref,
	msgBz *C.uint8_t,
	msgBzLen C.int,
	typeUrl *C.char,
	callbackFn *C.callback_fn,
	cErr **C.char,
) C.go_ref {
	// TODO_IN_THIS_COMMIT: design a consise way to register all message types.
	// TODO_CONSIDERATION: expose a method to add message types just in case.
	appStakeMsg := new(apptypes.MsgStakeApplication)
	msgTypeUrl := cosmostypes.MsgTypeURL(appStakeMsg)
	fmt.Printf(">>>> typeUrl: %s\n", msgTypeUrl)
	trimmedMsgTypeUrl := strings.Trim(msgTypeUrl, "/")
	fmt.Printf(">>>> trimmedMsgTypeUrl: %s\n", trimmedMsgTypeUrl)
	fmt.Printf(">>>> trimmedMsgTypeUrl == typeUrl: %v\n", trimmedMsgTypeUrl == string(C.GoString(typeUrl)))
	proto.RegisterType(appStakeMsg, trimmedMsgTypeUrl)

	//msgType, err := protoregistry.GlobalTypes.FindMessageByURL(C.GoString(typeUrl))
	// TODO_IN_THIS_COMMIT: double-check typeUrl - may need to convert to non-protoreflect type.
	fmt.Printf(">>>> typeUrl: %s\n", string(C.GoString(typeUrl)))
	msgType := proto.MessageType(C.GoString(typeUrl))
	if msgType == nil {
		*cErr = C.CString(fmt.Sprintf("unknown message type: %s", string(C.GoString(typeUrl))))
		return CNilRef
	}

	// TODO_CONSIDERATION: Could support a version of methods which receive a go context, created elsewhere..
	ctx := context.Background()

	txClient, err := GetGoMem[client.TxClient](txClientRef)
	if err != nil {
		*cErr = C.CString(err.Error())
		return CNilRef
	}

	//msg := msgType.New().Interface()
	msg := reflect.New(msgType.Elem()).Interface().(cosmosproto.Message)
	if err = cdc.Unmarshal(C.GoBytes(unsafe.Pointer(msgBz), msgBzLen), msg); err != nil {
		*cErr = C.CString(err.Error())
		return CNilRef
	}

	eitherAsyncErr := txClient.SignAndBroadcast(ctx, msg)
	err, errCh := eitherAsyncErr.SyncOrAsyncError()
	if err != nil {
		*cErr = C.CString(err.Error())
		return CNilRef
	}

	// TODO_IN_THIS_COMMIT: factor out & comment...
	go func() {
		if err = <-errCh; err != nil {
			//fmt.Println(">>> err ch returned err")
			*cErr = C.CString(err.Error())
			C.bridge_callback(callbackFn, nil, cErr)
		} else {
			//fmt.Println(">>> err ch closed")
			C.bridge_callback(callbackFn, nil, cErr)
		}
	}()

	return C.go_ref(SetGoMem(errCh))
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
