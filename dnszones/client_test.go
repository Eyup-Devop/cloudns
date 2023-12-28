package dnszones

import (
	"testing"

	cloudns "github.com/Eyup-Devop/cloudns"
	"github.com/stretchr/testify/assert"
)

var (
	authId       *string = cloudns.String("******")
	authPassword *string = cloudns.String("************")
)

func TestNameServerList(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	params := &cloudns.AvailableNameServersListParams{
		DetailedInfo: cloudns.Int(1),
	}
	nameServerList, err := AvailableNameServers(params)
	assert.Nil(t, err)
	assert.NotNil(t, nameServerList)
}

func TestRegisterDomainZone(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	params := &cloudns.RegisterDomainZoneParams{
		DomainName: cloudns.String("examplehedgus.com"),
		ZoneType:   cloudns.String("master"),
		NS:         []string{"ns1.hedgusdns.com", "ns2.hedgusdns.com", "ns3.hedgusdns.com"},
		//	MasterIP:   cloudns.String("111.111.111.111"),
	}
	registeredZoneResponse, err := RegisterDomainZone(params)
	assert.Nil(t, err)
	assert.NotNil(t, registeredZoneResponse)
}

func TestDeleteDomainZone(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	params := &cloudns.DeleteDomainZoneParams{
		DomainName: cloudns.String("examplehedgus.com"),
	}
	registeredZoneResponse, err := DeleteDomainZone(params)
	assert.Nil(t, err)
	assert.NotNil(t, registeredZoneResponse)
}

func TestListDomainZones(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	params := &cloudns.ListDomainZonesParams{
		Page:        cloudns.Int(1),
		RowsPerPage: cloudns.Int(20),
	}
	registeredZoneResponse, err := ListDomainZones(params)
	assert.Nil(t, err)
	assert.NotNil(t, registeredZoneResponse)
}

func TestGetDomainZonesPageCount(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	params := &cloudns.DomainZonesPagesCountParams{
		RowsPerPage: cloudns.Int(20),
	}
	pageCount, err := GetDomainZonesPagesCount(params)
	assert.Nil(t, err)
	assert.NotNil(t, pageCount)
}

func TestGetZonesStatistics(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	zoneStatistics, err := GetZoneStatistics()
	assert.Nil(t, err)
	assert.NotNil(t, zoneStatistics)
}

func TestGetZoneInformation(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	params := &cloudns.GetZoneInformationParams{
		DomainName: cloudns.String("hedgus.com"),
	}
	zoneStatistics, err := GetZoneInformation(params)
	assert.Nil(t, err)
	assert.NotNil(t, zoneStatistics)
}

func TestUpdateZone(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	params := &cloudns.UpdateZoneParams{
		DomainName: cloudns.String("hedgus.com"),
	}
	updateZone, err := UpdateZone(params)
	assert.Nil(t, err)
	assert.NotNil(t, updateZone)
}

func TestUpdateZoneStatus(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	params := &cloudns.UpdateZoneStatusParams{
		DomainName: cloudns.String("hedgus.com"),
	}
	updateZoneStatus, err := UpdateZoneStatus(params)
	assert.Nil(t, err)
	assert.NotNil(t, updateZoneStatus)
}

func TestIsUpdatedZone(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	params := &cloudns.IsUpdatedZoneParams{
		DomainName: cloudns.String("hedgus.com"),
	}
	isUpdatedZone, err := IsUpdatedZone(params)
	assert.Nil(t, err)
	assert.NotNil(t, isUpdatedZone)
}

func TestChangeZoneStatus(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	params := &cloudns.ChangeZoneStatusParams{
		DomainName: cloudns.String("hedgus1.com"),
		Status:     cloudns.Int(1),
	}
	response, err := ChangeZoneStatus(params)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}
