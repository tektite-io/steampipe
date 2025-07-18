package pluginmanager

import (
	"io"
	"log"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc"
	"github.com/turbot/steampipe-plugin-sdk/v5/logging"
	pb "github.com/turbot/steampipe/v2/pkg/pluginmanager_service/grpc/proto"
	pluginshared "github.com/turbot/steampipe/v2/pkg/pluginmanager_service/grpc/shared"
)

// PluginManagerClient is the client used by steampipe to access the plugin manager
type PluginManagerClient struct {
	manager            pluginshared.PluginManager
	pluginManagerState *State
	rawClient          *plugin.Client
}

func NewPluginManagerClient(pluginManagerState *State) (*PluginManagerClient, error) {
	res := &PluginManagerClient{
		pluginManagerState: pluginManagerState,
	}
	err := res.attachToPluginManager()
	if err != nil {
		log.Printf("[TRACE] failed to attach to plugin manager: %s", err.Error())
		return nil, err
	}

	return res, nil
}

func (c *PluginManagerClient) attachToPluginManager() error {
	// discard logging from the plugin client (plugin logs will still flow through)
	loggOpts := &hclog.LoggerOptions{Name: "plugin", Output: io.Discard}
	logger := logging.NewLogger(loggOpts)

	// construct a client using the plugin manager reaattach config
	c.rawClient = plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: pluginshared.Handshake,
		Plugins:         pluginshared.PluginMap,
		Reattach:        c.pluginManagerState.reattachConfig(),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
		Logger: logger,
	})

	// connect via RPC
	rpcClient, err := c.rawClient.Client()
	if err != nil {
		log.Printf("[TRACE] failed to connect to plugin manager: %s", err.Error())
		return err
	}

	// request the plugin
	raw, err := rpcClient.Dispense(pluginshared.PluginName)
	if err != nil {
		log.Printf("[TRACE] failed to retreive to plugin manager from running plugin process: %s", err.Error())
		return err
	}

	// cast to correct type
	pluginManager := raw.(pluginshared.PluginManager)
	c.manager = pluginManager

	return nil
}

func (c *PluginManagerClient) Get(req *pb.GetRequest) (*pb.GetResponse, error) {
	res, err := c.manager.Get(req)
	if err != nil {
		return nil, grpc.HandleGrpcError(err, "PluginManager", "Get")
	}
	return res, nil
}

func (c *PluginManagerClient) RefreshConnections(req *pb.RefreshConnectionsRequest) (*pb.RefreshConnectionsResponse, error) {
	res, err := c.manager.RefreshConnections(req)
	if err != nil {
		return nil, grpc.HandleGrpcError(err, "PluginManager", "RefreshConnections")
	}
	return res, nil
}

func (c *PluginManagerClient) Shutdown(req *pb.ShutdownRequest) (*pb.ShutdownResponse, error) {
	log.Printf("[DEBUG] PluginManagerClient.Shutdown start")
	defer log.Printf("[DEBUG] PluginManagerClient.Shutdown done")

	res, err := c.manager.Shutdown(req)
	if err != nil {
		return nil, grpc.HandleGrpcError(err, "PluginManager", "Get")
	}
	return res, nil
}
