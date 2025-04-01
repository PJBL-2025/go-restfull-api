package config

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/spf13/viper"
)

func InitMidtrans() *snap.Client {
	serverKey := viper.GetString("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		panic("MIDTRANS_SERVER_KEY is missing in config")
	}

	snapClient := snap.Client{}
	snapClient.New(serverKey, midtrans.Sandbox)

	return &snapClient
}
