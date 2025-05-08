package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// IIOTDevice represents an Industrial IoT device
type IIOTDevice struct {
	IP         string `json:"ip"`
	Name       string `json:"name"`
	DeviceType string `json:"device_type"`
	State      string `json:"state"` // active, blocked, deleted
}

type SmartContract struct {
	contractapi.Contract
}

// =======================
// Create Composite Key
// =======================

func (s *SmartContract) createDeviceKey(ctx contractapi.TransactionContextInterface, state, ip string) (string, error) {
	return ctx.GetStub().CreateCompositeKey("Device", []string{state, ip})
}

// ===========================
// Check if device exists
// ===========================

func (s *SmartContract) DeviceExists(ctx contractapi.TransactionContextInterface, state, ip string) (bool, error) {
	key, err := s.createDeviceKey(ctx, state, ip)
	if err != nil {
		return false, err
	}
	data, err := ctx.GetStub().GetState(key)
	if err != nil {
		return false, err
	}
	return data != nil, nil
}

// =======================
// InitLedger
// =======================

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	devices := []IIOTDevice{
		{IP: "192.168.0.1", Name: "TempSensor-A1", DeviceType: "Sensor", State: "active"},
		{IP: "192.168.0.2", Name: "PressureSensor-B2", DeviceType: "Sensor", State: "active"},
		{IP: "192.168.0.3", Name: "ValveController-C3", DeviceType: "Controller", State: "active"},
		{IP: "192.168.0.4", Name: "FlowMeter-D4", DeviceType: "Meter", State: "active"},
		{IP: "192.168.0.5", Name: "HumiditySensor-E5", DeviceType: "Sensor", State: "active"},
		{IP: "192.168.0.6", Name: "GasDetector-F6", DeviceType: "Detector", State: "active"},
		{IP: "192.168.0.7", Name: "VibrationSensor-G7", DeviceType: "Sensor", State: "active"},
		{IP: "192.168.0.8", Name: "PumpController-H8", DeviceType: "Controller", State: "active"},
		{IP: "192.168.0.9", Name: "Actuator-I9", DeviceType: "Actuator", State: "active"},
		{IP: "192.168.0.10", Name: "Motor-J10", DeviceType: "Motor", State: "active"},
		{IP: "192.168.0.11", Name: "PLC-K11", DeviceType: "Controller", State: "active"},
		{IP: "192.168.0.12", Name: "Camera-L12", DeviceType: "Sensor", State: "active"},
		{IP: "192.168.0.13", Name: "RFIDReader-M13", DeviceType: "Reader", State: "active"},
		{IP: "192.168.0.14", Name: "ThermalSensor-N14", DeviceType: "Sensor", State: "active"},
		{IP: "192.168.0.15", Name: "UltrasonicSensor-O15", DeviceType: "Sensor", State: "active"},
		{IP: "192.168.0.16", Name: "RoboticArm-P16", DeviceType: "Actuator", State: "active"},
		{IP: "192.168.0.17", Name: "ControlPanel-Q17", DeviceType: "Panel", State: "active"},
		{IP: "192.168.0.18", Name: "Breaker-R18", DeviceType: "Switch", State: "active"},
		{IP: "192.168.0.19", Name: "LaserSensor-S19", DeviceType: "Sensor", State: "active"},
		{IP: "192.168.0.20", Name: "InfraredSensor-T20", DeviceType: "Sensor", State: "active"},
	}

	for _, device := range devices {
		key, _ := s.createDeviceKey(ctx, device.State, device.IP)

		deviceJSON, err := json.Marshal(device)
		if err != nil {
			return fmt.Errorf("Failed to marshal device %s JSON: %v", device.IP, err)
		}

		err = ctx.GetStub().PutState(key, deviceJSON)
		if err != nil {
			return fmt.Errorf("Failed to put device %s to world state: %v", device.IP, err)
		}

		fmt.Printf("Device with IP %s and Name %s initialized with state %s.\n", device.IP, device.Name, device.State)
	}

	return nil
}

// =======================
// RegisterDevice
// =======================

func (s *SmartContract) RegisterDevice(ctx contractapi.TransactionContextInterface, ip, name, deviceType string) error {
	if ip == "" || name == "" || deviceType == "" {
		return fmt.Errorf("All fields must be provided")
	}

	for _, state := range []string{"blocked", "deleted", "active"} {
		exists, _ := s.DeviceExists(ctx, state, ip)
		if exists {
			return fmt.Errorf("Device with IP %s and Name %s already exists in %s state", ip, state)
		}
	}

	device := IIOTDevice{IP: ip, Name: name, DeviceType: deviceType, State: "active"}

	key, _ := s.createDeviceKey(ctx, "active", ip)

	deviceJSON, err := json.Marshal(device)
	if err != nil {
		return fmt.Errorf("Failed to marshal device JSON: %v", err)
	}

	if err := ctx.GetStub().PutState(key, deviceJSON); err != nil {
		return fmt.Errorf("Failed to save device: %v", err)
	}

	fmt.Printf("Device with IP %s and Name %s registered successfully.\n", ip, name)

	return nil
}

// =======================
// QueryDevicesByState
// =======================

func (s *SmartContract) QueryDevicesByState(ctx contractapi.TransactionContextInterface, state string) ([]*IIOTDevice, error) {
	iterator, err := ctx.GetStub().GetStateByPartialCompositeKey("Device", []string{state})
	if err != nil {
		return nil, fmt.Errorf("Failed to query devices: %v", err)
	}
	defer iterator.Close()

	var devices []*IIOTDevice
	for iterator.HasNext() {
		queryResponse, err := iterator.Next()
		if err != nil {
			return nil, fmt.Errorf("Failed to iterate over devices: %v", err)
		}

		var device IIOTDevice
		err = json.Unmarshal(queryResponse.Value, &device)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal device JSON: %v", err)
		}

		devices = append(devices, &device)
	}

	return devices, nil
}

// =======================
// BlockDevice
// =======================

func (s *SmartContract) BlockDevice(ctx contractapi.TransactionContextInterface, ip string) error {
	activeKey, _ := s.createDeviceKey(ctx, "active", ip)

	data, err := ctx.GetStub().GetState(activeKey)
	if err != nil || data == nil {
		return fmt.Errorf("Device with IP %s not found in active state", ip)
	}

	var device IIOTDevice
	_ = json.Unmarshal(data, &device)
	device.State = "blocked"

	// Update state
	blockedKey, _ := s.createDeviceKey(ctx, "blocked", ip)

	updated, _ := json.Marshal(device)

	ctx.GetStub().PutState(blockedKey, updated)
	ctx.GetStub().DelState(activeKey)

	fmt.Printf("Device with IP %s and Name %s has been blocked.\n", device.IP, device.Name)

	return nil
}

// =================================
// DeleteDevice (only if blocked)
// =================================

func (s *SmartContract) DeleteDevice(ctx contractapi.TransactionContextInterface, ip string) error {
	blockedKey, _ := s.createDeviceKey(ctx, "blocked", ip)

	data, err := ctx.GetStub().GetState(blockedKey)
	if err != nil || data == nil {
		return fmt.Errorf("Device with IP %s not found in blocked state", ip)
	}

	var device IIOTDevice
	_ = json.Unmarshal(data, &device)
	device.State = "deleted"

	deletedKey, _ := s.createDeviceKey(ctx, "deleted", ip)
	updated, _ := json.Marshal(device)
	ctx.GetStub().PutState(deletedKey, updated)
	ctx.GetStub().DelState(blockedKey)

	fmt.Printf("Device with IP %s and Name %s has been deleted.\n", device.IP, device.Name)
	return nil
}

// ============
// Main
// ============

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating chaincode: %v\n", err)
		return
	}
	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting chaincode: %v\n", err)
	}
}
