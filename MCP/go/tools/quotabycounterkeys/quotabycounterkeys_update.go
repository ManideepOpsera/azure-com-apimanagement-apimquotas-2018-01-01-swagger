package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Quotabycounterkeys_updateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		resourceGroupNameVal, ok := args["resourceGroupName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: resourceGroupName"), nil
		}
		resourceGroupName, ok := resourceGroupNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: resourceGroupName"), nil
		}
		serviceNameVal, ok := args["serviceName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: serviceName"), nil
		}
		serviceName, ok := serviceNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: serviceName"), nil
		}
		quotaCounterKeyVal, ok := args["quotaCounterKey"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: quotaCounterKey"), nil
		}
		quotaCounterKey, ok := quotaCounterKeyVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: quotaCounterKey"), nil
		}
		subscriptionIdVal, ok := args["subscriptionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: subscriptionId"), nil
		}
		subscriptionId, ok := subscriptionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: subscriptionId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["api-version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api-version=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.QuotaCounterValueContractProperties
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/quotas/%s%s", cfg.BaseURL, resourceGroupName, serviceName, quotaCounterKey, subscriptionId, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
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
		var result interface{}
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

func CreateQuotabycounterkeys_updateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_subscriptions_subscriptionId_resourceGroups_resourceGroupName_providers_Microsoft_ApiManagement_service_serviceName_quotas_quotaCounterKey",
		mcp.WithDescription("Updates all the quota counter values specified with the existing quota counter key to a value in the specified service instance. This should be used for reset of the quota counter values."),
		mcp.WithString("resourceGroupName", mcp.Required(), mcp.Description("The name of the resource group.")),
		mcp.WithString("serviceName", mcp.Required(), mcp.Description("The name of the API Management service.")),
		mcp.WithString("quotaCounterKey", mcp.Required(), mcp.Description("Quota counter key identifier.This is the result of expression defined in counter-key attribute of the quota-by-key policy.For Example, if you specify counter-key=\"boo\" in the policy, then it’s accessible by \"boo\" counter key. But if it’s defined as counter-key=\"@(\"b\"+\"a\")\" then it will be accessible by \"ba\" key")),
		mcp.WithString("api-version", mcp.Required(), mcp.Description("Version of the API to be used with the client request.")),
		mcp.WithString("subscriptionId", mcp.Required(), mcp.Description("Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms part of the URI for every service call.")),
		mcp.WithNumber("callsCount", mcp.Description("Input parameter: Number of times Counter was called.")),
		mcp.WithString("kbTransferred", mcp.Description("Input parameter: Data Transferred in KiloBytes.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Quotabycounterkeys_updateHandler(cfg),
	}
}
