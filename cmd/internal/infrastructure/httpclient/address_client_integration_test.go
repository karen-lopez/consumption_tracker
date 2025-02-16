package httpclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAddressByMeterIDIntegration(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"address": "address mock"}`))
	}))
	defer ts.Close()

	client := NewAddressClient(ts.Client(), ts.URL)

	address, err := client.GetAddressByMeterID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, "address mock", address)
}
