// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract AdvancedHomeAutomation is AccessControl {
    bytes32 public constant SUPER_ADMIN_ROLE = keccak256("SUPER_ADMIN_ROLE");
    bytes32 public constant ROOM_OWNER_ROLE = keccak256("ROOM_OWNER_ROLE");

    enum DeviceType { OnOff, Fan, Dimmer, RGB }

    struct Device {
        string name;
        uint256 pinNo;
        DeviceType dType;
        uint256 value; // 0-1 for OnOff, 0-5 for Fan, 0-100 for Dimmer, Hex for RGB
        bool exists;
    }

    struct Room {
        string name;
        string espIP; // To map which ESP controls this room
        uint256 deviceCount;
        mapping(uint256 => Device) devices;
        bool exists;
    }

    struct AccessRule {
        uint256 fromTimestamp; // 0 if no bound
        uint256 toTimestamp;   // 0 if no bound
        bool isActive;
    }

    mapping(uint256 => Room) public rooms;
    uint256 public roomCount;
    
    // RoomID => User => AccessRule
    mapping(uint256 => mapping(address => AccessRule)) public accessRules;

    // Logs for every action as per flowchart
    event StateChanged(uint256 indexed roomId, uint256 deviceId, uint256 newValue);
    event AccessUpdated(uint256 indexed roomId, address user, uint256 from, uint256 to);
    event RoleTransferred(bytes32 indexed role, address indexed from, address indexed to);

    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(SUPER_ADMIN_ROLE, msg.sender);
    }

    // --- Admin Management ---
    function addSuperAdmin(address newAdmin) external onlyRole(SUPER_ADMIN_ROLE) {
        _grantRole(SUPER_ADMIN_ROLE, newAdmin);
    }

    function transferSuperAdmin(address to) external onlyRole(SUPER_ADMIN_ROLE) {
        _grantRole(SUPER_ADMIN_ROLE, to);
        _revokeRole(SUPER_ADMIN_ROLE, msg.sender);
    }

    // --- Room & Device Setup (Super Admin Only) ---
    function createRoom(string memory _name, string memory _ip) external onlyRole(SUPER_ADMIN_ROLE) {
        roomCount++;
        rooms[roomCount].name = _name;
        rooms[roomCount].espIP = _ip;
        rooms[roomCount].exists = true;
    }

    function defineDevice(uint256 _roomId, string memory _name, uint256 _pin, DeviceType _type) external onlyRole(SUPER_ADMIN_ROLE) {
        uint256 dId = ++rooms[_roomId].deviceCount;
        rooms[_roomId].devices[dId] = Device(_name, _pin, _type, 0, true);
    }

    // --- Permission Layer ---
    function grantRoomAccess(uint256 _roomId, address _user, uint256 _start, uint256 _end, bool isOwner) external {
        require(hasRole(SUPER_ADMIN_ROLE, msg.sender) || 
               (hasRole(ROOM_OWNER_ROLE, msg.sender) && accessRules[_roomId][msg.sender].isActive), "Unauthorized");
        
        accessRules[_roomId][_user] = AccessRule(_start, _end, true);
        if(isOwner) _grantRole(ROOM_OWNER_ROLE, _user);
        emit AccessUpdated(_roomId, _user, _start, _end);
    }

    // --- Control Logic ---
    function operateDevice(uint256 _roomId, uint256 _deviceId, uint256 _value) external {
        AccessRule storage rule = accessRules[_roomId][msg.sender];
        bool timeValid = (rule.fromTimestamp == 0 || block.timestamp >= rule.fromTimestamp) && 
                         (rule.toTimestamp == 0 || block.timestamp <= rule.toTimestamp);
        
        require(hasRole(SUPER_ADMIN_ROLE, msg.sender) || (rule.isActive && timeValid), "Access Denied");

        rooms[_roomId].devices[_deviceId].value = _value;
        emit StateChanged(_roomId, _deviceId, _value);
    }
}