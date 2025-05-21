package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all configuration for the application
type Config struct {
	Database DatabaseConfig
	MQTT     MQTTConfig
	RabbitMQ RabbitMQConfig
	API      APIConfig
	Geofence GeofenceConfig
}

// DatabaseConfig holds PostgreSQL configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// MQTTConfig holds MQTT configuration
type MQTTConfig struct {
	Broker   string
	Port     int
	ClientID string
	Username string
	Password string
	Topic    string
}

// RabbitMQConfig holds RabbitMQ configuration
type RabbitMQConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Exchange string
	Queue    string
}

// APIConfig holds API server configuration
type APIConfig struct {
	Host string
	Port int
}

// GeofenceConfig holds geofence configuration
type GeofenceConfig struct {
	Radius float64 // in meters
	Points []GeoPoint
}

// GeoPoint represents a geographical point
type GeoPoint struct {
	Name      string
	Latitude  float64
	Longitude float64
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Database config
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %v", err)
	}

	// MQTT config
	mqttPort, err := strconv.Atoi(getEnv("MQTT_PORT", "1883"))
	if err != nil {
		return nil, fmt.Errorf("invalid MQTT_PORT: %v", err)
	}

	// RabbitMQ config
	rabbitPort, err := strconv.Atoi(getEnv("RABBITMQ_PORT", "5672"))
	if err != nil {
		return nil, fmt.Errorf("invalid RABBITMQ_PORT: %v", err)
	}

	// API config
	apiPort, err := strconv.Atoi(getEnv("API_PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("invalid API_PORT: %v", err)
	}

	// Geofence radius
	radius, err := strconv.ParseFloat(getEnv("GEOFENCE_RADIUS", "50"), 64)
	if err != nil {
		return nil, fmt.Errorf("invalid GEOFENCE_RADIUS: %v", err)
	}

	// Sample geofence points (in production, these would be loaded from a database or config file)
	// Here we're just creating a sample point for Jakarta
	samplePoints := []GeoPoint{
		{
			Name:      "Jakarta City Center",
			Latitude:  -6.2088,
			Longitude: 106.8456,
		},
	}

	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "postgres"),
			Port:     dbPort,
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "fleetdb"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		MQTT: MQTTConfig{
			Broker:   getEnv("MQTT_BROKER", "mosquitto"),
			Port:     mqttPort,
			ClientID: getEnv("MQTT_CLIENT_ID", "fleet-backend"),
			Username: getEnv("MQTT_USERNAME", ""),
			Password: getEnv("MQTT_PASSWORD", ""),
			Topic:    getEnv("MQTT_TOPIC", "/fleet/vehicle/+/location"),
		},
		RabbitMQ: RabbitMQConfig{
			Host:     getEnv("RABBITMQ_HOST", "rabbitmq"),
			Port:     rabbitPort,
			User:     getEnv("RABBITMQ_USER", "guest"),
			Password: getEnv("RABBITMQ_PASSWORD", "guest"),
			Exchange: getEnv("RABBITMQ_EXCHANGE", "fleet.events"),
			Queue:    getEnv("RABBITMQ_QUEUE", "geofence_alerts"),
		},
		API: APIConfig{
			Host: getEnv("API_HOST", "0.0.0.0"),
			Port: apiPort,
		},
		Geofence: GeofenceConfig{
			Radius: radius,
			Points: samplePoints,
		},
	}, nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
