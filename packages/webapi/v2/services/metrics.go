package services

import (
	"github.com/iotaledger/hive.go/core/logger"
	"github.com/iotaledger/wasp/packages/chains"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/webapi/v2/dto"
)

type MetricsService struct {
	log *logger.Logger

	chainProvider chains.Provider
}

func NewMetricsService(log *logger.Logger, chainProvider chains.Provider) *MetricsService {
	return &MetricsService{
		log: log,

		chainProvider: chainProvider,
	}
}

func (c *MetricsService) GetConnectionMetrics(chainID *isc.ChainID) *dto.MetricsReport {
	chain := c.chainProvider().Get(chainID)

	if chain == nil {
		return nil
	}

	metrics := chain.GetNodeConnectionMetrics()

	return &dto.MetricsReport{
		InAliasOutput:                   dto.MapMetricItem(metrics.GetInAliasOutput()),
		InOnLedgerRequest:               dto.MapMetricItem(metrics.GetInOnLedgerRequest()),
		InOutput:                        dto.MapMetricItem(metrics.GetInOutput()),
		InStateOutput:                   dto.MapMetricItem(metrics.GetInStateOutput()),
		InTxInclusionState:              dto.MapMetricItem(metrics.GetInTxInclusionState()),
		OutPublishGovernanceTransaction: dto.MapMetricItem(metrics.GetOutPublishGovernanceTransaction()),
		OutPullLatestOutput:             dto.MapMetricItem(metrics.GetOutPullLatestOutput()),
		OutPullOutputByID:               dto.MapMetricItem(metrics.GetOutPullOutputByID()),
		OutPullTxInclusionState:         dto.MapMetricItem(metrics.GetOutPullTxInclusionState()),
	}
}
