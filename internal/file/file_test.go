package file

import (
	testUtils "FileTimeChanger/internal/utils/test"
	"reflect"
	"testing"
	"time"
)

func TestNewData(t *testing.T) {
	type args struct {
		filepath     string
		modifiedTime time.Time
	}
	tests := []struct {
		name string
		args args
		want *Data
	}{
		{
			name: "Normal Data",
			args: args{
				filepath:     "test/data/blankFile.txt",
				modifiedTime: testUtils.EnsureParsedTime(t, TimeFormat, "02 Jan 2006 15:04 MST"),
			},
			want: &Data{
				filepath:     "test/data/blankFile.txt",
				modifiedTime: testUtils.EnsureParsedTime(t, TimeFormat, "02 Jan 2006 15:04 MST"),
			},
		},
		{
			name: "Blank filename",
			args: args{
				filepath:     "",
				modifiedTime: testUtils.EnsureParsedTime(t, TimeFormat, "02 Jan 2006 15:04 MST"),
			},
			want: &Data{
				filepath:     "",
				modifiedTime: testUtils.EnsureParsedTime(t, TimeFormat, "02 Jan 2006 15:04 MST"),
			},
		},
		{
			name: "Zero Time",
			args: args{
				filepath:     "",
				modifiedTime: time.Time{},
			},
			want: &Data{
				filepath:     "",
				modifiedTime: time.Time{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewData(tt.args.filepath, tt.args.modifiedTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDataBlank(t *testing.T) {
	tests := []struct {
		name string
		want *Data
	}{
		{
			name: "Assert data zero values",
			want: &Data{
				filepath:     "",
				modifiedTime: time.Time{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDataBlank(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDataBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_GetFilepath(t *testing.T) {
	type fields struct {
		filepath     string
		modifiedTime time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get test/data/blankFile.txt",
			fields: fields{
				filepath:     "test/data/blankFile.txt",
				modifiedTime: time.Now(),
			},
			want: "test/data/blankFile.txt",
		},
		{
			name: "Get empty string",
			fields: fields{
				filepath:     "",
				modifiedTime: time.Now(),
			},
			want: "",
		},
		{
			name: "Get file that does not exist",
			fields: fields{
				filepath:     "test/notFound",
				modifiedTime: time.Now(),
			},
			want: "test/notFound",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Data{
				filepath:     tt.fields.filepath,
				modifiedTime: tt.fields.modifiedTime,
			}
			if got := d.GetFilepath(); got != tt.want {
				t.Errorf("Data.GetFilepath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_GetModifiedTime(t *testing.T) {
	type fields struct {
		filepath     string
		modifiedTime time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "Get current time",
			fields: fields{
				filepath:     "",
				modifiedTime: time.Now(),
			},
			want: time.Now(),
		},
		{
			name: "Get zero time",
			fields: fields{
				filepath:     "",
				modifiedTime: time.Time{},
			},
			want: time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Data{
				filepath:     tt.fields.filepath,
				modifiedTime: tt.fields.modifiedTime,
			}
			if got := d.GetModifiedTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data.GetModifiedTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData_SetFilepath(t *testing.T) {
	type fields struct {
		filepath     string
		modifiedTime time.Time
	}
	type args struct {
		filepath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Data{
				filepath:     tt.fields.filepath,
				modifiedTime: tt.fields.modifiedTime,
			}
			d.SetFilepath(tt.args.filepath)
		})
	}
}

func TestData_SetModifiedTime(t *testing.T) {
	type fields struct {
		filepath     string
		modifiedTime time.Time
	}
	type args struct {
		modifiedTime time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Data{
				filepath:     tt.fields.filepath,
				modifiedTime: tt.fields.modifiedTime,
			}
			d.SetModifiedTime(tt.args.modifiedTime)
		})
	}
}

func TestData_ChangeFileTime(t *testing.T) {
	type fields struct {
		filepath     string
		modifiedTime time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Data{
				filepath:     tt.fields.filepath,
				modifiedTime: tt.fields.modifiedTime,
			}
			if err := d.ChangeFileTime(); (err != nil) != tt.wantErr {
				t.Errorf("Data.ChangeFileTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
