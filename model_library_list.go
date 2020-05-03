/*
 * AppBrain API
 *
 * API to interact with AppBrain.
 *
 * API version: 2.0
 * Contact: lilac@enby.ninja
 */

package appbrain

// LibraryList describes a list of known libraries used in apps
type LibraryList struct {
	// library information
	Libraries []Library `json:"libraries,omitempty"`
}
