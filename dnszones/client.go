// Package customer provides the dns zones APIs
package dnszones

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

func AvailableNameServers(params *cloudns.AvailableNameServersListParams) ([]*cloudns.AvailableNameServers, error) {
	return getC().AvailableNameServers(params)
}

func (c Client) AvailableNameServers(params *cloudns.AvailableNameServersListParams) ([]*cloudns.AvailableNameServers, error) {
	availableNameServersList := []*cloudns.AvailableNameServers{}
	err := c.B.Call(http.MethodPost, "/dns/available-name-servers.json", params, &availableNameServersList, false)
	return availableNameServersList, err
}

func RegisterDomainZone(params *cloudns.RegisterDomainZoneParams) (*cloudns.RegisterDomainZoneResponse, error) {
	return getC().RegisterDomainZone(params)
}

func (c Client) RegisterDomainZone(params *cloudns.RegisterDomainZoneParams) (*cloudns.RegisterDomainZoneResponse, error) {
	validateError := cloudns.ValidateParams(params)
	if validateError != nil {
		return nil, errors.New(strings.Join(validateError, ","))
	}
	registerDomainZoneResponse := &cloudns.RegisterDomainZoneResponse{}
	err := c.B.Call(http.MethodPost, "/dns/register.json", params, registerDomainZoneResponse, false)
	return registerDomainZoneResponse, err
}

func DeleteDomainZone(params *cloudns.DeleteDomainZoneParams) (*cloudns.DeleteDomainZoneResponse, error) {
	return getC().DeleteDomainZone(params)
}

func (c Client) DeleteDomainZone(params *cloudns.DeleteDomainZoneParams) (*cloudns.DeleteDomainZoneResponse, error) {
	validateError := cloudns.ValidateParams(params)
	if validateError != nil {
		return nil, errors.New(strings.Join(validateError, ","))
	}
	deleteDomainZoneResponse := &cloudns.DeleteDomainZoneResponse{}
	err := c.B.Call(http.MethodPost, "/dns/delete.json", params, deleteDomainZoneResponse, false)
	return deleteDomainZoneResponse, err
}

func ListDomainZones(params *cloudns.ListDomainZonesParams) (*cloudns.ListDomainZonesResponse, error) {
	return getC().ListDomainZone(params)
}

func (c Client) ListDomainZone(params *cloudns.ListDomainZonesParams) (*cloudns.ListDomainZonesResponse, error) {
	validateError := cloudns.ValidateParams(params)
	if validateError != nil {
		return nil, errors.New(strings.Join(validateError, ","))
	}
	pageCountParams := &cloudns.DomainZonesPagesCountParams{
		RowsPerPage: params.RowsPerPage,
	}
	pageCount, pageCountError := c.GetDomainZonesPagesCount(pageCountParams)
	if pageCount == nil {
		return nil, pageCountError
	}

	if *params.Page > *pageCount {
		return nil, errors.New("page is greater than the total pages")
	}

	result := &cloudns.ListDomainZonesResponse{
		Page:        params.Page,
		RowsPerPage: params.RowsPerPage,
		PageCount:   pageCount,
	}

	domainZonesList := []*cloudns.DomainZones{}
	err := c.B.Call(http.MethodPost, "/dns/list-zones.json", params, &domainZonesList, false)
	result.DomainZones = domainZonesList
	return result, err
}

func GetDomainZonesPagesCount(params *cloudns.DomainZonesPagesCountParams) (*int, error) {
	return getC().GetDomainZonesPagesCount(params)
}

func (c Client) GetDomainZonesPagesCount(params *cloudns.DomainZonesPagesCountParams) (*int, error) {
	validateError := cloudns.ValidateParams(params)
	if validateError != nil {
		return nil, errors.New(strings.Join(validateError, ","))
	}
	var result int
	err := c.B.Call(http.MethodPost, "/dns/get-pages-count.json", params, &result, true)
	return &result, err
}

func GetZoneStatistics() (*cloudns.ZoneStatisticsResponse, error) {
	return getC().GetZoneStatistics()
}

func (c Client) GetZoneStatistics() (*cloudns.ZoneStatisticsResponse, error) {
	zoneStatistics := &cloudns.ZoneStatisticsResponse{}
	params := &auth.Auth{}
	err := c.B.Call(http.MethodPost, "/dns/get-zones-stats.json", params, zoneStatistics, false)
	return zoneStatistics, err
}

func GetZoneInformation(params *cloudns.GetZoneInformationParams) (*cloudns.ZoneInformationResponse, error) {
	return getC().GetZoneInformation(params)
}

func (c Client) GetZoneInformation(params *cloudns.GetZoneInformationParams) (*cloudns.ZoneInformationResponse, error) {
	validateError := cloudns.ValidateParams(params)
	if validateError != nil {
		return nil, errors.New(strings.Join(validateError, ","))
	}
	zoneInformation := &cloudns.ZoneInformationResponse{}
	err := c.B.Call(http.MethodPost, "/dns/get-zone-info.json", params, zoneInformation, false)
	return zoneInformation, err
}

func UpdateZone(params *cloudns.UpdateZoneParams) (*cloudns.UpdateZoneResponse, error) {
	return getC().UpdateZone(params)
}

func (c Client) UpdateZone(params *cloudns.UpdateZoneParams) (*cloudns.UpdateZoneResponse, error) {
	validateError := cloudns.ValidateParams(params)
	if validateError != nil {
		return nil, errors.New(strings.Join(validateError, ","))
	}
	updateZoneResponse := &cloudns.UpdateZoneResponse{}
	err := c.B.Call(http.MethodPost, "/dns/update-zone.json", params, updateZoneResponse, false)
	return updateZoneResponse, err
}

func UpdateZoneStatus(params *cloudns.UpdateZoneStatusParams) ([]*cloudns.UpdateZoneStatusResponse, error) {
	return getC().UpdateZoneStatus(params)
}

func (c Client) UpdateZoneStatus(params *cloudns.UpdateZoneStatusParams) ([]*cloudns.UpdateZoneStatusResponse, error) {
	validateError := cloudns.ValidateParams(params)
	if validateError != nil {
		return nil, errors.New(strings.Join(validateError, ","))
	}
	updateZoneStatusResponse := []*cloudns.UpdateZoneStatusResponse{}
	err := c.B.Call(http.MethodPost, "/dns/update-status.json", params, &updateZoneStatusResponse, false)
	return updateZoneStatusResponse, err
}

func IsUpdatedZone(params *cloudns.IsUpdatedZoneParams) (*bool, error) {
	return getC().IsUpdatedZone(params)
}

func (c Client) IsUpdatedZone(params *cloudns.IsUpdatedZoneParams) (*bool, error) {
	validateError := cloudns.ValidateParams(params)
	if validateError != nil {
		return nil, errors.New(strings.Join(validateError, ","))
	}
	var result bool
	err := c.B.Call(http.MethodPost, "/dns/is-updated.json", params, &result, true)
	return &result, err
}

func ChangeZoneStatus(params *cloudns.ChangeZoneStatusParams) (*cloudns.ChangeZoneStatusResponse, error) {
	return getC().ChangeZoneStatus(params)
}

func (c Client) ChangeZoneStatus(params *cloudns.ChangeZoneStatusParams) (*cloudns.ChangeZoneStatusResponse, error) {
	validateError := cloudns.ValidateParams(params)
	if validateError != nil {
		return nil, errors.New(strings.Join(validateError, ","))
	}
	response := &cloudns.ChangeZoneStatusResponse{}
	err := c.B.Call(http.MethodPost, "/dns/change-status.json", params, response, false)
	return response, err
}

func getC() Client {
	return Client{cloudns.GetBackend()}
}
