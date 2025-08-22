package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/open-meteo-apis/mcp-server/config"
	"github.com/open-meteo-apis/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_v1_forecastHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["hourly"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("hourly=%v", val))
		}
		if val, ok := args["daily"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("daily=%v", val))
		}
		if val, ok := args["latitude"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("latitude=%v", val))
		}
		if val, ok := args["longitude"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("longitude=%v", val))
		}
		if val, ok := args["current_weather"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("current_weather=%v", val))
		}
		if val, ok := args["temperature_unit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("temperature_unit=%v", val))
		}
		if val, ok := args["wind_speed_unit"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("wind_speed_unit=%v", val))
		}
		if val, ok := args["timeformat"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("timeformat=%v", val))
		}
		if val, ok := args["timezone"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("timezone=%v", val))
		}
		if val, ok := args["past_days"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("past_days=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/forecast%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
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
		var result map[string]interface{}
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

func CreateGet_v1_forecastTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_forecast",
		mcp.WithDescription("7 day weather forecast for coordinates"),
		mcp.WithArray("hourly", mcp.Description("")),
		mcp.WithArray("daily", mcp.Description("")),
		mcp.WithString("latitude", mcp.Required(), mcp.Description("WGS84 coordinate")),
		mcp.WithString("longitude", mcp.Required(), mcp.Description("WGS84 coordinate")),
		mcp.WithBoolean("current_weather", mcp.Description("")),
		mcp.WithString("temperature_unit", mcp.Description("")),
		mcp.WithString("wind_speed_unit", mcp.Description("")),
		mcp.WithString("timeformat", mcp.Description("If format `unixtime` is selected, all time values are returned in UNIX epoch time in seconds. Please not that all time is then in GMT+0! For daily values with unix timestamp, please apply `utc_offset_seconds` again to get the correct date.")),
		mcp.WithString("timezone", mcp.Description("If `timezone` is set, all timestamps are returned as local-time and data is returned starting at 0:00 local-time. Any time zone name from the [time zone database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) is supported.")),
		mcp.WithNumber("past_days", mcp.Description("If `past_days` is set, yesterdays or the day before yesterdays data are also returned.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_v1_forecastHandler(cfg),
	}
}
