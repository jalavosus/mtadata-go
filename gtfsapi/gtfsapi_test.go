package gtfsapi_test

import (
	"context"
	"testing"
	"time"

	"github.com/jalavosus/mtadata/gtfsapi"

	_ "github.com/joho/godotenv/autoload"
)

func TestMtaApi_StationStatus(t *testing.T) {
	type args struct {
		gtfsStopId string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "R01",
			args:    args{"R13"},
			wantErr: false,
		},
	}

	testCtx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(testCtx, 10*time.Second)
			defer cancel()

			m := gtfsapi.NewMtaApi()

			if err := m.StationStatus(ctx, tt.args.gtfsStopId); (err != nil) != tt.wantErr {
				t.Errorf("StationStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
