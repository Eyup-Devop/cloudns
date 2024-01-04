// Package customer provides the records APIs
package records

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Eyup-Devop/cloudns"
	"github.com/Eyup-Devop/cloudns/auth"
)

type Client struct {
	B cloudns.Backend
}

func GetRecordsStatistics() (*cloudns.RecordsStatisticsResponse, error) {
	return getC().GetRecordsStatistics()
}

func (c Client) GetRecordsStatistics() (*cloudns.RecordsStatisticsResponse, error) {
	recordsStatistics := &cloudns.RecordsStatisticsResponse{}
	params := &auth.Auth{}
	err := c.B.Call(http.MethodPost, "/dns/get-records-stats.json", params, recordsStatistics, false)
	return recordsStatistics, err
}

func GetRecord(params *cloudns.RecordParams) (*cloudns.RecordResponse, error) {
	return getC().GetRecord(params)
}

func (c Client) GetRecord(params *cloudns.RecordParams) (*cloudns.RecordResponse, error) {
	validateError := cloudns.ValidateParams(params)
	if validateError != nil {
		return nil, errors.New(strings.Join(validateError, ","))
	}
	response := &cloudns.RecordResponse{}
	err := c.B.Call(http.MethodPost, "/dns/get-record.json", params, response, false)
	return response, err
}

func GetRecordList(params *cloudns.RecordsListParams) (*cloudns.RecordsListResponse, error) {
	return getC().GetRecordList(params)
}

func (c Client) GetRecordList(params *cloudns.RecordsListParams) (*cloudns.RecordsListResponse, error) {
	validateError := cloudns.ValidateParams(params)
	if validateError != nil {
		return nil, errors.New(strings.Join(validateError, ","))
	}
	result := &cloudns.RecordsListResponse{
		Page:        params.Page,
		RowsPerPage: params.RowsPerPage,
	}

	if params.RowsPerPage != nil && params.Page != nil {
		pageCountParams := &cloudns.RecordsPagesCountParams{
			DomainName:  params.DomainName,
			Host:        params.Host,
			Type:        params.Type,
			RowsPerPage: params.RowsPerPage,
		}
		pageCount, pageCountError := c.GetRecordsPageCount(pageCountParams)
		if pageCount == nil {
			return nil, pageCountError
		}

		result.PageCount = pageCount

		if *params.Page > *pageCount {
			return nil, errors.New("page is greater than the total pages")
		}
	}

	recordsList := map[string]*cloudns.Record{}
	err := c.B.Call(http.MethodPost, "/dns/records.json", params, &recordsList, false)
	result.Records = recordsList
	return result, err
}

func GetRecordsPageCount(params *cloudns.RecordsPagesCountParams) (*int, error) {
	return getC().GetRecordsPageCount(params)
}

func (c Client) GetRecordsPageCount(params *cloudns.RecordsPagesCountParams) (*int, error) {
	validateError := cloudns.ValidateParams(params)
	if validateError != nil {
		return nil, errors.New(strings.Join(validateError, ","))
	}
	var result int
	err := c.B.Call(http.MethodPost, "/dns/get-records-pages-count.json", params, &result, true)
	return &result, err
}

func AddRecord(params *cloudns.AddRecordParams) (*cloudns.AddRecordResponse, error) {
	return getC().AddRecord(params)
}

func (c Client) AddRecord(params *cloudns.AddRecordParams) (*cloudns.AddRecordResponse, error) {
	validateError := cloudns.ValidateParams(params)
	if validateError != nil {
		return nil, errors.New(strings.Join(validateError, ","))
	}
	response := &cloudns.AddRecordResponse{}
	err := c.B.Call(http.MethodPost, "/dns/add-record.json", params, response, false)
	return response, err
}

func DeleteRecord(params *cloudns.DeleteRecordParams) (*cloudns.DeleteRecordResponse, error) {
	return getC().DeleteRecord(params)
}

func (c Client) DeleteRecord(params *cloudns.DeleteRecordParams) (*cloudns.DeleteRecordResponse, error) {
	validateError := cloudns.ValidateParams(params)
	if validateError != nil {
		return nil, errors.New(strings.Join(validateError, ","))
	}
	response := &cloudns.DeleteRecordResponse{}
	err := c.B.Call(http.MethodPost, "/dns/delete-record.json", params, response, false)
	return response, err
}

func ModifyRecord(params *cloudns.ModifyRecordParams) (*cloudns.ModifyRecordResponse, error) {
	return getC().ModifyRecord(params)
}

func (c Client) ModifyRecord(params *cloudns.ModifyRecordParams) (*cloudns.ModifyRecordResponse, error) {
	validateError := cloudns.ValidateParams(params)
	if validateError != nil {
		return nil, errors.New(strings.Join(validateError, ","))
	}
	response := &cloudns.ModifyRecordResponse{}
	err := c.B.Call(http.MethodPost, "/dns/mod-record.json", params, response, false)
	return response, err
}

func getC() Client {
	return Client{cloudns.GetBackend()}
}
