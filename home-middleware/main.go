package main

import (
	"context"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/grandcat/zeroconf"
)

func main() {
	// 1. MQTT Connection (Local Mosquitto)
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("Go_Home_Middleware")
	mqttClient := mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("MQTT Connection Error:", token.Error())
	}
	fmt.Println("✅ MQTT: Connected to Raspberry Pi Broker.")

	// 2. Blockchain Connection (Anvil/Hardhat)
	// Using the IP of your machine running Anvil
	client, err := ethclient.Dial("ws://10.206.160.62:8545")
	if err != nil {
		log.Fatal("Blockchain Connection Error:", err)
	}

	contractAddr := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	instance, err := NewAdvancedHomeAutomation(contractAddr, client)
	if err != nil {
		log.Fatal("Contract Instance Error:", err)
	}

	// 3. mDNS Discovery (Finding the ESP32)
	go func() {
		resolver, _ := zeroconf.NewResolver(nil)
		for {
			entries := make(chan *zeroconf.ServiceEntry)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			
			err := resolver.Browse(ctx, "_homeautomation._tcp", "local.", entries)
			if err == nil {
				go func() {
					for entry := range entries {
						fmt.Printf("[DISCOVERY] 🛰️ ESP32 Found: %s | IP: %v\n", entry.Instance, entry.AddrIPv4)
					}
				}()
			}
			<-ctx.Done()
			cancel()
			time.Sleep(30 * time.Second) // Re-scan every 30s
		}
	}()

	// 4. Watch Events
	// Ensure the sink matches the generated struct in contract.go
	sink := make(chan *AdvancedHomeAutomationStateChanged)
	
	// Note: We use 3 arguments here to match your specific abigen signature
	sub, err := instance.WatchStateChanged(nil, sink) 
	if err != nil {
		log.Fatal("Event Subscription Error:", err)
	}

	fmt.Println("?? System Online: Watching for Authenticated Commands...")

	for {
		select {
		case err := <-sub.Err():
			// If we get the "length insufficient" error, don't just crash. 
			// Log it and wait. This usually means an ABI mismatch.
			fmt.Printf("⚠️  Blockchain Sync Error: %v\n", err)
			time.Sleep(2 * time.Second)
			continue
			
		case event := <-sink:
			// Safely extract the data
			rId := event.RoomId.Uint64()
			dId := event.DeviceId.Uint64()
			val := event.NewValue.Uint64()

			fmt.Printf("\n[BLOCKCHAIN] 🔔 Event Received!")
			fmt.Printf("\n📍 Room: %d | 💡 Device: %d | ⚡ Value: %d\n", rId, dId, val)

			// Publish to MQTT
			topic := fmt.Sprintf("home/room%d/device%d", rId, dId)
			payload := fmt.Sprintf("%d", val)
			
			mqttClient.Publish(topic, 1, false, payload)
			fmt.Printf("📡 MQTT -> Published '%s' to '%s'\n", payload, topic)
		}
	}
}