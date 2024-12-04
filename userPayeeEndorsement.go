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

// The User Payee Endorsement Format Record is conditional, and contains a user controlled number of fields.  The
// record is used based on clearing arrangements.  The Record can occur anywhere in the file based on those clearing
// arrangements, HOWEVER it is typically recommended that it appear in the checkDetail or ReturnDetail.
// The current implementation of the User Payee Endorsement Format Record is for the Concrete Type Only.
// For reader and writer implementation, please adjust based on specific clearing arrangements, or contact MOOV for
// your particular implementation.

// UserPayeeEndorsement Record
type UserPayeeEndorsement struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record.
	recordType string
	// OwnerIdentifierIndicator indicates the type of number represented in OwnerIdentifier
	OwnerIdentifierIndicator int `json:"ownerIdentifierIndicator"`
	// OwnerIdentifier is a number used by the organization that controls the definition and formatting of this record.
	OwnerIdentifier string `json:"ownerIdentifier"`
	// OwnerIdentifierModifier is a modifier which uniquely identifies the owner within the owning organization.
	OwnerIdentifierModifier string `json:"ownerIdentifierModifier"`
	// UserRecordFormatType uniquely identifies the particular format used to parse and interrogate this record.
	// Provides a means for differentiating user record data layouts. This field shall not be populated with 001
	// since this is reserved for the UserRecordFormatType 001 PayeeEndorsement.
	UserRecordFormatType string `json:"userRecordFormatType"`
	// FormatTypeVersionLevel is a code identifies the version of the UserRecordFormatType. Provides a means for
	// identifying different versions of a record layout.
	FormatTypeVersionLevel string `json:"formatTypeVersionLevel"`
	// LengthUserData is the number of characters or bytes contained in the UserData and must be greater than 0.
	LengthUserData string `json:"LengthUserData"`
	// PayeeName is the payee name to which the check is written.
	PayeeName string `json:"payeeName"`
	// EndorsementDate The year, month, and day that the immediate origin institution creates the file which
	// shall be in Eastern Time zone format. Other time zones may be used under clearing arrangements.
	// Format: YYYYMMDD, where: YYYY year, MM month, DD day
	// Values:
	// YYYY 1993 through 9999
	// MM 01 through 12
	// DD 01 through 31
	EndorsementDate time.Time `json:"endorsementDate"`
	// BankRoutingNumber is a number that identifies the institution or the organization where the item is being
	// deposited. (note: should be the routing number associated with the BankAccountNumber.
	// Format: TTTTAAAAC, where:
	// TTTT Federal Reserve Prefix
	// AAAA ABA Institution Identifier
	// C Check Digit
	// For a number that identifies a non-financial institution: NNNNNNNNN
	BankRoutingNumber string `json:"bankRoutingNumber"`
	// BankAccountNumber is the Bank Account Number of the endorsing organization.
	BankAccountNumber string `json:"bankAccountNumber"`
	// CustomerIdentifier is a number or code identifying the customer of the endorser (e.g., driver’s license number
	// or shopper number, etc.).
	CustomerIdentifier string `json:"customerIdentifier"`
	// CustomerContactInformation is Customer contact information (e.g., phone number, e-mail, addresses, etc).
	// Unique field data shall be separated by commas.
	CustomerContactInformation string `json:"customerContactInformation"`
	// StoreMerchantProcessingSiteNumber is a number or code identifying the merchant, store or processing site.
	StoreMerchantProcessingSiteNumber string `json:"storeMerchantProcessingSiteNumber"`
	// InternalControlSequenceNumber is a number or code defined by the customer for audit proposes
	// (i.e., this can include item sequence, pocket, pass information).
	InternalControlSequenceNumber string `json:"internalControlSequenceNumber"`
	// Time is the time associated with this transaction. The default time shall be in Eastern Time zone format. The
	// local time zone or a specific time zone may be used under clearing arrangements.
	// Format: hhmm, where: hh hour, mm minute
	// Values:
	// hh '00' through '23'
	// mm '00' through '59'
	Time time.Time `json:"time"`
	// OperatorNameInitials is the name or initials of the operator or clerk processing the item.
	OperatorName string `json:"operatorName"`
	// OperatorNumber is a number or code identifying the operator or clerk processing the item.
	OperatorNumber string `json:"operatorNumber"`
	// ManagerName is The name or initials of the manager or supervisor approving the transaction.
	ManagerName string `json:"managerName"`
	// ManagerNumber is a number or code identifying the manager or supervisor approving the transaction.
	ManagerNumber string `json:"managerNumber"`
	// EquipmentNumber is  number or code of the equipment/system used to process this transaction.
	EquipmentNumber string `json:"equipmentNumber"`
	// EndorsementIndicator is an indicator to identify the type of electronic payee endorsement associated with
	// this transaction.
	// Values:
	// 0: Endorsed in Blank–Instrument becomes payable to bearer
	// 1: For Deposit Only
	// 2: For Collection Only
	// 3: Anomalous Endorsement–Endorsement made by person who is not holder of instrument
	// 4: Restrictive Endorsement–Limiting to a particular person or situation
	// 5: Guaranteed Endorsement–Deposit to the account of within named payee absence of endorsement guaranteed by
	// the bank whose Routing Number appears in BankRoutingNumber
	// 9: Other
	EndorsementIndicator int `json:"endorsementIndicator"`
	// UserField is a field used at the discretion of users of this record.
	UserField string `json:"userField"`
	// validator is composed for ImageCashLetter data validation
	validator
	// converters is composed for  to golang Converters
	converters
}

// NewUserPayeeEndorsement returns a new UserPayeeEndorsement with default values for non exported fields
func NewUserPayeeEndorsement() *UserPayeeEndorsement {
	upe := &UserPayeeEndorsement{
		UserRecordFormatType: "001",
	}
	upe.setRecordType()
	return upe
}

func (upe *UserPayeeEndorsement) setRecordType() {
	if upe == nil {
		return
	}
	upe.recordType = "68"
}

// Parse takes the input record string and parses the UserPayeeEndorsement values
func (upe *UserPayeeEndorsement) Parse(record string) {
	if utf8.RuneCountInString(record) < 335 {
		return
	}

	// Character position 1-2, Always "68"
	upe.setRecordType()
	// 03-03
	upe.OwnerIdentifierIndicator = upe.parseNumField(record[2:3])
	// 04-12
	upe.OwnerIdentifier = upe.parseStringField(record[3:12])
	// 13-32
	upe.OwnerIdentifierModifier = upe.parseStringField(record[12:32])
	// 33-35
	upe.UserRecordFormatType = "001"
	// 36-38
	upe.FormatTypeVersionLevel = upe.parseStringField(record[35:38])
	// 39-45
	upe.LengthUserData = upe.parseStringField(record[38:45])
	// 46-95
	upe.PayeeName = upe.parseStringField(record[45:95])
	// 96–103
	upe.EndorsementDate = upe.parseYYYYMMDDDate(record[95:103])
	// 104–112
	upe.BankRoutingNumber = upe.parseStringField(record[103:112])
	// 113–132
	upe.BankAccountNumber = upe.parseStringField(record[112:132])
	// 133–152
	upe.CustomerIdentifier = upe.parseStringField(record[132:152])
	// 153–202
	upe.CustomerContactInformation = upe.parseStringField(record[152:202])
	// 203–210
	upe.StoreMerchantProcessingSiteNumber = upe.parseStringField(record[202:210])
	// 211–235
	upe.InternalControlSequenceNumber = upe.parseStringField(record[210:235])
	// 236–239
	upe.Time = upe.parseSimpleTime(record[235:239])
	// 240–269
	upe.OperatorName = upe.parseStringField(record[239:269])
	// 270–274
	upe.OperatorNumber = upe.parseStringField(record[269:274])
	// 275–304
	upe.ManagerName = upe.parseStringField(record[274:304])
	// 305–309
	upe.ManagerNumber = upe.parseStringField(record[304:309])
	// 310–324
	upe.EquipmentNumber = upe.parseStringField(record[309:324])
	// 325–325
	upe.EndorsementIndicator = upe.parseNumField(record[324:325])
	// 326-335
	upe.UserField = upe.parseStringField(record[325:335])
}

func (upe *UserPayeeEndorsement) UnmarshalJSON(data []byte) error {
	type Alias UserPayeeEndorsement
	aux := struct {
		*Alias
	}{
		(*Alias)(upe),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	upe.setRecordType()
	return nil
}

// String writes the UserPayeeEndorsement struct to a variable length string.
func (upe *UserPayeeEndorsement) String() string {
	if upe == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(335)
	buf.WriteString(upe.recordType)
	buf.WriteString(upe.OwnerIdentifierIndicatorField())
	buf.WriteString(upe.OwnerIdentifierField())
	buf.WriteString(upe.OwnerIdentifierModifierField())
	buf.WriteString(upe.UserRecordFormatTypeField())
	buf.WriteString(upe.FormatTypeVersionLevelField())
	buf.WriteString(upe.LengthUserDataField())
	buf.WriteString(upe.PayeeNameField())
	buf.WriteString(upe.EndorsementDateField())
	buf.WriteString(upe.BankRoutingNumberField())
	buf.WriteString(upe.BankAccountNumberField())
	buf.WriteString(upe.CustomerIdentifierField())
	buf.WriteString(upe.CustomerContactInformationField())
	buf.WriteString(upe.StoreMerchantProcessingSiteNumberField())
	buf.WriteString(upe.InternalControlSequenceNumberField())
	buf.WriteString(upe.TimeField())
	buf.WriteString(upe.OperatorNameField())
	buf.WriteString(upe.OperatorNumberField())
	buf.WriteString(upe.ManagerNameField())
	buf.WriteString(upe.ManagerNumberField())
	buf.WriteString(upe.EquipmentNumberField())
	buf.WriteString(upe.EndorsementIndicatorField())
	buf.WriteString(upe.UserFieldField())
	return buf.String()
}

// Validate performs imagecashletter format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (upe *UserPayeeEndorsement) Validate() error {
	if err := upe.fieldInclusion(); err != nil {
		return err
	}
	if upe.recordType != "68" {
		msg := fmt.Sprintf(msgRecordType, 68)
		return &FieldError{FieldName: "recordType", Value: upe.recordType, Msg: msg}
	}
	if upe.UserRecordFormatType != "001" {
		msg := msgInvalid
		return &FieldError{FieldName: "UserRecordFormatType", Value: upe.UserRecordFormatType, Msg: msg}
	}
	if err := upe.isNumeric(upe.FormatTypeVersionLevel); err != nil {
		return &FieldError{FieldName: "FormatTypeVersionLevel",
			Value: upe.FormatTypeVersionLevel, Msg: msgNumeric}
	}
	if err := upe.validateOwnerFields(); err != nil {
		return err
	}
	if upe.CustomerIdentifier != "" {
		if err := upe.isAlphanumericSpecial(upe.CustomerIdentifier); err != nil {
			return &FieldError{FieldName: "CustomerIdentifier", Value: upe.CustomerIdentifier, Msg: err.Error()}
		}
	}
	if upe.CustomerContactInformation != "" {
		if err := upe.isAlphanumericSpecial(upe.CustomerContactInformation); err != nil {
			return &FieldError{FieldName: "CustomerContactInformation",
				Value: upe.CustomerContactInformation, Msg: err.Error()}
		}
	}
	if upe.StoreMerchantProcessingSiteNumber != "" {
		if err := upe.isAlphanumericSpecial(upe.StoreMerchantProcessingSiteNumber); err != nil {
			return &FieldError{FieldName: "StoreMerchantProcessingSiteNumber",
				Value: upe.StoreMerchantProcessingSiteNumber, Msg: err.Error()}
		}
	}
	if upe.InternalControlSequenceNumber != "" {
		if err := upe.isAlphanumericSpecial(upe.InternalControlSequenceNumber); err != nil {
			return &FieldError{FieldName: "InternalControlSequenceNumber",
				Value: upe.InternalControlSequenceNumber, Msg: err.Error()}
		}
	}
	if upe.EndorsementIndicatorField() != "" {
		if err := upe.isEndorsementIndicator(upe.EndorsementIndicator); err != nil {
			return &FieldError{FieldName: "EndorsementIndicator",
				Value: upe.EndorsementIndicatorField(), Msg: err.Error()}
		}
	}

	if upe.UserField != "" {
		if err := upe.isAlphanumericSpecial(upe.UserField); err != nil {
			return &FieldError{FieldName: "UserField", Value: upe.UserField, Msg: err.Error()}
		}
	}

	if err := upe.validateNameNumberFields(); err != nil {
		return err
	}
	return nil

}

func (upe *UserPayeeEndorsement) validateNameNumberFields() error {
	if upe.PayeeName != "" {
		if err := upe.isAlphanumericSpecial(upe.PayeeName); err != nil {
			return &FieldError{FieldName: "PayeeName", Value: upe.PayeeName, Msg: err.Error()}
		}
	}
	if upe.BankRoutingNumber != "" {
		if err := upe.isNumeric(upe.BankRoutingNumber); err != nil {
			return &FieldError{FieldName: "BankRoutingNumber", Value: upe.BankRoutingNumber, Msg: err.Error()}
		}
	}
	if upe.BankAccountNumber != "" {
		if err := upe.isAlphanumericSpecial(upe.BankAccountNumber); err != nil {
			return &FieldError{FieldName: "BankAccountNumber", Value: upe.BankAccountNumber, Msg: err.Error()}
		}
	}
	if upe.OperatorName != "" {
		if err := upe.isAlphanumericSpecial(upe.OperatorName); err != nil {
			return &FieldError{FieldName: "OperatorName", Value: upe.OperatorName, Msg: err.Error()}
		}
	}
	if upe.OperatorNumber != "" {
		if err := upe.isAlphanumericSpecial(upe.OperatorNumber); err != nil {
			return &FieldError{FieldName: "OperatorNumber", Value: upe.OperatorNumber, Msg: err.Error()}
		}
	}
	if upe.ManagerName != "" {
		if err := upe.isAlphanumericSpecial(upe.ManagerName); err != nil {
			return &FieldError{FieldName: "ManagerName", Value: upe.ManagerName, Msg: err.Error()}
		}
	}
	if upe.ManagerNumber != "" {
		if err := upe.isAlphanumericSpecial(upe.ManagerNumber); err != nil {
			return &FieldError{FieldName: "ManagerNumber", Value: upe.ManagerNumber, Msg: err.Error()}
		}
	}
	if upe.EquipmentNumber != "" {
		if err := upe.isAlphanumericSpecial(upe.EquipmentNumber); err != nil {
			return &FieldError{FieldName: "EquipmentNumber", Value: upe.EquipmentNumber, Msg: err.Error()}
		}
	}

	return nil
}

func (upe *UserPayeeEndorsement) validateOwnerFields() error {

	if err := upe.isOwnerIdentifierIndicator(upe.OwnerIdentifierIndicator); err != nil {
		return &FieldError{FieldName: "OwnerIdentifierIndicator",
			Value: upe.OwnerIdentifierIndicatorField(), Msg: err.Error()}
	}
	if upe.OwnerIdentifierModifier != "" {
		if err := upe.isAlphanumericSpecial(upe.OwnerIdentifierModifier); err != nil {
			return &FieldError{FieldName: "OwnerIdentifierModifier", Value: upe.OwnerIdentifierModifier, Msg: err.Error()}
		}
	}

	switch upe.OwnerIdentifierIndicator {
	case 0:
		if upe.OwnerIdentifier != "" {
			return &FieldError{FieldName: "OwnerIdentifier", Value: upe.OwnerIdentifier, Msg: msgInvalid}
		}
	case 1, 2, 3:
		if err := upe.isNumeric(upe.OwnerIdentifier); err != nil {
			return &FieldError{FieldName: "OwnerIdentifier", Value: upe.OwnerIdentifier, Msg: err.Error()}
		}
	case 4:
		if err := upe.isAlphanumericSpecial(upe.OwnerIdentifier); err != nil {
			return &FieldError{FieldName: "OwnerIdentifier", Value: upe.OwnerIdentifier, Msg: err.Error()}
		}
	default:
	}
	if upe.LengthUserData != "0000290" {
		return &FieldError{FieldName: "LengthUserData", Value: upe.LengthUserData, Msg: msgInvalid}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the Electronic Exchange will be returned.
func (upe *UserPayeeEndorsement) fieldInclusion() error {
	if upe.recordType == "" {
		return &FieldError{FieldName: "recordType",
			Value: upe.recordType,
			Msg:   msgFieldInclusion + ", did you use UserPayeeEndorsement()?"}
	}
	if upe.UserRecordFormatType == "" {
		return &FieldError{FieldName: "UserRecordFormatType",
			Value: upe.UserRecordFormatType,
			Msg:   msgFieldInclusion + ", did you use UserPayeeEndorsement()?"}
	}
	if upe.FormatTypeVersionLevel == "" {
		return &FieldError{FieldName: "FormatTypeVersionLevel",
			Value: upe.FormatTypeVersionLevel,
			Msg:   msgFieldInclusion + ", did you use UserPayeeEndorsement()?"}
	}
	if upe.LengthUserData == "" {
		return &FieldError{FieldName: "LengthUserData",
			Value: upe.LengthUserData,
			Msg:   msgFieldInclusion + ", did you use UserPayeeEndorsement()?"}
	}
	return nil
}

// OwnerIdentifierIndicatorField gets the OwnerIdentifierIndicator field
func (upe *UserPayeeEndorsement) OwnerIdentifierIndicatorField() string {
	return upe.numericField(upe.OwnerIdentifierIndicator, 1)
}

// OwnerIdentifierField gets the OwnerIdentifier field
func (upe *UserPayeeEndorsement) OwnerIdentifierField() string {
	return upe.alphaField(upe.OwnerIdentifier, 9)
}

// OwnerIdentifierModifierField gets the OwnerIdentifierModifier field
func (upe *UserPayeeEndorsement) OwnerIdentifierModifierField() string {
	return upe.alphaField(upe.OwnerIdentifierModifier, 20)
}

// UserRecordFormatTypeField gets the UserRecordFormatType field
func (upe *UserPayeeEndorsement) UserRecordFormatTypeField() string {
	return upe.alphaField(upe.UserRecordFormatType, 3)
}

// FormatTypeVersionLevelField gets the FormatTypeVersionLevel field
func (upe *UserPayeeEndorsement) FormatTypeVersionLevelField() string {
	return upe.alphaField(upe.FormatTypeVersionLevel, 3)
}

// LengthUserDataField gets the LengthUserData field
func (upe *UserPayeeEndorsement) LengthUserDataField() string {
	return upe.alphaField(upe.LengthUserData, 7)
}

// PayeeNameField gets the PayeeName field
func (upe *UserPayeeEndorsement) PayeeNameField() string {
	return upe.alphaField(upe.PayeeName, 50)
}

// EndorsementDateField gets the EndorsementDate field
func (upe *UserPayeeEndorsement) EndorsementDateField() string {
	return upe.formatYYYYMMDDDate(upe.EndorsementDate)
}

// BankRoutingNumberField gets the BankRoutingNumber field
func (upe *UserPayeeEndorsement) BankRoutingNumberField() string {
	return upe.alphaField(upe.BankRoutingNumber, 9)
}

// BankAccountNumberField gets the BankAccountNumber field
func (upe *UserPayeeEndorsement) BankAccountNumberField() string {
	return upe.alphaField(upe.BankAccountNumber, 20)
}

// CustomerIdentifierField gets the CustomerIdentifier field
func (upe *UserPayeeEndorsement) CustomerIdentifierField() string {
	return upe.alphaField(upe.CustomerIdentifier, 20)
}

// CustomerContactInformationField gets the CustomerContactInformation field
func (upe *UserPayeeEndorsement) CustomerContactInformationField() string {
	return upe.alphaField(upe.CustomerContactInformation, 50)
}

// StoreMerchantProcessingSiteNumberField gets the StoreMerchantProcessingSiteNumber field
func (upe *UserPayeeEndorsement) StoreMerchantProcessingSiteNumberField() string {
	return upe.alphaField(upe.StoreMerchantProcessingSiteNumber, 8)
}

// InternalControlSequenceNumberField gets the InternalControlSequenceNumber field
func (upe *UserPayeeEndorsement) InternalControlSequenceNumberField() string {
	return upe.alphaField(upe.InternalControlSequenceNumber, 25)
}

// TimeField gets the Time field
func (upe *UserPayeeEndorsement) TimeField() string {
	return upe.formatSimpleTime(upe.Time)
}

// OperatorNameField gets the OperatorName field
func (upe *UserPayeeEndorsement) OperatorNameField() string {
	return upe.alphaField(upe.OperatorName, 30)
}

// OperatorNumberField gets the OperatorNumber field
func (upe *UserPayeeEndorsement) OperatorNumberField() string {
	return upe.alphaField(upe.OperatorNumber, 5)
}

// ManagerNameField gets the ManagerName field
func (upe *UserPayeeEndorsement) ManagerNameField() string {
	return upe.alphaField(upe.ManagerName, 30)
}

// ManagerNumberField gets the ManagerNumber field
func (upe *UserPayeeEndorsement) ManagerNumberField() string {
	return upe.alphaField(upe.ManagerNumber, 5)
}

// EquipmentNumberField gets the EquipmentNumber field
func (upe *UserPayeeEndorsement) EquipmentNumberField() string {
	return upe.alphaField(upe.EquipmentNumber, 15)
}

// EndorsementIndicatorField gets the EndorsementIndicator field
func (upe *UserPayeeEndorsement) EndorsementIndicatorField() string {
	return upe.numericField(upe.EndorsementIndicator, 1)
}

// UserFieldField gets the UserField field
func (upe *UserPayeeEndorsement) UserFieldField() string {
	return upe.alphaField(upe.UserField, 10)
}
