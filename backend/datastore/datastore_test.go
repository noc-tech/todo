package datastore

import (
	"reflect"
	"testing"
)

func TestNewDatastore(t *testing.T) {
	type args struct {
		todos []*Todo
	}
	tests := []struct {
		name string
		args args
		want StorerI
	}{
		{
			name: "return storer",
			args: args{},
			want: &DataStore{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDatastore(tt.args.todos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDatastore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataStore_TodoList(t *testing.T) {
	type fields struct {
		data []*Todo
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Todo
	}{
		{
			name: "return todo list",
			fields: fields{
				data: []*Todo{
					{"id", "Todo"},
				},
			},
			want: []*Todo{
				{"id", "Todo"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DataStore{
				data: tt.fields.data,
			}
			got := s.TodoList()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataStore.TodoList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataStore_TodoCreate(t *testing.T) {
	type fields struct {
		data []*Todo
	}
	type args struct {
		todo *Todo
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Todo
	}{
		{
			name: "create todo",
			fields: fields{
				data: []*Todo{
					{"id1", "Todo"},
				},
			},
			args: args{
				todo: &Todo{"id2", "Todo"},
			},
			want: &Todo{"id2", "Todo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DataStore{
				data: tt.fields.data,
			}
			got := s.TodoCreate(tt.args.todo)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataStore.TodoCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataStore_TodoDelete(t *testing.T) {
	type fields struct {
		data []*Todo
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delete todo",
			fields: fields{
				data: []*Todo{
					{"id1", "Todo"},
				},
			},
			args: args{
				id: "id1",
			},
			wantErr: false,
		},
		{
			name: "delete todo error",
			fields: fields{
				data: []*Todo{
					{"id2", "Todo"},
				},
			},
			args: args{
				id: "id1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DataStore{
				data: tt.fields.data,
			}
			if err := s.TodoDelete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DataStore.TodoDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
