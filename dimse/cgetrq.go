package dimse

import (
	"fmt"
	"io"

	"github.com/mlibanori/go-netdicom/commandset"
	"github.com/suyashkumar/dicom"
)

type CGetRq struct {
	AffectedSOPClassUID string
	MessageID           MessageID
	Priority            uint16
	CommandDataSetType  uint16
	Extra               []*dicom.Element // Unparsed elements
}

func (v *CGetRq) Encode(e io.Writer) error {
	elems := []*dicom.Element{}

	elem, err := NewElement(commandset.CommandField, v.CommandField())
	if err != nil {
		return fmt.Errorf("CGetRq.Encode: failed to create CommandField element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.AffectedSOPClassUID, v.AffectedSOPClassUID)
	if err != nil {
		return fmt.Errorf("CGetRq.Encode: failed to create AffectedSOPClassUID element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.MessageID, v.MessageID)
	if err != nil {
		return fmt.Errorf("CGetRq.Encode: failed to create MessageID element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.Priority, v.Priority)
	if err != nil {
		return fmt.Errorf("CGetRq.Encode: failed to create Priority element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.CommandDataSetType, v.CommandDataSetType)
	if err != nil {
		return fmt.Errorf("CGetRq.Encode: failed to create CommandDataSetType element: %w", err)
	}
	elems = append(elems, elem)

	elems = append(elems, v.Extra...)

	if err := EncodeElements(e, elems); err != nil {
		return fmt.Errorf("CGetRq.Encode: failed to encode elements: %w", err)
	}
	return nil
}

func (v *CGetRq) HasData() bool {
	return v.CommandDataSetType != CommandDataSetTypeNull
}

func (v *CGetRq) CommandField() uint16 {
	return CommandFieldCGetRq
}

func (v *CGetRq) GetMessageID() MessageID {
	return v.MessageID
}

func (v *CGetRq) GetStatus() *Status {
	return nil
}

func (v *CGetRq) String() string {
	return fmt.Sprintf("CGetRq{AffectedSOPClassUID:%v MessageID:%v Priority:%v CommandDataSetType:%v}}", v.AffectedSOPClassUID, v.MessageID, v.Priority, v.CommandDataSetType)
}

func (CGetRq) decode(d *MessageDecoder) (*CGetRq, error) {
	v := &CGetRq{}
	var err error

	v.AffectedSOPClassUID, err = d.GetString(commandset.AffectedSOPClassUID, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cGetRq.decode: failed to decode AffectedSOPClassUID: %w", err)
	}

	v.MessageID, err = d.GetUInt16(commandset.MessageID, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cGetRq.decode: failed to decode MessageID: %w", err)
	}

	v.Priority, err = d.GetUInt16(commandset.Priority, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cGetRq.decode: failed to decode Priority: %w", err)
	}

	v.CommandDataSetType, err = d.GetUInt16(commandset.CommandDataSetType, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cGetRq.decode: failed to decode CommandDataSetType: %w", err)
	}

	v.Extra = d.UnparsedElements()
	return v, nil
}
