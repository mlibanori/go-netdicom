package dimse

// Code generated from generate_dimse_messages.py. DO NOT EDIT.

import (
	"fmt"
	"io"
	"bytes"

	"github.com/mlibanori/go-netdicom/commandset"
	"github.com/mlibanori/go-netdicom/pdu"
	"github.com/suyashkumar/dicom"
	dicomtag "github.com/suyashkumar/dicom/pkg/tag"
)

const (
CommandFieldCStoreRq	uint16	= 0x0001
CommandFieldCStoreRsp uint16	= 0x8001
CommandFieldCFindRq		uint16	= 0x0020
CommandFieldCFindRsp	uint16	= 0x8020
CommandFieldCGetRq		uint16	= 0x0010
CommandFieldCGetRsp		uint16	= 0x8010
CommandFieldCMoveRq		uint16	= 0x0021
CommandFieldCMoveRsp 	uint16	= 0x8021
CommandFieldCEchoRq 	uint16	= 0x0030
CommandFieldCEchoRsp 	uint16	= 0x8030
)

func DecodeMessageForType(d *MessageDecoder, commandField uint16) (Message, error) {
	switch commandField {
	case CommandFieldCStoreRq:
		return CStoreRq{}.decode(d)
	case CommandFieldCStoreRsp:
		return CStoreRsp{}.decode(d)
	case CommandFieldCFindRq:
		return CFindRq{}.decode(d)
	case CommandFieldCFindRsp:
		return CFindRsp{}.decode(d)
	case CommandFieldCGetRq:
		return CGetRq{}.decode(d)
	case CommandFieldCGetRsp:
		return CGetRsp{}.decode(d)
	case CommandFieldCMoveRq:
		return CMoveRq{}.decode(d)
	case CommandFieldCMoveRsp:
		return CMoveRsp{}.decode(d)
	case CommandFieldCEchoRq:
		return CEchoRq{}.decode(d)
	case CommandFieldCEchoRsp:
		return CEchoRsp{}.decode(d)
	default:
		return nil, fmt.Errorf("Unknown DIMSE command 0x%x", commandField)
	}
}

// CommandAssembler is a helper that assembles a DIMSE command message and data
// payload from a sequence of P_DATA_TF PDUs.
type CommandAssembler struct {
	contextID      byte
	commandBytes   []byte
	command        Message
	dataBytes      []byte
	readAllCommand bool

	readAllData bool
}

func (commandAssembler *CommandAssembler) AddDataPDU(pdu *pdu.PDataTf) (byte, Message, []byte, error) {
	for _, item := range pdu.Items {
		if commandAssembler.contextID == 0 {
			commandAssembler.contextID = item.ContextID
		} else if commandAssembler.contextID != item.ContextID {
			return 0, nil, nil, fmt.Errorf("Mixed context: %d %d", commandAssembler.contextID, item.ContextID)
		}
		if item.Command {
			commandAssembler.commandBytes = append(commandAssembler.commandBytes, item.Value...)
			if item.Last {
				if commandAssembler.readAllCommand {
					return 0, nil, nil, fmt.Errorf("P_DATA_TF: found >1 command chunks with the Last bit set")
				}
				commandAssembler.readAllCommand = true
			}
		} else {
			commandAssembler.dataBytes = append(commandAssembler.dataBytes, item.Value...)
			if item.Last {
				if commandAssembler.readAllData {
					return 0, nil, nil, fmt.Errorf("P_DATA_TF: found >1 data chunks with the Last bit set")
				}
				commandAssembler.readAllData = true
			}
		}
	}
	if !commandAssembler.readAllCommand {
		return 0, nil, nil, nil
	}
	if commandAssembler.command == nil {
		ioReader := bytes.NewReader(commandAssembler.commandBytes)
		parser, err := dicom.Parse(ioReader,int64(ioReader.Len()), nil, dicom.SkipPixelData(), dicom.SkipMetadataReadOnNewParserInit(), dicom.SkipTransferSyntaxDetection())
		if err != nil {
			return 0, nil, nil, fmt.Errorf("P_DATA_TF: failed to parse command bytes: %w", err)
		}
		commandAssembler.command, err = ReadMessage(&parser)
		if err != nil {
			return 0, nil, nil, err
		}
	}
	if commandAssembler.command.HasData() && !commandAssembler.readAllData {
		return 0, nil, nil, nil
	}
	contextID := commandAssembler.contextID
	command := commandAssembler.command
	dataBytes := commandAssembler.dataBytes
	*commandAssembler = CommandAssembler{}
	return contextID, command, dataBytes, nil
	// TODO(saito) Verify that there's no unread items after the last command&data.
}

// ReadMessage constructs a typed Message object, given a set of
// dicom.Elements,
func ReadMessage(dataset *dicom.Dataset) (message Message, err error) {
	mDecoder := MessageDecoder{
		unparsed: make(map[dicomtag.Tag]*dicom.Element),
		parsed:   make(map[dicomtag.Tag]*dicom.Element),
	}
	for _, elem := range dataset.Elements {
		tag := elem.Tag
		mDecoder.unparsed[tag] = elem
	}
	commandField, err := mDecoder.GetUInt16(commandset.CommandField, RequiredElement)
	if err != nil {
		return nil, fmt.Errorf("ReadMessage: failed to get command field: %w", err)
	}
	return DecodeMessageForType(&mDecoder, uint16(commandField))
}

type MessageID = uint16


// Message defines the common interface for all DIMSE message types.
type Message interface {
	fmt.Stringer // Print human-readable description for debugging.
	Encode(io.Writer) error
	// GetMessageID extracts the message ID field.
	GetMessageID() MessageID
	// CommandField returns the command field value of this message.
	CommandField() uint16
	// GetStatus returns the the response status value. It is nil for request message
	// types, and non-nil for response message types.
	GetStatus() *Status
	// HasData is true if we expect P_DATA_TF packets after the command packets.
	HasData() bool
}

// Status represents a result of a DIMSE call.  P3.7 C defines list of status
// codes and error payloads.
type Status struct {
	// Status==StatusSuccess on success. A non-zero value on error.
	Status StatusCode

	// Optional error payloads.
	ErrorComment string // Encoded as (0000,0902)
}

// Success is an OK status for a call.
var Success = Status{Status: StatusSuccess}

type StatusCode uint16

const (
	StatusSuccess               StatusCode = 0
	StatusCancel                StatusCode = 0xFE00
	StatusSOPClassNotSupported  StatusCode = 0x0112
	StatusInvalidArgumentValue  StatusCode = 0x0115
	StatusInvalidAttributeValue StatusCode = 0x0106
	StatusInvalidObjectInstance StatusCode = 0x0117
	StatusUnrecognizedOperation StatusCode = 0x0211
	StatusNotAuthorized         StatusCode = 0x0124
	StatusPending               StatusCode = 0xff00

	// C-STORE-specific status codes. P3.4 GG4-1
	CStoreOutOfResources              StatusCode = 0xa700
	CStoreCannotUnderstand            StatusCode = 0xc000
	CStoreDataSetDoesNotMatchSOPClass StatusCode = 0xa900

	// C-FIND-specific status codes.
	CFindUnableToProcess StatusCode = 0xc000

	// C-MOVE/C-GET-specific status codes.
	CMoveOutOfResourcesUnableToCalculateNumberOfMatches StatusCode = 0xa701
	CMoveOutOfResourcesUnableToPerformSubOperations     StatusCode = 0xa702
	CMoveMoveDestinationUnknown                         StatusCode = 0xa801
	CMoveDataSetDoesNotMatchSOPClass                    StatusCode = 0xa900

	// Warning codes.
	StatusAttributeValueOutOfRange StatusCode = 0x0116
	StatusAttributeListError       StatusCode = 0x0107
)

// Helper class for extracting values from a list of DicomElement.
type MessageDecoder struct {
	unparsed  map[dicomtag.Tag]*dicom.Element
	parsed 		map[dicomtag.Tag]*dicom.Element
}

type isOptionalElement int

const (
	RequiredElement isOptionalElement = iota
	OptionalElement
)

func (d *MessageDecoder) UnparsedElements() []*dicom.Element {
	elems := make([]*dicom.Element, 0, len(d.unparsed))
	for _, elem := range d.unparsed {
		elems = append(elems, elem)
	}
	return elems
}

func (d *MessageDecoder) FindElementByTag(tag dicomtag.Tag) (*dicom.Element) {
	elem, found := d.unparsed[tag]
	if found {
		return elem
	}
	elem, found = d.parsed[tag]
	if found {
		return elem
	}
	return nil
}

func (d *MessageDecoder) MarkElementAsParsed(tag dicomtag.Tag){
	elem, found := d.unparsed[tag]
	if found {
		d.parsed[tag] = elem
		delete(d.unparsed, tag)
	}
}


func (d *MessageDecoder) GetStatus() (s Status, err error) {
	statusCode, err := d.GetUInt16(commandset.Status, RequiredElement)
	if err != nil {
		return s, fmt.Errorf("GetStatus: failed to get status code: %w", err)
	}
	s.Status = StatusCode(statusCode)
	s.ErrorComment, err = d.GetString(commandset.ErrorComment, OptionalElement)
	if err != nil {
		return s, fmt.Errorf("GetStatus: failed to get error comment: %w", err)
	}
	return s, nil
}

func (d *MessageDecoder) GetString(tag dicomtag.Tag, optional isOptionalElement) (string, error) {
	elem := d.FindElementByTag(tag)
	if elem == nil {
		if optional == RequiredElement {
			return "", fmt.Errorf("GetString: tag %s not found", tag.String())
		}
		return "", nil
	}
	if elem.Value == nil {
		return "", fmt.Errorf("GetString: tag %s has no value", tag.String())
	}
	rawValue := elem.Value.GetValue()
	if rawValue == nil {
		return "", fmt.Errorf("GetString: tag %s has a nil value", tag.String())
	}
	v, ok := rawValue.([]string)
	if !ok {
		return "", fmt.Errorf("GetString: failed to convert tag %s to []string, got %d", tag.String(), elem.Value.ValueType())
	}
	if len(v) == 0 {
		return "", nil
	}
	d.MarkElementAsParsed(tag)
	return v[0], nil
}

// Find an element with "tag", and extract a uint16 from it. Errors are reported in d.err.
func (d *MessageDecoder) GetUInt16(tag dicomtag.Tag, optional isOptionalElement) (uint16, error) {
	elem := d.FindElementByTag(tag)
	if elem == nil {
		if optional == RequiredElement {
			return 0, fmt.Errorf("GetUInt16: tag %s not found", tag.String())
		}
		return 0, nil
	}
	if elem == nil {
		return 0, fmt.Errorf("GetUInt16: tag %s not found", tag.String())
	}
	if elem.Value == nil {
		return 0, fmt.Errorf("GetUInt16: tag %s has no value", tag.String())
	}
	if elem.Value.ValueType() != dicom.Ints  {
		return 0, fmt.Errorf("GetUInt16: element %s is not an int, got %v", tag.String(), elem.Value.ValueType())
	}
	rawValue := elem.Value.GetValue()
	if rawValue == nil {
		return 0, fmt.Errorf("GetUInt16: tag %s has a nil value", tag.String())
	}
	v, ok := rawValue.([]int)
	if !ok {
		return 0, fmt.Errorf("GetUInt16: failed to convert tag %s to []int", tag.String())
	}
	if len(v) == 0 {
		return 0, nil
	}
	if v[0] < 0 || v[0] > 65535 {
		return 0, fmt.Errorf("GetUInt16: value %v is out of range for uint16", v)
	}
	d.MarkElementAsParsed(tag)
	return uint16(v[0]), nil
}

