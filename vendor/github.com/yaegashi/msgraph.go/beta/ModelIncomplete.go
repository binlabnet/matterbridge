// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// IncompleteData undocumented
type IncompleteData struct {
	// Object is the base model of IncompleteData
	Object
	// MissingDataBeforeDateTime undocumented
	MissingDataBeforeDateTime *time.Time `json:"missingDataBeforeDateTime,omitempty"`
	// WasThrottled undocumented
	WasThrottled *bool `json:"wasThrottled,omitempty"`
}
