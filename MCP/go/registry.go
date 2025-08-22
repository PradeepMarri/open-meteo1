package main

import (
	"github.com/open-meteo-apis/mcp-server/config"
	"github.com/open-meteo-apis/mcp-server/models"
	tools_weather_forecast_apis "github.com/open-meteo-apis/mcp-server/tools/weather_forecast_apis"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_weather_forecast_apis.CreateGet_v1_forecastTool(cfg),
	}
}
