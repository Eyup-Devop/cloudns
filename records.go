package cloudns

import "github.com/Eyup-Devop/cloudns/auth"

type RecordsStatisticsResponse struct {
	Count *int    `json:"count"`
	Limit *string `json:"limit"`
}

type RecordParams struct {
	auth.Auth  `json:",inline"`
	DomainName *string `json:"domain-name" validate:"required"`
	RecordId   *int    `json:"record-id" validate:"required"`
}

type RecordResponse struct {
	Id               *string `json:"id"`
	Type             *string `json:"type"`
	Host             *string `json:"host"`
	Record           *string `json:"record"`
	TTL              *string `json:"ttl"`
	Priority         *string `json:"priority"`
	Weight           *string `json:"weight,omitempty"`
	Port             *string `json:"port,omitempty"`
	DynamicUrlStatus *int    `json:"dynamicurl_status"`
	Failover         *string `json:"failover"`
	Status           *int    `json:"status"`
}

type RecordListParams struct {
	auth.Auth   `json:",inline"`
	DomainName  *string `json:"domain-name" validate:"required"`
	Host        *string `json:"host,omitempty"`
	HostLike    *string `json:"host-like,omitempty"`
	Type        *string `json:"type,omitempty" validate:"omitempty,record_types"`
	RowsPerPage *int    `json:"rows-per-page,omitempty" validate:"omitempty,rows_per_page"`
	Page        *int    `json:"page,omitempty"`
	OrderBy     *string `json:"order-by,omitempty" validate:"omitempty,order_by"`
}

type Record struct {
	Id               *string `json:"id"`
	Type             *string `json:"type"`
	Host             *string `json:"host"`
	Record           *string `json:"record"`
	DynamicUrlStatus *int    `json:"dynamicurl_status"`
	Priority         *string `json:"priority"`
	Weight           *string `json:"weight"`
	Port             *string `json:"port"`
	Failover         *string `json:"failover"`
	TTL              *string `json:"ttl"`
	Status           *int    `json:"status"`
}

type RecordListResponse struct {
	Page        *int               `json:"page"`
	RowsPerPage *int               `json:"rows-per-page"`
	PageCount   *int               `json:"page-count"`
	Records     map[string]*Record `json:"records"`
}

type RecordsPagesCountParams struct {
	auth.Auth   `json:",inline"`
	DomainName  *string `json:"domain-name" validate:"required"`
	Host        *string `json:"host,omitempty"`
	Type        *string `json:"type,omitempty" validate:"omitempty,record_types"`
	RowsPerPage *int    `json:"rows-per-page" validate:"required,rows_per_page"`
}

type AddRecordParams struct {
	auth.Auth  `json:",inline"`
	DomainName *string `json:"domain-name" validate:"required"`
	RecordType *string `json:"record-type" validate:"required,record_types"`
	Host       *string `json:"host" validate:"required"`
	Record     *string `json:"record" validate:"required"`
	TTL        *string `json:"ttl,omitempty" validate:"required,ttl"`
	Priority   *string `json:"priority,omitempty" validate:"omitempty,priority"`
	Weight     *string `json:"weight,omitempty" validate:"omitempty,weight"`
	Port       *string `json:"port,omitempty" validate:"omitempty,port"`
}

/*

frame	Integer	Optional	0 or 1 for Web redirects to disable or enable frame
frame-title	String	Optional	Title if frame is enabled in Web redirects
frame-keywords 	String	Optional	Keywords if frame is enabled in Web redirects
frame-description	String	Optional	Description if frame is enabled in Web redirects
mobile-meta	Integer	Optional	Mobile responsive meta tags if Web redirects with frame is enabled. Default value - 0.
save-path	Integer	Optional	0 or 1 for Web redirects
redirect-type 	Integer	Optional	301 or 302 for Web redirects if frame is disabled
mail	String	Optional	E-mail address for RP records
txt	String	Optional	Domain name for TXT record used in RP records
algorithm	Integer	Optional	Algorithm used to create the SSHFP fingerprint. Required for SSHFP records only.
fptype	Integer	Optional	Type of the SSHFP algorithm. Required for SSHFP records only.
status	Integer	Optional	Set to 1 to create the record active or to 0 to create it inactive. If omitted the record will be created active.
geodns-location	Integer	Optional	ID of a GeoDNS location for A, AAAA, CNAME, NAPTR or SRV record. The GeoDNS locations can be obtained with List GeoDNS locations
geodns-code	String	Optional	Code of a GeoDNS location for A, AAAA, CNAME, NAPTR or SRV record. The GeoDNS location codes can be obtained with List GeoDNS locations
caa_flag	Integer	 Optional	0 - Non critical or 128 - Critical
caa_type	 String	 Optional	Type of CAA record. The available flags are issue, issuewild, iodef.
caa_value	 String	 Optional	If caa_type is issue, caa_value can be hostname or ";". If caa_type is issuewild, it can be hostname or ";". If caa_type is iodef, it can be "mailto:someemail@address.tld, http://example.tld or http://example.tld.
tlsa_usage	 String	 Optional	(From 0 to 3) It shows the provided association that will be used.
tlsa_selector	 String	 Optional	(From 0 to 1) It specifies which part of the TLS certificate presented by the server will be matched against the association data
tlsa_matching_type	 String	 Optional	(From 0 to 2) It specifies how the certificate association is presented.
key-tag	Integer	 Optional	A numeric value used for identifying the referenced DS record.
algorithm	Integer	 Optional	The algorithm of the referenced DS record.
digest-type	Integer	 Optional	The cryptographic hash algorithm is used to create the Digest value.
order	String	 Optional	Specifies the order in which multiple NAPTR records must be processed (low to high).
pref	String	 Optional	Specifies the order (low to high) in which NAPTR records with equal Order values should be processed.
flag	Integer	Optional	Controls aspects of the rewriting and interpretation of the fields in the record.
params	String	Optional	Specifies the service parameters applicable to this delegation path.
regexp	String	Optional	Contains a substitution expression that is applied to the original string, held by the client in order to construct the next domain name to lookup.
replace	Integer	Optional	Specifies the next domain name (fully qualified) to query for depending on the potential values found in the flags field.
cert-type	Integer	Optional	Type of the Certificate/CRL.
cert-key-tag	Integer	Optional	A numeric value (0-65535), used the efficiently pick a CERT record.
cert-algorithm	Integer	Optional	Identifies the algorithm, used to produce a legitimate signature.
lat-deg	Integer	Optional	A numeric value(0-90), sets the latitude degrees.
lat-min	Integer	Optional	A numeric value(0-59), sets the latitude minutes. If omitted the default value is 0.
lat-sec	Integer	Optional	A numeric value(0-59), sets the latitude seconds. If omitted the default value is 0.
lat-dir	String	Optional	Sets the latitude direction. Possible values:
N - North
S - South
long-deg	Integer	Optional	A numeric value(0-180), sets the longitude degrees.
long-min	Integer	Optional	A numeric value(0-59), sets the longitude minutes. If omitted the default value is 0
long-sec	Integer	Optional	A numeric value(0-59), sets the longitude seconds. If omitted the default value is 0
long-dir	String	Optional	Sets the longitude direction. Possible values:
W - West
E - East
altitude	Integer	Optional	A numeric value(-100000.00 - 42849672.95), sets the altitude in meters.
size	Integer	Optional	A numeric value(0 - 90000000.00), sets the size in meters. If omitted the default value is 0.
h-precision	Integer	Optional	A numeric value(0 - 90000000.00), sets the horizontal precision in meters. If omitted the default value is 10000.
v-precision	Integer	Optional	A numeric value(0 - 90000000.00), sets the vertical precision in meters. If omitted the default value is 10.
cpu	String	Optional	The CPU of the server.
os	String	Optional	The operating system of the server.
*/

type AddRecordResponse struct {
	Status            string                `json:"status"`
	StatusDescription string                `json:"statusDescription"`
	Data              AddRecordResponseData `json:"data"`
}

type AddRecordResponseData struct {
	Id *int `json:"id"`
}

type DeleteRecordParams struct {
	auth.Auth  `json:",inline"`
	DomainName *string `json:"domain-name" validate:"required"`
	RecordId   *string `json:"record-id" validate:"required"`
}

type DeleteRecordResponse struct {
	Status            string `json:"status"`
	StatusDescription string `json:"statusDescription"`
}

type ModifyRecordParams struct {
	auth.Auth  `json:",inline"`
	DomainName *string `json:"domain-name" validate:"required"`
	RecordId   *string `json:"record-id" validate:"required"`
	Host       *string `json:"host" validate:"required"`
	Record     *string `json:"record" validate:"required"`
	TTL        *string `json:"ttl,omitempty" validate:"required,ttl"`
	Priority   *string `json:"priority,omitempty" validate:"omitempty,priority"`
	Weight     *string `json:"weight,omitempty" validate:"omitempty,weight"`
	Port       *string `json:"port,omitempty" validate:"omitempty,port"`
}

type ModifyRecordResponse struct {
	Status            string `json:"status"`
	StatusDescription string `json:"statusDescription"`
}
