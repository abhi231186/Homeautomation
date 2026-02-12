// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/access/AccessControl.sol";

contract AdvancedHomeAutomation is AccessControl {
    // Roles using Keccak256 hashes
    bytes32 public constant SUPER_ADMIN_ROLE = keccak256("SUPER_ADMIN_ROLE");
    bytes32 public constant ROOM_ADMIN_ROLE = keccak256("ROOM_ADMIN_ROLE");
    bytes32 public constant GUEST_ROLE = keccak256("GUEST_ROLE");

    enum DeviceType { OnOff, Fan, Dimmer, RGB }

    struct Device {
        string name;
        uint256 pinNo;
        DeviceType dType;
        uint256 value;
        bool exists;
    }

    struct Room {
        string name;
        string espIP;
        uint256 deviceCount;
        mapping(uint256 => Device) devices;
        bool exists;
    }

    struct AccessRule {
        uint256 fromTimestamp;
        uint256 toTimestamp;
        bool isActive;
    }

    mapping(uint256 => Room) public rooms;
    uint256 public roomCount;
    
    mapping(uint256 => mapping(address => AccessRule)) public accessRules;

    // --- CRITICAL CHANGE: NO INDEXED PARAMETERS ---
    // This ensures all 3 uint256s (32+32+32 = 96 bytes) are sent in the DATA field.
    event StateChanged(uint256 roomId, uint256 deviceId, uint256 newValue);
    
    // We keep these indexed for easier searching, but StateChanged is the one 
    // the middleware watches constantly.
    event AccessUpdated(uint256 indexed roomId, address indexed user, uint256 from, uint256 to, bytes32 role);
    event RoomCreated(uint256 indexed roomId, string name);

    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(SUPER_ADMIN_ROLE, msg.sender);
    }

    // --- 1. SUPER ADMIN MANAGEMENT ---
    function addSuperAdmin(address newAdmin) external onlyRole(SUPER_ADMIN_ROLE) {
        _grantRole(SUPER_ADMIN_ROLE, newAdmin);
    }

    function removeSuperAdmin(address admin) external onlyRole(SUPER_ADMIN_ROLE) {
        require(admin != msg.sender, "Cannot remove yourself");
        _revokeRole(SUPER_ADMIN_ROLE, admin);
    }

    // --- 2. ROOM & DEVICE SETUP ---
    function createRoom(string memory _name, string memory _ip) external onlyRole(SUPER_ADMIN_ROLE) {
        roomCount++;
        rooms[roomCount].name = _name;
        rooms[roomCount].espIP = _ip;
        rooms[roomCount].exists = true;
        emit RoomCreated(roomCount, _name);
    }

    function defineDevice(uint256 _roomId, string memory _name, uint256 _pin, DeviceType _type) external onlyRole(SUPER_ADMIN_ROLE) {
        require(rooms[_roomId].exists, "Room 404");
        uint256 dId = ++rooms[_roomId].deviceCount;
        rooms[_roomId].devices[dId] = Device({
            name: _name,
            pinNo: _pin,
            dType: _type,
            value: 0,
            exists: true
        });
    }

    // --- 3. PERMISSION LAYER ---
    function grantAccess(
        uint256 _roomId, 
        address _user, 
        uint256 _start, 
        uint256 _end, 
        bytes32 _role
    ) external {
        require(hasRole(SUPER_ADMIN_ROLE, msg.sender), "Only SuperAdmin can assign roles");
        
        accessRules[_roomId][_user] = AccessRule({
            fromTimestamp: _start,
            toTimestamp: _end,
            isActive: true
        });
        _grantRole(_role, _user);
        
        emit AccessUpdated(_roomId, _user, _start, _end, _role);
    }

    function revokeAccess(uint256 _roomId, address _user, bytes32 _role) external onlyRole(SUPER_ADMIN_ROLE) {
        accessRules[_roomId][_user].isActive = false;
        _revokeRole(_role, _user);
    }

    // --- 4. CONTROL LOGIC ---
    function operateDevice(uint256 _roomId, uint256 _deviceId, uint256 _value) external {
        require(rooms[_roomId].devices[_deviceId].exists, "Device 404");

        if (hasRole(SUPER_ADMIN_ROLE, msg.sender)) {
            _executeCommand(_roomId, _deviceId, _value);
            return;
        }

        AccessRule storage rule = accessRules[_roomId][msg.sender];
        require(rule.isActive, "Access Inactive");

        bool timeValid = (rule.fromTimestamp == 0 || block.timestamp >= rule.fromTimestamp) && 
                         (rule.toTimestamp == 0 || block.timestamp <= rule.toTimestamp);

        if (hasRole(ROOM_ADMIN_ROLE, msg.sender)) {
            _executeCommand(_roomId, _deviceId, _value);
        } else if (hasRole(GUEST_ROLE, msg.sender) && timeValid) {
            _executeCommand(_roomId, _deviceId, _value);
        } else {
            revert("Access Denied: Role/Time Invalid");
        }
    }

    function _executeCommand(uint256 _rId, uint256 _dId, uint256 _val) internal {
        rooms[_rId].devices[_dId].value = _val;
        emit StateChanged(_rId, _dId, _val);
    }

    function getDeviceStatus(uint256 _rId, uint256 _dId) external view returns (uint256) {
        return rooms[_rId].devices[_dId].value;
    }
}