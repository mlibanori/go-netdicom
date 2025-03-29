package dimse

import (
	"fmt"
	"io"

	"github.com/mlibanori/go-netdicom/commandset"
	"github.com/suyashkumar/dicom"
)

type CGetRsp struct {
	AffectedSOPClassUID            string
	MessageIDBeingRespondedTo      MessageID
	CommandDataSetType             uint16
	NumberOfRemainingSuboperations uint16
	NumberOfCompletedSuboperations uint16
	NumberOfFailedSuboperations    uint16
	NumberOfWarningSuboperations   uint16
	Status                         Status
	Extra                          []*dicom.Element // Unparsed elements
}

func (v *CGetRsp) Encode(e io.Writer) error {
	elems := []*dicom.Element{}

	elem, err := NewElement(commandset.CommandField, CommandFieldCGetRsp)
	if err != nil {
		return fmt.Errorf("CGetRsp.Encode: failed to create CommandField element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.AffectedSOPClassUID, v.AffectedSOPClassUID)
	if err != nil {
		return fmt.Errorf("CGetRsp.Encode: failed to create AffectedSOPClassUID element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.MessageIDBeingRespondedTo, v.MessageIDBeingRespondedTo)
	if err != nil {
		return fmt.Errorf("CGetRsp.Encode: failed to create MessageIDBeingRespondedTo element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.CommandDataSetType, v.CommandDataSetType)
	if err != nil {
		return fmt.Errorf("CGetRsp.Encode: failed to create CommandDataSetType element: %w", err)
	}
	elems = append(elems, elem)

	if v.NumberOfRemainingSuboperations != 0 {
		elem, err = NewElement(commandset.NumberOfRemainingSuboperations, v.NumberOfRemainingSuboperations)
		if err != nil {
			return fmt.Errorf("CGetRsp.Encode: failed to create NumberOfRemainingSuboperations element: %w", err)
		}
		elems = append(elems, elem)
	}

	if v.NumberOfCompletedSuboperations != 0 {
		elem, err = NewElement(commandset.NumberOfCompletedSuboperations, v.NumberOfCompletedSuboperations)
		if err != nil {
			return fmt.Errorf("CGetRsp.Encode: failed to create NumberOfCompletedSuboperations element: %w", err)
		}
		elems = append(elems, elem)
	}

	if v.NumberOfFailedSuboperations != 0 {
		elem, err = NewElement(commandset.NumberOfFailedSuboperations, v.NumberOfFailedSuboperations)
		if err != nil {
			return fmt.Errorf("CGetRsp.Encode: failed to create NumberOfFailedSuboperations element: %w", err)
		}
		elems = append(elems, elem)
	}

	if v.NumberOfWarningSuboperations != 0 {
		elem, err = NewElement(commandset.NumberOfWarningSuboperations, v.NumberOfWarningSuboperations)
		if err != nil {
			return fmt.Errorf("CGetRsp.Encode: failed to create NumberOfWarningSuboperations element: %w", err)
		}
		elems = append(elems, elem)
	}

	statusElems, err := NewStatusElements(v.Status)
	if err != nil {
		return fmt.Errorf("CGetRsp.Encode: failed to create Status elements: %w", err)
	}
	elems = append(elems, statusElems...)

	elems = append(elems, v.Extra...)

	if err := EncodeElements(e, elems); err != nil {
		return fmt.Errorf("CGetRsp.Encode: failed to encode elements: %w", err)
	}

	return nil
}

func (v *CGetRsp) HasData() bool {
	return v.CommandDataSetType != CommandDataSetTypeNull
}

func (v *CGetRsp) CommandField() uint16 {
	return CommandFieldCGetRsp
}

func (v *CGetRsp) GetMessageID() MessageID {
	return v.MessageIDBeingRespondedTo
}

func (v *CGetRsp) GetStatus() *Status {
	return &v.Status
}

func (v *CGetRsp) String() string {
	return fmt.Sprintf("CGetRsp{AffectedSOPClassUID:%v MessageIDBeingRespondedTo:%v CommandDataSetType:%v NumberOfRemainingSuboperations:%v NumberOfCompletedSuboperations:%v NumberOfFailedSuboperations:%v NumberOfWarningSuboperations:%v Status:%v}}", v.AffectedSOPClassUID, v.MessageIDBeingRespondedTo, v.CommandDataSetType, v.NumberOfRemainingSuboperations, v.NumberOfCompletedSuboperations, v.NumberOfFailedSuboperations, v.NumberOfWarningSuboperations, v.Status)
}

func (CGetRsp) decode(d *MessageDecoder) (*CGetRsp, error) {
	v := &CGetRsp{}
	var err error

	v.AffectedSOPClassUID, err = d.GetString(commandset.AffectedSOPClassUID, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("CGetRsp.decode: failed to decode AffectedSOPClassUID: %w", err)
	}

	v.MessageIDBeingRespondedTo, err = d.GetUInt16(commandset.MessageIDBeingRespondedTo, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("CGetRsp.decode: failed to decode MessageIDBeingRespondedTo: %w", err)
	}

	v.CommandDataSetType, err = d.GetUInt16(commandset.CommandDataSetType, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("CGetRsp.decode: failed to decode CommandDataSetType: %w", err)
	}

	v.NumberOfRemainingSuboperations, err = d.GetUInt16(commandset.NumberOfRemainingSuboperations, OptionalElement)
	if err != nil {
		return nil, fmt.Errorf("CGetRsp.decode: failed to decode NumberOfRemainingSuboperations: %w", err)
	}

	v.NumberOfCompletedSuboperations, err = d.GetUInt16(commandset.NumberOfCompletedSuboperations, OptionalElement)
	if err != nil {
		return nil, fmt.Errorf("CGetRsp.decode: failed to decode NumberOfCompletedSuboperations: %w", err)
	}

	v.NumberOfFailedSuboperations, err = d.GetUInt16(commandset.NumberOfFailedSuboperations, OptionalElement)
	if err != nil {
		return nil, fmt.Errorf("CGetRsp.decode: failed to decode NumberOfFailedSuboperations: %w", err)
	}

	v.NumberOfWarningSuboperations, err = d.GetUInt16(commandset.NumberOfWarningSuboperations, OptionalElement)
	if err != nil {
		return nil, fmt.Errorf("CGetRsp.decode: failed to decode NumberOfWarningSuboperations: %w", err)
	}

	v.Status, err = d.GetStatus()
	if err != nil {
		return nil, fmt.Errorf("CGetRsp.decode: failed to decode Status: %w", err)
	}

	v.Extra = d.UnparsedElements()
	return v, nil
}
