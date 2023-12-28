package records

import (
	"testing"

	cloudns "github.com/Eyup-Devop/cloudns"
	"github.com/stretchr/testify/assert"
)

var (
	authId       *string = cloudns.String("*****")
	authPassword *string = cloudns.String("**********")
)

func TestGetRecordsStatistics(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	recordsStatistics, err := GetRecordsStatistics()
	assert.Nil(t, err)
	assert.NotNil(t, recordsStatistics)
}

func TestGetRecord(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	params := &cloudns.RecordParams{
		DomainName: cloudns.String("pacificbanyo.com"),
		RecordId:   cloudns.Int(351698684),
	}
	response, err := GetRecord(params)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestListRecords(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	params := &cloudns.RecordsListParams{
		DomainName:  cloudns.String("hedgus.com"),
		Page:        cloudns.Int(1),
		RowsPerPage: cloudns.Int(50),
	}
	response, err := GetRecordList(params)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestGetRecordsPageCount(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	params := &cloudns.RecordsPagesCountParams{
		DomainName:  cloudns.String("hedgus.com"),
		RowsPerPage: cloudns.Int(20),
	}
	pageCount, err := GetRecordsPageCount(params)
	assert.Nil(t, err)
	assert.NotNil(t, pageCount)
}

func TestAddRecord(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	params := &cloudns.AddRecordParams{
		DomainName: cloudns.String("hedgus1.com"),
		RecordType: cloudns.String("A"),
		Host:       cloudns.String(""),
		Record:     cloudns.String("10.0.0.3"),
		TTL:        cloudns.Int(60),
	}
	response, err := AddRecord(params)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}
