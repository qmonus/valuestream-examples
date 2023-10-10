//go:build database

package database

import (
	"demo-backend-app/pkg/database/schema"
	"reflect"
	"testing"
)

func TestInitialize(t *testing.T) {
	tests := []struct {
		name string
		want schema.Todo
	}{
		{
			name: "first record",
			want: schema.Todo{
				Id:      1,
				Comment: "test-1",
				State:   0,
			},
		},
		{
			name: "second record",
			want: schema.Todo{
				Id:      2,
				Comment: "test-2",
				State:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := ConnectMySQL()
			if err != nil {
				t.Errorf("ConnectMySQL() error = %v", err)
				return
			}
			Initialize(db)
			// Write
			db.Create(&tt.want)
			// Read
			got := schema.Todo{}
			db.Last(&got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectMySQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
