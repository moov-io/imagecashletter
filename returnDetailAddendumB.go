// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package imagecashletter

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

// Errors specific to a ReturnDetailAddendumB Record

// ReturnDetailAddendumB Record
type ReturnDetailAddendumB struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// PayorBankName is short name of the institution by or through which the item is payable.
	PayorBankName string `json:"payorBankName"`
	// AuxiliaryOnUs identifies a code used on commercial checks at the discretion of the payor bank.
	AuxiliaryOnUs string `json:"auxiliaryOnUs"`
	// PayorBankSequenceNumber is a number that identifies the item at the payor bank.
	PayorBankSequenceNumber string `json:"payorBankSequenceNumber"`
	// PayorBankBusinessDate is The year, month, and day the payor bank processed the Return Record.
	// Format: YYYYMMDD, where: YYYY year, MM month, DD day
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	PayorBankBusinessDate time.Time `json:"payorBankBusinessDate"`
	// PayorAccountName is the account name from payor bank records.
	PayorAccountName string `json:"payorAccountName"`
	// validator is composed for imagecashletter data validation
	validator
	// converters is composed for imagecashletter to golang Converters
	converters
}

// NewReturnDetailAddendumB returns a new ReturnDetailAddendumB with default values for non exported fields
func NewReturnDetailAddendumB() ReturnDetailAddendumB {
	rdAddendumB := ReturnDetailAddendumB{}
	rdAddendumB.setRecordType()
	return rdAddendumB
}

func (rdAddendumB *ReturnDetailAddendumB) setRecordType() {
	if rdAddendumB == nil {
		return
	}
	rdAddendumB.recordType = "33"
}

// Parse takes the input record string and parses the ReturnDetailAddendumB values
func (rdAddendumB *ReturnDetailAddendumB) Parse(record string) {
	if utf8.RuneCountInString(record) < 80 {
		return // line too short
	}

	// Character position 1-2, Always "33"
	rdAddendumB.setRecordType()
	// 03-20
	rdAddendumB.PayorBankName = rdAddendumB.parseStringField(record[2:20])
	// 21-35
	rdAddendumB.AuxiliaryOnUs = rdAddendumB.parseStringField(record[20:35])
	// 36-50
	rdAddendumB.PayorBankSequenceNumber = rdAddendumB.parseStringField(record[35:50])
	// 51-58
	rdAddendumB.PayorBankBusinessDate = rdAddendumB.parseYYYYMMDDDate(record[50:58])
	// 59-80
	rdAddendumB.PayorAccountName = rdAddendumB.parseStringField(record[58:80])
}

func (rdAddendumB *ReturnDetailAddendumB) UnmarshalJSON(data []byte) error {
	type Alias ReturnDetailAddendumB
	aux := struct {
		// string is used here so we can check if the date is "" and
		// avoid trying to unmarshal that value into a time.Time
		PayorBankBusinessDate string `json:"payorBankBusinessDate"`
		*Alias
	}{
		Alias: (*Alias)(rdAddendumB),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	// The spec calls for this field to be formatted as YYYYMMDD, but we should handle
	// RFC3339 as well for convenience. Default to empty time.Time{} if no date was passed in.
	rdAddendumB.PayorBankBusinessDate = time.Time{}
	if aux.PayorBankBusinessDate != "" {
		parsed, err := time.Parse(time.RFC3339, aux.PayorBankBusinessDate)
		if err != nil {
			parsed, err = time.Parse("20060102", aux.PayorBankBusinessDate)
			if err != nil {
				return err
			}
		}
		rdAddendumB.PayorBankBusinessDate = parsed
	}
	rdAddendumB.setRecordType()
	return nil
}

func (rdAddendumB *ReturnDetailAddendumB) MarshalJSON() ([]byte, error) {
	type Alias ReturnDetailAddendumB
	if rdAddendumB.PayorBankBusinessDate.IsZero() {
		// put the empty string in instead of marshalling the zero value
		return json.Marshal(&struct {
			*Alias
			PayorBankBusinessDate string `json:"payorBankBusinessDate"`
		}{
			Alias:                 (*Alias)(rdAddendumB),
			PayorBankBusinessDate: "",
		})
	}

	// necessary to still use the alias to avoid infinite recursion
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(rdAddendumB),
	})
}

// String writes the ReturnDetailAddendumB struct to a string.
func (rdAddendumB *ReturnDetailAddendumB) String() string {
	if rdAddendumB == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(80)
	buf.WriteString(rdAddendumB.recordType)
	buf.WriteString(rdAddendumB.PayorBankNameField())
	buf.WriteString(rdAddendumB.AuxiliaryOnUsField())
	buf.WriteString(rdAddendumB.PayorBankSequenceNumberField())
	buf.WriteString(rdAddendumB.PayorBankBusinessDateField())
	buf.WriteString(rdAddendumB.PayorAccountNameField())
	return buf.String()
}

// Validate performs imagecashletter format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (rdAddendumB *ReturnDetailAddendumB) Validate() error {
	if err := rdAddendumB.fieldInclusion(); err != nil {
		return err
	}
	if rdAddendumB.recordType != "33" {
		msg := fmt.Sprintf(msgRecordType, 33)
		return &FieldError{FieldName: "recordType", Value: rdAddendumB.recordType, Msg: msg}
	}
	if err := rdAddendumB.isAlphanumericSpecial(rdAddendumB.PayorBankName); err != nil {
		return &FieldError{FieldName: "PayorBankName", Value: rdAddendumB.PayorBankName, Msg: err.Error()}
	}
	if err := rdAddendumB.isAlphanumericSpecial(rdAddendumB.PayorAccountName); err != nil {
		return &FieldError{FieldName: "PayorAccountName", Value: rdAddendumB.PayorAccountName, Msg: err.Error()}
	}

	if date := rdAddendumB.PayorBankBusinessDate; !date.IsZero() {
		// optional field - if present, year must be between 1993 and 9999
		if date.Year() < 1993 || date.Year() > 9999 {
			return &FieldError{FieldName: "PayorBankBusinessDate",
				Value: rdAddendumB.PayorBankBusinessDateField(), Msg: msgInvalidDate + ": year must be between 1993 and 9999"}
		}
	}

	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (rdAddendumB *ReturnDetailAddendumB) fieldInclusion() error {
	if rdAddendumB.recordType == "" {
		return &FieldError{FieldName: "recordType",
			Value: rdAddendumB.recordType,
			Msg:   msgFieldInclusion + ", did you use ReturnDetailAddendumB()?"}
	}
	if rdAddendumB.PayorBankSequenceNumberField() == "               " {
		return &FieldError{FieldName: "PayorBankSequenceNumber",
			Value: rdAddendumB.PayorBankSequenceNumber,
			Msg:   msgFieldInclusion + ", did you use ReturnDetailAddendumB()?"}
	}

	return nil
}

// PayorBankNameField gets the PayorBankName field
func (rdAddendumB *ReturnDetailAddendumB) PayorBankNameField() string {
	return rdAddendumB.alphaField(rdAddendumB.PayorBankName, 18)
}

// AuxiliaryOnUsField gets the AuxiliaryOnUs field
func (rdAddendumB *ReturnDetailAddendumB) AuxiliaryOnUsField() string {
	return rdAddendumB.nbsmField(rdAddendumB.AuxiliaryOnUs, 15)
}

// PayorBankSequenceNumberField gets the PayorBankSequenceNumber field
func (rdAddendumB *ReturnDetailAddendumB) PayorBankSequenceNumberField() string {
	return rdAddendumB.alphaField(rdAddendumB.PayorBankSequenceNumber, 15)
}

// PayorBankBusinessDateField gets the PayorBankBusinessDate in YYYYMMDD format
func (rdAddendumB *ReturnDetailAddendumB) PayorBankBusinessDateField() string {
	if rdAddendumB.PayorBankBusinessDate.IsZero() {
		return rdAddendumB.alphaField("", 8)
	}

	return rdAddendumB.formatYYYYMMDDDate(rdAddendumB.PayorBankBusinessDate)
}

// PayorAccountNameField gets the PayorAccountName field
func (rdAddendumB *ReturnDetailAddendumB) PayorAccountNameField() string {
	return rdAddendumB.alphaField(rdAddendumB.PayorAccountName, 22)
}
