package dimse

//go:generate ./generate_dimse_messages.py
//go:generate stringer -type StatusCode

// Implements message types defined in P3.7.
//
// http://dicom.nema.org/medical/dicom/current/output/pdf/part07.pdf

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"sort"

	"github.com/mlibanori/go-netdicom/commandset"
	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
)

// Encode the given elements. The elements are sorted in ascending tag order.
func EncodeElements(e io.Writer, elems []*dicom.Element) error {
	writer, err := dicom.NewWriter(e)
	if err != nil {
		return fmt.Errorf("EncodeElements: failed to create writer: %w", err)
	}
	writer.SetTransferSyntax(binary.LittleEndian, true)
	sort.Slice(elems, func(i, j int) bool {
		return elems[i].Tag.Compare(elems[j].Tag) < 0
	})
	for _, elem := range elems {
		if err := writer.WriteElement(elem); err != nil {
			return fmt.Errorf("EncodeElements: error writing element %s: %w", elem.Tag.String(), err)
		}

	}
	return nil
}

// Create a list of elements that represent the dimse status. The list contains
// multiple elements for non-ok status.
func NewStatusElements(s Status) ([]*dicom.Element, error) {
	statusElement, err := NewElement(commandset.Status, s.Status)
	if err != nil {
		return nil, fmt.Errorf("NewStatusElements: error creating status element with status %v: %w", s.Status, err)
	}
	elems := []*dicom.Element{statusElement}
	if s.ErrorComment != "" {
		errorCommentElement, err := NewElement(commandset.ErrorComment, s.ErrorComment)
		if err != nil {
			return nil, fmt.Errorf("NewStatusElements: error creating error comment element with comment %v: %w", s.ErrorComment, err)
		}
		elems = append(elems, errorCommentElement)
	}
	return elems, nil
}

func NewElement(tag tag.Tag, value any) (*dicom.Element, error) {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return dicom.NewElement(tag, []string{value.(string)})
	case reflect.Int:
		return dicom.NewElement(tag, []int{value.(int)})
	case reflect.Uint16, reflect.Uint8:
		return dicom.NewElement(tag, []int{int(v.Uint())})
	default:
		if v.CanConvert(reflect.TypeOf(uint16(0))) {
			return dicom.NewElement(tag, []int{int(v.Convert(reflect.TypeOf(uint16(0))).Uint())})
		}
		if v.CanConvert(reflect.TypeOf(uint8(0))) {
			return dicom.NewElement(tag, []int{int(v.Convert(reflect.TypeOf(uint8(0))).Uint())})
		}
		if v.CanConvert(reflect.TypeOf(int(0))) {
			return dicom.NewElement(tag, []int{int(v.Convert(reflect.TypeOf(int(0))).Int())})
		}
		return nil, fmt.Errorf("NewElement: unsupported type %T for tag %s", value, tag)
	}
}

// CommandDataSetTypeNull indicates that the DIMSE message has no data payload,
// when set in dicom.TagCommandDataSetType. Any other value indicates the
// existence of a payload.
const CommandDataSetTypeNull uint16 = 0x101

// CommandDataSetTypeNonNull indicates that the DIMSE message has a data
// payload, when set in dicom.TagCommandDataSetType.
const CommandDataSetTypeNonNull uint16 = 1

// StatusCode represents a DIMSE service response code, as defined in P3.7

// EncodeMessage serializes the given message. Errors are reported through e.Error()
func EncodeMessage(out io.Writer, v Message) error {
	writer, err := dicom.NewWriter(out)
	if err != nil {
		return fmt.Errorf("EncodeMessage: error creating writer: %w", err)
	}
	subEncoderBuffer := bytes.Buffer{}
	if err := v.Encode(&subEncoderBuffer); err != nil {
		return fmt.Errorf("EncodeMessage: error encoding message: %w", err)
	}
	// DIMSE messages are always encoded Implicit+LE. See P3.7 6.3.1.
	writer.SetTransferSyntax(binary.LittleEndian, true)
	element, err := NewElement(commandset.CommandGroupLength, subEncoderBuffer.Len())
	if err != nil {
		return fmt.Errorf("EncodeMessage: failed to create CommandGroupLength element: %w", err)
	}
	writer.WriteElement(element)
	out.Write(subEncoderBuffer.Bytes())
	return nil
}

// AddDataPDU is to be called for each P_DATA_TF PDU received from the
// network. If the fragment is marked as the last one, AddDataPDU returns
// <SOPUID, TransferSyntaxUID, payload, nil>.  If it needs more fragments, it
// returns <"", "", nil, nil>.  On error, it returns a non-nil error.
