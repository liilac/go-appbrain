/*
 * AppBrain API
 *
 * API to interact with AppBrain.
 *
 * API version: 2.0
 * Contact: lilac@enby.ninja
 */

package appbrain

// AppInfoList describes a list of apps from the AppBrain API
type AppInfoList struct {
	// List of apps
	Apps []AppInfo `json:"apps,omitempty"`
}
