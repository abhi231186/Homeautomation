package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/grandcat/zeroconf"
)

func main() {
	// 1. MQTT Connection
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	mqttClient := mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	fmt.Println("MQTT: ?? Connected to Mosquitto.")

	// 2. Blockchain Connection
	client, err := ethclient.Dial("ws://10.23.104.62:8545")
	if err != nil {
		log.Fatal(err)
	}

	contractAddr := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	instance, err := NewAdvancedHomeAutomation(contractAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Robust Discovery (The Eyes)
	go func() {
		resolver, _ := zeroconf.NewResolver(nil)
		for {
			entries := make(chan *zeroconf.ServiceEntry)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			
			err := resolver.Browse(ctx, "_homeautomation._tcp", "local.", entries)
			if err != nil {
				log.Println("Discovery Error:", err)
			}

			go func() {
				for entry := range entries {
					fmt.Printf("\n[DISCOVERY] ?? Found ESP32: %s | IP: %v\n", entry.Instance, entry.AddrIPv4)
				}
			}()

			<-ctx.Done()
			cancel()
			time.Sleep(15 * time.Second)
		}
	}()

	// 4. Watch Events (The Ears)
	sink := make(chan *AdvancedHomeAutomationStateChanged)
	sub, err := instance.WatchStateChanged(nil, sink, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("System Online: ?? Watching for commands...")

	for {
		select {
		case err := <-sub.Err():
			log.Println("Subscription Error (restarting...):", err)
			return // Let the Pi restart the script
		case event := <-sink:
			rId := event.RoomId.Uint64()
			dId := event.DeviceId.Uint64()
			val := event.NewValue.Uint64()

			fmt.Printf("\n[BLOCKCHAIN] Command: Room %d | Device %d | Value %d\n", rId, dId, val)

			topic := fmt.Sprintf("home/room%d/device%d", rId, dId)
			payload := fmt.Sprintf("%d", val)
			mqttClient.Publish(topic, 1, false, payload)
			fmt.Printf("MQTT: ?? Sent '%s' to '%s'\n", payload, topic)
		}
	}
}