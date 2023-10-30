package listBuckets

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var client *s3.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration err, " + err.Error())
	}

	client = s3.NewFromConfig(cfg)
}

func ListBuckets() ([]*string, error) {
	butckets := make([]*string, 0, 10)
	params := &s3.ListBucketsInput{}

	res, err := client.ListBuckets(context.TODO(), params)
	if err != nil {
		return nil, err
	}

	for _, bucket := range res.Buckets {
		butckets = append(butckets, bucket.Name)
	}

	return butckets, nil
}
