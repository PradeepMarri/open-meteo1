package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// HourlyResponse represents the HourlyResponse schema from the OpenAPI specification
type HourlyResponse struct {
	Shortwave_radiation []float64 `json:"shortwave_radiation,omitempty"`
	Cloud_cover []float64 `json:"cloud_cover,omitempty"`
	Soil_temperature_54cm []float64 `json:"soil_temperature_54cm,omitempty"`
	Temperature_2m []float64 `json:"temperature_2m,omitempty"`
	Apparent_temperature []float64 `json:"apparent_temperature,omitempty"`
	Relative_humidity_2m []float64 `json:"relative_humidity_2m,omitempty"`
	Direct_radiation []float64 `json:"direct_radiation,omitempty"`
	Soil_moisture_27_81cm []float64 `json:"soil_moisture_27_81cm,omitempty"`
	Cloud_cover_mid []float64 `json:"cloud_cover_mid,omitempty"`
	Diffuse_radiation []float64 `json:"diffuse_radiation,omitempty"`
	Freezing_level_height []float64 `json:"freezing_level_height,omitempty"`
	Precipitation []float64 `json:"precipitation,omitempty"`
	Wind_speed_180m []float64 `json:"wind_speed_180m,omitempty"`
	Cloud_cover_low []float64 `json:"cloud_cover_low,omitempty"`
	Soil_moisture_0_1cm []float64 `json:"soil_moisture_0_1cm,omitempty"`
	Weather_code []float64 `json:"weather_code,omitempty"`
	Wind_speed_10m []float64 `json:"wind_speed_10m,omitempty"`
	Soil_moisture_9_27cm []float64 `json:"soil_moisture_9_27cm,omitempty"`
	Evapotranspiration []float64 `json:"evapotranspiration,omitempty"`
	Soil_temperature_18cm []float64 `json:"soil_temperature_18cm,omitempty"`
	Soil_temperature_0cm []float64 `json:"soil_temperature_0cm,omitempty"`
	Wind_speed_120m []float64 `json:"wind_speed_120m,omitempty"`
	Wind_direction_10m []float64 `json:"wind_direction_10m,omitempty"`
	Soil_moisture_1_3cm []float64 `json:"soil_moisture_1_3cm,omitempty"`
	Direct_normal_irradiance []float64 `json:"direct_normal_irradiance,omitempty"`
	Wind_direction_120m []float64 `json:"wind_direction_120m,omitempty"`
	Cloud_cover_high []float64 `json:"cloud_cover_high,omitempty"`
	Pressure_msl []float64 `json:"pressure_msl,omitempty"`
	Wind_direction_180m []float64 `json:"wind_direction_180m,omitempty"`
	Wind_speed_80m []float64 `json:"wind_speed_80m,omitempty"`
	Snow_height []float64 `json:"snow_height,omitempty"`
	Dew_point_2m []float64 `json:"dew_point_2m,omitempty"`
	Soil_temperature_6cm []float64 `json:"soil_temperature_6cm,omitempty"`
	Wind_direction_80m []float64 `json:"wind_direction_80m,omitempty"`
	Wind_gusts_10m []float64 `json:"wind_gusts_10m,omitempty"`
	Time []string `json:"time"`
	Vapour_pressure_deficit []float64 `json:"vapour_pressure_deficit,omitempty"`
	Soil_moisture_3_9cm []float64 `json:"soil_moisture_3_9cm,omitempty"`
}

// CurrentWeather represents the CurrentWeather schema from the OpenAPI specification
type CurrentWeather struct {
	Wind_speed float64 `json:"wind_speed"`
	Temperature float64 `json:"temperature"`
	Time string `json:"time"`
	Weather_code int `json:"weather_code"`
	Wind_direction float64 `json:"wind_direction"`
}

// DailyResponse represents the DailyResponse schema from the OpenAPI specification
type DailyResponse struct {
	Shortwave_radiation_sum []float64 `json:"shortwave_radiation_sum,omitempty"`
	Apparent_temperature_max []float64 `json:"apparent_temperature_max,omitempty"`
	Apparent_temperature_min []float64 `json:"apparent_temperature_min,omitempty"`
	Et0_fao_evapotranspiration []float64 `json:"et0_fao_evapotranspiration,omitempty"`
	Weather_code []float64 `json:"weather_code,omitempty"`
	Precipitation_sum []float64 `json:"precipitation_sum,omitempty"`
	Uv_index_clear_sky_max []float64 `json:"uv_index_clear_sky_max,omitempty"`
	Uv_index_max []float64 `json:"uv_index_max,omitempty"`
	Wind_speed_10m_max []float64 `json:"wind_speed_10m_max,omitempty"`
	Time []string `json:"time"`
	Temperature_2m_min []float64 `json:"temperature_2m_min,omitempty"`
	Wind_direction_10m_dominant []float64 `json:"wind_direction_10m_dominant,omitempty"`
	Wind_gusts_10m_max []float64 `json:"wind_gusts_10m_max,omitempty"`
	Sunrise []float64 `json:"sunrise,omitempty"`
	Temperature_2m_max []float64 `json:"temperature_2m_max,omitempty"`
	Precipitation_hours []float64 `json:"precipitation_hours,omitempty"`
	Sunset []float64 `json:"sunset,omitempty"`
}
