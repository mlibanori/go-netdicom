package dimse

import (
	"fmt"
	"io"

	"github.com/mlibanori/go-netdicom/commandset"
	"github.com/suyashkumar/dicom"
)

type CStoreRq struct {
	AffectedSOPClassUID                  string
	MessageID                            MessageID
	Priority                             uint16
	CommandDataSetType                   uint16
	AffectedSOPInstanceUID               string
	MoveOriginatorApplicationEntityTitle string
	MoveOriginatorMessageID              MessageID
	Extra                                []*dicom.Element // Unparsed elements
}

func (v *CStoreRq) Encode(e io.Writer) error {
	elems := []*dicom.Element{}

	elem, err := NewElement(commandset.CommandField, v.CommandField())
	if err != nil {
		return fmt.Errorf("CStoreRq.Encode: failed to create CommandField element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.AffectedSOPClassUID, v.AffectedSOPClassUID)
	if err != nil {
		return fmt.Errorf("CStoreRq.Encode: failed to create AffectedSOPClassUID element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.MessageID, v.MessageID)
	if err != nil {
		return fmt.Errorf("CStoreRq.Encode: failed to create MessageID element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.Priority, v.Priority)
	if err != nil {
		return fmt.Errorf("CStoreRq.Encode: failed to create Priority element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.CommandDataSetType, v.CommandDataSetType)
	if err != nil {
		return fmt.Errorf("CStoreRq.Encode: failed to create CommandDataSetType element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.AffectedSOPInstanceUID, v.AffectedSOPInstanceUID)
	if err != nil {
		return fmt.Errorf("CStoreRq.Encode: failed to create AffectedSOPInstanceUID element: %w", err)
	}
	elems = append(elems, elem)

	if v.MoveOriginatorApplicationEntityTitle != "" {
		elem, err = NewElement(commandset.MoveOriginatorApplicationEntityTitle, v.MoveOriginatorApplicationEntityTitle)
		if err != nil {
			return fmt.Errorf("CStoreRq.Encode: failed to create MoveOriginatorApplicationEntityTitle element: %w", err)
		}
		elems = append(elems, elem)
	}

	if v.MoveOriginatorMessageID != 0 {
		elem, err = NewElement(commandset.MoveOriginatorMessageID, v.MoveOriginatorMessageID)
		if err != nil {
			return fmt.Errorf("CStoreRq.Encode: failed to create MoveOriginatorMessageID element: %w", err)
		}
		elems = append(elems, elem)
	}

	elems = append(elems, v.Extra...)
	if err := EncodeElements(e, elems); err != nil {
		return fmt.Errorf("CStoreRq.Encode: failed to encode elements: %w", err)
	}
	return nil
}

func (v *CStoreRq) HasData() bool {
	return v.CommandDataSetType != CommandDataSetTypeNull
}

func (v *CStoreRq) CommandField() uint16 {
	return CommandFieldCStoreRq
}

func (v *CStoreRq) GetMessageID() MessageID {
	return v.MessageID
}

func (v *CStoreRq) GetStatus() *Status {
	return nil
}

func (v *CStoreRq) String() string {
	return fmt.Sprintf("CStoreRq{AffectedSOPClassUID:%v MessageID:%v Priority:%v CommandDataSetType:%v AffectedSOPInstanceUID:%v MoveOriginatorApplicationEntityTitle:%v MoveOriginatorMessageID:%v}}", v.AffectedSOPClassUID, v.MessageID, v.Priority, v.CommandDataSetType, v.AffectedSOPInstanceUID, v.MoveOriginatorApplicationEntityTitle, v.MoveOriginatorMessageID)
}

func (CStoreRq) decode(d *MessageDecoder) (*CStoreRq, error) {
	v := &CStoreRq{}
	var err error

	v.AffectedSOPClassUID, err = d.GetString(commandset.AffectedSOPClassUID, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cStoreRq.decode: failed to decode AffectedSOPClassUID: %w", err)
	}

	v.MessageID, err = d.GetUInt16(commandset.MessageID, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cStoreRq.decode: failed to decode MessageID: %w", err)
	}

	v.Priority, err = d.GetUInt16(commandset.Priority, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cStoreRq.decode: failed to decode Priority: %w", err)
	}

	v.CommandDataSetType, err = d.GetUInt16(commandset.CommandDataSetType, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cStoreRq.decode: failed to decode CommandDataSetType: %w", err)
	}

	v.AffectedSOPInstanceUID, err = d.GetString(commandset.AffectedSOPInstanceUID, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cStoreRq.decode: failed to decode AffectedSOPInstanceUID: %w", err)
	}

	v.MoveOriginatorApplicationEntityTitle, err = d.GetString(commandset.MoveOriginatorApplicationEntityTitle, OptionalElement)
	if err != nil {
		return nil, fmt.Errorf("cStoreRq.decode: failed to decode MoveOriginatorApplicationEntityTitle: %w", err)
	}

	v.MoveOriginatorMessageID, err = d.GetUInt16(commandset.MoveOriginatorMessageID, OptionalElement)
	if err != nil {
		return nil, fmt.Errorf("cStoreRq.decode: failed to decode MoveOriginatorMessageID: %w", err)
	}

	v.Extra = d.UnparsedElements()
	return v, nil
}
