// Code generated by swagger-doc. DO NOT EDIT.

package v1alpha1

func (VirtualMachineFlavor) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "VirtualMachineFlavor resource contains quantitative and resource related VirtualMachine configuration\nthat can be used by multiple VirtualMachine resources.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+genclient",
		"spec": "Required spec describing the flavor",
	}
}

func (VirtualMachineFlavorList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineFlavorList is a list of VirtualMachineFlavor resources.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
	}
}

func (VirtualMachineClusterFlavor) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "VirtualMachineClusterFlavor is a cluster scoped version of VirtualMachineFlavor resource.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+genclient\n+genclient:nonNamespaced",
		"spec": "Required spec describing the flavor",
	}
}

func (VirtualMachineClusterFlavorList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineClusterFlavorList is a list of VirtualMachineClusterFlavor resources.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
	}
}

func (VirtualMachineFlavorSpecRevision) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "VirtualMachineFlavorSpecRevision is used to capture a versioned copy of VirtualMachineFlavorSpec within a ControllerRevision\n\n**WARNING** This struct should never change ensuring it can always be unmarshaled from a ControllerRevision",
		"apiVersion": "The specific version of VirtualMachineFlavorSpec that is contained below",
		"spec":       "This is a marshaled version of the VirtualMachineFlavorSpec\n+listType=atomic",
	}
}

func (VirtualMachineFlavorSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                "VirtualMachineFlavorSpec is a description of the VirtualMachineFlavor or VirtualMachineClusterFlavor.\n\nCPU and Memory are required attributes with both requiring that their Guest attribute is defined, ensuring a number of vCPUs and amount of RAM is always provided by each flavor.",
		"cpu":             "Required CPU related attributes of the flavor.",
		"memory":          "Required Memory related attributes of the flavor.",
		"gpus":            "Optionally defines any GPU devices associated with the flavor.\n\n+optional\n+listType=atomic",
		"hostDevices":     "Optionally defines any HostDevices associated with the flavor.\n\n+optional\n+listType=atomic",
		"ioThreadsPolicy": "Optionally defines the IOThreadsPolicy to be used by the flavor.\n\n+optional",
		"launchSecurity":  "Optionally defines the LaunchSecurity to be used by the flavor.\n\n+optional",
	}
}

func (CPUFlavor) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                      "CPUFlavor contains the CPU related configuration of a given VirtualMachineFlavorSpec.\n\nGuest is a required attribute and defines the number of vCPUs to be exposed to the guest by the flavor.",
		"guest":                 "Required number of vCPUs to expose to the guest.\n\nThe resulting CPU topology being derived from the optional PreferredCPUTopology attribute of CPUPreferences that itself defaults to PreferCores.",
		"model":                 "Model specifies the CPU model inside the VMI.\nList of available models https://github.com/libvirt/libvirt/tree/master/src/cpu_map.\nIt is possible to specify special cases like \"host-passthrough\" to get the same CPU as the node\nand \"host-model\" to get CPU closest to the node one.\nDefaults to host-model.\n+optional",
		"dedicatedCPUPlacement": "DedicatedCPUPlacement requests the scheduler to place the VirtualMachineInstance on a node\nwith enough dedicated pCPUs and pin the vCPUs to it.\n+optional",
		"numa":                  "NUMA allows specifying settings for the guest NUMA topology\n+optional",
		"isolateEmulatorThread": "IsolateEmulatorThread requests one more dedicated pCPU to be allocated for the VMI to place\nthe emulator thread on it.\n+optional",
		"realtime":              "Realtime instructs the virt-launcher to tune the VMI for lower latency, optional for real time workloads\n+optional",
	}
}

func (MemoryFlavor) SwaggerDoc() map[string]string {
	return map[string]string{
		"":          "MemoryFlavor contains the Memory related configuration of a given VirtualMachineFlavorSpec.\n\nGuest is a required attribute and defines the amount of RAM to be exposed to the guest by the flavor.",
		"guest":     "Required amount of memory which is visible inside the guest OS.",
		"hugepages": "Optionally enables the use of hugepages for the VirtualMachineInstance instead of regular memory.\n+optional",
	}
}

func (VirtualMachinePreference) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "VirtualMachinePreference resource contains optional preferences related to the VirtualMachine.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+genclient",
		"spec": "Required spec describing the preferences",
	}
}

func (VirtualMachinePreferenceList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "VirtualMachinePreferenceList is a list of VirtualMachinePreference resources.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "+listType=set",
	}
}

func (VirtualMachineClusterPreference) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "VirtualMachineClusterPreference is a cluster scoped version of the VirtualMachinePreference resource.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+genclient\n+genclient:nonNamespaced",
		"spec": "Required spec describing the preferences",
	}
}

func (VirtualMachineClusterPreferenceList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "VirtualMachineClusterPreferenceList is a list of VirtualMachineClusterPreference resources.\n\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "+listType=set",
	}
}

func (VirtualMachinePreferenceSpecRevision) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "VirtualMachinePreferenceSpecRevision is used to capture a versioned copy of VirtualMachinePreferenceSpec within a ControllerRevision\n\n**WARNING** This struct should never change ensuring it can always be unmarshaled from a ControllerRevision",
		"apiVersion": "The specific version of VirtualMachinePreferenceSpec that is contained below",
		"spec":       "This is a marshaled version of the VirtualMachinePreferenceSpec\n+listType=atomic",
	}
}

func (VirtualMachinePreferenceSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "VirtualMachinePreferenceSpec is a description of the VirtualMachinePreference or VirtualMachineClusterPreference.",
		"clock":    "Clock optionally defines preferences associated with the Clock attribute of a VirtualMachineInstance DomainSpec\n\n+optional",
		"cpu":      "CPU optionally defines preferences associated with the CPU attribute of a VirtualMachineInstance DomainSpec\n\n+optional",
		"devices":  "Devices optionally defines preferences associated with the Devices attribute of a VirtualMachineInstance DomainSpec\n\n+optional",
		"features": "Features optionally defines preferences associated with the Features attribute of a VirtualMachineInstance DomainSpec\n\n+optional",
		"firmware": "Firmware optionally defines preferences associated with the Firmware attribute of a VirtualMachineInstance DomainSpec\n\n+optional",
		"machine":  "Machine optionally defines preferences associated with the Machine attribute of a VirtualMachineInstance DomainSpec\n\n+optional",
	}
}

func (CPUPreferences) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                     "CPUPreferences contains various optional CPU preferences.",
		"preferredCPUTopology": "PreferredCPUTopology optionally defines the preferred guest visible CPU topology, defaults to PreferSockets.\n\n+optional",
	}
}

func (DevicePreferences) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                                    "DevicePreferences contains various optional Device preferences.",
		"preferredAutoattachGraphicsDevice":   "PreferredAutoattachGraphicsDevice optionally defines the preferred value of AutoattachGraphicsDevice\n\n+optional",
		"preferredAutoattachMemBalloon":       "PreferredAutoattachMemBalloon optionally defines the preferred value of AutoattachMemBalloon\n\n+optional",
		"preferredAutoattachPodInterface":     "PreferredAutoattachPodInterface optionally defines the preferred value of AutoattachPodInterface\n\n+optional",
		"preferredAutoattachSerialConsole":    "PreferredAutoattachSerialConsole optionally defines the preferred value of AutoattachSerialConsole\n\n+optional",
		"preferredDisableHotplug":             "PreferredDisableHotplug optionally defines the preferred value of DisableHotplug\n\n+optional",
		"preferredVirtualGPUOptions":          "PreferredVirtualGPUOptions optionally defines the preferred value of VirtualGPUOptions\n\n+optional",
		"preferredSoundModel":                 "PreferredSoundModel optionally defines the preferred model for Sound devices.\n\n+optional",
		"preferredUseVirtioTransitional":      "PreferredUseVirtioTransitional optionally defines the preferred value of UseVirtioTransitional\n\n+optional",
		"preferredInputBus":                   "PreferredInputBus optionally defines the preferred bus for Input devices.\n\n+optional",
		"preferredInputType":                  "PreferredInputType optionally defines the preferred type for Input devices.\n\n+optional",
		"preferredDiskBus":                    "PreferredDiskBus optionally defines the preferred bus for Disk Disk devices.\n\n+optional",
		"preferredLunBus":                     "PreferredLunBus optionally defines the preferred bus for Lun Disk devices.\n\n+optional",
		"preferredCdromBus":                   "PreferredCdromBus optionally defines the preferred bus for Cdrom Disk devices.\n\n+optional",
		"preferredDiskDedicatedIoThread":      "PreferredDedicatedIoThread optionally enables dedicated IO threads for Disk devices.\n\n+optional",
		"preferredDiskCache":                  "PreferredCache optionally defines the DriverCache to be used by Disk devices.\n\n+optional",
		"preferredDiskIO":                     "PreferredIo optionally defines the QEMU disk IO mode to be used by Disk devices.\n\n+optional",
		"preferredDiskBlockSize":              "PreferredBlockSize optionally defines the block size of Disk devices.\n\n+optional",
		"preferredInterfaceModel":             "PreferredInterfaceModel optionally defines the preferred model to be used by Interface devices.\n\n+optional",
		"preferredRng":                        "PreferredRng optionally defines the preferred rng device to be used.\n\n+optional",
		"preferredBlockMultiQueue":            "PreferredBlockMultiQueue optionally enables the vhost multiqueue feature for virtio disks.\n\n+optional",
		"preferredNetworkInterfaceMultiQueue": "PreferredNetworkInterfaceMultiQueue optionally enables the vhost multiqueue feature for virtio interfaces.\n\n+optional",
		"preferredTPM":                        "PreferredTPM optionally defines the preferred TPM device to be used.\n\n+optional",
	}
}

func (FeaturePreferences) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                    "FeaturePreferences contains various optional defaults for Features.",
		"preferredAcpi":       "PreferredAcpi optionally enables the ACPI feature\n\n+optional",
		"preferredApic":       "PreferredApic optionally enables and configures the APIC feature\n\n+optional",
		"preferredHyperv":     "PreferredHyperv optionally enables and configures HyperV features\n\n+optional",
		"preferredKvm":        "PreferredKvm optionally enables and configures KVM features\n\n+optional",
		"preferredPvspinlock": "PreferredPvspinlock optionally enables the Pvspinlock feature\n\n+optional",
		"preferredSmm":        "PreferredSmm optionally enables the SMM feature\n\n+optional",
	}
}

func (FirmwarePreferences) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                       "FirmwarePreferences contains various optional defaults for Firmware.",
		"preferredUseBios":       "PreferredUseBios optionally enables BIOS\n\n+optional",
		"preferredUseBiosSerial": "PreferredUseBiosSerial optionally transmitts BIOS output over the serial.\n\nRequires PreferredUseBios to be enabled.\n\n+optional",
		"preferredUseEfi":        "PreferredUseEfi optionally enables EFI\n\n+optional",
		"preferredUseSecureBoot": "PreferredUseSecureBoot optionally enables SecureBoot and the OVMF roms will be swapped for SecureBoot-enabled ones.\n\nRequires PreferredUseEfi and PreferredSmm to be enabled.\n\n+optional",
	}
}

func (MachinePreferences) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                     "MachinePreferences contains various optional defaults for Machine.",
		"preferredMachineType": "PreferredMachineType optionally defines the preferred machine type to use.\n\n+optional",
	}
}

func (ClockPreferences) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                     "ClockPreferences contains various optional defaults for Clock.",
		"preferredClockOffset": "ClockOffset allows specifying the UTC offset or the timezone of the guest clock.\n\n+optional",
		"preferredTimer":       "Timer specifies whih timers are attached to the vmi.\n\n+optional",
	}
}
