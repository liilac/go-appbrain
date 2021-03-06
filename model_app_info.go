/*
 * AppBrain API
 *
 * API to interact with AppBrain.
 *
 * API version: 2.0
 * Contact: lilac@enby.ninja
 */

package appbrain

// AppInfo describes details of an app from the AppBrain API
type AppInfo struct {
	// package
	Pkg string `json:"package,omitempty"`
	// most recent version
	VersionString string `json:"versionString,omitempty"`
	// most recent version code
	VersionCode int32 `json:"versionCode,omitempty"`
	// name of the app
	Name string `json:"name,omitempty"`
	// short one-line description of the app
	ShortDescription string `json:"shortDescription,omitempty"`
	// description of the app
	Description string `json:"description,omitempty"`
	// market category of the app
	MarketCategory string `json:"marketCategory,omitempty"`
	// price of the app
	Price string `json:"price,omitempty"`
	// size of the APK
	ApkSize int64 `json:"apkSize,omitempty"`
	// maturity level of the app
	MaturityLevel string `json:"maturityLevel,omitempty"`
	// minimum supported SDK version
	MinSdkVersion int32 `json:"minSdkVersion,omitempty"`
	// maximum supported SDK version
	MaxSdkVersion int32 `json:"maxSdkVersion,omitempty"`
	// whether the app can be installed on external storage
	SupportsApp2SD string `json:"supportsApp2SD,omitempty"`
	// timestamp of the app's launch in seconds since epoch
	LaunchTime int32 `json:"launchTime,omitempty"`
	// timestamp of the app's latest update in seconds since epoch
	LastAppUpdateTime int32 `json:"lastAppUpdateTime,omitempty"`
	// timestamp of the last refresh of the app information in seconds since epoch
	InfoRefreshTime int32 `json:"infoRefreshTime,omitempty"`
	// Google play category of the number of downloads
	DownloadsCategory string `json:"downloadsCategory,omitempty"`
	// estimated number of downloads
	EstimatedDownloads int64 `json:"estimatedDownloads,omitempty"`
	// estimated number of downloads in the last 30 days
	EstimatedRecentDownloads int64 `json:"estimatedRecentDownloads,omitempty"`
	// average rating
	Rating float64 `json:"rating,omitempty"`
	// average rating in the last 30 days
	RecentRating float64 `json:"recentRating,omitempty"`
	// number of ratings
	RatingCount int32 `json:"ratingCount,omitempty"`
	// number of 1-star ratings
	Rating1StarCount int32 `json:"rating1StarCount,omitempty"`
	// number of 2-star ratings
	Rating2StarCount int32 `json:"rating2StarCount,omitempty"`
	// number of 3-star ratings
	Rating3StarCount int32 `json:"rating3StarCount,omitempty"`
	// number of 4-star ratings
	Rating4StarCount int32 `json:"rating4StarCount,omitempty"`
	// number of 5-star ratings
	Rating5StarCount int32 `json:"rating5StarCount,omitempty"`
	// number of +1s
	PlusOneCount int32 `json:"plusOneCount,omitempty"`
	// number of comments
	CommentCount int32 `json:"commentCount,omitempty"`
	// URL of the app's icon
	IconURL string `json:"iconUrl,omitempty"`
	// URL of the app's feature graphic
	FeatureGraphicURL string `json:"featureGraphicUrl,omitempty"`
	// YouTube URL of the app's promo video
	YouTubeURL string `json:"youtubeUrl,omitempty"`
	// name of the developer of the app
	DeveloperName string `json:"developerName,omitempty"`
	// email address of the developer
	DeveloperEmail string `json:"developerEmail,omitempty"`
	// address of the developer
	DeveloperAddress string `json:"developerAddress,omitempty"`
	// official website of the app
	Website string `json:"website,omitempty"`
	// URLs of the app's screenshots
	ScreenshotUrls []string `json:"screenshotUrls,omitempty"`
	// requested permissions
	Permissions []string `json:"permissions,omitempty"`
	// libraries used in the app
	Libraries []string `json:"libraries,omitempty"`
}
