// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The ``rabbitmq..Exchange`` resource creates and manages an exchange.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as rabbitmq from "@pulumi/rabbitmq";
 * 
 * const testVHost = new rabbitmq.VHost("test", {});
 * const guest = new rabbitmq.Permissions("guest", {
 *     permissions: {
 *         configure: ".*",
 *         read: ".*",
 *         write: ".*",
 *     },
 *     user: "guest",
 *     vhost: testVHost.name,
 * });
 * const testExchange = new rabbitmq.Exchange("test", {
 *     settings: {
 *         autoDelete: true,
 *         durable: false,
 *         type: "fanout",
 *     },
 *     vhost: guest.vhost,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-rabbitmq/blob/master/website/docs/r/exchange.html.markdown.
 */
export class Exchange extends pulumi.CustomResource {
    /**
     * Get an existing Exchange resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExchangeState, opts?: pulumi.CustomResourceOptions): Exchange {
        return new Exchange(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'rabbitmq:index/exchange:Exchange';

    /**
     * Returns true if the given object is an instance of Exchange.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Exchange {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Exchange.__pulumiType;
    }

    /**
     * The name of the exchange.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The settings of the exchange. The structure is
     * described below.
     */
    public readonly settings!: pulumi.Output<outputs.ExchangeSettings>;
    /**
     * The vhost to create the resource in.
     */
    public readonly vhost!: pulumi.Output<string | undefined>;

    /**
     * Create a Exchange resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExchangeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExchangeArgs | ExchangeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ExchangeState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["settings"] = state ? state.settings : undefined;
            inputs["vhost"] = state ? state.vhost : undefined;
        } else {
            const args = argsOrState as ExchangeArgs | undefined;
            if (!args || args.settings === undefined) {
                throw new Error("Missing required property 'settings'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["settings"] = args ? args.settings : undefined;
            inputs["vhost"] = args ? args.vhost : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Exchange.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Exchange resources.
 */
export interface ExchangeState {
    /**
     * The name of the exchange.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The settings of the exchange. The structure is
     * described below.
     */
    readonly settings?: pulumi.Input<inputs.ExchangeSettings>;
    /**
     * The vhost to create the resource in.
     */
    readonly vhost?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Exchange resource.
 */
export interface ExchangeArgs {
    /**
     * The name of the exchange.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The settings of the exchange. The structure is
     * described below.
     */
    readonly settings: pulumi.Input<inputs.ExchangeSettings>;
    /**
     * The vhost to create the resource in.
     */
    readonly vhost?: pulumi.Input<string>;
}