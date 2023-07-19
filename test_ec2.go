package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.Background(),
		// Specify the shared configuration profile to load.
		config.WithSharedConfigProfile(os.Args[1]),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Region loaded from credentials file.
	fmt.Println("Region:", cfg.Region)

	svc := ec2.New(ec2.Options{
		Credentials: cfg.Credentials,
		Region:      cfg.Region,
	})
	input := &ec2.DescribeInstancesInput{}
	result, err := svc.DescribeInstances(context.Background(), input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("instances: %+v\n", result)
}
