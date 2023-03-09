// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package checkly

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/checkly/pulumi-checkly/sdk/go/checkly"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := checkly.NewPrivateLocation(ctx, "location", &checkly.PrivateLocationArgs{
// 			SlugName: pulumi.String("new-private-location"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type PrivateLocation struct {
	pulumi.CustomResourceState

	// Icon assigned to the private location.
	Icon pulumi.StringPtrOutput `pulumi:"icon"`
	// Private location API keys.
	Keys pulumi.StringArrayOutput `pulumi:"keys"`
	// The private location name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Valid slug name.
	SlugName pulumi.StringOutput `pulumi:"slugName"`
}

// NewPrivateLocation registers a new resource with the given unique name, arguments, and options.
func NewPrivateLocation(ctx *pulumi.Context,
	name string, args *PrivateLocationArgs, opts ...pulumi.ResourceOption) (*PrivateLocation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SlugName == nil {
		return nil, errors.New("invalid value for required argument 'SlugName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource PrivateLocation
	err := ctx.RegisterResource("checkly:index/privateLocation:PrivateLocation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateLocation gets an existing PrivateLocation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateLocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLocationState, opts ...pulumi.ResourceOption) (*PrivateLocation, error) {
	var resource PrivateLocation
	err := ctx.ReadResource("checkly:index/privateLocation:PrivateLocation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateLocation resources.
type privateLocationState struct {
	// Icon assigned to the private location.
	Icon *string `pulumi:"icon"`
	// Private location API keys.
	Keys []string `pulumi:"keys"`
	// The private location name.
	Name *string `pulumi:"name"`
	// Valid slug name.
	SlugName *string `pulumi:"slugName"`
}

type PrivateLocationState struct {
	// Icon assigned to the private location.
	Icon pulumi.StringPtrInput
	// Private location API keys.
	Keys pulumi.StringArrayInput
	// The private location name.
	Name pulumi.StringPtrInput
	// Valid slug name.
	SlugName pulumi.StringPtrInput
}

func (PrivateLocationState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLocationState)(nil)).Elem()
}

type privateLocationArgs struct {
	// Icon assigned to the private location.
	Icon *string `pulumi:"icon"`
	// The private location name.
	Name *string `pulumi:"name"`
	// Valid slug name.
	SlugName string `pulumi:"slugName"`
}

// The set of arguments for constructing a PrivateLocation resource.
type PrivateLocationArgs struct {
	// Icon assigned to the private location.
	Icon pulumi.StringPtrInput
	// The private location name.
	Name pulumi.StringPtrInput
	// Valid slug name.
	SlugName pulumi.StringInput
}

func (PrivateLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLocationArgs)(nil)).Elem()
}

type PrivateLocationInput interface {
	pulumi.Input

	ToPrivateLocationOutput() PrivateLocationOutput
	ToPrivateLocationOutputWithContext(ctx context.Context) PrivateLocationOutput
}

func (*PrivateLocation) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLocation)(nil)).Elem()
}

func (i *PrivateLocation) ToPrivateLocationOutput() PrivateLocationOutput {
	return i.ToPrivateLocationOutputWithContext(context.Background())
}

func (i *PrivateLocation) ToPrivateLocationOutputWithContext(ctx context.Context) PrivateLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLocationOutput)
}

// PrivateLocationArrayInput is an input type that accepts PrivateLocationArray and PrivateLocationArrayOutput values.
// You can construct a concrete instance of `PrivateLocationArrayInput` via:
//
//          PrivateLocationArray{ PrivateLocationArgs{...} }
type PrivateLocationArrayInput interface {
	pulumi.Input

	ToPrivateLocationArrayOutput() PrivateLocationArrayOutput
	ToPrivateLocationArrayOutputWithContext(context.Context) PrivateLocationArrayOutput
}

type PrivateLocationArray []PrivateLocationInput

func (PrivateLocationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateLocation)(nil)).Elem()
}

func (i PrivateLocationArray) ToPrivateLocationArrayOutput() PrivateLocationArrayOutput {
	return i.ToPrivateLocationArrayOutputWithContext(context.Background())
}

func (i PrivateLocationArray) ToPrivateLocationArrayOutputWithContext(ctx context.Context) PrivateLocationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLocationArrayOutput)
}

// PrivateLocationMapInput is an input type that accepts PrivateLocationMap and PrivateLocationMapOutput values.
// You can construct a concrete instance of `PrivateLocationMapInput` via:
//
//          PrivateLocationMap{ "key": PrivateLocationArgs{...} }
type PrivateLocationMapInput interface {
	pulumi.Input

	ToPrivateLocationMapOutput() PrivateLocationMapOutput
	ToPrivateLocationMapOutputWithContext(context.Context) PrivateLocationMapOutput
}

type PrivateLocationMap map[string]PrivateLocationInput

func (PrivateLocationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateLocation)(nil)).Elem()
}

func (i PrivateLocationMap) ToPrivateLocationMapOutput() PrivateLocationMapOutput {
	return i.ToPrivateLocationMapOutputWithContext(context.Background())
}

func (i PrivateLocationMap) ToPrivateLocationMapOutputWithContext(ctx context.Context) PrivateLocationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLocationMapOutput)
}

type PrivateLocationOutput struct{ *pulumi.OutputState }

func (PrivateLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLocation)(nil)).Elem()
}

func (o PrivateLocationOutput) ToPrivateLocationOutput() PrivateLocationOutput {
	return o
}

func (o PrivateLocationOutput) ToPrivateLocationOutputWithContext(ctx context.Context) PrivateLocationOutput {
	return o
}

type PrivateLocationArrayOutput struct{ *pulumi.OutputState }

func (PrivateLocationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateLocation)(nil)).Elem()
}

func (o PrivateLocationArrayOutput) ToPrivateLocationArrayOutput() PrivateLocationArrayOutput {
	return o
}

func (o PrivateLocationArrayOutput) ToPrivateLocationArrayOutputWithContext(ctx context.Context) PrivateLocationArrayOutput {
	return o
}

func (o PrivateLocationArrayOutput) Index(i pulumi.IntInput) PrivateLocationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrivateLocation {
		return vs[0].([]*PrivateLocation)[vs[1].(int)]
	}).(PrivateLocationOutput)
}

type PrivateLocationMapOutput struct{ *pulumi.OutputState }

func (PrivateLocationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateLocation)(nil)).Elem()
}

func (o PrivateLocationMapOutput) ToPrivateLocationMapOutput() PrivateLocationMapOutput {
	return o
}

func (o PrivateLocationMapOutput) ToPrivateLocationMapOutputWithContext(ctx context.Context) PrivateLocationMapOutput {
	return o
}

func (o PrivateLocationMapOutput) MapIndex(k pulumi.StringInput) PrivateLocationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrivateLocation {
		return vs[0].(map[string]*PrivateLocation)[vs[1].(string)]
	}).(PrivateLocationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLocationInput)(nil)).Elem(), &PrivateLocation{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLocationArrayInput)(nil)).Elem(), PrivateLocationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLocationMapInput)(nil)).Elem(), PrivateLocationMap{})
	pulumi.RegisterOutputType(PrivateLocationOutput{})
	pulumi.RegisterOutputType(PrivateLocationArrayOutput{})
	pulumi.RegisterOutputType(PrivateLocationMapOutput{})
}
