// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Checkly
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Checkly = Pulumi.Checkly;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var test_trigger_check = new Checkly.TriggerCheck("test-trigger-check", new Checkly.TriggerCheckArgs
    ///         {
    ///             CheckId = "c1ff95c5-d7f6-4a90-9ce2-1e605f117592",
    ///         });
    ///         this.Test_trigger_check_url = test_trigger_check.Url;
    ///     }
    /// 
    ///     [Output("test-trigger-check-url")]
    ///     public Output&lt;string&gt; Test_trigger_check_url { get; set; }
    /// }
    /// ```
    /// </summary>
    [ChecklyResourceType("checkly:index/triggerCheck:TriggerCheck")]
    public partial class TriggerCheck : Pulumi.CustomResource
    {
        /// <summary>
        /// The id of the check that you want to attach the trigger to.
        /// </summary>
        [Output("checkId")]
        public Output<string> CheckId { get; private set; } = null!;

        /// <summary>
        /// The token value created to trigger the check
        /// </summary>
        [Output("token")]
        public Output<string> Token { get; private set; } = null!;

        /// <summary>
        /// The request URL to trigger the check run.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a TriggerCheck resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TriggerCheck(string name, TriggerCheckArgs args, CustomResourceOptions? options = null)
            : base("checkly:index/triggerCheck:TriggerCheck", name, args ?? new TriggerCheckArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TriggerCheck(string name, Input<string> id, TriggerCheckState? state = null, CustomResourceOptions? options = null)
            : base("checkly:index/triggerCheck:TriggerCheck", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/checkly/pulumi-checkly/releases/${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TriggerCheck resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TriggerCheck Get(string name, Input<string> id, TriggerCheckState? state = null, CustomResourceOptions? options = null)
        {
            return new TriggerCheck(name, id, state, options);
        }
    }

    public sealed class TriggerCheckArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the check that you want to attach the trigger to.
        /// </summary>
        [Input("checkId", required: true)]
        public Input<string> CheckId { get; set; } = null!;

        /// <summary>
        /// The token value created to trigger the check
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// The request URL to trigger the check run.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public TriggerCheckArgs()
        {
        }
    }

    public sealed class TriggerCheckState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the check that you want to attach the trigger to.
        /// </summary>
        [Input("checkId")]
        public Input<string>? CheckId { get; set; }

        /// <summary>
        /// The token value created to trigger the check
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// The request URL to trigger the check run.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public TriggerCheckState()
        {
        }
    }
}
