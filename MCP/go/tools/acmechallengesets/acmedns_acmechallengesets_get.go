package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/acme-dns-api/mcp-server/config"
	"github.com/acme-dns-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Acmedns_acmechallengesets_getHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		rootDomainVal, ok := args["rootDomain"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: rootDomain"), nil
		}
		rootDomain, ok := rootDomainVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: rootDomain"), nil
		}
		queryParams := make([]string, 0)
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("access_token=%s", cfg.BearerToken))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("key=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("oauth_token=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/acmeChallengeSets/%s%s", cfg.BaseURL, rootDomain, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.AcmeChallengeSet
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateAcmedns_acmechallengesets_getTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_acmeChallengeSets_rootDomain",
		mcp.WithDescription("Gets the ACME challenge set for a given domain name. Domain names must be provided in Punycode."),
		mcp.WithString("rootDomain", mcp.Required(), mcp.Description("Required. SLD + TLD domain name to list challenges. For example, this would be \"google.com\" for any FQDN under \"google.com\". That includes challenges for \"subdomain.google.com\". This MAY be Unicode or Punycode.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Acmedns_acmechallengesets_getHandler(cfg),
	}
}
