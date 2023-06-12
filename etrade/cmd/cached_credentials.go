package cmd

import (
	"encoding/json"
	"golang.org/x/exp/slog"
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

func LoadCachedCredentialsFromFile(filename string, logger *slog.Logger) (*CachedCredentials, error) {
	file, err := os.Open(filename)
	if file != nil {
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				logger.Error(err.Error())
			}
		}(file)
	}
	if err != nil {
		return nil, err
	}
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

func SaveCachedCredentialsToFile(filename string, credentials *CachedCredentials, logger *slog.Logger) error {
	dirPath := filepath.Dir(filename)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return err
	}

	file, err := os.Create(filename)
	if file != nil {
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				logger.Error(err.Error())
			}
		}(file)
	}
	if err != nil {
		return err
	}
	return SaveCachedCredentials(file, credentials)
}
