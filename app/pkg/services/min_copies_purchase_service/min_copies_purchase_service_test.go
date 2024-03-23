package minCopiesPurchaseService

import (
	"bufio"
	"flexeraCodeTest/app/pkg/constants"
	"flexeraCodeTest/app/pkg/models"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/go-playground/assert"
)

func Test_minApplicationRequiredByUserWithDeviceMapping(t *testing.T) {
	type args struct {
		usersDeviceCountMap map[string]map[string]int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Testing if both device count is same for a user",
			args: args{
				usersDeviceCountMap: map[string]map[string]int{"1": {"DESKTOP": 1, "LAPTOP": 1}},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Testing if device count is same for one user, and not same for the other",
			args: args{
				usersDeviceCountMap: map[string]map[string]int{"1": {"DESKTOP": 1, "LAPTOP": 1}, "2": {"DESKTOP": 2, "LAPTOP": 0}},
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "Testing if device count is different for both users",
			args: args{
				usersDeviceCountMap: map[string]map[string]int{"1": {"DESKTOP": 0, "LAPTOP": 1}, "2": {"DESKTOP": 1, "LAPTOP": 0}},
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := minApplicationRequiredByUserWithDeviceMapping(tt.args.usersDeviceCountMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("minApplicationRequiredByUserWithDeviceMapping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("minApplicationRequiredByUserWithDeviceMapping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_minApplicationRequiredByUserWithDeviceMapping(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minApplicationRequiredByUserWithDeviceMapping(map[string]map[string]int{"1": {"DESKTOP": 1, "LAPTOP": 1}, "2": {"DESKTOP": 2, "LAPTOP": 0}})
	}
}

func Test_checkAndUpdateUserDevicesCountMap(t *testing.T) {
	type args struct {
		duplicateRecordCheckerMap map[string]int
		record                    *models.RecordRows
		userDevicesCountMap       map[string]map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[string]map[string]int
	}{
		{
			name: "Testing if unique record of devices exist for a user",
			args: args{
				duplicateRecordCheckerMap: map[string]int{"1": 1},
				record: &models.RecordRows{
					ComputerID:    "2",
					UserID:        "1",
					ApplicationID: "374",
					ComputerType:  constants.DESKTOP,
					Comment:       "Exported from System A",
				},
				userDevicesCountMap: map[string]map[string]int{
					"1": {
						"LAPTOP": 1,
					},
				},
			},
			want: map[string]map[string]int{
				"1": {
					"LAPTOP":  1,
					"DESKTOP": 1,
				},
			},
		},
		{
			name: "Testing if unique record of devices exist for one user and same for another user",
			args: args{
				duplicateRecordCheckerMap: map[string]int{"1": 1, "2": 1, "3": 1},
				record: &models.RecordRows{
					ComputerID:    "4",
					UserID:        "2",
					ApplicationID: "374",
					ComputerType:  constants.DESKTOP,
					Comment:       "Exported from System A",
				},
				userDevicesCountMap: map[string]map[string]int{
					"1": {
						"LAPTOP":  1,
						"DESKTOP": 1,
					},
					"2": {
						"DESKTOP": 1,
					},
				},
			},
			want: map[string]map[string]int{
				"1": {
					"LAPTOP":  1,
					"DESKTOP": 1,
				},
				"2": {
					"DESKTOP": 2,
				},
			},
		},
		{
			name: "Testing if unique record of one devices exist for one user and duplicate record exists for another user",
			args: args{
				duplicateRecordCheckerMap: map[string]int{"1": 1, "2": 1},
				record: &models.RecordRows{
					ComputerID:    "2",
					UserID:        "2",
					ApplicationID: "374",
					ComputerType:  "desktop",
					Comment:       "Exported from System B",
				},
				userDevicesCountMap: map[string]map[string]int{
					"1": {
						"LAPTOP": 1,
					},
					"2": {
						"DESKTOP": 1,
					},
				},
			},
			want: map[string]map[string]int{
				"1": {
					"LAPTOP": 1,
				},
				"2": {
					"DESKTOP": 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkAndUpdateUserDevicesCountMap(tt.args.duplicateRecordCheckerMap, tt.args.record, tt.args.userDevicesCountMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkAndUpdateUserDevicesCountMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_checkAndUpdateUserDevicesCountMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkAndUpdateUserDevicesCountMap(map[string]int{"1": 1, "2": 1, "3": 1}, &models.RecordRows{
			ComputerID:    "4",
			UserID:        "2",
			ApplicationID: "374",
			ComputerType:  constants.DESKTOP,
			Comment:       "Exported from System A",
		}, map[string]map[string]int{
			"1": {
				"LAPTOP":  1,
				"DESKTOP": 1,
			},
			"2": {
				"DESKTOP": 1,
			},
		})
	}
}

func Test_MinCopiesPurchase(t *testing.T) {
	tests := []struct {
		name string
		args [][]string
		want int
	}{
		{"sample_1", [][]string{
			{"ComputerID", "UserID", "ApplicationID", "ComputerType", "Comment"},
			{"1", "1", "374", "LAPTOP", "Exported from System A"},
			{"2", "1", "374", "DESKTOP", "Exported from System A"},
		}, 1},
		{"sample_2", [][]string{
			{"ComputerID", "UserID", "ApplicationID", "ComputerType", "Comment"},
			{"1", "1", "374", "LAPTOP", "Exported from System A"},
			{"2", "1", "374", "DESKTOP", "Exported from System A"},
			{"3", "2", "374", "DESKTOP", "Exported from System A"},
			{"4", "2", "374", "DESKTOP", "Exported from System A"},
		}, 3},
		{"sample_3", [][]string{
			{"ComputerID", "UserID", "ApplicationID", "ComputerType", "Comment"},
			{"1", "1", "374", "LAPTOP", "Exported from System A"},
			{"2", "2", "374", "DESKTOP", "Exported from System A"},
			{"2", "2", "374", "desktop", "Exported from System B"},
		}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileName := fmt.Sprintf("%s.csv", tt.name)
			writeSampleToCsvFile(fileName, tt.args)
			got, err := MinCopiesPurchase(fileName)
			assert.Equal(t, err, nil)
			assert.Equal(t, tt.want, got)
			deleteCsvFile(fileName)
		})
	}
}

func writeSampleToCsvFile(fileName string, args [][]string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	for _, val := range args {
		_, err := writer.WriteString(strings.Join(val, ",") + "\n")
		if err != nil {
			log.Fatal("Error writing to csv: ", err)
		}
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func deleteCsvFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		} else {
			log.Fatal("Error deleting file:", err)
		}
	}
}

func Benchmark_MinCopiesPurchase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinCopiesPurchase("../test_data/sample.csv")
	}
}
