# \InformationApi

All URIs are relative to *https://api.appbrain.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrowseApps**](InformationApi.md#BrowseApps) | **Get** /info/browse | Apps information
[**GetApp**](InformationApi.md#GetApp) | **Get** /info/getapp | App information
[**GetCountries**](InformationApi.md#GetCountries) | **Get** /info/getcountries | Countries information
[**GetLibraries**](InformationApi.md#GetLibraries) | **Get** /info/getlibraries | Libraries information
[**SearchApps**](InformationApi.md#SearchApps) | **Get** /info/search | Apps information


# **BrowseApps**
> AppInfoList BrowseApps(ctx, optional)
Apps information

Returns a list of apps and everything about those apps. More information about usage and charges is available at <a href='/info/help/api/pricing.html'>AppBrain API pricing</a>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InformationApiBrowseAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InformationApiBrowseAppsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **optional.String**| sort order of the results (default: POPULAR) | 
 **filter** | **optional.String**| optional filter of the results | 
 **category** | **optional.String**| optional filter for the Google Play category | 
 **offset** | **optional.Int32**| offset of the browse results (default 0; max 500) | 
 **limit** | **optional.Int32**| number of the browse results (default 10; max 50) | 

### Return type

[**AppInfoList**](AppInfoList.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApp**
> AppInfo GetApp(ctx, package_)
App information

Returns all information about an app. More information about usage and charges is available at <a href='/info/help/api/pricing.html'>AppBrain API pricing</a>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **package_** | **string**| package | 

### Return type

[**AppInfo**](AppInfo.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCountries**
> CountryList GetCountries(ctx, )
Countries information

Returns a list of countries.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CountryList**](CountryList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLibraries**
> LibraryList GetLibraries(ctx, )
Libraries information

Returns a list of known libraries used in apps.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LibraryList**](LibraryList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchApps**
> AppInfoList SearchApps(ctx, optional)
Apps information

Returns a list of apps and everything about those apps. More information about usage and charges is available at <a href='/info/help/api/pricing.html'>AppBrain API pricing</a>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InformationApiSearchAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InformationApiSearchAppsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| search query | 
 **sort** | **optional.String**| sort order of the results (default: RELEVANCY) | 
 **filter** | **optional.String**| optional filter of the results | 
 **category** | **optional.String**| optional filter for the Google Play category | 
 **offset** | **optional.Int32**| offset of the browse results (default 0; max 500) | 
 **limit** | **optional.Int32**| number of the browse results (default 10; max 50) | 

### Return type

[**AppInfoList**](AppInfoList.md)

### Authorization

[apikey](../README.md#apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

