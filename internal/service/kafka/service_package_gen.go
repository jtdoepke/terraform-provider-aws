// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package kafka

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	kafka_sdkv2 "github.com/aws/aws-sdk-go-v2/service/kafka"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	kafka_sdkv1 "github.com/aws/aws-sdk-go/service/kafka"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceBrokerNodes,
			TypeName: "aws_msk_broker_nodes",
		},
		{
			Factory:  DataSourceCluster,
			TypeName: "aws_msk_cluster",
		},
		{
			Factory:  DataSourceConfiguration,
			TypeName: "aws_msk_configuration",
		},
		{
			Factory:  DataSourceVersion,
			TypeName: "aws_msk_kafka_version",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceCluster,
			TypeName: "aws_msk_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceConfiguration,
			TypeName: "aws_msk_configuration",
		},
		{
			Factory:  ResourceScramSecretAssociation,
			TypeName: "aws_msk_scram_secret_association",
		},
		{
			Factory:  ResourceServerlessCluster,
			TypeName: "aws_msk_serverless_cluster",
			Name:     "Serverless Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceVPCConnection,
			TypeName: "aws_msk_vpc_connection",
			Name:     "Vpc Connection",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Kafka
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*kafka_sdkv1.Kafka, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return kafka_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*kafka_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return kafka_sdkv2.NewFromConfig(cfg, func(o *kafka_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.EndpointResolver = kafka_sdkv2.EndpointResolverFromURL(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
