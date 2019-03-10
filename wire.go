//+build wireinject

package main

import "github.com/google/wire"

func InitializeServer(apiKey string) Server {
	wire.Build(NewWeatherServiceWire, NewServerWire)
	return new(server)
}
