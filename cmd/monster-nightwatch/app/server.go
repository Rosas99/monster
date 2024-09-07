package app

import (
	genericapiserver "k8s.io/apiserver/pkg/server"

	"github.com/rosas99/monster/cmd/monster-nightwatch/app/options"
	"github.com/rosas99/monster/internal/nightwatch"
	"github.com/rosas99/monster/pkg/app"
)

const commandDesc = `The sms server is a standard, specification-compliant demo 
example of the monster service.

Find more monster-sms information at:
    https://"github.com/rosas99/monster/blob/master/docs/guide/en-US/cmd/monster-nightwatch.md`

// NewApp creates an App object with default parameters.
func NewApp(name string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp(name, "Launch a monster nightwatch server",
		app.WithDescription(commandDesc),
		app.WithOptions(opts),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
	return func() error {
		cfg, err := opts.Config()
		if err != nil {
			return err
		}

		return Run(cfg, genericapiserver.SetupSignalHandler())
	}
}

// Run runs the specified APIServer. This should never exit.
func Run(c *nightwatch.Config, stopCh <-chan struct{}) error {
	server, err := c.Complete().New()
	if err != nil {
		return err
	}

	return server.Run(stopCh)
}