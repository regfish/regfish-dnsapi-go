package regfishapi

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	err := godotenv.Load(os.ExpandEnv(".env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	apiToken := os.Getenv("RF_API_KEY")
	client := NewClient(apiToken)
	assert.NotNil(t, client)
	assert.Equal(t, "https://api.regfish.de", client.BaseURL)

	t.Run("Testing DNS records", func(t *testing.T) {
		t.Run("Create a new record", func(t *testing.T) {
			// Create a new record
			record := Record{
				Name: "go-client-test1.example.com.",
				Type: "A",
				Data: "10.2.3.4",
				TTL:  60,
			}
			_, err := client.CreateRecord(record)
			assert.Nil(t, err)
		})

		RecordID := 0
		t.Run("Update an existing record", func(t *testing.T) {
			// Update an existing record
			record := Record{
				ID:   1,
				Name: "go-client-test1.example.com.",
				Type: "A",
				Data: "10.2.3.5",
				TTL:  61,
			}
			res, err := client.UpdateRecord(record)
			RecordID = res.ID
			assert.Nil(t, err)
		})

		t.Run("Delete an existing record", func(t *testing.T) {
			// Delete an existing record
			err := client.DeleteRecord(RecordID)
			assert.Nil(t, err)
		})
	})
}
