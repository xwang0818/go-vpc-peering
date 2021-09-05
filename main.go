package main

import (
    "github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
    "github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
    		// VPC
    		_, err := ec2.NewVpc(ctx, "xwvpc", &ec2.VpcArgs{
    			CidrBlock: pulumi.String("10.125.0.0/16"),
    		})
        if err != nil {
            return err
        }
        return nil
    })
}
