package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_quotabycounterkeys "github.com/apimanagementclient/mcp-server/tools/quotabycounterkeys"
	tools_quotabyperiodkeys "github.com/apimanagementclient/mcp-server/tools/quotabyperiodkeys"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_quotabycounterkeys.CreateQuotabycounterkeys_listbyserviceTool(cfg),
		tools_quotabycounterkeys.CreateQuotabycounterkeys_updateTool(cfg),
		tools_quotabyperiodkeys.CreateQuotabyperiodkeys_getTool(cfg),
		tools_quotabyperiodkeys.CreateQuotabyperiodkeys_updateTool(cfg),
	}
}
