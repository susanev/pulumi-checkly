// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@checkly/pulumi";
 *
 * const test_trigger_group = new checkly.TriggerCheckGroup("test-trigger-group", {groupId: "215"});
 * export const test_trigger_group_url = test_trigger_group.url;
 * ```
 */
export class TriggerCheckGroup extends pulumi.CustomResource {
    /**
     * Get an existing TriggerCheckGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TriggerCheckGroupState, opts?: pulumi.CustomResourceOptions): TriggerCheckGroup {
        return new TriggerCheckGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'checkly:index/triggerCheckGroup:TriggerCheckGroup';

    /**
     * Returns true if the given object is an instance of TriggerCheckGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TriggerCheckGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TriggerCheckGroup.__pulumiType;
    }

    /**
     * The id of the group that you want to attach the trigger to.
     */
    public readonly groupId!: pulumi.Output<number>;
    /**
     * The token value created to trigger the group
     */
    public readonly token!: pulumi.Output<string>;
    /**
     * The request URL to trigger the group run.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a TriggerCheckGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TriggerCheckGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TriggerCheckGroupArgs | TriggerCheckGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TriggerCheckGroupState | undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as TriggerCheckGroupArgs | undefined;
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["token"] = args ? args.token : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TriggerCheckGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TriggerCheckGroup resources.
 */
export interface TriggerCheckGroupState {
    /**
     * The id of the group that you want to attach the trigger to.
     */
    groupId?: pulumi.Input<number>;
    /**
     * The token value created to trigger the group
     */
    token?: pulumi.Input<string>;
    /**
     * The request URL to trigger the group run.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TriggerCheckGroup resource.
 */
export interface TriggerCheckGroupArgs {
    /**
     * The id of the group that you want to attach the trigger to.
     */
    groupId: pulumi.Input<number>;
    /**
     * The token value created to trigger the group
     */
    token?: pulumi.Input<string>;
    /**
     * The request URL to trigger the group run.
     */
    url?: pulumi.Input<string>;
}
