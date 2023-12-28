package cloudns

import (
	"fmt"
	"slices"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func ValidateParams[T any](value T) []string {
	err := validate.Struct(value)
	var result []string
	if err != nil {
		result = make([]string, 0)
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			result = append(result, e.Translate(trans))
		}
	}

	return result
}

func ValidatorInit() {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ = uni.GetTranslator("en")

	validate = validator.New()
	validate.RegisterValidation("master_ip", MasterIpType)
	validate.RegisterValidation("zonetype", ZoneType)
	validate.RegisterValidation("rows_per_page", RowsPerPage)
	validate.RegisterValidation("zone_status", ZoneStatus)
	validate.RegisterValidation("record_types", RecordTypes)
	validate.RegisterValidation("order_by", OrderBy)
	validate.RegisterValidation("ttl", Ttl)
	validate.RegisterValidation("priority", Priority)
	validate.RegisterValidation("weight", Weight)
	validate.RegisterValidation("port", Port)
	en_translations.RegisterDefaultTranslations(validate, trans)
	RegisterValidationErrors(trans)
}

func ZoneType(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	zoneTypes := []string{"master", "slave", "parked", "geodns"}

	return slices.Contains(zoneTypes, value)
}

func MasterIpType(fl validator.FieldLevel) bool {
	if fl.Field().String() != "" {
		if fl.Parent().FieldByName("ZoneType").Elem().String() != "slave" {
			return false
		}
	}
	return true
}

func RowsPerPage(fl validator.FieldLevel) bool {
	value := fmt.Sprintf("%d", fl.Field().Int())
	rowsPerPage := []string{"10", "20", "30", "50", "100"}

	return slices.Contains(rowsPerPage, value)
}

func ZoneStatus(fl validator.FieldLevel) bool {
	value := fl.Field().Int()
	if value != 0 && value != 1 {
		return false
	}
	return true
}

func RecordTypes(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	recordTypes := []string{
		"A", "AAAA", "MX", "CNAME", "TXT", "SPF", "NS", "SRV", "WR", "RP", "SSHFP",
		"ALIAS", "CAA", "TLSA", "CERT", "DS", "PTR", "NAPTR", "HINFO", "LOC", "DNAME", "SMIMEA", "OPENPGPKEY",
	}

	return slices.Contains(recordTypes, value)
}

func OrderBy(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	orderBy := []string{"host", "record-type", "points-to", "ttl"}

	return slices.Contains(orderBy, value)
}

func Ttl(fl validator.FieldLevel) bool {
	value := fmt.Sprintf("%d", fl.Field().Int())
	ttl := []string{"60", "300", "900", "1800", "3600", "21600", "43200", "86400", "172800", "259200", "604800", "1209600", "2592000"}

	return slices.Contains(ttl, value)
}

func Priority(fl validator.FieldLevel) bool {
	if fl.Parent().FieldByName("Type").Elem().String() == "MX" ||
		fl.Parent().FieldByName("Type").Elem().String() == "SRV" {
		return true
	}
	return false
}

func Weight(fl validator.FieldLevel) bool {
	return fl.Parent().FieldByName("Type").Elem().String() == "SRV"
}

func Port(fl validator.FieldLevel) bool {
	return fl.Parent().FieldByName("Type").Elem().String() == "SRV"
}

func RegisterValidationErrors(trans ut.Translator) {
	validate.RegisterTranslation("port", trans, func(ut ut.Translator) error {
		return ut.Add("port", "port can be used for SRV type!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("port")
		return t
	})

	validate.RegisterTranslation("weight", trans, func(ut ut.Translator) error {
		return ut.Add("weight", "weight can be used for SRV type!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("weight")
		return t
	})

	validate.RegisterTranslation("priority", trans, func(ut ut.Translator) error {
		return ut.Add("priority", "priority can be used for MX or SRV types!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("priority")
		return t
	})

	validate.RegisterTranslation("ttl", trans, func(ut ut.Translator) error {
		return ut.Add("ttl", "valid ttl must be used!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("ttl")
		return t
	})

	validate.RegisterTranslation("order_by", trans, func(ut ut.Translator) error {
		return ut.Add("order_by", "order by can be host, record-type, points-to and ttl!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("order_by")
		return t
	})

	validate.RegisterTranslation("record_types", trans, func(ut ut.Translator) error {
		return ut.Add("record_types", "valid record types must be used!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("record_types")
		return t
	})

	validate.RegisterTranslation("zone_status", trans, func(ut ut.Translator) error {
		return ut.Add("zone_status", "status must be 1 or 0!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("zone_status")
		return t
	})

	validate.RegisterTranslation("master_ip", trans, func(ut ut.Translator) error {
		return ut.Add("master_ip", "zone type must be set 'slave' for masterip!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("master_ip")
		return t
	})

	validate.RegisterTranslation("zonetype", trans, func(ut ut.Translator) error {
		return ut.Add("zonetype", "{0} must have one of master, slave, parked, geodns!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("zonetype", "zonetype")
		return t
	})

	validate.RegisterTranslation("rows_per_page", trans, func(ut ut.Translator) error {
		return ut.Add("rows_per_page", "{0} must have one of 10, 20, 30, 50, or 100!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("rows_per_page", "rows-per-page")
		return t
	})
}
