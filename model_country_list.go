/*
 * AppBrain API
 *
 * API to interact with AppBrain.
 *
 * API version: 2.0
 * Contact: lilac@enby.ninja
 */

package appbrain

// CountryList describes a list of countries for a promotion campaign
type CountryList struct {
	// country information
	Countries []CountryInfo `json:"countries,omitempty"`
}
