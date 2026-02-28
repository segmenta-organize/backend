package models

type AppConfig struct {
	// Server configuration
	BackendPort  string

	// Database configuration
	DBHost       string
    DBPort       string
    DBUser       string
    DBPassword   string
    DBName       string
    DBSSLMode    string
	DBConnection string
	
	// JWT configuration
	JWTSecret    string
}

type DBConfig struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    DBSSLMode  string
	DBConnection string
}