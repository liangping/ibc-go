package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	denom              = "atom"
	amount             = "100"
	largeAmount        = "18446744073709551616"                                                           // one greater than largest uint64 (^uint64(0))
	invalidLargeAmount = "115792089237316195423570985008687907853269984665640564039457584007913129639936" // 2^256
)

// TestFungibleTokenPacketDataValidateBasic tests ValidateBasic for FungibleTokenPacketData
func TestFungibleTokenPacketDataValidateBasic(t *testing.T) {
	testCases := []struct {
		name       string
		packetData FungibleTokenSwapPacketData
		expPass    bool
	}{
		// {"valid packet", NewFungibleTokenSwapPacketData(denom, amount, addr1, addr2), true},
		// {"valid packet with large amount", NewFungibleTokenSwapPacketData(denom, largeAmount, addr1, addr2), true},
		// {"invalid denom", NewFungibleTokenSwapPacketData("", amount, addr1, addr2), false},
		// {"invalid empty amount", NewFungibleTokenSwapPacketData(denom, "", addr1, addr2), false},
		// {"invalid zero amount", NewFungibleTokenSwapPacketData(denom, "0", addr1, addr2), false},
		// {"invalid negative amount", NewFungibleTokenSwapPacketData(denom, "-1", addr1, addr2), false},
		// {"invalid large amount", NewFungibleTokenSwapPacketData(denom, invalidLargeAmount, addr1, addr2), false},
		// {"missing sender address", NewFungibleTokenSwapPacketData(denom, amount, emptyAddr, addr2), false},
		// {"missing recipient address", NewFungibleTokenSwapPacketData(denom, amount, addr1, emptyAddr), false},
	}

	for i, tc := range testCases {
		err := tc.packetData.ValidateBasic()
		if tc.expPass {
			require.NoError(t, err, "valid test case %d failed: %v", i, err)
		} else {
			require.Error(t, err, "invalid test case %d passed: %s", i, tc.name)
		}
	}
}
