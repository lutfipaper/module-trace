package lib

import (
	"github.com/lutfipaper/module-trace/interfaces"
	// "github.com/lutfiharidha/Signoz-go/interfaces"
)

var Version string
var Commit string

type Modules struct {
	Option interfaces.Option
	Signoz *SignozOpenTelemetry
}

func NewLib() interfaces.Tracing {
	return &Modules{}
}

func (c *Modules) New() interfaces.Tracing {
	return NewLib()
}

func (c *Modules) Init(option interfaces.Option) {

	c.Option = option

	if c.Option.Config.Signoz.Enable {
		c.Signoz = NewSignozOpenTelemetry(c.Option)
		_ = c.Signoz.Setup()
	}

}

func (c *Modules) Closing() (err error) {
	if c.Signoz != nil {
		err = c.Signoz.Closing()
	}
	return err
}
