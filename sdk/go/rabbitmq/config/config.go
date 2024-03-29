// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"github.com/pulumi/pulumi/sdk/go/pulumi/config"
)

func GetCacertFile(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "rabbitmq:cacertFile")
	if err == nil {
		return v
	}
	if dv, ok := getEnvOrDefault("", nil, "RABBITMQ_CACERT_FILE").(string); ok {
		return dv
	}
	return v
}

func GetEndpoint(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "rabbitmq:endpoint")
	if err == nil {
		return v
	}
	if dv, ok := getEnvOrDefault("", nil, "RABBITMQ_MANAGEMENT_ENDPOINT").(string); ok {
		return dv
	}
	return v
}

func GetInsecure(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "rabbitmq:insecure")
	if err == nil {
		return v
	}
	if dv, ok := getEnvOrDefault(false, parseEnvBool, "RABBITMQ_INSECURE").(bool); ok {
		return dv
	}
	return v
}

func GetPassword(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "rabbitmq:password")
	if err == nil {
		return v
	}
	if dv, ok := getEnvOrDefault("", nil, "RABBITMQ_PASSWORD").(string); ok {
		return dv
	}
	return v
}

func GetUsername(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "rabbitmq:username")
	if err == nil {
		return v
	}
	if dv, ok := getEnvOrDefault("", nil, "RABBITMQ_USERNAME").(string); ok {
		return dv
	}
	return v
}
