package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// RotateChallengesRequest represents the RotateChallengesRequest schema from the OpenAPI specification
type RotateChallengesRequest struct {
	Accesstoken string `json:"accessToken,omitempty"` // Required. ACME DNS access token. This is a base64 token secret that is procured from the Google Domains website. It authorizes ACME TXT record updates for a domain.
	Keepexpiredrecords bool `json:"keepExpiredRecords,omitempty"` // Keep records older than 30 days that were used for previous requests.
	Recordstoadd []AcmeTxtRecord `json:"recordsToAdd,omitempty"` // ACME TXT record challenges to add. Supports multiple challenges on the same FQDN.
	Recordstoremove []AcmeTxtRecord `json:"recordsToRemove,omitempty"` // ACME TXT record challenges to remove.
}

// AcmeChallengeSet represents the AcmeChallengeSet schema from the OpenAPI specification
type AcmeChallengeSet struct {
	Record []AcmeTxtRecord `json:"record,omitempty"` // The ACME challenges on the requested domain represented as individual TXT records.
}

// AcmeTxtRecord represents the AcmeTxtRecord schema from the OpenAPI specification
type AcmeTxtRecord struct {
	Updatetime string `json:"updateTime,omitempty"` // Output only. The time when this record was last updated. This will be in UTC time.
	Digest string `json:"digest,omitempty"` // Holds the ACME challenge data put in the TXT record. This will be checked to be a valid TXT record data entry.
	Fqdn string `json:"fqdn,omitempty"` // The domain/subdomain for the record. In a request, this MAY be Unicode or Punycode. In a response, this will be in Unicode. The fqdn MUST contain the root_domain field on the request.
}
