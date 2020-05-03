/*
 * AppBrain API
 *
 * API to interact with AppBrain.
 *
 * API version: 2.0
 * Contact: lilac@enby.ninja
 */

package appbrain

// PromotionCampaign describes a promotion campaign on the AppBrain platform
type PromotionCampaign struct {
	// identifier of the promotion campaign
	Campaign string `json:"campaign,omitempty"`
	// state of the promotion campaign
	State string `json:"state,omitempty"`
	// name/description of the promotion campaign
	Name string `json:"name,omitempty"`
	// promotion title of the promotion campaign
	Title string `json:"title,omitempty"`
	// promotion subtitle of the promotion campaign
	Subtitle string `json:"subtitle,omitempty"`
	// daily budget of the promotion campaign in dollars
	DailyBudget float64 `json:"dailyBudget,omitempty"`
	// total remaining budget the promotion campaign in dollars
	RemainingBudget float64 `json:"remainingBudget,omitempty"`
	// attribution partner that is used for install tracking
	InstallTracking string `json:"installTracking,omitempty"`
	// clickthrough URL of the promotion campaign
	ClickThroughURL string `json:"clickthroughUrl,omitempty"`
	// list of settings for each country where the campaign is running
	Countries []Country `json:"countries,omitempty"`
	// list of publisher subIDs to allow for this campaign
	AllowedSubIds []string `json:"allowedSubIds,omitempty"`
	// list of publisher subIDs to block for this campaign
	BlockedSubIds []string `json:"blockedSubIds,omitempty"`
}
