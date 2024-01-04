package records

import (
	"testing"

	cloudns "github.com/Eyup-Devop/cloudns"
	"github.com/stretchr/testify/assert"
)

var (
	authId       *string = cloudns.String("****")
	authPassword *string = cloudns.String("************")
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
		DomainName: cloudns.String("hedgus.com"),
		RecordId:   cloudns.String("415185328"),
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
		DomainName: cloudns.String("hedgus.com"),
		RecordType: cloudns.String("SRV"),
		Host:       cloudns.String("_service._protocol"),
		Record:     cloudns.String("alias.hedgus.com"),
		Priority:   cloudns.String("10"),
		Weight:     cloudns.String("20"),
		Port:       cloudns.String("80"),
		TTL:        cloudns.String("60"),
	}
	response, err := AddRecord(params)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestDeleteRecord(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	params := &cloudns.DeleteRecordParams{
		DomainName: cloudns.String("hedgus.com"),
		RecordId:   cloudns.String("222222222"),
	}
	response, err := DeleteRecord(params)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}

func TestModifyRecord(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	params := &cloudns.ModifyRecordParams{
		DomainName: cloudns.String("hedgus.com"),
		RecordId:   cloudns.String("22222222"),
		Host:       cloudns.String("_service._protocol"),
		Record:     cloudns.String("alias.hedgus.com"),
		Priority:   cloudns.String("30"),
		Weight:     cloudns.String("20"),
		Port:       cloudns.String("80"),
		TTL:        cloudns.String("60"),
	}
	response, err := ModifyRecord(params)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}
