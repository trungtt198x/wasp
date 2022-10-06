package chain

import (
	"net/http"

	"github.com/iotaledger/wasp/packages/webapi/v2/models"

	"github.com/iotaledger/wasp/packages/kv/dict"

	"github.com/pangpanglabs/echoswagger/v2"

	loggerpkg "github.com/iotaledger/hive.go/core/logger"
	"github.com/iotaledger/wasp/packages/webapi/v2/interfaces"
)

type Controller struct {
	log *loggerpkg.Logger

	chainService     interfaces.ChainService
	nodeService      interfaces.CommitteeService
	offLedgerService interfaces.OffLedgerService
	registryService  interfaces.RegistryService
	vmService        interfaces.VMService
}

func NewChainController(log *loggerpkg.Logger, chainService interfaces.ChainService, nodeService interfaces.CommitteeService, offLedgerService interfaces.OffLedgerService, registryService interfaces.RegistryService, vmService interfaces.VMService) interfaces.APIController {
	return &Controller{
		log:              log,
		chainService:     chainService,
		nodeService:      nodeService,
		offLedgerService: offLedgerService,
		registryService:  registryService,
		vmService:        vmService,
	}
}

func (c *Controller) Name() string {
	return "chains"
}

func (c *Controller) RegisterExampleData(mock interfaces.Mocker) {
	mock.AddModel(&models.ChainInfoResponse{})
	mock.AddModel(&models.ContractListResponse{})
	mock.AddModel(&models.CommitteeInfoResponse{})
	mock.AddModel(&models.ChainListResponse{})
	mock.AddModel(&models.ContractInfoResponse{})
}

func (c *Controller) RegisterPublic(publicAPI echoswagger.ApiGroup, mocker interfaces.Mocker) {
	dictExample := dict.Dict{
		"key1": []byte("value1"),
	}.JSONDict()

	publicAPI.POST(":chainID/contract/:contractName/:functionName", c.callViewByContractName).
		AddParamBody(dictExample, "body", "Parameters", false).
		AddParamPath("", "chainID", "ChainID (Bech32)").
		AddParamPath("", "contractName", "Contract Name").
		AddParamPath("", "functionName", "Function name").
		AddResponse(http.StatusOK, "Result", dictExample, nil).
		SetResponseContentType("application/json").
		SetSummary("Call a view function on a contract by name")

	publicAPI.POST(":chainID/contract/:contractHName/:functionHName", c.callViewByHName).
		AddParamBody(dictExample, "body", "Parameters", false).
		AddParamPath("", "chainID", "ChainID (Bech32").
		AddParamPath("", "contractHName", "Contract Hname").
		AddParamPath("", "functionHName", "Function Hname").
		AddResponse(http.StatusOK, "Result", dictExample, nil).
		SetResponseContentType("application/json").
		SetSummary("Call a view function on a contract by Hname")

	publicAPI.POST(":chainID/request", c.handleNewRequest).
		AddParamBody(
			models.OffLedgerRequestBody{Request: "base64 string"},
			"body",
			"Offledger request as JSON. Request encoded in base64.",
			false).
		AddParamPath("", "chainID", "ChainID (Bech32)").
		AddResponse(http.StatusAccepted, "Request submitted", nil, nil).
		SetResponseContentType("application/json").
		SetSummary("Post an off-ledger request")
}

func (c *Controller) RegisterAdmin(adminAPI echoswagger.ApiGroup, mocker interfaces.Mocker) {
	adminAPI.GET("", c.getChainList).
		AddResponse(http.StatusOK, "A list of all available chains.", mocker.GetMockedStruct(models.ChainListResponse{}), nil).
		SetOperationId("getChainList").
		SetResponseContentType("application/json").
		SetSummary("Get a list of all chains.")

	adminAPI.PUT("", c.saveChain).
		AddParamBody(&SaveChainRecordRequest{}, "body", "The save chain request", true).
		AddResponse(http.StatusNotModified, "ChainService was not saved", nil, nil).
		AddResponse(http.StatusOK, "ChainService was saved", nil, nil).
		SetOperationId("deactivateChain").
		SetSummary("Deactivate a chain")

	adminAPI.POST(":chainID/activate", c.activateChain).
		AddParamPath("", "chainID", "ChainID (Bech32)").
		AddResponse(http.StatusNotModified, "ChainService was not activated", nil, nil).
		AddResponse(http.StatusOK, "ChainService was successfully activated", nil, nil).
		SetOperationId("activateChain").
		SetSummary("Activate a chain")

	adminAPI.POST(":chainID/deactivate", c.deactivateChain).
		AddParamPath("", "chainID", "ChainID (Bech32)").
		AddResponse(http.StatusNotModified, "ChainService was not deactivated", nil, nil).
		AddResponse(http.StatusOK, "ChainService was successfully deactivated", nil, nil).
		SetOperationId("deactivateChain").
		SetSummary("Deactivate a chain")

	adminAPI.GET(":chainID", c.getChainInfo).
		AddParamPath("", "chainID", "ChainID (Bech32)").
		AddResponse(http.StatusOK, "Information about a specific chain.", mocker.GetMockedStruct(models.ChainInfoResponse{}), nil).
		SetOperationId("getChainInfo").
		SetResponseContentType("application/json").
		SetSummary("Get information about a specific chain.")

	adminAPI.GET(":chainID/committee", c.getCommitteeInfo).
		AddParamPath("", "chainID", "ChainID (Bech32)").
		AddResponse(http.StatusOK, "A list of all nodes tied to the chain.", mocker.GetMockedStruct(models.CommitteeInfoResponse{}), nil).
		SetOperationId("getChainCommitteeInfo").
		SetResponseContentType("application/json").
		SetSummary("Get basic chain info.")

	adminAPI.GET(":chainID/contracts", c.getContracts).
		AddParamPath("", "chainID", "ChainID (Bech32)").
		AddResponse(http.StatusOK, "A list of all available contracts.", mocker.GetMockedStruct(models.ContractListResponse{}), nil).
		SetOperationId("getChainContracts").
		SetResponseContentType("application/json").
		SetSummary("Get all available chain contracts.")
}
