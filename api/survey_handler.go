package api

import (
	"context"
	"fmt"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/signer"
	"github.com/block-vision/sui-go-sdk/sui"
	"github.com/block-vision/sui-go-sdk/utils"
)

func SurveyHandler() {

	var ctx = context.Background()
	var cli = sui.NewSuiClient(constant.SuiTestnetEndpoint)

	//rsp, err := cli.SuiGetMoveFunctionArgTypes(ctx, models.GetMoveFunctionArgTypesRequest{
	//	Package:  "0x33e8022a5dee1677500089c38af2681f7638df3a53a31b4645a32649f9eab035",
	//	Module:   "suisurvey",
	//	Function: "create_survey",
	//})

	//rsp, err := cli.SuiGetNormalizedMoveModulesByPackage(ctx, models.GetNormalizedMoveModulesByPackageRequest{
	//	Package: "0x33e8022a5dee1677500089c38af2681f7638df3a53a31b4645a32649f9eab035",
	//})

	signerAccount, err := signer.NewSignertWithMnemonic("input your mnemonic")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	priKey := signerAccount.PriKey
	fmt.Printf("signerAccount.Address: %s\n", signerAccount.Address)

	rsp, err := cli.TransferSui(ctx, models.TransferSuiRequest{
		Signer:      signerAccount.Address,
		SuiObjectId: "0x33e8022a5dee1677500089c38af2681f7638df3a53a31b4645a32649f9eab035",
		GasBudget:   "100000000",
		Recipient:   "0x4ae8be62692d1bbf892b657ee78a59954240ee0525f20a5b5687a70995cf0eff",
		Amount:      "1",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// see the successful transaction url: https://suivision.xyz/txblock/C7iYsH4tU5RdY1KBeNax4mCBn3XLZ5UswsuDpKrVkcH6
	rsp2, err := cli.SignAndExecuteTransactionBlock(ctx, models.SignAndExecuteTransactionBlockRequest{
		TxnMetaData: rsp,
		PriKey:      priKey,
		Options: models.SuiTransactionBlockOptions{
			ShowInput:    true,
			ShowRawInput: true,
			ShowEffects:  true,
		},
		RequestType: "WaitForLocalExecution",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp2)

	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	utils.PrettyPrint(rsp)

}

// blobId string, name string, expiration string, minParticipants string, maxParticipants string, reward string, state string, clock string, contractInteraction string
func SurveyCallbackHandler() {
	var ctx = context.Background()
	var cli = sui.NewSuiClient(constant.SuiTestnetEndpoint)

	signerAccount, err := signer.NewSignertWithMnemonic("real rally fatigue section pool indicate million final evidence original axis yellow")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	priKey := signerAccount.PriKey
	fmt.Printf("signerAccount.Address: %s\n", signerAccount.Address)

	gasObj := "0x34116c28dd03071580835790817c13c24f3cf9b2bf7b7544ce326d5a5322d905"

	rsp, err := cli.MoveCall(ctx, models.MoveCallRequest{
		Signer:          signerAccount.Address,
		PackageObjectId: "0x33e8022a5dee1677500089c38af2681f7638df3a53a31b4645a32649f9eab035",
		Module:          "suisurvey",
		Function:        "create_survey",
		TypeArguments:   []interface{}{},
		Arguments: []interface{}{
			"marshall",
			"0",   //in ms
			"0",   // Replace nil with 0 or a default value if U64 is expected
			"100", // Replace nil with 0 or a default value if U64 is expected
			"0",   // Replace nil with 0 or a default value if U64 is expected
			"0x33e8022a5dee1677500089c38af2681f7638df3a53a31b4645a32649f9eab035", // Replace nil with 0 or a default value if U64 is expected
			"0",     // Replace nil with 0 or a default value if U64 is expected
			"since", // Replace nil with 0 or a default value if U64 is expected
			0,
		},
		Gas:       gasObj,
		GasBudget: "100000000",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// see the successful transaction url: https://explorer.sui.io/txblock/CD5hFB4bWFThhb6FtvKq3xAxRri72vsYLJAVd7p9t2sR?network=testnet
	rsp2, err := cli.SignAndExecuteTransactionBlock(ctx, models.SignAndExecuteTransactionBlockRequest{
		TxnMetaData: rsp,
		PriKey:      priKey,
		// only fetch the effects field
		Options: models.SuiTransactionBlockOptions{
			ShowInput:    true,
			ShowRawInput: true,
			ShowEffects:  true,
		},
		RequestType: "WaitForLocalExecution",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	utils.PrettyPrint(rsp2)
}
