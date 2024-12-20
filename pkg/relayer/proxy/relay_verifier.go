package proxy

import (
	"context"

	servicetypes "github.com/pokt-network/poktroll/x/service/types"
	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
)

// VerifyRelayRequest is a shared method used by RelayServers to check the relay
// request signature and session validity.
func (rp *relayerProxy) VerifyRelayRequest(
	ctx context.Context,
	relayRequest *servicetypes.RelayRequest,
	supplierServiceId string,
) error {
	// Get the block height at which the relayRequest should be processed.
	// Check if the relayRequest is on time or within the session's grace period
	// before attempting to verify the relayRequest signature.
	sessionBlockHeight, err := rp.getTargetSessionBlockHeight(ctx, relayRequest)
	if err != nil {
		return err
	}

	// Verify the relayRequest metadata, signature, session header and other
	// basic validation.
	if err = rp.ringCache.VerifyRelayRequestSignature(ctx, relayRequest); err != nil {
		return err
	}

	meta := relayRequest.GetMeta()

	// Extract the session header for usage below.
	// ringCache.VerifyRelayRequestSignature already verified the header's validaity.
	sessionHeader := meta.SessionHeader

	// Application address is used to verify the relayRequest signature.
	// It is guaranteed to be present in the relayRequest since the signature
	// has already been verified.
	appAddress := sessionHeader.GetApplicationAddress()

	rp.logger.Debug().
		Fields(map[string]any{
			"session_id":                sessionHeader.GetSessionId(),
			"application_address":       appAddress,
			"service_id":                sessionHeader.GetServiceId(),
			"supplier_operator_address": meta.GetSupplierOperatorAddress(),
		}).
		Msg("verifying relay request session")

	// TODO_TECHDEBT(@red-0ne): Optimize this so we don't have to query the session for every relay request.
	// Query for the current session to check if relayRequest sessionId matches the current session.
	session, err := rp.sessionQuerier.GetSession(
		ctx,
		appAddress,
		supplierServiceId,
		sessionBlockHeight,
	)
	if err != nil {
		return err
	}

	// Session validity can be checked via a basic ID comparison due to the reasons below.
	//
	// Since the retrieved sessionId was in terms of:
	// - the current block height and sessionGracePeriod (which are not provided by the relayRequest)
	// - serviceId (which is not provided by the relayRequest)
	// - applicationAddress (which is used to verify the relayRequest signature)
	if session.SessionId != sessionHeader.GetSessionId() {
		return ErrRelayerProxyInvalidSession.Wrapf(
			"session ID mismatch, expecting: %+v, got: %+v",
			session.GetSessionId(),
			relayRequest.Meta.GetSessionHeader().GetSessionId(),
		)
	}

	// Check if the relayRequest is allowed to be served by the relayer proxy.
	_, isSupplierOperatorAddressPresent := rp.OperatorAddressToSigningKeyNameMap[meta.GetSupplierOperatorAddress()]
	if !isSupplierOperatorAddressPresent {
		return ErrRelayerProxyMissingSupplierOperatorAddress
	}

	for _, supplier := range session.Suppliers {
		// Verify if the supplier operator address in the session matches the one in the relayRequest.
		if isSupplierOperatorAddressPresent && supplier.OperatorAddress == meta.GetSupplierOperatorAddress() {
			return nil
		}
	}

	return ErrRelayerProxyInvalidSupplier
}

// getTargetSessionBlockHeight returns the block height at which the session
// for the given relayRequest should be processed. If the session is within the
// grace period, the session's end block height is returned. Otherwise,
// the current block height is returned.
// If the session has expired, then return an error.
func (rp *relayerProxy) getTargetSessionBlockHeight(
	ctx context.Context,
	relayRequest *servicetypes.RelayRequest,
) (sessionHeight int64, err error) {
	currentHeight := rp.blockClient.LastBlock(ctx).Height()
	sessionEndHeight := relayRequest.Meta.SessionHeader.GetSessionEndBlockHeight()

	// TODO_MAINNET(#543): We don't really want to have to query the params for every method call.
	// Once `ModuleParamsClient` is implemented, use its replay observable's `#Last()` method
	// to get the most recently (asynchronously) observed (and cached) value.
	// TODO_MAINNET(@bryanchriswhite, #543): We also don't really want to use the current value of the params.
	// Instead, we should be using the value that the params had for the session given by sessionEndHeight.
	sharedParams, err := rp.sharedQuerier.GetParams(ctx)
	if err != nil {
		return 0, err
	}

	// Check if the RelayRequest's session has expired.
	if sessionEndHeight < currentHeight {
		// Do not process the `RelayRequest` if the session has expired and the current
		// block height is outside the session's grace period.
		if !sharedtypes.IsGracePeriodElapsed(sharedParams, sessionEndHeight, currentHeight) {
			// The RelayRequest's session has expired but is still within the
			// grace period so process it as if the session is still active.
			return sessionEndHeight, nil
		}

		return 0, ErrRelayerProxyInvalidSession.Wrapf(
			"session expired, expecting: %d, got: %d",
			sessionEndHeight,
			currentHeight,
		)
	}

	// The RelayRequest's session is active so return the current block height.
	return currentHeight, nil
}
