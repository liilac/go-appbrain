# \PromotionCampaignsAPI

All URIs are relative to *https://api.appbrain.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](PromotionCampaignsAPI.md#Get) | **Get** /campaigns/get | Retrieves a promotion campaign
[**List**](PromotionCampaignsAPI.md#List) | **Get** /campaigns/list | Lists all promotion campaigns
[**Update**](PromotionCampaignsAPI.md#Update) | **Post** /campaigns/update | Creates or updates a promotion campaign


# **Get**
> PromotionCampaign Get(ctx, campaign)
Retrieves a promotion campaign

Retrieves the details of a promotion campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaign** | **string**| campaign | 

### Return type

[**PromotionCampaign**](PromotionCampaign.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> PromotionCampaignList List(ctx, )
Lists all promotion campaigns

Returns a list of identifiers of all promotion campaigns.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PromotionCampaignList**](PromotionCampaignList.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> string Update(ctx, body, optional)
Creates or updates a promotion campaign

Creates a new promotion campaign or updates an existing one.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PromotionCampaignUpdates**](PromotionCampaignUpdates.md)| the properties of the campaign that must be updated | 
 **optional** | ***PromotionCampaignsAPIUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PromotionCampaignsAPIUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lowbidaction** | **optional.String**| what to do if a country bid is lower than the minimum allowed bid (ERROR: throw an error, RAISE: raise the country bid, STOP: stop campaign in the country) | 

### Return type

**string**

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

