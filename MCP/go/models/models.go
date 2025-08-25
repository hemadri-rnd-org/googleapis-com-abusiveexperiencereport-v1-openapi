package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ViolatingSitesResponse represents the ViolatingSitesResponse schema from the OpenAPI specification
type ViolatingSitesResponse struct {
	Violatingsites []SiteSummaryResponse `json:"violatingSites,omitempty"` // The list of violating sites.
}

// SiteSummaryResponse represents the SiteSummaryResponse schema from the OpenAPI specification
type SiteSummaryResponse struct {
	Reviewedsite string `json:"reviewedSite,omitempty"` // The name of the reviewed site, e.g. `google.com`.
	Underreview bool `json:"underReview,omitempty"` // Whether the site is currently under review.
	Abusivestatus string `json:"abusiveStatus,omitempty"` // The site's Abusive Experience Report status.
	Enforcementtime string `json:"enforcementTime,omitempty"` // The time at which [enforcement](https://support.google.com/webtools/answer/7538608) against the site began or will begin. Not set when the filter_status is OFF.
	Filterstatus string `json:"filterStatus,omitempty"` // The site's [enforcement status](https://support.google.com/webtools/answer/7538608).
	Lastchangetime string `json:"lastChangeTime,omitempty"` // The time at which the site's status last changed.
	Reporturl string `json:"reportUrl,omitempty"` // A link to the full Abusive Experience Report for the site. Not set in ViolatingSitesResponse. Note that you must complete the [Search Console verification process](https://support.google.com/webmasters/answer/9008080) for the site before you can access the full report.
}
