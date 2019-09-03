// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ``.Exchange`` resource creates and manages an exchange.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rabbitmq/blob/master/website/docs/r/exchange.html.markdown.
type Exchange struct {
	s *pulumi.ResourceState
}

// NewExchange registers a new resource with the given unique name, arguments, and options.
func NewExchange(ctx *pulumi.Context,
	name string, args *ExchangeArgs, opts ...pulumi.ResourceOpt) (*Exchange, error) {
	if args == nil || args.Settings == nil {
		return nil, errors.New("missing required argument 'Settings'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["settings"] = nil
		inputs["vhost"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["settings"] = args.Settings
		inputs["vhost"] = args.Vhost
	}
	s, err := ctx.RegisterResource("rabbitmq:index/exchange:Exchange", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Exchange{s: s}, nil
}

// GetExchange gets an existing Exchange resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExchange(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ExchangeState, opts ...pulumi.ResourceOpt) (*Exchange, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
		inputs["settings"] = state.Settings
		inputs["vhost"] = state.Vhost
	}
	s, err := ctx.ReadResource("rabbitmq:index/exchange:Exchange", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Exchange{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Exchange) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Exchange) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The name of the exchange.
func (r *Exchange) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The settings of the exchange. The structure is
// described below.
func (r *Exchange) Settings() *pulumi.Output {
	return r.s.State["settings"]
}

// The vhost to create the resource in.
func (r *Exchange) Vhost() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vhost"])
}

// Input properties used for looking up and filtering Exchange resources.
type ExchangeState struct {
	// The name of the exchange.
	Name interface{}
	// The settings of the exchange. The structure is
	// described below.
	Settings interface{}
	// The vhost to create the resource in.
	Vhost interface{}
}

// The set of arguments for constructing a Exchange resource.
type ExchangeArgs struct {
	// The name of the exchange.
	Name interface{}
	// The settings of the exchange. The structure is
	// described below.
	Settings interface{}
	// The vhost to create the resource in.
	Vhost interface{}
}