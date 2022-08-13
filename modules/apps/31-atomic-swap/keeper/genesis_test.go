package keeper_test

import (
	"github.com/cosmos/ibc-go/v5/modules/apps/31-atomic-swap/types"
)

func (suite *KeeperTestSuite) TestGenesis() {

	genesis := suite.chainA.GetSimApp().TransferKeeper.ExportGenesis(suite.chainA.GetContext())

	suite.Require().Equal(types.PortID, genesis.PortId)
	suite.Require().NotPanics(func() {
		suite.chainA.GetSimApp().TransferKeeper.InitGenesis(suite.chainA.GetContext(), *genesis)
	})
}
