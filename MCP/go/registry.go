package main

import (
	"github.com/abusive-experience-report-api/mcp-server/config"
	"github.com/abusive-experience-report-api/mcp-server/models"
	tools_sites "github.com/abusive-experience-report-api/mcp-server/tools/sites"
	tools_violatingsites "github.com/abusive-experience-report-api/mcp-server/tools/violatingsites"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_sites.CreateAbusiveexperiencereport_sites_getTool(cfg),
		tools_violatingsites.CreateAbusiveexperiencereport_violatingsites_listTool(cfg),
	}
}
