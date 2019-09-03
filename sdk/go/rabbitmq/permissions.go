// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ``.Permissions`` resource creates and manages a user's set of
// permissions.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-rabbitmq/blob/master/website/docs/r/permissions.html.markdown.
type Permissions struct {
	s *pulumi.ResourceState
}

// NewPermissions registers a new resource with the given unique name, arguments, and options.
func NewPermissions(ctx *pulumi.Context,
	name string, args *PermissionsArgs, opts ...pulumi.ResourceOpt) (*Permissions, error) {
	if args == nil || args.Permissions == nil {
		return nil, errors.New("missing required argument 'Permissions'")
	}
	if args == nil || args.User == nil {
		return nil, errors.New("missing required argument 'User'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["permissions"] = nil
		inputs["user"] = nil
		inputs["vhost"] = nil
	} else {
		inputs["permissions"] = args.Permissions
		inputs["user"] = args.User
		inputs["vhost"] = args.Vhost
	}
	s, err := ctx.RegisterResource("rabbitmq:index/permissions:Permissions", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Permissions{s: s}, nil
}

// GetPermissions gets an existing Permissions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPermissions(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PermissionsState, opts ...pulumi.ResourceOpt) (*Permissions, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["permissions"] = state.Permissions
		inputs["user"] = state.User
		inputs["vhost"] = state.Vhost
	}
	s, err := ctx.ReadResource("rabbitmq:index/permissions:Permissions", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Permissions{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Permissions) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Permissions) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The settings of the permissions. The structure is
// described below.
func (r *Permissions) Permissions() *pulumi.Output {
	return r.s.State["permissions"]
}

// The user to apply the permissions to.
func (r *Permissions) User() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["user"])
}

// The vhost to create the resource in.
func (r *Permissions) Vhost() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vhost"])
}

// Input properties used for looking up and filtering Permissions resources.
type PermissionsState struct {
	// The settings of the permissions. The structure is
	// described below.
	Permissions interface{}
	// The user to apply the permissions to.
	User interface{}
	// The vhost to create the resource in.
	Vhost interface{}
}

// The set of arguments for constructing a Permissions resource.
type PermissionsArgs struct {
	// The settings of the permissions. The structure is
	// described below.
	Permissions interface{}
	// The user to apply the permissions to.
	User interface{}
	// The vhost to create the resource in.
	Vhost interface{}
}
