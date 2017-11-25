package main

type Config struct {
	AuthToken string `json:"authToken"`
	AdminToken string `json:"adminToken"`
}

func NewConfig(AuthToken string, AdminToken string) *Config {
	return &Config{
		AuthToken: AuthToken,
		AdminToken: AdminToken,
	}
}
