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
// 		_, err := checkly.NewEnvironmentVariable(ctx, "variable-1", &checkly.EnvironmentVariableArgs{
// 			Key:    pulumi.String("API_KEY"),
// 			Locked: pulumi.Bool(true),
// 			Value:  pulumi.String("loZd9hOGHDUrGvmW"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = checkly.NewEnvironmentVariable(ctx, "variable-2", &checkly.EnvironmentVariableArgs{
// 			Key:   pulumi.String("API_URL"),
// 			Value: pulumi.String("http://localhost:3000"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type EnvironmentVariable struct {
	pulumi.CustomResourceState

	Key    pulumi.StringOutput  `pulumi:"key"`
	Locked pulumi.BoolPtrOutput `pulumi:"locked"`
	Value  pulumi.StringOutput  `pulumi:"value"`
}

// NewEnvironmentVariable registers a new resource with the given unique name, arguments, and options.
func NewEnvironmentVariable(ctx *pulumi.Context,
	name string, args *EnvironmentVariableArgs, opts ...pulumi.ResourceOption) (*EnvironmentVariable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource EnvironmentVariable
	err := ctx.RegisterResource("checkly:index/environmentVariable:EnvironmentVariable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironmentVariable gets an existing EnvironmentVariable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironmentVariable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentVariableState, opts ...pulumi.ResourceOption) (*EnvironmentVariable, error) {
	var resource EnvironmentVariable
	err := ctx.ReadResource("checkly:index/environmentVariable:EnvironmentVariable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvironmentVariable resources.
type environmentVariableState struct {
	Key    *string `pulumi:"key"`
	Locked *bool   `pulumi:"locked"`
	Value  *string `pulumi:"value"`
}

type EnvironmentVariableState struct {
	Key    pulumi.StringPtrInput
	Locked pulumi.BoolPtrInput
	Value  pulumi.StringPtrInput
}

func (EnvironmentVariableState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentVariableState)(nil)).Elem()
}

type environmentVariableArgs struct {
	Key    string `pulumi:"key"`
	Locked *bool  `pulumi:"locked"`
	Value  string `pulumi:"value"`
}

// The set of arguments for constructing a EnvironmentVariable resource.
type EnvironmentVariableArgs struct {
	Key    pulumi.StringInput
	Locked pulumi.BoolPtrInput
	Value  pulumi.StringInput
}

func (EnvironmentVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentVariableArgs)(nil)).Elem()
}

type EnvironmentVariableInput interface {
	pulumi.Input

	ToEnvironmentVariableOutput() EnvironmentVariableOutput
	ToEnvironmentVariableOutputWithContext(ctx context.Context) EnvironmentVariableOutput
}

func (*EnvironmentVariable) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentVariable)(nil)).Elem()
}

func (i *EnvironmentVariable) ToEnvironmentVariableOutput() EnvironmentVariableOutput {
	return i.ToEnvironmentVariableOutputWithContext(context.Background())
}

func (i *EnvironmentVariable) ToEnvironmentVariableOutputWithContext(ctx context.Context) EnvironmentVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVariableOutput)
}

// EnvironmentVariableArrayInput is an input type that accepts EnvironmentVariableArray and EnvironmentVariableArrayOutput values.
// You can construct a concrete instance of `EnvironmentVariableArrayInput` via:
//
//          EnvironmentVariableArray{ EnvironmentVariableArgs{...} }
type EnvironmentVariableArrayInput interface {
	pulumi.Input

	ToEnvironmentVariableArrayOutput() EnvironmentVariableArrayOutput
	ToEnvironmentVariableArrayOutputWithContext(context.Context) EnvironmentVariableArrayOutput
}

type EnvironmentVariableArray []EnvironmentVariableInput

func (EnvironmentVariableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnvironmentVariable)(nil)).Elem()
}

func (i EnvironmentVariableArray) ToEnvironmentVariableArrayOutput() EnvironmentVariableArrayOutput {
	return i.ToEnvironmentVariableArrayOutputWithContext(context.Background())
}

func (i EnvironmentVariableArray) ToEnvironmentVariableArrayOutputWithContext(ctx context.Context) EnvironmentVariableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVariableArrayOutput)
}

// EnvironmentVariableMapInput is an input type that accepts EnvironmentVariableMap and EnvironmentVariableMapOutput values.
// You can construct a concrete instance of `EnvironmentVariableMapInput` via:
//
//          EnvironmentVariableMap{ "key": EnvironmentVariableArgs{...} }
type EnvironmentVariableMapInput interface {
	pulumi.Input

	ToEnvironmentVariableMapOutput() EnvironmentVariableMapOutput
	ToEnvironmentVariableMapOutputWithContext(context.Context) EnvironmentVariableMapOutput
}

type EnvironmentVariableMap map[string]EnvironmentVariableInput

func (EnvironmentVariableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnvironmentVariable)(nil)).Elem()
}

func (i EnvironmentVariableMap) ToEnvironmentVariableMapOutput() EnvironmentVariableMapOutput {
	return i.ToEnvironmentVariableMapOutputWithContext(context.Background())
}

func (i EnvironmentVariableMap) ToEnvironmentVariableMapOutputWithContext(ctx context.Context) EnvironmentVariableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVariableMapOutput)
}

type EnvironmentVariableOutput struct{ *pulumi.OutputState }

func (EnvironmentVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentVariable)(nil)).Elem()
}

func (o EnvironmentVariableOutput) ToEnvironmentVariableOutput() EnvironmentVariableOutput {
	return o
}

func (o EnvironmentVariableOutput) ToEnvironmentVariableOutputWithContext(ctx context.Context) EnvironmentVariableOutput {
	return o
}

type EnvironmentVariableArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentVariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EnvironmentVariable)(nil)).Elem()
}

func (o EnvironmentVariableArrayOutput) ToEnvironmentVariableArrayOutput() EnvironmentVariableArrayOutput {
	return o
}

func (o EnvironmentVariableArrayOutput) ToEnvironmentVariableArrayOutputWithContext(ctx context.Context) EnvironmentVariableArrayOutput {
	return o
}

func (o EnvironmentVariableArrayOutput) Index(i pulumi.IntInput) EnvironmentVariableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EnvironmentVariable {
		return vs[0].([]*EnvironmentVariable)[vs[1].(int)]
	}).(EnvironmentVariableOutput)
}

type EnvironmentVariableMapOutput struct{ *pulumi.OutputState }

func (EnvironmentVariableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EnvironmentVariable)(nil)).Elem()
}

func (o EnvironmentVariableMapOutput) ToEnvironmentVariableMapOutput() EnvironmentVariableMapOutput {
	return o
}

func (o EnvironmentVariableMapOutput) ToEnvironmentVariableMapOutputWithContext(ctx context.Context) EnvironmentVariableMapOutput {
	return o
}

func (o EnvironmentVariableMapOutput) MapIndex(k pulumi.StringInput) EnvironmentVariableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EnvironmentVariable {
		return vs[0].(map[string]*EnvironmentVariable)[vs[1].(string)]
	}).(EnvironmentVariableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentVariableInput)(nil)).Elem(), &EnvironmentVariable{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentVariableArrayInput)(nil)).Elem(), EnvironmentVariableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentVariableMapInput)(nil)).Elem(), EnvironmentVariableMap{})
	pulumi.RegisterOutputType(EnvironmentVariableOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableMapOutput{})
}
