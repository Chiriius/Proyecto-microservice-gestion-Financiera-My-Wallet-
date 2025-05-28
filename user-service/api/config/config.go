package config

import (
	"log"
	"os"

	vault "github.com/hashicorp/vault/api"
)

type Config struct {
	Address string
	Token   string
	Secrets map[string]string
}

func Load() *Config {
	config := vault.DefaultConfig()
	config.Address = getEnv("VAULT_ADDRESS", "http://127.0.0.1:8200")

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Error creando el cliente de Vault: %v", err)
	}

	token := getEnv("VAULT_TOKEN", "dev-only-token")
	client.SetToken(token)

	secrets, err := getSecrets(client, "secret/data/mywallet")
	if err != nil {
		log.Fatalf("Error recuperando secretos desde Vault: %v", err)
	}

	return &Config{
		Address: config.Address,
		Token:   token,
		Secrets: secrets,
	}
}

func getSecrets(client *vault.Client, path string) (map[string]string, error) {
	secret, err := client.Logical().Read(path)
	if err != nil {
		return nil, err
	}

	if secret == nil || secret.Data["data"] == nil {
		return nil, err
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return nil, err
	}

	secrets := make(map[string]string)
	for key, value := range data {
		secrets[key] = value.(string)
	}

	return secrets, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
