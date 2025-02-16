package httpclient

import (
	"consumption_tracker/cmd/internal/infrastructure/dtos"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type AddressClient struct {
	client  *http.Client
	baseURL string
}

func NewAddressClient(client *http.Client, baseURL string) *AddressClient {
	return &AddressClient{client: client, baseURL: baseURL}
}

func (a *AddressClient) GetAddressByMeterID(ctx context.Context, meterID int) (string, error) {
	params := url.Values{}
	params.Add("meter_id", fmt.Sprintf("%d", meterID))
	req, requestErr := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/meter?%s", a.baseURL, params.Encode()), nil)
	if requestErr != nil {
		return "", requestErr
	}

	resp, clientErr := a.client.Do(req)
	if clientErr != nil {
		return "", clientErr
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get address: status code %d", resp.StatusCode)
	}

	meterAddress := dtos.MeterAddress{}
	if decoderErr := json.NewDecoder(resp.Body).Decode(&meterAddress); decoderErr != nil {
		return "", decoderErr
	}
	return meterAddress.Address, nil
}
