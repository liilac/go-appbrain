# PromotionCampaignUpdates

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaign** | **string** | identifier of the promotion campaign (i.e. the package of the promoted app, eventually followed by a dash and an index for multiple campaigns per app.) | [default to null]
**State** | **string** | state of the promotion campaign | [optional] [default to null]
**Name** | **string** | name/description of the promotion campaign | [optional] [default to null]
**Title** | **string** | promotion title of the promotion campaign | [optional] [default to null]
**Subtitle** | **string** | promotion subtitle of the promotion campaign | [optional] [default to null]
**DailyBudget** | **float64** | daily budget of the promotion campaign in dollars (a negative value indicates no budget limit) | [optional] [default to null]
**RemainingBudget** | **float64** | total remaining budget the promotion campaign in dollars (a negative value indicates no budget limit) | [optional] [default to null]
**InstallTracking** | **string** | attribution partner that is used for install tracking | [optional] [default to null]
**ClickThroughURL** | **string** | clickthrough URL of the promotion campaign | [optional] [default to null]
**Countries** | [**[]Country**](Country.md) | list of settings for each country where the campaign is running; setting this list overwrites all country settings for the promotion campaign (can&#39;t be used simultaneously with updateCountries or removeCountries) | [optional] [default to null]
**UpdateCountries** | [**[]Country**](Country.md) | updates the provided settings for the countries (can&#39;t be used simultaneously with countries) | [optional] [default to null]
**RemoveCountries** | **[]string** | removes the provided countries from the promotion campaign (can&#39;t be used simultaneously with countries) | [optional] [default to null]
**AllowedSubIds** | **[]string** | list of publisher subIDs to allow for this campaign; setting this list overwrites all existing allowed subIDs (can&#39;t be used simultaneously with addAllowedSubIds) | [optional] [default to null]
**AddAllowedSubIds** | **[]string** | adds publisher subIDs to the list of allowed subIDs (can&#39;t be used simultaneously with allowedSubIds) | [optional] [default to null]
**BlockedSubIds** | **[]string** | list of publisher subIDs to block for this campaign; setting this list overwrites all existing blocked subIDs (can&#39;t be used simultaneously with addBlockedSubIds) | [optional] [default to null]
**AddBlockedSubIds** | **[]string** | adds publisher subIDs to the list of blocked subIDs (can&#39;t be used simultaneously with blockedSubIds) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


