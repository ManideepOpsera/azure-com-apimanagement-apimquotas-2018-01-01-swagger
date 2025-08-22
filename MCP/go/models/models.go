package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// QuotaCounterCollection represents the QuotaCounterCollection schema from the OpenAPI specification
type QuotaCounterCollection struct {
	Count int64 `json:"count,omitempty"` // Total record count number across all pages.
	Nextlink string `json:"nextLink,omitempty"` // Next page link if any.
	Value []QuotaCounterContract `json:"value,omitempty"` // Quota counter values.
}

// QuotaCounterContract represents the QuotaCounterContract schema from the OpenAPI specification
type QuotaCounterContract struct {
	Periodendtime string `json:"periodEndTime"` // The date of the end of Counter Period. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Periodkey string `json:"periodKey"` // Identifier of the Period for which the counter was collected. Must not be empty.
	Periodstarttime string `json:"periodStartTime"` // The date of the start of Counter Period. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Value QuotaCounterValueContractProperties `json:"value,omitempty"` // Quota counter value details.
	Counterkey string `json:"counterKey"` // The Key value of the Counter. Must not be empty.
}

// QuotaCounterValueContract represents the QuotaCounterValueContract schema from the OpenAPI specification
type QuotaCounterValueContract struct {
	Value QuotaCounterValueContractProperties `json:"value,omitempty"` // Quota counter value details.
}

// QuotaCounterValueContractProperties represents the QuotaCounterValueContractProperties schema from the OpenAPI specification
type QuotaCounterValueContractProperties struct {
	Callscount int `json:"callsCount,omitempty"` // Number of times Counter was called.
	Kbtransferred float64 `json:"kbTransferred,omitempty"` // Data Transferred in KiloBytes.
}
