package ports

import "context"

type AddressService interface {
	GetAddressByMeterID(ctx context.Context, meterID int) (string, error)
}
