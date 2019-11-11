// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"osc-api/client/api_log"
	"osc-api/client/client_gateway"
	"osc-api/client/dhcp_option"
	"osc-api/client/direct_link"
	"osc-api/client/direct_link_interface"
	"osc-api/client/flexible_gpu"
	"osc-api/client/health"
	"osc-api/client/image"
	"osc-api/client/internet_service"
	"osc-api/client/keypair"
	"osc-api/client/listener"
	"osc-api/client/load_balancer"
	"osc-api/client/load_balancer_policy"
	"osc-api/client/location"
	"osc-api/client/nat_service"
	"osc-api/client/net"
	"osc-api/client/net_access_point"
	"osc-api/client/net_peering"
	"osc-api/client/nic"
	"osc-api/client/product_type"
	"osc-api/client/public_ip"
	"osc-api/client/public_ip_range"
	"osc-api/client/quota"
	"osc-api/client/route"
	"osc-api/client/route_table"
	"osc-api/client/security_group"
	"osc-api/client/security_group_rule"
	"osc-api/client/snapshot"
	"osc-api/client/subnet"
	"osc-api/client/subregion"
	"osc-api/client/tag"
	"osc-api/client/task"
	"osc-api/client/virtual_gateway"
	"osc-api/client/vm"
	"osc-api/client/volume"
	"osc-api/client/vpn_connection"
)

// Default nr3 d s o u t s c a l e HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new nr3 d s o u t s c a l e HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Nr3DSOUTSCALE {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new nr3 d s o u t s c a l e HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Nr3DSOUTSCALE {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new nr3 d s o u t s c a l e client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Nr3DSOUTSCALE {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Nr3DSOUTSCALE)
	cli.Transport = transport

	cli.APILog = api_log.New(transport, formats)

	cli.ClientGateway = client_gateway.New(transport, formats)

	cli.DhcpOption = dhcp_option.New(transport, formats)

	cli.DirectLink = direct_link.New(transport, formats)

	cli.DirectLinkInterface = direct_link_interface.New(transport, formats)

	cli.FlexibleGpu = flexible_gpu.New(transport, formats)

	cli.Health = health.New(transport, formats)

	cli.Image = image.New(transport, formats)

	cli.InternetService = internet_service.New(transport, formats)

	cli.Keypair = keypair.New(transport, formats)

	cli.Listener = listener.New(transport, formats)

	cli.LoadBalancer = load_balancer.New(transport, formats)

	cli.LoadBalancerPolicy = load_balancer_policy.New(transport, formats)

	cli.Location = location.New(transport, formats)

	cli.NatService = nat_service.New(transport, formats)

	cli.Net = net.New(transport, formats)

	cli.NetAccessPoint = net_access_point.New(transport, formats)

	cli.NetPeering = net_peering.New(transport, formats)

	cli.Nic = nic.New(transport, formats)

	cli.ProductType = product_type.New(transport, formats)

	cli.PublicIP = public_ip.New(transport, formats)

	cli.PublicIPRange = public_ip_range.New(transport, formats)

	cli.Quota = quota.New(transport, formats)

	cli.Route = route.New(transport, formats)

	cli.RouteTable = route_table.New(transport, formats)

	cli.SecurityGroup = security_group.New(transport, formats)

	cli.SecurityGroupRule = security_group_rule.New(transport, formats)

	cli.Snapshot = snapshot.New(transport, formats)

	cli.Subnet = subnet.New(transport, formats)

	cli.Subregion = subregion.New(transport, formats)

	cli.Tag = tag.New(transport, formats)

	cli.Task = task.New(transport, formats)

	cli.VirtualGateway = virtual_gateway.New(transport, formats)

	cli.VM = vm.New(transport, formats)

	cli.Volume = volume.New(transport, formats)

	cli.VpnConnection = vpn_connection.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Nr3DSOUTSCALE is a client for nr3 d s o u t s c a l e
type Nr3DSOUTSCALE struct {
	APILog *api_log.Client

	ClientGateway *client_gateway.Client

	DhcpOption *dhcp_option.Client

	DirectLink *direct_link.Client

	DirectLinkInterface *direct_link_interface.Client

	FlexibleGpu *flexible_gpu.Client

	Health *health.Client

	Image *image.Client

	InternetService *internet_service.Client

	Keypair *keypair.Client

	Listener *listener.Client

	LoadBalancer *load_balancer.Client

	LoadBalancerPolicy *load_balancer_policy.Client

	Location *location.Client

	NatService *nat_service.Client

	Net *net.Client

	NetAccessPoint *net_access_point.Client

	NetPeering *net_peering.Client

	Nic *nic.Client

	ProductType *product_type.Client

	PublicIP *public_ip.Client

	PublicIPRange *public_ip_range.Client

	Quota *quota.Client

	Route *route.Client

	RouteTable *route_table.Client

	SecurityGroup *security_group.Client

	SecurityGroupRule *security_group_rule.Client

	Snapshot *snapshot.Client

	Subnet *subnet.Client

	Subregion *subregion.Client

	Tag *tag.Client

	Task *task.Client

	VirtualGateway *virtual_gateway.Client

	VM *vm.Client

	Volume *volume.Client

	VpnConnection *vpn_connection.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Nr3DSOUTSCALE) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.APILog.SetTransport(transport)

	c.ClientGateway.SetTransport(transport)

	c.DhcpOption.SetTransport(transport)

	c.DirectLink.SetTransport(transport)

	c.DirectLinkInterface.SetTransport(transport)

	c.FlexibleGpu.SetTransport(transport)

	c.Health.SetTransport(transport)

	c.Image.SetTransport(transport)

	c.InternetService.SetTransport(transport)

	c.Keypair.SetTransport(transport)

	c.Listener.SetTransport(transport)

	c.LoadBalancer.SetTransport(transport)

	c.LoadBalancerPolicy.SetTransport(transport)

	c.Location.SetTransport(transport)

	c.NatService.SetTransport(transport)

	c.Net.SetTransport(transport)

	c.NetAccessPoint.SetTransport(transport)

	c.NetPeering.SetTransport(transport)

	c.Nic.SetTransport(transport)

	c.ProductType.SetTransport(transport)

	c.PublicIP.SetTransport(transport)

	c.PublicIPRange.SetTransport(transport)

	c.Quota.SetTransport(transport)

	c.Route.SetTransport(transport)

	c.RouteTable.SetTransport(transport)

	c.SecurityGroup.SetTransport(transport)

	c.SecurityGroupRule.SetTransport(transport)

	c.Snapshot.SetTransport(transport)

	c.Subnet.SetTransport(transport)

	c.Subregion.SetTransport(transport)

	c.Tag.SetTransport(transport)

	c.Task.SetTransport(transport)

	c.VirtualGateway.SetTransport(transport)

	c.VM.SetTransport(transport)

	c.Volume.SetTransport(transport)

	c.VpnConnection.SetTransport(transport)

}
