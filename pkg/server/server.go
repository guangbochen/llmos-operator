package server

import (
	"context"

	dlserver "github.com/rancher/dynamiclistener/server"
	steve "github.com/rancher/steve/pkg/server"
	"github.com/rancher/wrangler/v2/pkg/k8scheck"
	"github.com/rancher/wrangler/v2/pkg/ratelimit"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"

	"github.com/llmos-ai/llmos-controller/pkg/api"
	"github.com/llmos-ai/llmos-controller/pkg/api/auth"
	"github.com/llmos-ai/llmos-controller/pkg/config"
	"github.com/llmos-ai/llmos-controller/pkg/controller"
	"github.com/llmos-ai/llmos-controller/pkg/data"
	"github.com/llmos-ai/llmos-controller/pkg/database"
	"github.com/llmos-ai/llmos-controller/pkg/generated/ent"
	sconfig "github.com/llmos-ai/llmos-controller/pkg/server/config"
	"github.com/llmos-ai/llmos-controller/pkg/server/ui"
	"github.com/llmos-ai/llmos-controller/pkg/settings"
)

type APIServer struct {
	ctx             context.Context
	kubeconfig      string
	httpListenPort  int
	httpsListenPort int
	threadiness     int
	namespace       string
	skipAuth        bool

	mgmt        *sconfig.Management
	steveServer *steve.Server
	restConfig  *rest.Config
	entClient   *ent.Client
}

// Options define the api server options
type Options struct {
	Context         context.Context
	HTTPListenPort  int
	HTTPSListenPort int
	Threadiness     int
	SkipAuth        bool

	config.CommonOptions
}

func NewServer(o Options) (*APIServer, error) {
	s := &APIServer{
		ctx:             o.Context,
		kubeconfig:      o.KubeConfig,
		httpListenPort:  o.HTTPListenPort,
		httpsListenPort: o.HTTPSListenPort,
		threadiness:     o.Threadiness,
		namespace:       o.Namespace,
		skipAuth:        o.SkipAuth,
	}

	clientConfig, err := GetConfig(s.kubeconfig)
	if err != nil {
		return nil, err
	}

	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	restConfig.RateLimiter = ratelimit.None
	s.restConfig = restConfig

	err = k8scheck.Wait(s.ctx, *restConfig)
	if err != nil {
		return nil, err
	}

	// register DB client & DB schema
	entClient, err := registerDBSchema(s.ctx)
	if err != nil {
		return nil, err
	}
	if entClient != nil {
		s.entClient = entClient
	}

	serverOptions, err := s.setDefaults(restConfig)
	if err != nil {
		return nil, err
	}

	if err = data.Init(s.mgmt); err != nil {
		return nil, err
	}

	// register the controller
	if err = controller.Register(s.ctx, s.mgmt, s.threadiness); err != nil {
		return nil, err
	}

	// set up a new api server
	s.steveServer, err = steve.New(o.Context, restConfig, serverOptions)
	if err != nil {
		return nil, err
	}

	// configure the api ui
	ui.ConfigureAPIUI(s.steveServer.APIServer)

	// register api schemas
	if err = api.Register(s.ctx, s.mgmt, s.steveServer); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *APIServer) setDefaults(cfg *rest.Config) (*steve.Options, error) {
	var err error
	opts := &steve.Options{}

	// set up the management config
	s.mgmt, err = sconfig.SetupManagement(s.ctx, cfg, s.namespace, s.entClient)
	if err != nil {
		return nil, err
	}

	// define the next handler after the mgmt is setup
	r := NewRouter(s.mgmt)
	opts.Next = r.Routes()

	// define auth middleware
	if !s.skipAuth {
		auth := auth.NewMiddleware(s.mgmt)
		opts.AuthMiddleware = auth.AuthMiddleware
	}

	return opts, nil
}

func (s *APIServer) ListenAndServe(opts *dlserver.ListenOpts) error {
	return s.steveServer.ListenAndServe(s.ctx, s.httpsListenPort, s.httpListenPort, opts)
}

func registerDBSchema(ctx context.Context) (*ent.Client, error) {
	if settings.DatabaseURL.Get() != "" {
		return database.RunAutoMigrate(ctx)
	}
	logrus.Warnf("DatabaseURL is not set, skipping database auto migration")
	return nil, nil
}
