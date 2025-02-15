package postgresql

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetConsumption(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating db : %v", err)
	}
	defer db.Close()

	repo := NewPostgresqlRepository(db)

	rows := sqlmock.NewRows([]string{"meter_id", "active_energy", "reactive_energy", "capacitive_reactive", "solar", "date"}).
		AddRow(1, 1000, 200, 50, 300, time.Now())

	mock.ExpectQuery(`SELECT meter_id, active_energy, reactive_energy, capacitive_reactive, solar, date 
        FROM energy_consumption
        WHERE meter_id = \$1 AND date BETWEEN \$2 AND \$3
        ORDER BY date ASC`).
		WithArgs(1, "2023-01-01", "2023-01-31").
		WillReturnRows(rows)

	consumptions, err := repo.GetConsumption(context.Background(), 1, "2023-01-01", "2023-01-31")

	assert.NoError(t, err)
	assert.NotNil(t, consumptions)
	assert.Equal(t, 1, len(consumptions))
	assert.NoError(t, mock.ExpectationsWereMet())
}
