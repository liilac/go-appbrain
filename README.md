# Go API client for swagger

API to interact with AppBrain.

## Overview
This API client interacts with the [AppBrain API](https://www.appbrain.com/info/help/api/appbrain-api.html).

The client was generated using [swagger-codegen](https://github.com/swagger-api/swagger-codegen) and then manually cleaned up and refined to meet code quality standards.

- API version: 2.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
`go get` the packace and import using the URL of this git repository
````golang
import "github.com/liilac/go-appbrain"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.appbrain.com/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*InformationAPI* | [**BrowseApps**](docs/InformationAPI.md#browseapps) | **Get** /info/browse | Apps information
*InformationAPI* | [**GetApp**](docs/InformationAPI.md#getapp) | **Get** /info/getapp | App information
*InformationAPI* | [**GetCountries**](docs/InformationAPI.md#getcountries) | **Get** /info/getcountries | Countries information
*InformationAPI* | [**GetLibraries**](docs/InformationAPI.md#getlibraries) | **Get** /info/getlibraries | Libraries information
*InformationAPI* | [**SearchApps**](docs/InformationAPI.md#searchapps) | **Get** /info/search | Apps information
*PromotionCampaignsAPI* | [**Get**](docs/PromotionCampaignsAPI.md#get) | **Get** /campaigns/get | Retrieves a promotion campaign
*PromotionCampaignsAPI* | [**List**](docs/PromotionCampaignsAPI.md#list) | **Get** /campaigns/list | Lists all promotion campaigns
*PromotionCampaignsAPI* | [**Update**](docs/PromotionCampaignsAPI.md#update) | **Post** /campaigns/update | Creates or updates a promotion campaign


## Documentation For Models

 - [AppInfo](docs/AppInfo.md)
 - [AppInfoList](docs/AppInfoList.md)
 - [Country](docs/Country.md)
 - [CountryInfo](docs/CountryInfo.md)
 - [CountryList](docs/CountryList.md)
 - [Library](docs/Library.md)
 - [LibraryList](docs/LibraryList.md)
 - [PromotionCampaign](docs/PromotionCampaign.md)
 - [PromotionCampaignList](docs/PromotionCampaignList.md)
 - [PromotionCampaignUpdates](docs/PromotionCampaignUpdates.md)


## Documentation For Authorization

## apikey
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

Contact Lilac Kapul at lilac@enby.ninja

