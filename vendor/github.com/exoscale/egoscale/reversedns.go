package egoscale

import (
	"net"
)

// ReverseDNS represents the PTR record linked with an IPAddress or IP6Address belonging to a Virtual Machine or a Public IP Address (Elastic IP) instance
type ReverseDNS struct {
	DomainName       string `json:"domainname,omitempty" doc:"the domain name of the PTR record"`
	IP6Address       net.IP `json:"ip6address,omitempty" doc:"the IPv6 address linked with the PTR record (mutually exclusive with ipaddress)"`
	IPAddress        net.IP `json:"ipaddress,omitempty" doc:"the IPv4 address linked with the PTR record (mutually exclusive with ip6address)"`
	NicID            string `json:"nicid,omitempty" doc:"the virtual machine default NIC ID"`
	PublicIPID       string `json:"publicipid,omitempty" doc:"the public IP address ID"`
	VirtualMachineID string `json:"virtualmachineid,omitempty" doc:"the virtual machine ID"`
}

// DeleteReverseDNSFromPublicIPAddress is a command to create/delete the PTR record of a public IP address
type DeleteReverseDNSFromPublicIPAddress struct {
	ID string `json:"id,omitempty" doc:"the ID of the public IP address"`
	_  bool   `name:"deleteReverseDnsFromPublicIpAddress" description:"delete the PTR DNS record from the public IP address"`
}

func (*DeleteReverseDNSFromPublicIPAddress) response() interface{} {
	return new(booleanResponse)
}

// DeleteReverseDNSFromVirtualMachine is a command to create/delete the PTR record(s) of a virtual machine
type DeleteReverseDNSFromVirtualMachine struct {
	ID string `json:"id,omitempty" doc:"the ID of the virtual machine"`
	_  bool   `name:"deleteReverseDnsFromVirtualMachine" description:"Delete the PTR DNS record(s) from the virtual machine"`
}

func (*DeleteReverseDNSFromVirtualMachine) response() interface{} {
	return new(booleanResponse)
}

// QueryReverseDNSForPublicIPAddress is a command to create/query the PTR record of a public IP address
type QueryReverseDNSForPublicIPAddress struct {
	ID string `json:"id,omitempty" doc:"the ID of the public IP address"`
	_  bool   `name:"queryReverseDnsForPublicIpAddress" description:"Query the PTR DNS record for the public IP address"`
}

func (*QueryReverseDNSForPublicIPAddress) response() interface{} {
	return new(IPAddress)
}

// QueryReverseDNSForVirtualMachine is a command to create/query the PTR record(s) of a virtual machine
type QueryReverseDNSForVirtualMachine struct {
	ID string `json:"id,omitempty" doc:"the ID of the virtual machine"`
	_  bool   `name:"queryReverseDnsForVirtualMachine" description:"Query the PTR DNS record(s) for the virtual machine"`
}

func (*QueryReverseDNSForVirtualMachine) response() interface{} {
	return new(VirtualMachine)
}

// UpdateReverseDNSForPublicIPAddress is a command to create/update the PTR record of a public IP address
type UpdateReverseDNSForPublicIPAddress struct {
	DomainName string `json:"domainname,omitempty" doc:"the domain name for the PTR record. It must have a valid TLD"`
	ID         string `json:"id,omitempty" doc:"the ID of the public IP address"`
	_          bool   `name:"updateReverseDnsForPublicIpAddress" description:"Update/create the PTR DNS record for the public IP address"`
}

func (*UpdateReverseDNSForPublicIPAddress) response() interface{} {
	return new(IPAddress)
}

// UpdateReverseDNSForVirtualMachine is a command to create/update the PTR record(s) of a virtual machine
type UpdateReverseDNSForVirtualMachine struct {
	DomainName string `json:"domainname,omitempty" doc:"the domain name for the PTR record(s). It must have a valid TLD"`
	ID         string `json:"id,omitempty" doc:"the ID of the virtual machine"`
	_          bool   `name:"updateReverseDnsForVirtualMachine" description:"Update/create the PTR DNS record(s) for the virtual machine"`
}

func (*UpdateReverseDNSForVirtualMachine) response() interface{} {
	return new(VirtualMachine)
}
