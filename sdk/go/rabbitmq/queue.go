// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ``.Queue`` resource creates and manages a queue.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rabbitmq/blob/master/website/docs/r/queue.html.markdown.
type Queue struct {
	s *pulumi.ResourceState
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOpt) (*Queue, error) {
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
	s, err := ctx.RegisterResource("rabbitmq:index/queue:Queue", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Queue{s: s}, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.ID, state *QueueState, opts ...pulumi.ResourceOpt) (*Queue, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
		inputs["settings"] = state.Settings
		inputs["vhost"] = state.Vhost
	}
	s, err := ctx.ReadResource("rabbitmq:index/queue:Queue", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Queue{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Queue) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Queue) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The name of the queue.
func (r *Queue) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The settings of the queue. The structure is
// described below.
func (r *Queue) Settings() *pulumi.Output {
	return r.s.State["settings"]
}

// The vhost to create the resource in.
func (r *Queue) Vhost() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vhost"])
}

// Input properties used for looking up and filtering Queue resources.
type QueueState struct {
	// The name of the queue.
	Name interface{}
	// The settings of the queue. The structure is
	// described below.
	Settings interface{}
	// The vhost to create the resource in.
	Vhost interface{}
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// The name of the queue.
	Name interface{}
	// The settings of the queue. The structure is
	// described below.
	Settings interface{}
	// The vhost to create the resource in.
	Vhost interface{}
}