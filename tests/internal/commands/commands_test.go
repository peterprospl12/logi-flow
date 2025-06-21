package commands_test

import (
	"encoding/json"
	"github.com/peterprospl12/logi-flow/internal/commands"
	"reflect"
	"testing"
)

func TestCreateOrder_Serialization(t *testing.T) {
	t.Run("Should correctly serialize the structure to JSON format", func(t *testing.T) {
		order := commands.CreateOrder{
			ShipperID: "shipper-xyz-123",
			From:      "Spain",
			To:        "Poland",
			Price:     2500.50,
		}
		expectedJSON := `{"shipper_id":"shipper-xyz-123","from":"Spain","to":"Poland","price":2500.5}`

		jsonData, err := json.Marshal(order)

		if err != nil {
			t.Fatalf("Error during serialization to JSON: %v", err)
		}

		if string(jsonData) != expectedJSON {
			t.Errorf("Received JSON differs from expected.\nReceived: %s\nExpected: %s", string(jsonData), expectedJSON)
		}
	})

	t.Run("should correctly deserialize JSON to the structure", func(t *testing.T) {
		jsonData := []byte(`{"shipper_id":"shipper-abc-456","from":"Port X","to":"Port Y","price":1999.99}`)
		expectedOrder := commands.CreateOrder{
			ShipperID: "shipper-abc-456",
			From:      "Port X",
			To:        "Port Y",
			Price:     1999.99,
		}

		var actualOrder commands.CreateOrder
		err := json.Unmarshal(jsonData, &actualOrder)

		if err != nil {
			t.Fatalf("Error during deserialization from JSON: %v", err)
		}

		if !reflect.DeepEqual(expectedOrder, actualOrder) {
			t.Errorf("Received struct differs from expected.\nReceived: %+v\nExpected: %+v", actualOrder, expectedOrder)
		}
	})
}
