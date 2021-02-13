package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
)

func main() {
	// opts, _ := openstack.AuthOptionsFromEnv()

	opts := gophercloud.AuthOptions{
		IdentityEndpoint: os.Getenv("OS_AUTH_URL"),
		TokenID:          os.Getenv("OS_TOKEN"),
		TenantID:         os.Getenv("OS_PROJECT_ID"),
	}
	provider, _ := openstack.AuthenticatedClient(opts)
	client, err := openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if err != nil {
		fmt.Printf("%s", err)
	}
	allPages, _ := images.List(client, images.ListOpts{}).AllPages()
	allImages, _ := images.ExtractImages(allPages)
	for _, image := range allImages {
		imageJSON, _ := json.MarshalIndent(image, "", " ")
		fmt.Printf("%s\n", imageJSON)
	}
}
