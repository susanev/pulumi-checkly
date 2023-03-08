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
    ///         var dashboard1 = new Checkly.Dashboard("dashboard1", new Checkly.DashboardArgs
    ///         {
    ///             CustomDomain = "status.example.com",
    ///             CustomUrl = "checkly",
    ///             Header = "Public dashboard",
    ///             HideTags = false,
    ///             Logo = "https://www.checklyhq.com/logo.png",
    ///             Paginate = false,
    ///             PaginationRate = 30,
    ///             RefreshRate = 60,
    ///             Tags = 
    ///             {
    ///                 "production",
    ///             },
    ///             Width = "FULL",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [ChecklyResourceType("checkly:index/dashboard:Dashboard")]
    public partial class Dashboard : Pulumi.CustomResource
    {
        /// <summary>
        /// Determines how many checks to show per page.
        /// </summary>
        [Output("checksPerPage")]
        public Output<int?> ChecksPerPage { get; private set; } = null!;

        /// <summary>
        /// A custom user domain, e.g. 'status.example.com'. See the docs on updating your DNS and SSL usage.
        /// </summary>
        [Output("customDomain")]
        public Output<string?> CustomDomain { get; private set; } = null!;

        /// <summary>
        /// A subdomain name under 'checklyhq.com'. Needs to be unique across all users.
        /// </summary>
        [Output("customUrl")]
        public Output<string> CustomUrl { get; private set; } = null!;

        /// <summary>
        /// HTML &lt;meta&gt; description for the dashboard.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A URL pointing to an image file to use as browser favicon.
        /// </summary>
        [Output("favicon")]
        public Output<string?> Favicon { get; private set; } = null!;

        /// <summary>
        /// A piece of text displayed at the top of your dashboard.
        /// </summary>
        [Output("header")]
        public Output<string?> Header { get; private set; } = null!;

        /// <summary>
        /// Show or hide the tags on the dashboard.
        /// </summary>
        [Output("hideTags")]
        public Output<bool?> HideTags { get; private set; } = null!;

        /// <summary>
        /// A link to for the dashboard logo.
        /// </summary>
        [Output("link")]
        public Output<string?> Link { get; private set; } = null!;

        /// <summary>
        /// A URL pointing to an image file to use for the dashboard logo.
        /// </summary>
        [Output("logo")]
        public Output<string?> Logo { get; private set; } = null!;

        /// <summary>
        /// Determines if pagination is on or off.
        /// </summary>
        [Output("paginate")]
        public Output<bool?> Paginate { get; private set; } = null!;

        /// <summary>
        /// How often to trigger pagination in seconds. Possible values `30`, `60` and `300`.
        /// </summary>
        [Output("paginationRate")]
        public Output<int?> PaginationRate { get; private set; } = null!;

        /// <summary>
        /// How often to refresh the dashboard in seconds. Possible values `60`, '300' and `600`.
        /// </summary>
        [Output("refreshRate")]
        public Output<int?> RefreshRate { get; private set; } = null!;

        /// <summary>
        /// A list of one or more tags that filter which checks to display on the dashboard.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Set when to use AND operator for fetching dashboard tags.
        /// </summary>
        [Output("useTagsAndOperator")]
        public Output<bool?> UseTagsAndOperator { get; private set; } = null!;

        /// <summary>
        /// Determines whether to use the full screen or focus in the center. Possible values `FULL` and `960PX`.
        /// </summary>
        [Output("width")]
        public Output<string?> Width { get; private set; } = null!;


        /// <summary>
        /// Create a Dashboard resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dashboard(string name, DashboardArgs args, CustomResourceOptions? options = null)
            : base("checkly:index/dashboard:Dashboard", name, args ?? new DashboardArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dashboard(string name, Input<string> id, DashboardState? state = null, CustomResourceOptions? options = null)
            : base("checkly:index/dashboard:Dashboard", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/checkly",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Dashboard resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dashboard Get(string name, Input<string> id, DashboardState? state = null, CustomResourceOptions? options = null)
        {
            return new Dashboard(name, id, state, options);
        }
    }

    public sealed class DashboardArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines how many checks to show per page.
        /// </summary>
        [Input("checksPerPage")]
        public Input<int>? ChecksPerPage { get; set; }

        /// <summary>
        /// A custom user domain, e.g. 'status.example.com'. See the docs on updating your DNS and SSL usage.
        /// </summary>
        [Input("customDomain")]
        public Input<string>? CustomDomain { get; set; }

        /// <summary>
        /// A subdomain name under 'checklyhq.com'. Needs to be unique across all users.
        /// </summary>
        [Input("customUrl", required: true)]
        public Input<string> CustomUrl { get; set; } = null!;

        /// <summary>
        /// HTML &lt;meta&gt; description for the dashboard.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A URL pointing to an image file to use as browser favicon.
        /// </summary>
        [Input("favicon")]
        public Input<string>? Favicon { get; set; }

        /// <summary>
        /// A piece of text displayed at the top of your dashboard.
        /// </summary>
        [Input("header")]
        public Input<string>? Header { get; set; }

        /// <summary>
        /// Show or hide the tags on the dashboard.
        /// </summary>
        [Input("hideTags")]
        public Input<bool>? HideTags { get; set; }

        /// <summary>
        /// A link to for the dashboard logo.
        /// </summary>
        [Input("link")]
        public Input<string>? Link { get; set; }

        /// <summary>
        /// A URL pointing to an image file to use for the dashboard logo.
        /// </summary>
        [Input("logo")]
        public Input<string>? Logo { get; set; }

        /// <summary>
        /// Determines if pagination is on or off.
        /// </summary>
        [Input("paginate")]
        public Input<bool>? Paginate { get; set; }

        /// <summary>
        /// How often to trigger pagination in seconds. Possible values `30`, `60` and `300`.
        /// </summary>
        [Input("paginationRate")]
        public Input<int>? PaginationRate { get; set; }

        /// <summary>
        /// How often to refresh the dashboard in seconds. Possible values `60`, '300' and `600`.
        /// </summary>
        [Input("refreshRate")]
        public Input<int>? RefreshRate { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of one or more tags that filter which checks to display on the dashboard.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Set when to use AND operator for fetching dashboard tags.
        /// </summary>
        [Input("useTagsAndOperator")]
        public Input<bool>? UseTagsAndOperator { get; set; }

        /// <summary>
        /// Determines whether to use the full screen or focus in the center. Possible values `FULL` and `960PX`.
        /// </summary>
        [Input("width")]
        public Input<string>? Width { get; set; }

        public DashboardArgs()
        {
        }
    }

    public sealed class DashboardState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines how many checks to show per page.
        /// </summary>
        [Input("checksPerPage")]
        public Input<int>? ChecksPerPage { get; set; }

        /// <summary>
        /// A custom user domain, e.g. 'status.example.com'. See the docs on updating your DNS and SSL usage.
        /// </summary>
        [Input("customDomain")]
        public Input<string>? CustomDomain { get; set; }

        /// <summary>
        /// A subdomain name under 'checklyhq.com'. Needs to be unique across all users.
        /// </summary>
        [Input("customUrl")]
        public Input<string>? CustomUrl { get; set; }

        /// <summary>
        /// HTML &lt;meta&gt; description for the dashboard.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A URL pointing to an image file to use as browser favicon.
        /// </summary>
        [Input("favicon")]
        public Input<string>? Favicon { get; set; }

        /// <summary>
        /// A piece of text displayed at the top of your dashboard.
        /// </summary>
        [Input("header")]
        public Input<string>? Header { get; set; }

        /// <summary>
        /// Show or hide the tags on the dashboard.
        /// </summary>
        [Input("hideTags")]
        public Input<bool>? HideTags { get; set; }

        /// <summary>
        /// A link to for the dashboard logo.
        /// </summary>
        [Input("link")]
        public Input<string>? Link { get; set; }

        /// <summary>
        /// A URL pointing to an image file to use for the dashboard logo.
        /// </summary>
        [Input("logo")]
        public Input<string>? Logo { get; set; }

        /// <summary>
        /// Determines if pagination is on or off.
        /// </summary>
        [Input("paginate")]
        public Input<bool>? Paginate { get; set; }

        /// <summary>
        /// How often to trigger pagination in seconds. Possible values `30`, `60` and `300`.
        /// </summary>
        [Input("paginationRate")]
        public Input<int>? PaginationRate { get; set; }

        /// <summary>
        /// How often to refresh the dashboard in seconds. Possible values `60`, '300' and `600`.
        /// </summary>
        [Input("refreshRate")]
        public Input<int>? RefreshRate { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of one or more tags that filter which checks to display on the dashboard.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Set when to use AND operator for fetching dashboard tags.
        /// </summary>
        [Input("useTagsAndOperator")]
        public Input<bool>? UseTagsAndOperator { get; set; }

        /// <summary>
        /// Determines whether to use the full screen or focus in the center. Possible values `FULL` and `960PX`.
        /// </summary>
        [Input("width")]
        public Input<string>? Width { get; set; }

        public DashboardState()
        {
        }
    }
}
