// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ``.Policy`` resource creates and manages policies for exchanges
// and queues.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rabbitmq/blob/master/website/docs/r/policy.html.markdown.
type Policy struct {
	s *pulumi.ResourceState
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOpt) (*Policy, error) {
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	if args == nil || args.Vhost == nil {
		return nil, errors.New("missing required argument 'Vhost'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["policy"] = nil
		inputs["vhost"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["policy"] = args.Policy
		inputs["vhost"] = args.Vhost
	}
	s, err := ctx.RegisterResource("rabbitmq:index/policy:Policy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Policy{s: s}, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PolicyState, opts ...pulumi.ResourceOpt) (*Policy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
		inputs["policy"] = state.Policy
		inputs["vhost"] = state.Vhost
	}
	s, err := ctx.ReadResource("rabbitmq:index/policy:Policy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Policy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Policy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Policy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The name of the policy.
func (r *Policy) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The settings of the policy. The structure is
// described below.
func (r *Policy) Policy() *pulumi.Output {
	return r.s.State["policy"]
}

// The vhost to create the resource in.
func (r *Policy) Vhost() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vhost"])
}

// Input properties used for looking up and filtering Policy resources.
type PolicyState struct {
	// The name of the policy.
	Name interface{}
	// The settings of the policy. The structure is
	// described below.
	Policy interface{}
	// The vhost to create the resource in.
	Vhost interface{}
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// The name of the policy.
	Name interface{}
	// The settings of the policy. The structure is
	// described below.
	Policy interface{}
	// The vhost to create the resource in.
	Vhost interface{}
}
