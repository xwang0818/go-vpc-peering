package main

import (
    "github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
    "github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        // VPC
        xwVpc, err := ec2.NewVpc(ctx, "xwVpc", &ec2.VpcArgs{
            CidrBlock: pulumi.String("10.223.0.0/16"),
            Tags: pulumi.StringMap{
                "Name": pulumi.String("xw-vpc-pulumi-operater-test"),
            },
        })
        if err != nil {
            return err
        }
        // Subnet
        _, err = ec2.NewSubnet(ctx, "xwSubnet", &ec2.SubnetArgs{
            VpcId:            xwVpc.ID(),
            CidrBlock:        pulumi.String("10.222.1.0/24"),
            AvailabilityZone: pulumi.String("us-east-2a"),
            Tags: pulumi.StringMap{
                "Name": pulumi.String("xw-subnet-pulumi-operater-test"),
            },
        })
        if err != nil {
            return err
        }
        return nil
    })
}
