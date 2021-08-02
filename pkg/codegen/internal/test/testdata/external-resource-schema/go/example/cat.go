// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Cat struct {
	pulumi.CustomResourceState

	Name pulumi.StringPtrOutput `pulumi:"name"`
}

// NewCat registers a new resource with the given unique name, arguments, and options.
func NewCat(ctx *pulumi.Context,
	name string, args *CatArgs, opts ...pulumi.ResourceOption) (*Cat, error) {
	if args == nil {
		args = &CatArgs{}
	}

	var resource Cat
	err := ctx.RegisterResource("example::Cat", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCat gets an existing Cat resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCat(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CatState, opts ...pulumi.ResourceOption) (*Cat, error) {
	var resource Cat
	err := ctx.ReadResource("example::Cat", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cat resources.
type catState struct {
}

type CatState struct {
}

func (CatState) ElementType() reflect.Type {
	return reflect.TypeOf((*catState)(nil)).Elem()
}

type catArgs struct {
	Age *int `pulumi:"age"`
	Pet *Pet `pulumi:"pet"`
}

// The set of arguments for constructing a Cat resource.
type CatArgs struct {
	Age pulumi.IntPtrInput
	Pet PetPtrInput
}

func (CatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*catArgs)(nil)).Elem()
}

type CatInput interface {
	pulumi.Input

	ToCatOutput() CatOutput
	ToCatOutputWithContext(ctx context.Context) CatOutput
}

func (*Cat) ElementType() reflect.Type {
	return reflect.TypeOf((*Cat)(nil))
}

func (i *Cat) ToCatOutput() CatOutput {
	return i.ToCatOutputWithContext(context.Background())
}

func (i *Cat) ToCatOutputWithContext(ctx context.Context) CatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatOutput)
}

func (i *Cat) ToCatPtrOutput() CatPtrOutput {
	return i.ToCatPtrOutputWithContext(context.Background())
}

func (i *Cat) ToCatPtrOutputWithContext(ctx context.Context) CatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatPtrOutput)
}

type CatPtrInput interface {
	pulumi.Input

	ToCatPtrOutput() CatPtrOutput
	ToCatPtrOutputWithContext(ctx context.Context) CatPtrOutput
}

type catPtrType CatArgs

func (*catPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Cat)(nil))
}

func (i *catPtrType) ToCatPtrOutput() CatPtrOutput {
	return i.ToCatPtrOutputWithContext(context.Background())
}

func (i *catPtrType) ToCatPtrOutputWithContext(ctx context.Context) CatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatPtrOutput)
}

// CatArrayInput is an input type that accepts CatArray and CatArrayOutput values.
// You can construct a concrete instance of `CatArrayInput` via:
//
//          CatArray{ CatArgs{...} }
type CatArrayInput interface {
	pulumi.Input

	ToCatArrayOutput() CatArrayOutput
	ToCatArrayOutputWithContext(context.Context) CatArrayOutput
}

type CatArray []CatInput

func (CatArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cat)(nil)).Elem()
}

func (i CatArray) ToCatArrayOutput() CatArrayOutput {
	return i.ToCatArrayOutputWithContext(context.Background())
}

func (i CatArray) ToCatArrayOutputWithContext(ctx context.Context) CatArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatArrayOutput)
}

// CatMapInput is an input type that accepts CatMap and CatMapOutput values.
// You can construct a concrete instance of `CatMapInput` via:
//
//          CatMap{ "key": CatArgs{...} }
type CatMapInput interface {
	pulumi.Input

	ToCatMapOutput() CatMapOutput
	ToCatMapOutputWithContext(context.Context) CatMapOutput
}

type CatMap map[string]CatInput

func (CatMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cat)(nil)).Elem()
}

func (i CatMap) ToCatMapOutput() CatMapOutput {
	return i.ToCatMapOutputWithContext(context.Background())
}

func (i CatMap) ToCatMapOutputWithContext(ctx context.Context) CatMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CatMapOutput)
}

type CatOutput struct{ *pulumi.OutputState }

func (CatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Cat)(nil))
}

func (o CatOutput) ToCatOutput() CatOutput {
	return o
}

func (o CatOutput) ToCatOutputWithContext(ctx context.Context) CatOutput {
	return o
}

func (o CatOutput) ToCatPtrOutput() CatPtrOutput {
	return o.ToCatPtrOutputWithContext(context.Background())
}

func (o CatOutput) ToCatPtrOutputWithContext(ctx context.Context) CatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Cat) *Cat {
		return &v
	}).(CatPtrOutput)
}

type CatPtrOutput struct{ *pulumi.OutputState }

func (CatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cat)(nil))
}

func (o CatPtrOutput) ToCatPtrOutput() CatPtrOutput {
	return o
}

func (o CatPtrOutput) ToCatPtrOutputWithContext(ctx context.Context) CatPtrOutput {
	return o
}

func (o CatPtrOutput) Elem() CatOutput {
	return o.ApplyT(func(v *Cat) Cat {
		if v != nil {
			return *v
		}
		var ret Cat
		return ret
	}).(CatOutput)
}

type CatArrayOutput struct{ *pulumi.OutputState }

func (CatArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Cat)(nil))
}

func (o CatArrayOutput) ToCatArrayOutput() CatArrayOutput {
	return o
}

func (o CatArrayOutput) ToCatArrayOutputWithContext(ctx context.Context) CatArrayOutput {
	return o
}

func (o CatArrayOutput) Index(i pulumi.IntInput) CatOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Cat {
		return vs[0].([]Cat)[vs[1].(int)]
	}).(CatOutput)
}

type CatMapOutput struct{ *pulumi.OutputState }

func (CatMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Cat)(nil))
}

func (o CatMapOutput) ToCatMapOutput() CatMapOutput {
	return o
}

func (o CatMapOutput) ToCatMapOutputWithContext(ctx context.Context) CatMapOutput {
	return o
}

func (o CatMapOutput) MapIndex(k pulumi.StringInput) CatOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Cat {
		return vs[0].(map[string]Cat)[vs[1].(string)]
	}).(CatOutput)
}

func init() {
	pulumi.RegisterOutputType(CatOutput{})
	pulumi.RegisterOutputType(CatPtrOutput{})
	pulumi.RegisterOutputType(CatArrayOutput{})
	pulumi.RegisterOutputType(CatMapOutput{})
}
