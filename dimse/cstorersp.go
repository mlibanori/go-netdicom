package dimse

import (
	"fmt"
	"io"

	"github.com/mlibanori/go-netdicom/commandset"
	"github.com/suyashkumar/dicom"
)

type CStoreRsp struct {
	AffectedSOPClassUID       string
	MessageIDBeingRespondedTo MessageID
	CommandDataSetType        uint16
	AffectedSOPInstanceUID    string
	Status                    Status
	Extra                     []*dicom.Element // Unparsed elements
}

func (v *CStoreRsp) Encode(e io.Writer) error {
	elems := []*dicom.Element{}

	elem, err := NewElement(commandset.CommandField, v.CommandField())
	if err != nil {
		return fmt.Errorf("CStoreRsp.Encode: failed to create CommandField element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.AffectedSOPClassUID, v.AffectedSOPClassUID)
	if err != nil {
		return fmt.Errorf("CStoreRsp.Encode: failed to create AffectedSOPClassUID element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.MessageIDBeingRespondedTo, v.MessageIDBeingRespondedTo)
	if err != nil {
		return fmt.Errorf("CStoreRsp.Encode: failed to create MessageIDBeingRespondedTo element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.CommandDataSetType, v.CommandDataSetType)
	if err != nil {
		return fmt.Errorf("CStoreRsp.Encode: failed to create CommandDataSetType element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.AffectedSOPInstanceUID, v.AffectedSOPInstanceUID)
	if err != nil {
		return fmt.Errorf("CStoreRsp.Encode: failed to create AffectedSOPInstanceUID element: %w", err)
	}
	elems = append(elems, elem)

	statusElems, err := NewStatusElements(v.Status)
	if err != nil {
		return fmt.Errorf("CStoreRsp.Encode: failed to create Status elements: %w", err)
	}
	elems = append(elems, statusElems...)

	elems = append(elems, v.Extra...)

	if err := EncodeElements(e, elems); err != nil {
		return fmt.Errorf("CStoreRsp.Encode: failed to encode elements: %w", err)
	}
	return nil
}

func (v *CStoreRsp) HasData() bool {
	return v.CommandDataSetType != CommandDataSetTypeNull
}

func (v *CStoreRsp) CommandField() uint16 {
	return CommandFieldCStoreRsp
}

func (v *CStoreRsp) GetMessageID() MessageID {
	return v.MessageIDBeingRespondedTo
}

func (v *CStoreRsp) GetStatus() *Status {
	return &v.Status
}

func (v *CStoreRsp) String() string {
	return fmt.Sprintf("CStoreRsp{AffectedSOPClassUID:%v MessageIDBeingRespondedTo:%v CommandDataSetType:%v AffectedSOPInstanceUID:%v Status:%v}}", v.AffectedSOPClassUID, v.MessageIDBeingRespondedTo, v.CommandDataSetType, v.AffectedSOPInstanceUID, v.Status)
}

func (CStoreRsp) decode(d *MessageDecoder) (*CStoreRsp, error) {
	v := &CStoreRsp{}
	var err error

	v.AffectedSOPClassUID, err = d.GetString(commandset.AffectedSOPClassUID, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cStoreRsp.decode: failed to decode AffectedSOPClassUID: %w", err)
	}

	v.MessageIDBeingRespondedTo, err = d.GetUInt16(commandset.MessageIDBeingRespondedTo, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cStoreRsp.decode: failed to decode MessageIDBeingRespondedTo: %w", err)
	}

	v.CommandDataSetType, err = d.GetUInt16(commandset.CommandDataSetType, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cStoreRsp.decode: failed to decode CommandDataSetType: %w", err)
	}

	v.AffectedSOPInstanceUID, err = d.GetString(commandset.AffectedSOPInstanceUID, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cStoreRsp.decode: failed to decode AffectedSOPInstanceUID: %w", err)
	}

	v.Status, err = d.GetStatus()
	if err != nil {
		return nil, fmt.Errorf("cStoreRsp.decode: failed to decode Status: %w", err)
	}

	v.Extra = d.UnparsedElements()
	return v, nil
}
