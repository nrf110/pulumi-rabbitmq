# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Policy(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the policy.
    """
    policy: pulumi.Output[dict]
    """
    The settings of the policy. The structure is
    described below.
    
      * `applyTo` (`str`)
      * `definition` (`dict`)
      * `pattern` (`str`)
      * `priority` (`float`)
    """
    vhost: pulumi.Output[str]
    """
    The vhost to create the resource in.
    """
    def __init__(__self__, resource_name, opts=None, name=None, policy=None, vhost=None, __props__=None, __name__=None, __opts__=None):
        """
        The ``.Policy`` resource creates and manages policies for exchanges
        and queues.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the policy.
        :param pulumi.Input[dict] policy: The settings of the policy. The structure is
               described below.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
        
        The **policy** object supports the following:
        
          * `applyTo` (`pulumi.Input[str]`)
          * `definition` (`pulumi.Input[dict]`)
          * `pattern` (`pulumi.Input[str]`)
          * `priority` (`pulumi.Input[float]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-rabbitmq/blob/master/website/docs/r/policy.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['name'] = name
            if policy is None:
                raise TypeError("Missing required property 'policy'")
            __props__['policy'] = policy
            if vhost is None:
                raise TypeError("Missing required property 'vhost'")
            __props__['vhost'] = vhost
        super(Policy, __self__).__init__(
            'rabbitmq:index/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, name=None, policy=None, vhost=None):
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the policy.
        :param pulumi.Input[dict] policy: The settings of the policy. The structure is
               described below.
        :param pulumi.Input[str] vhost: The vhost to create the resource in.
        
        The **policy** object supports the following:
        
          * `applyTo` (`pulumi.Input[str]`)
          * `definition` (`pulumi.Input[dict]`)
          * `pattern` (`pulumi.Input[str]`)
          * `priority` (`pulumi.Input[float]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-rabbitmq/blob/master/website/docs/r/policy.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["name"] = name
        __props__["policy"] = policy
        __props__["vhost"] = vhost
        return Policy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
