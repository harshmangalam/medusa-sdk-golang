package schema

import "time"

// A Reason for why a given product is returned. A Return Reason can be used on Return Items in order to indicate why a Line Item was returned.

type ReturnReason struct {
	// The value to identify the reason by.
	Value string `json:"value"`

	// A text that can be displayed to the Customer as a reason.
	Label string `json:"Label"`

	// The cart's ID
	Id string `json:"id"`

	// A description of the Reason.
	Description string `json:"description"`

	// The ID of the parent reason.
	ParentReturnReason *ReturnReason `json:"parent_return_reason"`

	// A Reason for why a given product is returned. A Return Reason can be used on Return Items in order to indicate why a Line Item was returned.

	ParentReturnReasonId string `json:"parent_return_reason_id"`

	// A Reason for why a given product is returned. A Return Reason can be used on Return Items in order to indicate why a Line Item was returned.

	ReturnReasonChildren *ReturnReason `json:"return_reason_children"`

	// The date with timezone at which the resource was created.
	CreatedAt *time.Time `json:"created_at"`

	// The date with timezone at which the resource was updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// The date with timezone at which the resource was deleted.
	DeletedAt *time.Time `json:"deleted_at"`

	// An optional key-value map with additional details
	Metadata map[string]any `json:"metadata"`
}
