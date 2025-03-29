package dimse

import (
	"fmt"
	"io"

	"github.com/mlibanori/go-netdicom/commandset"
	"github.com/suyashkumar/dicom"
)

type CEchoRsp struct {
	MessageIDBeingRespondedTo MessageID
	CommandDataSetType        uint16
	Status                    Status
	Extra                     []*dicom.Element // Unparsed elements
}

func (v *CEchoRsp) Encode(e io.Writer) error {
	elems := []*dicom.Element{}

	elem, err := NewElement(commandset.CommandField, v.CommandField())
	if err != nil {
		return fmt.Errorf("CEchoRsp.Encode: failed to create CommandField element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.MessageIDBeingRespondedTo, v.MessageIDBeingRespondedTo)
	if err != nil {
		return fmt.Errorf("CEchoRsp.Encode: failed to create MessageIDBeingRespondedTo element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.CommandDataSetType, v.CommandDataSetType)
	if err != nil {
		return fmt.Errorf("CEchoRsp.Encode: failed to create CommandDataSetType element: %w", err)
	}
	elems = append(elems, elem)

	statusElems, err := NewStatusElements(v.Status)
	if err != nil {
		return fmt.Errorf("CEchoRsp.Encode: failed to create Status elements: %w", err)
	}
	elems = append(elems, statusElems...)

	elems = append(elems, v.Extra...)

	if err := EncodeElements(e, elems); err != nil {
		return fmt.Errorf("CEchoRsp.Encode: failed to encode elements: %w", err)
	}

	return nil
}

func (v *CEchoRsp) HasData() bool {
	return v.CommandDataSetType != CommandDataSetTypeNull
}

func (v *CEchoRsp) CommandField() uint16 {
	return CommandFieldCEchoRsp
}

func (v *CEchoRsp) GetMessageID() MessageID {
	return v.MessageIDBeingRespondedTo
}

func (v *CEchoRsp) GetStatus() *Status {
	return &v.Status
}

func (v *CEchoRsp) String() string {
	return fmt.Sprintf("CEchoRsp{MessageIDBeingRespondedTo:%v CommandDataSetType:%v Status:%v}}", v.MessageIDBeingRespondedTo, v.CommandDataSetType, v.Status)
}

func (CEchoRsp) decode(d *MessageDecoder) (*CEchoRsp, error) {
	v := &CEchoRsp{}
	var err error

	v.MessageIDBeingRespondedTo, err = d.GetUInt16(commandset.MessageIDBeingRespondedTo, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cEchoRsp.decode: failed to decode MessageIDBeingRespondedTo: %w", err)
	}

	v.CommandDataSetType, err = d.GetUInt16(commandset.CommandDataSetType, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("cEchoRsp.decode: failed to decode CommandDataSetType: %w", err)
	}

	v.Status, err = d.GetStatus()
	if err != nil {
		return nil, fmt.Errorf("cEchoRsp.decode: failed to decode Status: %w", err)
	}

	v.Extra = d.UnparsedElements()
	return v, nil
}
