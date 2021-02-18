package broadcaster_test

//import (
//	"os"
//	"testing"
//
//	"github.com/stretchr/testify/require"
//
//	sdk "github.com/cosmos/cosmos-sdk/types"
//
//	"github.com/maticnetwork/heimdall/app"
//	"github.com/maticnetwork/heimdall/helper"
//	hmCommon "github.com/maticnetwork/heimdall/types/common"
//	checkpointTypes "github.com/maticnetwork/heimdall/x/checkpoint/types"
//	"github.com/spf13/viper"
//	"github.com/stretchr/testify/assert"
//)
//
//// Parallel test - to check BroadcastToHeimdall synchronization
//func TestBroadcastToHeimdall(t *testing.T) {
//	t.Parallel()
//	cdc, _ := app.MakeCodecs()
//	// cli context
//	tendermintNode := "tcp://localhost:26657"
//	viper.Set(helper.NodeFlag, tendermintNode)
//	viper.Set("log_level", "info")
//	// cliCtx := cliContext.NewCLIContext().WithCodec(cdc)
//	// cliCtx.BroadcastMode = client.BroadcastSync
//	// cliCtx.TrustNode = true
//
//	helper.InitHeimdallConfig(os.ExpandEnv("$HOME/.heimdalld"))
//	_txBroadcaster := NewTxBroadcaster(cdc)
//
//	testData := []checkpointTypes.MsgCheckpoint{
//		{Proposer: helper.GetAddressStr(), StartBlock: 0, EndBlock: 63, RootHash: hmCommon.HexToHeimdallHash("0x5bd83f679c8ce7c48d6fa52ce41532fcacfbbd99d5dab415585f397bf44a0b6e").Bytes(), AccountRootHash: hmCommon.HexToHeimdallHash("0xd10b5c16c25efe0b0f5b3d75038834223934ae8c2ec2b63a62bbe42aa21e2d2d").Bytes()},
//		{Proposer: helper.GetAddressStr(), StartBlock: 64, EndBlock: 1024, RootHash: hmCommon.HexToHeimdallHash("0x5bd83f679c8ce7c48d6fa52ce41532fcacfbbd99d5dab415585f397bf44a0b6e").Bytes(), AccountRootHash: hmCommon.HexToHeimdallHash("0xd10b5c16c25efe0b0f5b3d75038834223934ae8c2ec2b63a62bbe42aa21e2d2d").Bytes()},
//		{Proposer: helper.GetAddressStr(), StartBlock: 1025, EndBlock: 2048, RootHash: hmCommon.HexToHeimdallHash("0x5bd83f679c8ce7c48d6fa52ce41532fcacfbbd99d5dab415585f397bf44a0b6e").Bytes(), AccountRootHash: hmCommon.HexToHeimdallHash("0xd10b5c16c25efe0b0f5b3d75038834223934ae8c2ec2b63a62bbe42aa21e2d2d").Bytes()},
//		{Proposer: helper.GetAddressStr(), StartBlock: 2049, EndBlock: 3124, RootHash: hmCommon.HexToHeimdallHash("0x5bd83f679c8ce7c48d6fa52ce41532fcacfbbd99d5dab415585f397bf44a0b6e").Bytes(), AccountRootHash: hmCommon.HexToHeimdallHash("0xd10b5c16c25efe0b0f5b3d75038834223934ae8c2ec2b63a62bbe42aa21e2d2d").Bytes()},
//	}
//
//	for index, test := range testData {
//		t.Run(string(rune(index)), func(t *testing.T) {
//			// create and send checkpoint message
//			accAddr, err := sdk.AccAddressFromHex(test.Proposer)
//			require.NoError(t, err)
//			msg := checkpointTypes.NewMsgCheckpointBlock(
//				accAddr,
//				test.StartBlock,
//				test.EndBlock,
//				hmCommon.BytesToHeimdallHash(test.RootHash),
//				hmCommon.BytesToHeimdallHash(test.AccountRootHash),
//				"15001",
//			)
//
//			err = _txBroadcaster.BroadcastToHeimdall(&msg)
//			assert.Empty(t, err, "Error broadcasting tx to heimdall", err)
//		})
//	}
//}
