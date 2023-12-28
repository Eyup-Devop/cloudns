package cloudns

import "github.com/Eyup-Devop/cloudns/auth"

type AvailableNameServersListParams struct {
	auth.Auth    `json:",inline"`
	DetailedInfo *int `json:"detailed-info,omitempty"`
}

type AvailableNameServers struct {
	Id            *string `json:"id"`
	Type          *string `json:"type"`
	Name          *string `json:"name"`
	Ip4           *string `json:"ip4"`
	Ip6           *string `json:"ip6"`
	Location      *string `json:"location"`
	LocationCc    *string `json:"location_cc"`
	DdosProtected *int    `json:"ddos_protected"`
}

type RegisterDomainZoneParams struct {
	auth.Auth  `json:",inline"`
	DomainName *string  `json:"domain-name" validate:"required"`
	ZoneType   *string  `json:"zone-type" validate:"zonetype"`
	NS         []string `json:"ns"`
	MasterIP   *string  `json:"master-ip,omitempty" validate:"omitempty,master_ip"`
}

type RegisterDomainZoneResponse struct {
	Status            string `json:"status"`
	StatusDescription string `json:"statusDescription"`
}

type DeleteDomainZoneParams struct {
	auth.Auth  `json:",inline"`
	DomainName *string `json:"domain-name" validate:"required"`
}

type DeleteDomainZoneResponse struct {
	Status            string `json:"status"`
	StatusDescription string `json:"statusDescription"`
}

type ListDomainZonesParams struct {
	auth.Auth   `json:",inline"`
	Page        *int    `json:"page" validate:"required"`
	RowsPerPage *int    `json:"rows-per-page" validate:"required,rows_per_page"`
	Search      *string `json:"search,omitempty"`
	GroupId     *int    `json:"group-id,omitempty"`
}

type DomainZones struct {
	Name      *string `json:"name,omitempty"`
	Type      *string `json:"type,omitempty"`
	Group     *string `json:"group,omitempty"`
	HasBulk   *bool   `json:"hasBulk,omitempty"`
	Zone      *string `json:"zone,omitempty"`
	Status    *string `json:"status,omitempty"`
	Serial    *string `json:"serial,omitempty"`
	IsUpdated *int    `json:"isUpdated,omitempty"`
}

type ListDomainZonesResponse struct {
	Page        *int           `json:"page"`
	RowsPerPage *int           `json:"rows-per-page"`
	PageCount   *int           `json:"page-count"`
	DomainZones []*DomainZones `json:"domain-zones"`
}

type DomainZonesPagesCountParams struct {
	auth.Auth   `json:",inline"`
	RowsPerPage *int `json:"rows-per-page" validate:"required,rows_per_page"`
}

type ZoneStatisticsResponse struct {
	Count *string `json:"count"`
	Limit *string `json:"limit"`
}

type GetZoneInformationParams struct {
	auth.Auth  `json:",inline"`
	DomainName *string `json:"domain-name" validate:"required"`
}

type ZoneInformationResponse struct {
	Name   *string `json:"name"`
	Type   *string `json:"type"`
	Zone   *string `json:"zone"`
	Status *string `json:"status"`
}

type UpdateZoneParams struct {
	auth.Auth  `json:",inline"`
	DomainName *string `json:"domain-name" validate:"required"`
}

type UpdateZoneResponse struct {
	Status            string `json:"status"`
	StatusDescription string `json:"statusDescription"`
}

type UpdateZoneStatusParams struct {
	auth.Auth  `json:",inline"`
	DomainName *string `json:"domain-name" validate:"required"`
}

type UpdateZoneStatusResponse struct {
	Server  *string `json:"server"`
	Ip4     *string `json:"ip4"`
	Ip6     *string `json:"ip6"`
	Updated *bool   `json:"updated"`
}

type IsUpdatedZoneParams struct {
	auth.Auth  `json:",inline"`
	DomainName *string `json:"domain-name" validate:"required"`
}

type ChangeZoneStatusParams struct {
	auth.Auth  `json:",inline"`
	DomainName *string `json:"domain-name" validate:"required"`
	Status     *int    `json:"status" validate:"omitempty,zone_status"`
}

type ChangeZoneStatusResponse struct {
	Status            string `json:"status"`
	StatusDescription string `json:"statusDescription"`
}
