package converter

import (
	"encoding/xml"
	"fmt"
	"runtime"
	"testing"

	v1 "kubevirt.io/api/core/v1"
	"kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/api"
	archconverter "kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/converter/arch"
	"kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/device/hostdevice/generic"
	"sigs.k8s.io/yaml"
)

const (
	failingVM = `apiVersion: kubevirt.io/v1
kind: VirtualMachineInstance
metadata:
  annotations:
    harvesterhci.io/sshNames: '[]'
    kubevirt.io/latest-observed-api-version: v1
    kubevirt.io/storage-observed-api-version: v1
    kubevirt.io/vm-generation: "3"
  creationTimestamp: "2026-02-05T00:00:49Z"
  finalizers:
  - kubevirt.io/virtualMachineControllerFinalize
  - foregroundDeleteVirtualMachine
  - wrangler.cattle.io/VMController.BackfillObservedNetworkMacAddress
  - wrangler.cattle.io/harvester-lb-vmi-controller
  - wrangler.cattle.io/virtual-machine-deletion
  generation: 1236
  labels:
    harvesterhci.io/vmName: pci-test
    kubevirt.io/nodeName: dell-140-tink-system
  name: pci-test
  namespace: default
  ownerReferences:
  - apiVersion: kubevirt.io/v1
    blockOwnerDeletion: true
    controller: true
    kind: VirtualMachine
    name: pci-test
    uid: 4c5dc9ff-110c-4617-92e1-bfe115f2900f
  resourceVersion: "23471404"
  uid: f1e44f1b-5e75-4e46-96f6-e667cb344aa1
spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: network.harvesterhci.io/mgmt
            operator: In
            values:
            - "true"
  architecture: amd64
  domain:
    cpu:
      cores: 4
      maxSockets: 1
      model: host-model
      sockets: 1
      threads: 1
    devices:
      disks:
      - bootOrder: 1
        disk:
          bus: virtio
        name: disk-0
      - bootOrder: 2
        disk:
          bus: scsi
        name: hot-plug-test
      - disk:
          bus: virtio
        name: cloudinitdisk
      hostDevices:
      - deviceName: intel.com/82599_ETHERNET_CONTROLLER_VIRTUAL_FUNCTION
        name: dell-140-tink-system-000003101
      inputs:
      - bus: usb
        name: tablet
        type: tablet
      interfaces:
      - bridge: {}
        macAddress: 4a:ac:50:63:ea:b2
        model: virtio
        name: default
    features:
      acpi:
        enabled: true
    firmware:
      serial: d1ef78da-ac1b-4b32-a4f9-4b1c9c3a388f
      uuid: e231fb9a-3b31-4186-a037-e76ab700f9db
    machine:
      type: q35
    memory:
      guest: 8Gi
      maxGuest: 32Gi
    resources:
      limits:
        cpu: "4"
        memory: 8Gi
      requests:
        cpu: 250m
        memory: 8Gi
  evictionStrategy: LiveMigrateIfPossible
  hostname: pci-test
  networks:
  - multus:
      networkName: default/workload
    name: default
  terminationGracePeriodSeconds: 120
  volumes:
  - name: disk-0
    persistentVolumeClaim:
      claimName: pci-test-disk-0-jzf2i
  - name: hot-plug-test
    persistentVolumeClaim:
      claimName: test-volume
      hotpluggable: true
  - cloudInitNoCloud:
      networkDataSecretRef:
        name: pci-test-prach
      secretRef:
        name: pci-test-prach
    name: cloudinitdisk`

	workingVMI = `apiVersion: kubevirt.io/v1
kind: VirtualMachineInstance
metadata:
  annotations:
    harvesterhci.io/sshNames: '[]'
    kubevirt.io/latest-observed-api-version: v1
    kubevirt.io/storage-observed-api-version: v1
    kubevirt.io/vm-generation: "1"
  creationTimestamp: "2026-02-04T22:37:33Z"
  finalizers:
  - kubevirt.io/virtualMachineControllerFinalize
  - foregroundDeleteVirtualMachine
  - wrangler.cattle.io/VMController.BackfillObservedNetworkMacAddress
  - wrangler.cattle.io/harvester-lb-vmi-controller
  - wrangler.cattle.io/virtual-machine-deletion
  generation: 41
  labels:
    harvesterhci.io/vmName: pci-test-2
    kubevirt.io/nodeName: dell-140-tink-system
  name: pci-test-2
  namespace: default
  ownerReferences:
  - apiVersion: kubevirt.io/v1
    blockOwnerDeletion: true
    controller: true
    kind: VirtualMachine
    name: pci-test-2
    uid: cb6c43e7-32e0-49a4-b5cf-0f69b1c56ce3
  resourceVersion: "23602234"
  uid: cacd47a6-6a3a-43be-a219-da741f5c007d
spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: network.harvesterhci.io/mgmt
            operator: In
            values:
            - "true"
  architecture: amd64
  domain:
    cpu:
      cores: 4
      maxSockets: 1
      model: host-model
      sockets: 1
      threads: 1
    devices:
      disks:
      - bootOrder: 1
        disk:
          bus: virtio
        name: disk-0
      - disk:
          bus: virtio
        name: cloudinitdisk
      hostDevices:
      - deviceName: intel.com/82599_ETHERNET_CONTROLLER_VIRTUAL_FUNCTION
        name: dell-140-tink-system-000003103
      inputs:
      - bus: usb
        name: tablet
        type: tablet
      interfaces:
      - bridge: {}
        macAddress: fa:08:fc:b6:42:3a
        model: virtio
        name: default
    features:
      acpi:
        enabled: true
    firmware:
      serial: 1af50482-9547-49c9-80ef-f784ee595e5f
      uuid: c74cee74-ff21-4607-8681-d72906eb9e25
    machine:
      type: q35
    memory:
      guest: 8Gi
      maxGuest: 32Gi
    resources:
      limits:
        cpu: "4"
        memory: 8Gi
      requests:
        cpu: 250m
        memory: 8Gi
  evictionStrategy: LiveMigrateIfPossible
  hostname: pci-test-2
  networks:
  - multus:
      networkName: default/workload
    name: default
  terminationGracePeriodSeconds: 120
  volumes:
  - name: disk-0
    persistentVolumeClaim:
      claimName: pci-test-2-disk-0-ribys
  - cloudInitNoCloud:
      networkDataSecretRef:
        name: pci-test-2-uf233
      secretRef:
        name: pci-test-2-uf233
    name: cloudinitdisk`
)

func Test_FailedVMIConversion(t *testing.T) {
	domain := &api.Domain{}
	vmi, err := generateVMI([]byte(failingVM))
	if err != nil {
		t.Fatalf("error generating vmi: %v", err)
	}
	c := &ConverterContext{
		Architecture:   archconverter.NewConverter(runtime.GOARCH),
		VirtualMachine: vmi,
		AllowEmulation: true,
	}

	pciPool := newAddressPoolStub()
	pciPool.AddResource("intel.com/82599_ETHERNET_CONTROLLER_VIRTUAL_FUNCTION", "0000:03:00.0")
	mdevPool := newAddressPoolStub()
	usbPool := newAddressPoolStub()
	genericHostDevices, err := generic.CreateHostDevicesFromPools(vmi.Spec.Domain.Devices.HostDevices, pciPool, mdevPool, usbPool)
	if err != nil {
		t.Fatalf("error creating devices :%v", err)
	}
	c.GenericHostDevices = genericHostDevices
	err = Convert_v1_VirtualMachineInstance_To_api_Domain(vmi, domain, c)
	if err != nil {
		t.Fatalf("error generating domain api spec: %v", err)
	}
	xmlBytes, err := xml.Marshal(domain.Spec)
	if err != nil {
		t.Fatalf("error generating domain xml %v", err)
	}
	t.Log(string(xmlBytes))
}

func Test_VMIConversion(t *testing.T) {
	domain := &api.Domain{}
	vmi, err := generateVMI([]byte(workingVMI))
	if err != nil {
		t.Fatalf("error generating vmi: %v", err)
	}
	c := &ConverterContext{
		Architecture:   archconverter.NewConverter(runtime.GOARCH),
		VirtualMachine: vmi,
		AllowEmulation: true,
	}

	pciPool := newAddressPoolStub()
	pciPool.AddResource("intel.com/82599_ETHERNET_CONTROLLER_VIRTUAL_FUNCTION", "0000:03:00.0")
	mdevPool := newAddressPoolStub()
	usbPool := newAddressPoolStub()
	genericHostDevices, err := generic.CreateHostDevicesFromPools(vmi.Spec.Domain.Devices.HostDevices, pciPool, mdevPool, usbPool)
	if err != nil {
		t.Fatalf("error creating devices :%v", err)
	}
	c.GenericHostDevices = genericHostDevices
	err = Convert_v1_VirtualMachineInstance_To_api_Domain(vmi, domain, c)
	if err != nil {
		t.Fatalf("error generating domain api spec: %v", err)
	}
	xmlBytes, err := xml.Marshal(domain.Spec)
	if err != nil {
		t.Fatalf("error generating domain xml %v", err)
	}
	t.Log(string(xmlBytes))
}
func generateVMI(data []byte) (*v1.VirtualMachineInstance, error) {
	vmi := &v1.VirtualMachineInstance{}
	err := yaml.Unmarshal(data, vmi)
	return vmi, err
}

func Error(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("error returned: %v", err)
	}
}

type stubAddressPool struct {
	addresses map[string][]string
}

func newAddressPoolStub() *stubAddressPool {
	return &stubAddressPool{addresses: make(map[string][]string)}
}

func (p *stubAddressPool) AddResource(resource string, addresses ...string) {
	p.addresses[resource] = addresses
}

func (p *stubAddressPool) Pop(resource string) (string, error) {
	addresses, exists := p.addresses[resource]
	if !exists {
		return "", fmt.Errorf("no resource: %s", resource)
	}
	if len(addresses) == 0 {
		return "", fmt.Errorf("pool is empty")
	}

	address := addresses[0]
	p.addresses[resource] = addresses[1:]

	return address, nil
}
