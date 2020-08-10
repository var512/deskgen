package desktop

import (
	"testing"
)

const (
	validName   = "Test"
	invalidName = ""
)

func TestTypeAndName(t *testing.T) {
	typeKeyTests := []struct {
		typeKey string
		name    string
		wantErr bool
	}{
		// WantErr = true.
		{wantErr: true},
		{typeKey: "", name: validName, wantErr: true},
		{typeKey: " ", name: validName, wantErr: true},
		{typeKey: "0", name: validName, wantErr: true},
		{typeKey: "1", name: validName, wantErr: true},
		{typeKey: "False", name: validName, wantErr: true},
		{typeKey: "True", name: validName, wantErr: true},
		{typeKey: "0123456789", name: validName, wantErr: true},
		{typeKey: "Invalid", name: validName, wantErr: true},
		{typeKey: "application", name: validName, wantErr: true},
		{typeKey: "directory", name: validName, wantErr: true},
		{typeKey: "Application", name: invalidName, wantErr: true},
		{typeKey: "Directory", name: invalidName, wantErr: true},
		// WantErr = false.
		{typeKey: "Application", name: validName, wantErr: false},
		{typeKey: "Directory", name: validName, wantErr: false},
	}

	for _, tt := range typeKeyTests {
		t.Run("Type "+tt.typeKey+" Name "+tt.name, func(t *testing.T) {
			_, err := NewEntry(tt.typeKey, tt.name)

			if err != nil && !tt.wantErr {
				t.Errorf("got error, want nil: Type=%#v, Name=%#v", tt.typeKey, tt.name)
			}

			if err == nil && tt.wantErr {
				t.Errorf("got nil, want error: Type=%#v, Name=%#v", tt.typeKey, tt.name)
			}
		})
	}
}
