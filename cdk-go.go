package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type CdkGoStackProps struct {
	awscdk.StackProps
}

func NewCdkGoStack(scope constructs.Construct, id string, props *CdkGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	var publicsub = []awsec2.SubnetConfiguration{
		Name: 
	}
	var privatesub[]awsec2.SubnetConfiguration 

	awsec2.NewVpc(stack, jsii.String("MyVpc"), &awsec2.VpcProps{
		Cidr:   jsii.String("10.0.0.0/16"),
		MaxAzs: jsii.Number(2),
		
		SubnetConfiguration: 
		[
				{
					Name: jsii.String("PublicSubnet"),
					
				}
		]
		
	})

	// awsec2.NewPublicSubnet(stack, jsii.String("PublicSubnet"), &awsec2.PublicSubnetProps{
	// 	AvailabilityZone: jsii.String("us-east-1a"),
	// 	CidrBlock:        jsii.String("10.0.1.0/24"),
	// 	VpcId:            vpc.VpcId(),
	// })

	// awsec2.NewPrivateSubnet(stack, jsii.String("PrivateSubnet"), &awsec2.PrivateSubnetProps{
	// 	AvailabilityZone: jsii.String("us-east-1b"),
	// 	CidrBlock:        jsii.String("10.0.2.0/24"),
	// 	VpcId:            vpc.VpcId(),
	// })

	// awsec2.NewInstance(stack, jsii.String("NatInstance"), &awsec2.InstanceProps{
	// 	InstanceType: awsec2.NewInstanceType(jsii.String("t2.nano")),
	// 	MachineImage: awsec2.NewAmazonLinuxImage(nil),
	// 	Vpc:          vpc,
	// 	VpcSubnets:   &awsec2.SubnetSelection{SubnetType: "PUBLIC"},
	// })

	// awsec2.NewInstance(stack, jsii.String("PrivateInstance"), &awsec2.InstanceProps{
	// 	InstanceType: awsec2.NewInstanceType(jsii.String("t2.nano")),
	// 	VpcSubnets: ,
	// })

	// awsec2.NewCfnRouteTable(stack,jsii.String("RT"), &awsec2.CfnRouteTableProps{
	// 	VpcId: vpc.VpcId(),
	// 	Tags:  ,
	// 	Vpc: vpc,

	// })

	// awsec2.NewCfnInternetGateway(stack, jsii.String("IGW"), &awsec2.CfnInternetGatewayProps{
	// 	Tags: ,
	// })

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewCdkGoStack(app, "CdkGoStack", &CdkGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
