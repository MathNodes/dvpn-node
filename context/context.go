package context

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
	hubtypes "github.com/sentinel-official/hub/types"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/dvpn-node/lite"
	"github.com/sentinel-official/dvpn-node/types"
)

type Context struct {
	logger    tmlog.Logger
	service   types.Service
	bandwidth *hubtypes.Bandwidth
	client    *lite.Client
	config    *types.Config
	location  *types.GeoIPLocation
	router    *mux.Router
	sessions  *types.Sessions
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) WithBandwidth(v *hubtypes.Bandwidth) *Context { c.bandwidth = v; return c }
func (c *Context) WithClient(v *lite.Client) *Context           { c.client = v; return c }
func (c *Context) WithConfig(v *types.Config) *Context          { c.config = v; return c }
func (c *Context) WithLocation(v *types.GeoIPLocation) *Context { c.location = v; return c }
func (c *Context) WithLogger(v tmlog.Logger) *Context           { c.logger = v; return c }
func (c *Context) WithRouter(v *mux.Router) *Context            { c.router = v; return c }
func (c *Context) WithService(v types.Service) *Context         { c.service = v; return c }
func (c *Context) WithSessions(v *types.Sessions) *Context      { c.sessions = v; return c }

func (c *Context) Address() hubtypes.NodeAddress       { return c.Operator().Bytes() }
func (c *Context) Bandwidth() *hubtypes.Bandwidth      { return c.bandwidth }
func (c *Context) Client() *lite.Client                { return c.client }
func (c *Context) Config() *types.Config               { return c.config }
func (c *Context) IntervalSetSessions() time.Duration  { return c.Config().Node.IntervalSetSessions }
func (c *Context) IntervalUpdateStatus() time.Duration { return c.Config().Node.IntervalUpdateStatus }
func (c *Context) ListenOn() string                    { return c.Config().Node.ListenOn }
func (c *Context) Location() *types.GeoIPLocation      { return c.location }
func (c *Context) Log() tmlog.Logger                   { return c.logger }
func (c *Context) Moniker() string                     { return c.Config().Node.Moniker }
func (c *Context) Operator() sdk.AccAddress            { return c.client.FromAddress() }
func (c *Context) RemoteURL() string                   { return c.Config().Node.RemoteURL }
func (c *Context) Router() *mux.Router                 { return c.router }
func (c *Context) Service() types.Service              { return c.service }
func (c *Context) Sessions() *types.Sessions           { return c.sessions }

func (c *Context) IntervalUpdateSessions() time.Duration {
	return c.Config().Node.IntervalUpdateSessions
}

func (c *Context) Provider() hubtypes.ProvAddress {
	if c.Config().Node.Provider == "" {
		return nil
	}

	address, err := hubtypes.ProvAddressFromBech32(c.Config().Node.Provider)
	if err != nil {
		panic(err)
	}

	return address
}

func (c *Context) Price() sdk.Coins {
	if c.Config().Node.Price == "" {
		return nil
	}

	coins, err := sdk.ParseCoinsNormalized(c.Config().Node.Price)
	if err != nil {
		panic(err)
	}

	return coins
}