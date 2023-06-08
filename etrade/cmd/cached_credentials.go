package cmd

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"
)

type CachedCredentials struct {
	AccessToken  string    `json:"accessToken"`
	AccessSecret string    `json:"accessSecret"`
	LastUpdated  time.Time `json:"lastUpdated"`
}

func LoadCachedCredentials(reader io.Reader) (*CachedCredentials, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var credentials CachedCredentials
	if err := json.Unmarshal(bytes, &credentials); err != nil {
		return nil, err
	}
	return &credentials, nil
}

func LoadCachedCredentialsFromFile(filename string) (*CachedCredentials, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return LoadCachedCredentials(file)
}

func SaveCachedCredentials(writer io.Writer, credentials *CachedCredentials) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(credentials); err != nil {
		return err
	}
	return nil
}

func SaveCachedCredentialsToFile(filename string, credentials *CachedCredentials) error {
	dirPath := filepath.Dir(filename)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return SaveCachedCredentials(file, credentials)
}
