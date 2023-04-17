package config

const (
	// AdafruitIOUsername is the username of the Adafruit IO account
	AdafruitIOUsername = "luutodinh"
	// AdafruitBaseURL is the base URL of the Adafruit IO API
	AdafruitBaseURL = "https://io.adafruit.com/api/v2/"
)

const (
	// TemperatureFeedName is the name of the temperature feed
	TemperatureFeedName = "cambiennhietdo"
	// HumidityFeedName is the name of the humidity feed
	HumidityFeedName = "cambiendoam"
	// InfraredFeedName is the name of the infrared feed
	InfraredFeedName = "phathiennguoi"
	// TemperatureThresholdFeedName is the name of the temperature threshold feed
	TemperatureThresholdFeedName = "nguongnhietdo"
	// HumidityThresholdFeedName is the name of the humidity threshold feed
	HumidityThresholdFeedName = "nguongdoam"
	// NebulizerFeedName is the name of the nebulizer feed
	NebulizerFeedName = "nutbomnuoc"
	// FanFeedName is the name of the fan feed
	FanFeedName = "nutquat"
)

func GetAdafruitIOFeedURL(feedName string) string {
	return AdafruitBaseURL + AdafruitIOUsername + "/feeds/" + feedName + "/data?limit=1&include=value,id,created_at"
}
