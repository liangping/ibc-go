package types

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
)

// define constants used for testing
const (
	validPort        = "testportid"
	invalidPort      = "(invalidport1)"
	invalidShortPort = "p"
	// 195 characters
	invalidLongPort = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis eros neque, ultricies vel ligula ac, convallis porttitor elit. Maecenas tincidunt turpis elit, vel faucibus nisl pellentesque sodales"

	validChannel        = "testchannel"
	invalidChannel      = "(invalidchannel1)"
	invalidShortChannel = "invalid"
	invalidLongChannel  = "invalidlongchannelinvalidlongchannelinvalidlongchannelinvalidlongchannel"
)

var (
	addr1        = sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String()
	addr2        = sdk.AccAddress("testaddr2").String()
	counterparty = "osmo1jxv0u20scum4trha72c7ltfgfqef6nsc638lza"
	emptyAddr    string
	sendCoin     = sdk.NewCoin("uatom", sdk.NewInt(1000000))
	demandCoin   = sdk.NewCoin("uosmo", sdk.NewInt(5000000))

	// ibcCoin          = sdk.NewCoin("ibc/7F1D3FCF4AE79E1554D670D1AD949A9BA4E4A3C76C63093E17E446A46061A7A2", sdk.NewInt(100))
	// invalidIBCCoin   = sdk.NewCoin("ibc/7F1D3FCF4AE79E1554", sdk.NewInt(100))
	// invalidDenomCoin = sdk.Coin{Denom: "0atom", Amount: sdk.NewInt(100)}
	// zeroCoin         = sdk.Coin{Denom: "atoms", Amount: sdk.NewInt(0)}

	timeoutHeight = clienttypes.NewHeight(0, 10)
)

// TestMsgTransferRoute tests Route for MsgTransfer
func TestMsgSwapRoute(t *testing.T) {
	msg := NewMsgSwap(validPort, validChannel, sendCoin, demandCoin, addr1, addr2, counterparty, timeoutHeight, 0)

	require.Equal(t, RouterKey, msg.Route())
}

// TestMsgTransferType tests Type for MsgTransfer
func TestMsgTransferType(t *testing.T) {
	msg := NewMsgSwap(validPort, validChannel, sendCoin, demandCoin, addr1, addr2, counterparty, timeoutHeight, 0)

	require.Equal(t, "swap", msg.Type())
}

func TestMsgSwapGetSignBytes(t *testing.T) {
	msg := NewMsgSwap(validPort, validChannel, sendCoin, demandCoin, addr1, addr2, counterparty, timeoutHeight, 0)
	expected := fmt.Sprintf(`{"type":"cosmos-sdk/MsgSwap","value":{"receiver":"%s","sender":"%s","source_channel":"testchannel","source_port":"testportid","timeout_height":{"revision_height":"10"},"token":{"amount":"100","denom":"atom"}}}`, addr2, addr1)
	require.NotPanics(t, func() {
		res := msg.GetSignBytes()
		require.Equal(t, expected, string(res))
	})
}

// TestMsgTransferValidation tests ValidateBasic for MsgTransfer
func TestMsgSwapValidation(t *testing.T) {
	testCases := []struct {
		name    string
		msg     *MsgSwap
		expPass bool
	}{
		{"valid msg with base denom", NewMsgSwap(validPort, validChannel, sendCoin, demandCoin, addr1, addr2, counterparty, timeoutHeight, 0), true},
		// {"valid msg with trace hash", NewMsgSwap(validPort, validChannel, ibcCoin, receiveCoin, addr1, addr2, counterparty, timeoutHeight, 0), true},
		// {"invalid ibc denom", NewMsgSwap(validPort, validChannel, invalidIBCCoin, addr1, addr2, counterparty, timeoutHeight, 0), false},
		// {"too short port id", NewMsgSwap(invalidShortPort, validChannel, coin, addr1, addr2, counterparty, timeoutHeight, 0), false},
		// {"too long port id", NewMsgSwap(invalidLongPort, validChannel, coin, addr1, addr2, counterparty, timeoutHeight, 0), false},
		// {"port id contains non-alpha", NewMsgSwap(invalidPort, validChannel, coin, addr1, addr2, counterparty, timeoutHeight, 0), false},
		// {"too short channel id", NewMsgSwap(validPort, invalidShortChannel, coin, addr1, addr2, counterparty, timeoutHeight, 0), false},
		// {"too long channel id", NewMsgSwap(validPort, invalidLongChannel, coin, addr1, addr2, counterparty, timeoutHeight, 0), false},
		// {"channel id contains non-alpha", NewMsgSwap(validPort, invalidChannel, coin, addr1, addr2, counterparty, timeoutHeight, 0), false},
		// {"invalid denom", NewMsgSwap(validPort, validChannel, invalidDenomCoin, addr1, addr2, counterparty, timeoutHeight, 0), false},
		// {"zero coin", NewMsgSwap(validPort, validChannel, zeroCoin, addr1, addr2, counterparty, timeoutHeight, 0), false},
		// {"missing sender address", NewMsgSwap(validPort, validChannel, coin, emptyAddr, addr2, counterparty, timeoutHeight, 0), false},
		// {"missing recipient address", NewMsgSwap(validPort, validChannel, coin, addr1, "", counterparty, timeoutHeight, 0), false},
		// {"empty coin", NewMsgSwap(validPort, validChannel, sdk.Coin{}, addr1, addr2, counterparty, timeoutHeight, 0), false},
	}

	for i, tc := range testCases {
		err := tc.msg.ValidateBasic()
		if tc.expPass {
			require.NoError(t, err, "valid test case %d failed: %s", i, tc.name)
		} else {
			require.Error(t, err, "invalid test case %d passed: %s", i, tc.name)
		}
	}
}

// TestMsgTransferGetSigners tests GetSigners for MsgTransfer
func TestMsgTransferGetSigners(t *testing.T) {
	addr := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())

	msg := NewMsgSwap(validPort, validChannel, sendCoin, demandCoin, addr.String(), addr2, counterparty, timeoutHeight, 0)
	res := msg.GetSigners()

	require.Equal(t, []sdk.AccAddress{addr}, res)
}
