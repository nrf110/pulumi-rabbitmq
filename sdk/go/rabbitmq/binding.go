// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ``.Binding`` resource creates and manages a binding relationship
// between a queue an exchange.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rabbitmq/blob/master/website/docs/r/binding.html.markdown.
type Binding struct {
	s *pulumi.ResourceState
}

// NewBinding registers a new resource with the given unique name, arguments, and options.
func NewBinding(ctx *pulumi.Context,
	name string, args *BindingArgs, opts ...pulumi.ResourceOpt) (*Binding, error) {
	if args == nil || args.Destination == nil {
		return nil, errors.New("missing required argument 'Destination'")
	}
	if args == nil || args.DestinationType == nil {
		return nil, errors.New("missing required argument 'DestinationType'")
	}
	if args == nil || args.Source == nil {
		return nil, errors.New("missing required argument 'Source'")
	}
	if args == nil || args.Vhost == nil {
		return nil, errors.New("missing required argument 'Vhost'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["arguments"] = nil
		inputs["destination"] = nil
		inputs["destinationType"] = nil
		inputs["routingKey"] = nil
		inputs["source"] = nil
		inputs["vhost"] = nil
	} else {
		inputs["arguments"] = args.Arguments
		inputs["destination"] = args.Destination
		inputs["destinationType"] = args.DestinationType
		inputs["routingKey"] = args.RoutingKey
		inputs["source"] = args.Source
		inputs["vhost"] = args.Vhost
	}
	inputs["propertiesKey"] = nil
	s, err := ctx.RegisterResource("rabbitmq:index/binding:Binding", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Binding{s: s}, nil
}

// GetBinding gets an existing Binding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBinding(ctx *pulumi.Context,
	name string, id pulumi.ID, state *BindingState, opts ...pulumi.ResourceOpt) (*Binding, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arguments"] = state.Arguments
		inputs["destination"] = state.Destination
		inputs["destinationType"] = state.DestinationType
		inputs["propertiesKey"] = state.PropertiesKey
		inputs["routingKey"] = state.RoutingKey
		inputs["source"] = state.Source
		inputs["vhost"] = state.Vhost
	}
	s, err := ctx.ReadResource("rabbitmq:index/binding:Binding", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Binding{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Binding) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Binding) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Additional key/value arguments for the binding.
func (r *Binding) Arguments() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["arguments"])
}

// The destination queue or exchange.
func (r *Binding) Destination() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["destination"])
}

// The type of destination (queue or exchange).
func (r *Binding) DestinationType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["destinationType"])
}

// A unique key to refer to the binding.
func (r *Binding) PropertiesKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["propertiesKey"])
}

// A routing key for the binding.
func (r *Binding) RoutingKey() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["routingKey"])
}

// The source exchange.
func (r *Binding) Source() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["source"])
}

// The vhost to create the resource in.
func (r *Binding) Vhost() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vhost"])
}

// Input properties used for looking up and filtering Binding resources.
type BindingState struct {
	// Additional key/value arguments for the binding.
	Arguments interface{}
	// The destination queue or exchange.
	Destination interface{}
	// The type of destination (queue or exchange).
	DestinationType interface{}
	// A unique key to refer to the binding.
	PropertiesKey interface{}
	// A routing key for the binding.
	RoutingKey interface{}
	// The source exchange.
	Source interface{}
	// The vhost to create the resource in.
	Vhost interface{}
}

// The set of arguments for constructing a Binding resource.
type BindingArgs struct {
	// Additional key/value arguments for the binding.
	Arguments interface{}
	// The destination queue or exchange.
	Destination interface{}
	// The type of destination (queue or exchange).
	DestinationType interface{}
	// A routing key for the binding.
	RoutingKey interface{}
	// The source exchange.
	Source interface{}
	// The vhost to create the resource in.
	Vhost interface{}
}