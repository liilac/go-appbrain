/*
 * AppBrain API
 *
 * API to interact with AppBrain.
 *
 * API version: 2.0
 * Contact: lilac@enby.ninja
 */

package appbrain

// PromotionCampaignUpdates describes proposed updates for a promotion campaign
type PromotionCampaignUpdates struct {
	// identifier of the promotion campaign (i.e. the package of the promoted app, eventually followed by a dash and an index for multiple campaigns per app.)
	Campaign string `json:"campaign"`
	// state of the promotion campaign
	State string `json:"state,omitempty"`
	// name/description of the promotion campaign
	Name string `json:"name,omitempty"`
	// promotion title of the promotion campaign
	Title string `json:"title,omitempty"`
	// promotion subtitle of the promotion campaign
	Subtitle string `json:"subtitle,omitempty"`
	// daily budget of the promotion campaign in dollars (a negative value indicates no budget limit)
	DailyBudget float64 `json:"dailyBudget,omitempty"`
	// total remaining budget the promotion campaign in dollars (a negative value indicates no budget limit)
	RemainingBudget float64 `json:"remainingBudget,omitempty"`
	// attribution partner that is used for install tracking
	InstallTracking string `json:"installTracking,omitempty"`
	// clickthrough URL of the promotion campaign
	ClickThroughURL string `json:"clickthroughUrl,omitempty"`
	// list of settings for each country where the campaign is running; setting this list overwrites all country settings for the promotion campaign (can't be used simultaneously with updateCountries or removeCountries)
	Countries []Country `json:"countries,omitempty"`
	// updates the provided settings for the countries (can't be used simultaneously with countries)
	UpdateCountries []Country `json:"updateCountries,omitempty"`
	// removes the provided countries from the promotion campaign (can't be used simultaneously with countries)
	RemoveCountries []string `json:"removeCountries,omitempty"`
	// list of publisher subIDs to allow for this campaign; setting this list overwrites all existing allowed subIDs (can't be used simultaneously with addAllowedSubIds)
	AllowedSubIds []string `json:"allowedSubIds,omitempty"`
	// adds publisher subIDs to the list of allowed subIDs (can't be used simultaneously with allowedSubIds)
	AddAllowedSubIds []string `json:"addAllowedSubIds,omitempty"`
	// list of publisher subIDs to block for this campaign; setting this list overwrites all existing blocked subIDs (can't be used simultaneously with addBlockedSubIds)
	BlockedSubIds []string `json:"blockedSubIds,omitempty"`
	// adds publisher subIDs to the list of blocked subIDs (can't be used simultaneously with blockedSubIds)
	AddBlockedSubIds []string `json:"addBlockedSubIds,omitempty"`
}
