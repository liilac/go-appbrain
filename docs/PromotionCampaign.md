# PromotionCampaign

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaign** | **string** | identifier of the promotion campaign | [optional] [default to null]
**State** | **string** | state of the promotion campaign | [optional] [default to null]
**Name** | **string** | name/description of the promotion campaign | [optional] [default to null]
**Title** | **string** | promotion title of the promotion campaign | [optional] [default to null]
**Subtitle** | **string** | promotion subtitle of the promotion campaign | [optional] [default to null]
**DailyBudget** | **float64** | daily budget of the promotion campaign in dollars | [optional] [default to null]
**RemainingBudget** | **float64** | total remaining budget the promotion campaign in dollars | [optional] [default to null]
**InstallTracking** | **string** | attribution partner that is used for install tracking | [optional] [default to null]
**ClickThroughURL** | **string** | clickthrough URL of the promotion campaign | [optional] [default to null]
**Countries** | [**[]Country**](Country.md) | list of settings for each country where the campaign is running | [optional] [default to null]
**AllowedSubIds** | **[]string** | list of publisher subIDs to allow for this campaign | [optional] [default to null]
**BlockedSubIds** | **[]string** | list of publisher subIDs to block for this campaign | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


