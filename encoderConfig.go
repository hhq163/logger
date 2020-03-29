package logger

const (
	// unix timestamp EpochTime
	TimeFormat_TimeStamp = "timestamp"
	// 2016-01-02T00:00:00
	TimeFormat_ISO8601 = "iso8691"

	ShortCaller    = "shortCaller"
	FullPathCaller = "fullPathCaller"
)

type EncoderConfig struct {
	TimeFormat   string
	EncodeCaller string
}
