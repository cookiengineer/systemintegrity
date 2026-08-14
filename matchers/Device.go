package matchers

import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strings"

type Device struct {
	Name   string `json:"name"`
	Vendor string `json:"vendor"`
	Device string `json:"device"`
	Bus    string `json:"bus"`
}

func NewDevice() Device {

	var device Device

	device.Name = "any"
	device.Vendor = "any"
	device.Device = "any"
	device.Bus = "any"

	return device

}

func ToDevice(value string) Device {

	var device Device

	device.Name = "any"
	device.Vendor = "any"
	device.Device = "any"
	device.Bus = "any"

	device.SetName(value)

	return device

}

func (device *Device) IsIdentical(value Device) bool {

	var result bool

	if device.Name == value.Name &&
		device.Vendor == value.Vendor &&
		device.Device == value.Device &&
		device.Bus == value.Bus {
		result = true
	}

	return result

}

func (device *Device) IsValid() bool {

	var result bool

	if device.Name != "any" || device.Vendor != "any" || device.Device != "any" || device.Bus != "any" {
		result = true
	}

	return result

}

func (device *Device) Matches(name string, vendor string, device_ string, bus string) bool {
	return device.MatchesName(name) && device.MatchesVendor(vendor) && device.MatchesDevice(device_) && device.MatchesBus(bus)
}

func (device *Device) MatchesBus(value string) bool {

	var result bool

	if device.Bus == value {
		result = true
	} else if device.Bus == "any" {
		result = true
	}

	return result

}

func (device *Device) MatchesDevice(value string) bool {

	var result bool

	if device.Device == value {
		result = true
	} else if device.Device == "any" {
		result = true
	}

	return result

}

func (device *Device) MatchesName(value string) bool {

	var result bool

	if device.Name == value {
		result = true
	} else if device.Name == "any" {
		result = true
	}

	return result

}

func (device *Device) MatchesVendor(value string) bool {

	var result bool

	if device.Vendor == value {
		result = true
	} else if device.Vendor == "any" {
		result = true
	}

	return result

}

func (device *Device) SetBus(value string) {

	if value == "any" {
		device.Bus = "any"
	} else if value == "hid" {
		device.Bus = "hid"
	} else if value == "i2c" {
		device.Bus = "i2c"
	} else if value == "pci" {
		device.Bus = "pci"
	} else if value == "scsi" {
		device.Bus = "scsi"
	} else if value == "usb" {
		device.Bus = "usb"
	} else if value == "other" {
		device.Bus = "other"
	}

}

func (device *Device) SetDevice(value string) {
	device.Device = strings.TrimSpace(value)
}

func (device *Device) SetName(value string) {
	device.Name = strings.TrimSpace(value)
}

func (device *Device) SetVendor(value string) {
	device.Vendor = strings.TrimSpace(value)
}

func (device *Device) Hash() string {

	var hash string

	if device.Name != "" {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			device.Name,
			device.Vendor,
			device.Device,
			device.Bus,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
