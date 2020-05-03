/*
 * AppBrain API
 *
 * API to interact with AppBrain.
 *
 * API version: 2.0
 * Contact: lilac@enby.ninja
 */

package appbrain

// PromotionCampaignList describes a list of promotion campaigns
type PromotionCampaignList struct {
	// identifiers of existing promotion campaigns
	Campaigns []string `json:"campaigns,omitempty"`
}
