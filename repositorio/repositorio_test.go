package repositorio

import (
	m "DesafioTecnico/server/model"
	"reflect"
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

func TestRead(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantMc  m.CryptoCurrency
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMc, err := Read(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMc, tt.wantMc) {
				t.Errorf("Read() = %v, want %v", gotMc, tt.wantMc)
			}
		})
	}
}

func TestReadAll(t *testing.T) {
	type args struct {
		sortParam string
		ascending bool
	}
	tests := []struct {
		name    string
		args    args
		wantObj []m.CryptoCurrency
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotObj, err := ReadAll(tt.args.sortParam, tt.args.ascending)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotObj, tt.wantObj) {
				t.Errorf("ReadAll() = %v, want %v", gotObj, tt.wantObj)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
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
			if err := Update(tt.args.mc); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		id string
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
			if err := Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
