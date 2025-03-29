package dimse

import (
	"fmt"
	"io"

	"github.com/mlibanori/go-netdicom/commandset"
	"github.com/suyashkumar/dicom"
)

type CEchoRq struct {
	MessageID          MessageID
	CommandDataSetType uint16
	Extra              []*dicom.Element // Unparsed elements
}

func (v *CEchoRq) Encode(e io.Writer) error {
	elems := []*dicom.Element{}
	elem, err := NewElement(commandset.CommandField, v.CommandField())
	if err != nil {
		return fmt.Errorf("CEchoRq.Encode: failed to create CommandField element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.MessageID, v.MessageID)
	if err != nil {
		return fmt.Errorf("CEchoRq.Encode: failed to create MessageID element: %w", err)
	}
	elems = append(elems, elem)

	elem, err = NewElement(commandset.CommandDataSetType, v.CommandDataSetType)
	if err != nil {
		return fmt.Errorf("CEchoRq.Encode: failed to create CommandDataSetType element: %w", err)
	}
	elems = append(elems, elem)
	elems = append(elems, v.Extra...)
	if err := EncodeElements(e, elems); err != nil {
		return fmt.Errorf("CEchoRq.Encode: failed to encode elements: %w", err)
	}
	return nil
}

func (v *CEchoRq) HasData() bool {
	return v.CommandDataSetType != CommandDataSetTypeNull
}

func (v *CEchoRq) CommandField() uint16 {
	return CommandFieldCEchoRq
}

func (v *CEchoRq) GetMessageID() MessageID {
	return v.MessageID
}

func (v *CEchoRq) GetStatus() *Status {
	return nil
}

func (v *CEchoRq) String() string {
	return fmt.Sprintf("CEchoRq{MessageID:%v CommandDataSetType:%v}}", v.MessageID, v.CommandDataSetType)
}

func (CEchoRq) decode(d *MessageDecoder) (*CEchoRq, error) {
	v := &CEchoRq{}
	var err error
	v.MessageID, err = d.GetUInt16(commandset.MessageID, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("CEchoRq.decode: failed to get MessageID: %w", err)
	}

	v.CommandDataSetType, err = d.GetUInt16(commandset.CommandDataSetType, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("CEchoRq.decode: failed to get CommandDataSetType: %w", err)
	}
	v.Extra = d.UnparsedElements()
	return v, nil
}
