package repositorio

import (
	m "DesafioTecnico/server/model"
	"testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		mc m.CryptoCurrency
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create(tt.args.mc); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
