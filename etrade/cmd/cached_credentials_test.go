package cmd

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestLoadCachedCredentialsSucceedsWithGoodJson(t *testing.T) {
	jsonData := `{
  "accessToken": "TestToken",
  "accessSecret": "TestSecret",
  "lastUpdated": "2021-02-18T21:54:42.123Z"
}`
	reader := strings.NewReader(jsonData)
	cc, err := LoadCachedCredentials(reader)
	assert.Nil(t, err)
	assert.Equal(t, "TestToken", cc.AccessToken)
	assert.Equal(t, "TestSecret", cc.AccessSecret)
	lastUpdated, _ := time.Parse(time.RFC3339, "2021-02-18T21:54:42.123Z")
	assert.Equal(t, lastUpdated, cc.LastUpdated)
}

func TestLoadCachedCredentialsSucceedsWithMissingFields(t *testing.T) {
	jsonData := `{
}`
	reader := strings.NewReader(jsonData)
	cc, err := LoadCachedCredentials(reader)
	assert.Nil(t, err)
	assert.Equal(t, "", cc.AccessToken)
	assert.Equal(t, "", cc.AccessSecret)
	assert.Equal(t, time.Time{}, cc.LastUpdated)
}

func TestLoadCachedCredentialsSucceedsWithExtraFields(t *testing.T) {
	jsonData := `{
  "accessToken": "TestToken",
  "accessSecret": "TestSecret",
  "lastUpdated": "2021-02-18T21:54:42.123Z",
  "unexpectedField": "TestUnexpected"
}`
	reader := strings.NewReader(jsonData)
	cc, err := LoadCachedCredentials(reader)
	assert.Nil(t, err)
	assert.Equal(t, "TestToken", cc.AccessToken)
	assert.Equal(t, "TestSecret", cc.AccessSecret)
	lastUpdated, _ := time.Parse(time.RFC3339, "2021-02-18T21:54:42.123Z")
	assert.Equal(t, lastUpdated, cc.LastUpdated)
}

func TestLoadCachedCredentialsFailsWithBadJson(t *testing.T) {
	// Malformed JSON
	jsonData := `{
  "accessToken": TestToken,
  "accessSecret":
  "lastUpdated": "2021-02-18T21:54:42.123Z"
}`
	reader := strings.NewReader(jsonData)
	cc, err := LoadCachedCredentials(reader)
	assert.NotNil(t, err)
	assert.Nil(t, cc)
}

func TestSaveCachedCredentialsSucceeds(t *testing.T) {
	lastUpdated, _ := time.Parse(time.RFC3339, "2021-02-18T21:54:42.123Z")
	credentials := CachedCredentials{
		AccessToken:  "TestToken",
		AccessSecret: "TestSecret",
		LastUpdated:  lastUpdated,
	}
	expectedJson := `{
  "accessToken": "TestToken",
  "accessSecret": "TestSecret",
  "lastUpdated": "2021-02-18T21:54:42.123Z"
}` + "\n"

	actualJson := strings.Builder{}
	err := SaveCachedCredentials(&actualJson, &credentials)
	assert.Nil(t, err)

	assert.Equal(t, expectedJson, actualJson.String())
}
