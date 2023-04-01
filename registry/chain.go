package main

import (
	"fmt"
	"os"
	"sync"

	pb "github.com/Anmol1696/starship/registry/registry"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibctm "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	lens "github.com/strangelove-ventures/lens/client"
	"github.com/strangelove-ventures/lens/client/query"
	"go.uber.org/zap"
)

type ChainClients []*ChainClient

// NewChainClients returns a list of chain clients from a list of strings
func NewChainClients(logger *zap.Logger, chainIDs, rpcAddrs []string, home string) (ChainClients, error) {
	// make home if non existant
	_ = os.MkdirAll(home, 0755)

	var clients []*ChainClient
	for i := range chainIDs {
		client, err := NewChainClient(logger, chainIDs[i], rpcAddrs[i], home)

		if err != nil {
			logger.Error("unable to create client for chain",
				zap.String("chain_id", chainIDs[i]),
				zap.String("rpc_addr", rpcAddrs[i]),
				zap.Error(err),
			)
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

// GetChainClient returns a chain client pointer for the given chain id
func (cc ChainClients) GetChainClient(chainID string) (*ChainClient, error) {
	for _, client := range cc {
		if client.chainConfig.ChainID == chainID {
			return client, nil
		}
	}

	return nil, fmt.Errorf("not found: client chain id %s", chainID)
}

type ChainClient struct {
	logger      *zap.Logger
	chainConfig *lens.ChainClientConfig

	client *lens.ChainClient

	mu           sync.Mutex
	chainIBCInfo []*ChainIBCInfo
}

func NewChainClient(logger *zap.Logger, chainID, rpcAddr, home string) (*ChainClient, error) {
	ccc := &lens.ChainClientConfig{
		ChainID:        chainID,
		RPCAddr:        rpcAddr,
		KeyringBackend: "test",
		Debug:          true,
		Timeout:        "20s",
		SignModeStr:    "direct",
	}
	client, err := lens.NewChainClient(logger, ccc, home, os.Stdin, os.Stdout)
	if err != nil {
		return nil, err
	}

	chainClient := &ChainClient{
		logger:       logger,
		chainConfig:  ccc,
		client:       client,
		chainIBCInfo: nil,
	}

	// Cache initial values, the best effort
	_, _ = chainClient.GetCachedChainInfo()

	return chainClient, nil
}

// getChainSeedPeers returns the nodes for the genesis node
func (c *ChainClient) getChainSeedPeers() ([]*pb.Peer, error) {

}

func (c *ChainClient) getChainPersistencePeers() ([]*pb.Peer, error) {

}

// getChannelPort returns the chains and the counterparty info
func (c *ChainClient) getChannelsPorts() ([]ChannelsInfo, error) {
	querier := query.Query{Client: c.client, Options: query.DefaultOptions()}

	c.logger.Info("created a querier object, making the first query")
	channels, err := querier.Ibc_Channels()
	if err != nil {
		return nil, err
	}
	c.logger.Info("channels queried from the upstream", zap.Any("channels", channels))

	var channelsInfo []ChannelsInfo
	for _, channel := range channels.Channels {
		// use connection hop to get connections info
		if len(channel.ConnectionHops) != 1 {
			return nil, fmt.Errorf("number of connections not 1")
		}

		ci := ChannelsInfo{
			ChannelPort: ChannelPort{
				ChannelId: channel.ChannelId,
				PortId:    channel.PortId,
			},
			Counterparty: ChannelPort{
				ChannelId: channel.Counterparty.ChannelId,
				PortId:    channel.Counterparty.PortId,
			},
			ConnectionId: channel.ConnectionHops[0],
			Ordering:     channel.Ordering.String(),
			Version:      channel.Version,
		}

		channelsInfo = append(channelsInfo, ci)
	}

	return channelsInfo, nil
}

func (c *ChainClient) getConnectionClient(connectionId string) (*ConnectionInfo, error) {
	querier := query.Query{Client: c.client, Options: query.DefaultOptions()}

	connection, err := querier.Ibc_Connection(connectionId)
	if err != nil {
		return nil, err
	}

	c.logger.Info("connection queried from the upstream",
		zap.String("connection_id", connectionId),
		zap.Any("connection", connection))

	return &ConnectionInfo{
		ConnectionClient: ConnectionClient{
			ConnectionId: connectionId,
			ClientId:     connection.Connection.ClientId,
		},
		Counterparty: ConnectionClient{
			ConnectionId: connection.Connection.Counterparty.ConnectionId,
			ClientId:     connection.Connection.Counterparty.ClientId,
		},
	}, nil
}

// GetIBCClients will fetch all the IBC channels for the chain
func (c *ChainClient) getChainIdFromClient(clientId string) (string, error) {
	querier := query.Query{Client: c.client, Options: query.DefaultOptions()}

	c.logger.Info("making query for chain id from client id", zap.String("clientId", clientId))
	state, err := querier.Ibc_ClientState(clientId)
	if err != nil {
		return "", err
	}
	c.logger.Info("query ibc client state", zap.String("clientId", clientId), zap.Any("state", state))

	clientState, err := clienttypes.UnpackClientState(state.ClientState)
	if err != nil {
		return "", err
	}

	if clientState.ClientType() != exported.Tendermint {
		return "", fmt.Errorf("client state type not %s", exported.Tendermint)
	}

	cs, ok := clientState.(*ibctm.ClientState)
	if !ok {
		return "", fmt.Errorf("unable to convert client state to lightclient ClientState, client: %s", clientState.String())
	}

	return cs.ChainId, nil
}

// GetChainInfo will fetch all the IBC channels for the chain
func (c *ChainClient) GetChainInfo() ([]*ChainIBCInfo, error) {
	channelsInfo, err := c.getChannelsPorts()
	if err != nil {
		return nil, err
	}

	var chainIBCInfos []*ChainIBCInfo
	for _, channelInfo := range channelsInfo {
		connectionInfo, err := c.getConnectionClient(channelInfo.ConnectionId)
		if err != nil {
			return nil, err
		}

		cpChainId, err := c.getChainIdFromClient(connectionInfo.ClientId)
		if err != nil {
			return nil, err
		}

		cii := &ChainIBCInfo{
			IBCInfo: IBCInfo{
				ChainId:      c.chainConfig.ChainID,
				ChannelId:    channelInfo.ChannelId,
				PortId:       channelInfo.PortId,
				ConnectionId: connectionInfo.ConnectionId,
				ClientId:     connectionInfo.ClientId,
			},
			Counterparty: IBCInfo{
				ChainId:      cpChainId,
				ChannelId:    channelInfo.Counterparty.ChannelId,
				PortId:       channelInfo.Counterparty.PortId,
				ConnectionId: connectionInfo.Counterparty.ConnectionId,
				ClientId:     connectionInfo.Counterparty.ClientId,
			},
			Ordering: channelInfo.Ordering,
			Version:  channelInfo.Version,
			State:    channelInfo.State,
		}

		chainIBCInfos = append(chainIBCInfos, cii)
	}

	return chainIBCInfos, nil
}

// GetCachedChainInfo will return cached chain info, if no cache, then will cache the info
func (c *ChainClient) GetCachedChainInfo() ([]*ChainIBCInfo, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// try and fetch from cache
	if c.chainIBCInfo != nil {
		return c.chainIBCInfo, nil
	}

	// set cache if unset
	chainInfo, err := c.GetChainInfo()
	if err != nil {
		c.logger.Error("unable to cache value for chain info", zap.Error(err))
		return nil, err
	}
	c.chainIBCInfo = chainInfo

	return c.chainIBCInfo, nil
}
