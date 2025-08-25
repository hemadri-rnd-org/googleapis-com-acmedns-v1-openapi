package main

import (
	"github.com/acme-dns-api/mcp-server/config"
	"github.com/acme-dns-api/mcp-server/models"
	tools_acmechallengesets "github.com/acme-dns-api/mcp-server/tools/acmechallengesets"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_acmechallengesets.CreateAcmedns_acmechallengesets_getTool(cfg),
		tools_acmechallengesets.CreateAcmedns_acmechallengesets_rotatechallengesTool(cfg),
	}
}
