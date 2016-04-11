package client

import (
	"time"

	dm "github.com/NEU-SNS/ReverseTraceroute/datamodel"
	"github.com/NEU-SNS/ReverseTraceroute/log"
	"github.com/NEU-SNS/ReverseTraceroute/vpservice/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type client struct {
	context.Context
	pb.VPServiceClient
}

// VPSource is the inteface to something that gives vps
type VPSource interface {
	GetVPs() (*dm.VPReturn, error)
	GetOneVPPerSite() (*dm.VPReturn, error)
	GetSpoofers(addr, max uint32) ([]*dm.VantagePoint, error)
}

// New returns a VPSource
func New(ctx context.Context, cc *grpc.ClientConn) VPSource {
	return client{Context: ctx, VPServiceClient: pb.NewVPServiceClient(cc)}
}

func (c client) GetVPs() (*dm.VPReturn, error) {
	ctx, cancel := context.WithTimeout(c.Context, time.Second*30)
	defer cancel()
	vpr, err := c.VPServiceClient.GetVPs(ctx, &dm.VPRequest{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return vpr, nil
}

func (c client) GetOneVPPerSite() (*dm.VPReturn, error) {
	ctx, cancel := context.WithTimeout(c.Context, time.Second*30)
	defer cancel()
	vpr, err := c.VPServiceClient.GetVPs(ctx, &dm.VPRequest{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	set := make(map[string]*dm.VantagePoint)
	vps := vpr.GetVps()
	for _, vp := range vps {
		set[vp.Site] = vp
	}
	var ret []*dm.VantagePoint
	for _, val := range set {
		ret = append(ret, val)
	}
	vpr.Vps = ret
	return vpr, nil
}

func (c client) GetSpoofers(addr, max uint32) ([]*dm.VantagePoint, error) {
	ctx, cancel := context.WithTimeout(c.Context, time.Second*30)
	defer cancel()
	arg := &dm.SpooferRequest{
		Addr: addr,
		Max:  max,
	}
	sr, err := c.VPServiceClient.GetSpoofers(ctx, arg)
	if err != nil {
		return nil, err
	}
	return sr.Spoofers, nil
}
