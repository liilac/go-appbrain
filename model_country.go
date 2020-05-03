/*
 * AppBrain API
 *
 * API to interact with AppBrain.
 *
 * API version: 2.0
 * Contact: lilac@enby.ninja
 */

package appbrain

// Country defines the settings for a country in a promotion campaign
type Country struct {
	// the two-letter ISO country code
	Code string `json:"code,omitempty"`
	// the maximum cost-per-install for a country in dollars
	Cpi float64 `json:"cpi,omitempty"`
}
