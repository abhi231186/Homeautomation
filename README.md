# **Blockchain-Based Home Automation System - Complete Project Specification**

## **📋 Project Definition**

A **decentralized home automation platform** that leverages blockchain technology to provide secure, transparent, and immutable access control for IoT devices. Users can manage smart home appliances through a web interface while physical device control is synchronized via blockchain events and MQTT protocol.

---

## **🎯 Project Objective**

Create a trustless, decentralized system where:
- **Ownership is cryptographically proven** via wallet addresses
- **Access permissions are immutable** and time-bound on-chain
- **Device state changes are auditable** through blockchain events
- **Real-time hardware control** is achieved through event-driven architecture

---

## **🏗️ Architecture Overview**

### **Three-Tier Architecture:**

```
┌─────────────────────────────────────────────────────────────┐
│                     USER INTERFACE LAYER                     │
│  (Next.js 14 + React + Wagmi + Mantine + Tailwind CSS)     │
└────────────────────┬────────────────────────────────────────┘
                     │ Web3 RPC Calls
                     │ (Read/Write Contract)
┌────────────────────▼────────────────────────────────────────┐
│                   BLOCKCHAIN LAYER                           │
│   (Solidity Smart Contract on BSC Testnet/Mainnet)         │
│   - HomeAutomation.sol                                       │
│   - Access Control Logic                                     │
│   - State Management                                         │
└────────────────────┬────────────────────────────────────────┘
                     │ Event Emissions
                     │ (stateChange, multiStateChange)
┌────────────────────▼────────────────────────────────────────┐
│                  MIDDLEWARE LAYER                            │
│         (Go Service - Event Listener)                        │
│   - Listens to blockchain events                            │
│   - Decodes transaction data                                 │
│   - Publishes to MQTT broker                                 │
└────────────────────┬────────────────────────────────────────┘
                     │ MQTT Messages
                     │ (JSON payload with pin/state)
┌────────────────────▼────────────────────────────────────────┐
│                   HARDWARE LAYER                             │
│   (IoT Devices - ESP32/Arduino/Raspberry Pi)                │
│   - Subscribe to MQTT topics                                 │
│   - Control GPIO pins                                        │
│   - Execute physical state changes                           │
└─────────────────────────────────────────────────────────────┘
```

---

## **💻 Technology Stack**

### **1. Smart Contract Layer**
- **Language:** Solidity ^0.8.19
- **Framework:** Foundry (Forge, Cast, Anvil)
- **Blockchain:** Binance Smart Chain (BSC) Testnet/Mainnet
- **Development:** 
  - forge-std for testing
  - OpenZeppelin patterns for security

### **2. Frontend Application**
- **Framework:** Next.js 14 (App Router)
- **Language:** JavaScript/JSX
- **Web3 Integration:**
  - Wagmi 2.x (React hooks for Ethereum)
  - Viem (TypeScript Ethereum library)
  - @web3modal/wagmi (Wallet connection UI)
- **UI Libraries:**
  - Mantine 7.x (Component library)
  - Tailwind CSS (Utility-first styling)
  - @tabler/icons-react (Icons)
  - Material-UI/Joy UI (Additional components)
- **State Management:**
  - React Context API
  - TanStack Query (React Query)
- **Notifications:** react-hot-toast

### **3. Middleware Service**
- **Language:** Go 1.21+
- **Blockchain Client:** go-ethereum (geth)
- **MQTT Client:** eclipse/paho.mqtt.golang
- **Features:**
  - WebSocket connection to RPC
  - ABI encoding/decoding
  - Real-time block monitoring

### **4. Hardware Layer** (Not in current code but implied)
- **Devices:** ESP32, Arduino, Raspberry Pi
- **Protocol:** MQTT
- **Libraries:** PubSubClient (Arduino), paho-mqtt (Python)

---

## **🔐 Smart Contract Architecture**

### **Core Data Structures:**

```solidity
// User roles hierarchy
superAdmin (deployer)
  └── Room Owners (per room)
       └── Allowed Users (per appliance, time-bound)

// Room structure
struct Room {
    uint roomId;
    string roomName;
    uint noOfAppliances;
}

// Appliance structure
struct appliances {
    uint roomId;
    uint applianceId;
    uint pinNo;                          // GPIO pin number
    bool state;                          // ON/OFF for simple devices
    string details;                      // Name/description
    bool ismultiparameters;              // Fan speed, dimmer, etc.
    uint multiparams;                    // Current multi-param value
    uint8[5] multiparamPinNo;           // Pins for complex devices
    mapping(address => allowedAddresses) // Time-bound access
}

// Access control with time windows
struct allowedAddresses {
    bool isallowed;
    uint startTime;     // Unix timestamp
    uint endTime;       // 0 = no expiry
}
```

### **Access Control Modifiers:**

1. **`isSuperAdmin`** - Only contract deployer
2. **`isOwner(roomId)`** - SuperAdmin OR room owner
3. **`onlyAllowedAddress(roomId, applianceId)`** - SuperAdmin OR owner OR time-bound user

### **Key Functions:**

**Admin Functions:**
- `addRoom(roomName)` - Create new room
- `addRoomOwner(roomId, owner)` - Grant room ownership
- `removeRoomOwner(roomId, owner)` - Revoke ownership

**Room Owner Functions:**
- `addAppliances(...)` - Add device to room
- `editAppliances(...)` - Modify device config
- `editRoom(roomId, newName)` - Rename room
- `removeApplianceAccess(roomId, applianceId, user)` - Revoke user access

**User Functions:**
- `changeState(roomId, applianceId)` - Toggle simple device
- `changeState(roomId, applianceId, newState)` - Set multi-param value

**View Functions:**
- `getAllRooms()` - Get all rooms
- `getRoomCount()` - Total rooms
- `getApplianceCount(roomId)` - Appliances in room
- `getSwitchState(...)` - Get device state
- `checkAllowedAddress(...)` - Verify access

---

## **📱 Frontend Application Flow**

### **Page Structure:**

```
/                       → Landing/Login (Wallet Connect)
/dashboard              → User overview
/admin                  → List owned rooms
/room                   → All rooms summary
/room/[roomid]          → Appliance control panel
/user                   → User profile
```

### **Component Hierarchy:**

```
RootLayout (layout.js)
├── MantineProvider
├── WagmiProvider (Web3ModalProvider)
│   ├── QueryClientProvider
│   └── OwnedRoomsProvider (Context)
│       ├── Navbar
│       │   ├── Logo
│       │   ├── LinksGroup (Dynamic room links)
│       │   └── UserButton
│       └── Page Components
```

### **User Flows:**

#### **Flow 1: Wallet Connection**
1. User visits `/` (redirects to `/login`)
2. Clicks "Connect" → Web3Modal shows wallet options
3. Selects wallet (MetaMask, WalletConnect, etc.)
4. Signs message for authentication
5. Redirected to `/dashboard`

#### **Flow 2: View Owned Rooms**
1. `OwnedRoomsProvider` fetches:
   - `getRoomCount()` from contract
   - For each room: `owner[roomId][userAddress]` check
   - `rooms[roomId]` for room details
2. Navbar dynamically generates room links
3. User sees owned rooms in sidebar

#### **Flow 3: Control Appliance**
1. User navigates to `/room/[roomid]`
2. Page fetches:
   - Room details from context
   - For each appliance:
     - `getApplianceDetails()` (name)
     - `getSwitchState()` (current state)
     - `checkIsMultiparams()` (device type)
3. User clicks toggle button
4. Frontend calls `writeContract()`:
   ```javascript
   writeContract({
     address: CONTRACT_ADDRESS,
     abi: abi,
     functionName: 'changeState',
     args: [roomId, applianceId]
   })
   ```
5. Transaction submitted → Wallet confirmation
6. Transaction mined → Event emitted
7. Frontend refetches state and shows toast

#### **Flow 4: Add New Room** (SuperAdmin)
1. Admin calls `addRoom("Living Room")`
2. Contract creates room with `roomId = rooms.length`
3. Sets admin as room owner
4. Emits `AddRoomOwner` event
5. Frontend context refetches room list

#### **Flow 5: Grant Appliance Access**
1. Room owner calls `addAppliances()` with:
   - `roomId`, `details` (name), `pinNo`
   - `allowedAddresses[]` (user addresses)
   - `allowedAddresses[]` (time windows)
   - `ismultiparams`, `multiparamPinNo`
2. Contract stores appliance data
3. Maps each user to access struct
4. User can now control appliance if within time window

---

## **⚙️ Middleware Service Flow**

### **Go Service Responsibilities:**

```go
1. Connect to BSC WebSocket RPC
2. Subscribe to new blocks
3. For each new block:
   a. Get all transactions
   b. Filter: tx.To() == CONTRACT_ADDRESS
   c. Check function selector (changeState methods)
   d. Decode input arguments:
      - roomId
      - applianceId
      - (optional) newState for multi-params
   e. Read current state from contract:
      - getSwitchState(roomId, applianceId)
      - getPinNo(roomId, applianceId)
      - checkIsMultiparams(...)
      - getMultiparamsState(...)
   f. Construct MQTT payload:
      {
        "pin": <GPIO_PIN>,
        "mode": 1/0,
        "ismultistate": bool,
        "multistate": <value>,
        "roomId": <id>,
        "contract": <address>
      }
   g. Publish to MQTT broker
```

### **MQTT Message Format:**

```json
{
  "pin": 14,
  "mode": 1,
  "ismultistate": false,
  "multistate": 0,
  "roomId": 0,
  "contract": "0x..."
}
```

### **Hardware Response:**
- ESP32 subscribed to topic receives JSON
- Parses `pin` and `mode`
- Sets `digitalWrite(pin, mode)`
- Device physically turns ON/OFF

---

## **🌊 Complete System Flow (End-to-End)**

### **Example: User Turns On Living Room Light**

```
1. USER ACTION:
   User clicks toggle in browser → Frontend

2. FRONTEND:
   writeContract('changeState', [0, 2]) → Sends transaction

3. WALLET:
   User confirms transaction → Broadcasts to network

4. BLOCKCHAIN:
   Transaction mined in block #12345678
   ├── Contract executes changeState()
   ├── Updates state: data[0][2].state = true
   └── Emits: stateChange(0, 2, true)

5. GO MIDDLEWARE:
   Detects new block → Scans transactions
   ├── Finds changeState call
   ├── Decodes: roomId=0, applianceId=2
   ├── Reads: getSwitchState(0,2) → true
   ├── Reads: getPinNo(0,2) → 14
   └── Publishes MQTT: {"pin":14, "mode":1, ...}

6. MQTT BROKER:
   Receives message → Forwards to subscribers

7. HARDWARE (ESP32):
   Receives MQTT message
   ├── Parses JSON
   ├── digitalWrite(14, HIGH)
   └── LED/Relay ON → Light turns ON physically

8. FRONTEND (Optional):
   Listens to events or refetches state
   └── Shows "Light is ON" with green toggle
```

---

## **🎨 Key Features**

### **1. Decentralized Access Control**
- No central server can deny access
- Ownership verified cryptographically
- Time-bound guest access (start/end timestamps)
- Immutable audit trail

### **2. Multi-Parameter Device Support**
- Simple ON/OFF switches
- Variable devices (fan speed, dimmer brightness)
- Multiple GPIO pins per device

### **3. Role-Based Permissions**
- **SuperAdmin**: Full control, can add rooms/owners
- **Room Owner**: Manage appliances, grant access
- **Allowed User**: Control specific appliances in time window

### **4. Real-Time Synchronization**
- Blockchain events trigger hardware updates
- Frontend polls or subscribes to state changes
- MQTT ensures low-latency device control

### **5. Wallet-Based Authentication**
- No passwords to remember
- Sign-in with MetaMask, WalletConnect, etc.
- Web3 native UX

### **6. Responsive UI**
- Mobile-first design
- Drawer navigation on small screens
- Real-time state updates
- Toast notifications for transactions

---

## **🔒 Security Considerations**

1. **Smart Contract:**
   - Access modifiers prevent unauthorized changes
   - Time-based access expires automatically
   - No reentrancy risks (no external calls in state changes)
   - Events for all critical operations

2. **Frontend:**
   - Wallet signature required for all writes
   - Read-only operations via public RPC
   - No private keys stored in browser

3. **Middleware:**
   - Only reads blockchain (no write access)
   - MQTT can use TLS + authentication
   - Service runs server-side, not exposed to users

4. **Hardware:**
   - MQTT credentials managed per device
   - GPIO pins isolated per appliance
   - Physical failsafes (circuit breakers)

---

## **📊 Data Flow Diagram**

```
┌──────────┐     Transaction      ┌────────────┐
│  User    │────────────────────→│ Blockchain │
│ (Wallet) │                      │  Contract  │
└──────────┘                      └─────┬──────┘
     ↑                                  │
     │ Read State                       │ Event Emission
     │                                  │
┌────┴──────┐                    ┌──────▼─────┐
│ Frontend  │                    │    Go      │
│  Next.js  │                    │ Middleware │
└───────────┘                    └──────┬─────┘
                                        │
                                        │ MQTT Publish
                                        │
                                  ┌─────▼──────┐
                                  │   MQTT     │
                                  │   Broker   │
                                  └─────┬──────┘
                                        │
                                        │ Subscribe
                                        │
                                  ┌─────▼──────┐
                                  │  Hardware  │
                                  │  (ESP32)   │
                                  └────────────┘
```

---

## **🚀 Deployment Strategy**

### **Phase 1: Smart Contract**
1. Configure Foundry with BSC testnet RPC
2. Add deployment script
3. Deploy from superAdmin wallet
4. Verify contract on BSCScan
5. Note deployed address

### **Phase 2: Frontend**
1. Update `config.js` with contract address
2. Update ABI from compiled contract
3. Configure environment variables
4. Deploy to Vercel/Netlify
5. Connect custom domain

### **Phase 3: Middleware**
1. Build Go binary
2. Deploy to VPS/cloud instance
3. Configure environment variables
4. Run as systemd service
5. Monitor logs for errors

### **Phase 4: Hardware**
1. Flash firmware to devices
2. Configure WiFi + MQTT credentials
3. Map GPIO pins to appliances
4. Physical installation

---

## **📈 Scalability Considerations**

1. **Blockchain:**
   - BSC has ~3s block time
   - Gas costs are low (~$0.01-0.10 per transaction)
   - Can migrate to Layer 2 if needed

2. **Frontend:**
   - Static site generation for fast loads
   - RPC calls can use rate-limited public nodes or private Infura/Alchemy

3. **Middleware:**
   - Horizontally scalable (multiple Go instances)
   - Can filter by room/contract to reduce load
   - MQTT broker handles thousands of connections

4. **Hardware:**
   - Each device subscribes to specific topics
   - MQTT QoS levels ensure delivery
   - Local fallback if internet drops

---

## **🛠️ Future Enhancements**

1. **Smart Contract:**
   - Transfer superAdmin ownership
   - Batch operations for gas savings
   - ERC-721/1155 for room/appliance NFTs
   - Delegation patterns for sub-admins

2. **Frontend:**
   - Energy usage analytics
   - Scheduling/automation rules
   - Voice control integration
   - Mobile native apps (React Native)

3. **Middleware:**
   - Support multiple contracts/chains
   - Caching layer for state reads
   - WebSocket API for frontend

4. **Hardware:**
   - Sensor data back to blockchain
   - Mesh networking for offline operation
   - Energy monitoring

---

## **📁 Project Structure**

```
blockchain-home-automation/
├── smart-contracts/          # Solidity contracts and tests
│   ├── src/
│   │   └── HomeAutomation.sol
│   ├── test/
│   │   └── HomeAutomation.t.sol
│   ├── script/
│   │   └── Deploy.s.sol
│   ├── foundry.toml
│   └── .env.example
│
├── frontend/                 # Next.js web application
│   ├── app/
│   │   ├── layout.js
│   │   ├── page.js
│   │   ├── login/
│   │   ├── dashboard/
│   │   ├── room/
│   │   └── admin/
│   ├── components/
│   ├── context/
│   ├── config.js
│   ├── package.json
│   └── .env.local.example
│
├── middleware/               # Go event listener service
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── .env.example
│
├── hardware/                 # IoT device firmware (optional)
│   └── esp32/
│       └── mqtt_controller.ino
│
└── README.md                 # This file
```

---

## **🎓 Learning Resources**

- **Solidity:** https://docs.soliditylang.org/
- **Foundry:** https://book.getfoundry.sh/
- **Next.js:** https://nextjs.org/docs
- **Wagmi:** https://wagmi.sh/
- **Go Ethereum:** https://geth.ethereum.org/docs/
- **MQTT:** https://mqtt.org/

---

## **📝 License**

This project is open-source and available under the MIT License.

---

## **🤝 Contributing**

Contributions, issues, and feature requests are welcome! Feel free to check the issues page.

---

## **👨‍💻 Author**

Your Name / Team Name

---

**Built with ❤️ using Blockchain, Web3, and IoT technologies**
# Homeautomation
