/*
 * AppBrain API
 *
 * API to interact with AppBrain.
 *
 * API version: 2.0
 * Contact: lilac@enby.ninja
 */

package appbrain

// Library describes a library known to be used in apps
type Library struct {
	// identifier for the library
	ID string `json:"id,omitempty"`
	// name of the library
	Name string `json:"name,omitempty"`
	// description of the library
	Description string `json:"description,omitempty"`
	// types of the library
	Types []string `json:"types,omitempty"`
	// tags associated with the library
	Tags []string `json:"tags,omitempty"`
	// official URL of the library
	OfficialURL string `json:"officialUrl,omitempty"`
	// AppBrain URL of the library
	AppBrainURL string `json:"appBrainUrl,omitempty"`
}
