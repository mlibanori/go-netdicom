package commandset

import (
	"github.com/suyashkumar/dicom/pkg/tag"
)

var (
	CommandGroupLength                   = tag.Tag{Group: 0x0000, Element: 0x0000}
	AffectedSOPClassUID                  = tag.Tag{Group: 0x0000, Element: 0x0002}
	RequestedSOPClassUID                 = tag.Tag{Group: 0x0000, Element: 0x0003}
	CommandField                         = tag.Tag{Group: 0x0000, Element: 0x0100}
	MessageID                            = tag.Tag{Group: 0x0000, Element: 0x0110}
	MessageIDBeingRespondedTo            = tag.Tag{Group: 0x0000, Element: 0x0120}
	MoveDestination                      = tag.Tag{Group: 0x0000, Element: 0x0600}
	Priority                             = tag.Tag{Group: 0x0000, Element: 0x0700}
	CommandDataSetType                   = tag.Tag{Group: 0x0000, Element: 0x0800}
	Status                               = tag.Tag{Group: 0x0000, Element: 0x0900}
	OffendingElement                     = tag.Tag{Group: 0x0000, Element: 0x0901}
	ErrorComment                         = tag.Tag{Group: 0x0000, Element: 0x0902}
	ErrorID                              = tag.Tag{Group: 0x0000, Element: 0x0903}
	AffectedSOPInstanceUID               = tag.Tag{Group: 0x0000, Element: 0x1000}
	RequestedSOPInstanceUID              = tag.Tag{Group: 0x0000, Element: 0x1001}
	EventTypeID                          = tag.Tag{Group: 0x0000, Element: 0x1002}
	AttributeIdentifierList              = tag.Tag{Group: 0x0000, Element: 0x1005}
	ActionTypeID                         = tag.Tag{Group: 0x0000, Element: 0x1008}
	NumberOfRemainingSuboperations       = tag.Tag{Group: 0x0000, Element: 0x1020}
	NumberOfCompletedSuboperations       = tag.Tag{Group: 0x0000, Element: 0x1021}
	NumberOfFailedSuboperations          = tag.Tag{Group: 0x0000, Element: 0x1022}
	NumberOfWarningSuboperations         = tag.Tag{Group: 0x0000, Element: 0x1023}
	MoveOriginatorApplicationEntityTitle = tag.Tag{Group: 0x0000, Element: 0x1030}
	MoveOriginatorMessageID              = tag.Tag{Group: 0x0000, Element: 0x1031}
)

var tagInfos = []tag.Info{
	{
		Name:    "CommandGroupLength",
		Tag:     CommandGroupLength,
		Keyword: "Command Group Length",
		VRs:     []string{"UL"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "AffectedSOPClassUID",
		Tag:     AffectedSOPClassUID,
		Keyword: "Affected SOP Class UID",
		VRs:     []string{"UI"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "RequestedSOPClassUID",
		Tag:     RequestedSOPClassUID,
		Keyword: "Requested SOP Class UID",
		VRs:     []string{"UI"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "CommandField",
		Tag:     CommandField,
		Keyword: "Command Field",
		VRs:     []string{"US"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "MessageID",
		Tag:     MessageID,
		Keyword: "Message ID",
		VRs:     []string{"US"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "MessageIDBeingRespondedTo",
		Tag:     MessageIDBeingRespondedTo,
		Keyword: "Message ID Being Responded To",
		VRs:     []string{"US"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "MoveDestination",
		Tag:     MoveDestination,
		Keyword: "Move Destination",
		VRs:     []string{"AE"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "Priority",
		Tag:     Priority,
		Keyword: "Priority",
		VRs:     []string{"US"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "CommandDataSetType",
		Tag:     CommandDataSetType,
		Keyword: "Command Data Set Type",
		VRs:     []string{"US"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "Status",
		Tag:     Status,
		Keyword: "Status",
		VRs:     []string{"US"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "OffendingElement",
		Tag:     OffendingElement,
		Keyword: "Offending Element",
		VRs:     []string{"AT"},
		VM:      "1-n",
		Retired: false,
	},
	{
		Name:    "ErrorComment",
		Tag:     ErrorComment,
		Keyword: "Error Comment",
		VRs:     []string{"LO"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "ErrorID",
		Tag:     ErrorID,
		Keyword: "Error ID",
		VRs:     []string{"US"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "AffectedSOPInstanceUID",
		Tag:     AffectedSOPInstanceUID,
		Keyword: "Affected SOP Instance UID",
		VRs:     []string{"UI"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "RequestedSOPInstanceUID",
		Tag:     RequestedSOPInstanceUID,
		Keyword: "Requested SOP Instance UID",
		VRs:     []string{"UI"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "EventTypeID",
		Tag:     EventTypeID,
		Keyword: "Event Type ID",
		VRs:     []string{"US"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "AttributeIdentifierList",
		Tag:     AttributeIdentifierList,
		Keyword: "Attribute Identifier List",
		VRs:     []string{"AT"},
		VM:      "1-n",
		Retired: false,
	},
	{
		Name:    "ActionTypeID",
		Tag:     ActionTypeID,
		Keyword: "Action Type ID",
		VRs:     []string{"US"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "NumberOfRemainingSuboperations",
		Tag:     NumberOfRemainingSuboperations,
		Keyword: "Number of Remaining Suboperations",
		VRs:     []string{"US"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "NumberOfCompletedSuboperations",
		Tag:     NumberOfCompletedSuboperations,
		Keyword: "Number of Completed Suboperations",
		VRs:     []string{"US"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "NumberOfFailedSuboperations",
		Tag:     NumberOfFailedSuboperations,
		Keyword: "Number of Failed Suboperations",
		VRs:     []string{"US"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "NumberOfWarningSuboperations",
		Tag:     NumberOfWarningSuboperations,
		Keyword: "Number of Warning Suboperations",
		VRs:     []string{"US"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "MoveOriginatorApplicationEntityTitle",
		Tag:     MoveOriginatorApplicationEntityTitle,
		Keyword: "Move Originator Application Entity Title",
		VRs:     []string{"AE"},
		VM:      "1",
		Retired: false,
	},
	{
		Name:    "MoveOriginatorMessageID",
		Tag:     MoveOriginatorMessageID,
		Keyword: "Move Originator Message ID",
		VRs:     []string{"US"},
		VM:      "1",
		Retired: false,
	},
}

func Init() {
	for _, info := range tagInfos {
		tag.Add(info, false)
	}
}
