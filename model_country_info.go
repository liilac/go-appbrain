/*
 * AppBrain API
 *
 * API to interact with AppBrain.
 *
 * API version: 2.0
 * Contact: lilac@enby.ninja
 */

package appbrain

// CountryInfo describes a country used in a promotion campaign
type CountryInfo struct {
	// country code
	Code string `json:"code,omitempty"`
	// name of the country
	Name string `json:"name,omitempty"`
	// minimum cost-per-install bid for app promotions in this country
	MinCpiBid float64 `json:"minCpiBid,omitempty"`
}
