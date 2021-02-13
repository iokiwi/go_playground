package main

import (
	"fmt"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/containerinfra/v1/clusters"
)

func main() {

	opts := gophercloud.AuthOptions{
		IdentityEndpoint: os.Getenv("OS_AUTH_URL"),
		TenantID:         os.Getenv("OS_PROJECT_ID"),
		TokenID:          os.Getenv("OS_TOKEN"),
		// DomainID: "default",
		// Username: os.Getenv("OS_USERNAME"),
		// Password: os.Getenv("OS_PASSWORD"),
	}
	provider, _ := openstack.AuthenticatedClient(opts)

	client, _ := openstack.NewContainerInfraV1(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})

	// masterCount := 1
	// nodeCount := 1
	// createTimeout := 30
	clusterName := "simon-go-test-5"
	boolTrue := true
	// boolFalse := false

	createOpts := clusters.CreateOpts{
		ClusterTemplateID: "24f658cf-6c27-409b-b736-47a9516cac0d",
		// CreateTimeout:     &createTimeout,
		// DiscoveryURL:      "",
		// FlavorID:          "mc2.small",
		// KeyPair:           "my_keypair",
		// Labels:            map[string]string{},
		// MasterCount:       &masterCount,
		// MasterFlavorID:    "m1.small",
		Name:            clusterName,
		MasterLBEnabled: &boolTrue,
		// NodeCount:         &nodeCount,
	}

	_, err := clusters.Create(client, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	cluster, err := clusters.Get(client, clusterName).Extract()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", cluster)

}
