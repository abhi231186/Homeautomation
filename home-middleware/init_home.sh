#!/bin/bash

# --- CONFIGURATION ---
RPC="http://127.0.0.1:8545"
# Use your SuperAdmin Private Key (Account #0 from Anvil)
PK="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
# PASTE YOUR NEW CONTRACT ADDRESS HERE (From Step 2)
CONTRACT="0x5FbDB2315678afecb367f032d93F642f64180aa3"

# Role Hashes (Must match your Solidity code)
ROOM_ADMIN_ROLE="0x93786134b2f8a848c467e4125b281f6d90045618451551555155515551555155" # ROOM_ADMIN_ROLE
GUEST_ROLE="0x822448342468303f83196c8ee9347895e5461230113c4155b1115598144214f5"      # GUEST_ROLE

echo "--- PHASE 1: ROOM DETECTION & CREATION ---"
# createRoom(name, ip)
cast send $CONTRACT "createRoom(string,string)" "LivingRoom" "10.206.160.62" --rpc-url $RPC --private-key $PK

echo "--- PHASE 2: DEFINING HARDWARE COMPONENTS ---"
# defineDevice(roomId, name, gpio, type) | Type 0:Bulb, 1:Fan, 2:RGB

echo "Defining Fan (Device 1)..."
cast send $CONTRACT "defineDevice(uint256,string,uint256,uint8)" 1 "Main Fan" 14 1 --rpc-url $RPC --private-key $PK

echo "Defining Bulb (Device 2)..."
cast send $CONTRACT "defineDevice(uint256,string,uint256,uint8)" 1 "Main Light" 33 0 --rpc-url $RPC --private-key $PK

echo "Defining Smart Plug (Device 3)..."
cast send $CONTRACT "defineDevice(uint256,string,uint256,uint8)" 1 "Smart Plug" 18 0 --rpc-url $RPC --private-key $PK

echo "Defining RGB Strip (Device 4)..."
cast send $CONTRACT "defineDevice(uint256,string,uint256,uint8)" 1 "Mood Lighting" 5 2 --rpc-url $RPC --private-key $PK

echo "--- PHASE 3: SETTING UP GUEST ACCESS (ABAC) ---"
# Granting Guest access for 1 hour starting from now
CURRENT_TIME=$(date +%s)
EXPIRY_TIME=$((CURRENT_TIME + 3600))
GUEST_ADDR="0x70997970C51812dc3A010C7d01b50e0d17dc79C8"

echo "Granting 1-Hour Guest Access to $GUEST_ADDR..."
# grantAccess(roomId, user, start, end, roleHash)
cast send $CONTRACT "grantAccess(uint256,address,uint256,uint256,bytes32)" 1 $GUEST_ADDR $CURRENT_TIME $EXPIRY_TIME $GUEST_ROLE --rpc-url $RPC --private-key $PK

echo "--- FLOW COMPLETE: SYSTEM READY FOR CONTROL ---"