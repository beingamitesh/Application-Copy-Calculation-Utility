package helpers

import (
	"flexeraCodeTest/app/pkg/models"
	"reflect"
	"testing"
)

func TestOpenFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Testing if valid file exists and is openable",
			args: args{
				fileName: "../test_data/sample.csv",
			},
			wantErr: false,
		},
		{
			name: "Testing if non-valid file is given as input",
			args: args{
				fileName: "sample2.csv",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := OpenFile(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestFetchRecords(t *testing.T) {
	type args struct {
		text          string
		applicationId string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.RecordRows
		wantErr bool
	}{
		{
			name: "Testing if number of columns is not equal to 5",
			args: args{
				text:          "1,1,374,LAPTOP,Exported from System A,Rs. 378203",
				applicationId: "374",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Testing if application id is not equal to config value",
			args: args{
				text:          "1,1,374,LAPTOP,Exported from System A",
				applicationId: "375",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchRecords(tt.args.text, tt.args.applicationId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
