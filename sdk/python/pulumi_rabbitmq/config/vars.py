# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

__config__ = pulumi.Config('rabbitmq')

cacert_file = __config__.get('cacertFile') or utilities.get_env('RABBITMQ_CACERT_FILE')

endpoint = __config__.get('endpoint') or utilities.get_env('RABBITMQ_MANAGEMENT_ENDPOINT')

insecure = __config__.get('insecure') or utilities.get_env_bool('RABBITMQ_INSECURE')

password = __config__.get('password') or utilities.get_env('RABBITMQ_PASSWORD')

username = __config__.get('username') or utilities.get_env('RABBITMQ_USERNAME')

