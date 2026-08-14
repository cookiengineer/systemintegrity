package structs

import "strings"

type Boot struct {
	Bootloader         string `json:"bootloader"`
	Mode               string `json:"mode"`
	Kernel             string `json:"kernel"`
	KernelArchitecture string `json:"kernel_architecture"`
	KernelVersion      string `json:"kernel_version"`
	Initramfs          string `json:"initramfs"`
	SecureBoot         string `json:"secure_boot"`
	ESP                string `json:"esp"`
}

func NewBoot() Boot {

	var boot Boot

	boot.Bootloader = ""
	boot.Mode = ""
	boot.Kernel = ""
	boot.KernelArchitecture = ""
	boot.KernelVersion = ""
	boot.Initramfs = ""
	boot.SecureBoot = ""
	boot.ESP = ""

	return boot

}

func (boot *Boot) IsValid() bool {

	var result bool

	if boot.Kernel != "" || boot.KernelVersion != "" || boot.Bootloader != "" || boot.Mode != "" {
		result = true
	}

	return result

}

func (boot *Boot) SetBootloader(value string) {

	value = strings.TrimSpace(value)

	if value == "grub" || value == "Grub" || value == "GRUB" {
		boot.Bootloader = "grub"
	} else if value == "systemd-boot" || value == "systemd" || value == "Systemd" || value == "SYSTEMD" {
		boot.Bootloader = "systemd-boot"
	} else if value == "uefi" || value == "UEFI" {
		boot.Bootloader = "uefi"
	} else if value != "" {
		boot.Bootloader = strings.ToLower(value)
	}

}

func (boot *Boot) SetMode(value string) {

	value = strings.ToLower(strings.TrimSpace(value))

	if value == "uefi" || value == "efi" {
		boot.Mode = "uefi"
	} else if value == "bios" || value == "legacy" {
		boot.Mode = "bios"
	}

}

func (boot *Boot) SetKernel(value string) {
	boot.Kernel = strings.TrimSpace(value)
}

func (boot *Boot) SetKernelArchitecture(value string) {
	boot.KernelArchitecture = strings.TrimSpace(value)
}

func (boot *Boot) SetKernelVersion(value string) {
	boot.KernelVersion = strings.TrimSpace(value)
}

func (boot *Boot) SetInitramfs(value string) {
	boot.Initramfs = strings.TrimSpace(value)
}

func (boot *Boot) SetSecureBoot(value string) {

	value = strings.ToLower(strings.TrimSpace(value))

	if value == "enabled" || value == "yes" || value == "true" || value == "1" {
		boot.SecureBoot = "enabled"
	} else if value == "disabled" || value == "no" || value == "false" || value == "0" {
		boot.SecureBoot = "disabled"
	}

}

func (boot *Boot) SetESP(value string) {

	value = strings.TrimSpace(value)

	if strings.HasPrefix(value, "/") {
		boot.ESP = value
	}

}
