package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/objectstorage/v1/containers"
	"github.com/gophercloud/gophercloud/openstack/objectstorage/v1/objects"
)

func createContainer(client *gophercloud.ServiceClient, name string, versionsLocation string) {

	fmt.Printf("Creating container: %s\n", name)

	createOpts := containers.CreateOpts{
		StoragePolicy: "nz-hlz-1--o1--sr-r3",
		// ContentType:   "application/json",
	}

	if versionsLocation != "" {
		createOpts.VersionsLocation = versionsLocation
	}

	_, err := containers.Create(client, name, createOpts).Extract()
	if err != nil {
		panic(err)
	}
}

func createObject(client *gophercloud.ServiceClient, container string, name string, content string, ifNoneMatch string) {

	fmt.Printf("Creating object: %s | %s\n", container, name)

	createOpts := objects.CreateOpts{
		ContentType: "text/plain",
		Content:     strings.NewReader(content),
		IfNoneMatch: ifNoneMatch,
		DeleteAfter: 60,
	}

	_, err := objects.Create(client, container, name, createOpts).Extract()
	if err != nil {
		fmt.Printf("%s\n", err)
	}

}

func deleteObject(client *gophercloud.ServiceClient, container string, name string) {

	fmt.Printf("Deleting object: %s | %s\n", container, name)

	deleteOpts := objects.DeleteOpts{}
	objects.Delete(client, container, name, deleteOpts)
}

func showObject(client *gophercloud.ServiceClient, container string, name string) {

	fmt.Printf("Showing object: %s | %s\n", container, name)

	downloadOpts := objects.DownloadOpts{
		Newest: true,
	}

	object := objects.Download(client, container, name, downloadOpts)
	if object.Err != nil {
		fmt.Printf("%s\n", object.Err)
	}

	// if "ExtractContent" method is not called, the HTTP connection will remain consumed
	content, err := object.ExtractContent()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%s\n", content)

}

func listObjects(client *gophercloud.ServiceClient, container string) {

	fmt.Printf("Listing objects in container: %s\n", container)

	allPages, err := objects.List(client, container, objects.ListOpts{}).AllPages()
	if err != nil {
		fmt.Printf("%s", err)
	}

	allObjects, err := objects.ExtractNames(allPages)
	if err != nil {
		fmt.Printf("%s", err)
	}

	for _, object := range allObjects {
		fmt.Printf("%+v\n", object)
	}

}

func listContainers(client *gophercloud.ServiceClient) {

	fmt.Printf("Listing containers\n")

	allPages, _ := containers.List(client, containers.ListOpts{}).AllPages()
	allContainers, _ := containers.ExtractInfo(allPages)

	for _, container := range allContainers {
		fmt.Printf("%+v\n", container)
	}
}

func showContainer(client *gophercloud.ServiceClient, containerName string) {

	fmt.Printf("Showing container %s\n", containerName)

	container := containers.Get(client, containerName, containers.GetOpts{})
	// jsonData, _ := json.MarshalIndent(container, " ", "")
	fmt.Printf("%s\n", container)
}

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

	client, _ := openstack.NewObjectStorageV1(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})

	// container := "test-container-2"
	container := "swift-bug-test"
	// containerArchive := container + "-archive"
	object := "test-object-8"

	// Create versioned container and container archive
	// createContainer(client, containerArchive, "")
	// showContainer(client, containerArchive)

	// createContainer(client, container, containerArchive)
	showContainer(client, container)

	// Create test object
	createObject(client, container, object, "foo\n", "*")
	showObject(client, container, object)

	// Create test object
	// createObject(client, container, object, "bar\n", "*")
	// showObject(client, container, object)

	// List containers
	// listObjects(client, container)
	// listObjects(client, containerArchive)

	// 	// Create new version of test object
	// 	createObject(client, container, object, "baz", "")
	// 	showObject(client, container, object)

	// 	// List containers
	// 	listObjects(client, container)
	// 	listObjects(client, containerArchive)

	// Delete test object
	// deleteObject(client, container, object)
	// fmt.Println("Waiting")
	// time.Sleep(5 * time.Second)

	// // List containers
	// listObjects(client, container)
	// listObjects(client, containerArchive)

	// createObject(client, container, object, "baz\n", "*")
	// showObject(client, container, object)

	// c := 0
	// for c < 1000 {
	// 	c += 1
	// 	showObject(client, container, object)
	// 	// List containers
	// 	listObjects(client, container)
	// 	// listObjects(client, containerArchive)
	// 	time.Sleep(2 * time.Second)
	// }

	// 	// Delete test object
	// 	deleteObject(client, container, object)

	// 	// List containers
	// 	listObjects(client, container)
	// 	listObjects(client, containerArchive)

	// 	// Create test object
	// 	createObject(client, container, object, "foo", "*")

	// 	// List containers
	// 	listObjects(client, container)
	// 	listObjects(client, containerArchive)
}
