package main

type StockConfig struct {
	Name      string
	Symbol    string
	TrackTime int // seconds ?
}

type Keys struct {
	SengridApiKey string
}

type Transport struct {
	TargetEmail string
	TargetPhone string
}

type Config struct {
	Stocks []StockConfig
	Keys
	Transport
}

func loadConfig() Config {
	return Config{}
}
