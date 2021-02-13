package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/containerinfra/v1/clustertemplates"
)

func main() {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: os.Getenv("OS_AUTH_URL"),
		TokenID:          os.Getenv("OS_TOKEN"),
		TenantID:         os.Getenv("OS_PROJECT_ID"),
	}
	provider, _ := openstack.AuthenticatedClient(opts)

	client, _ := openstack.NewContainerInfraV1(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})

	labelMap := make(map[string]string)
	labelMap["auto_healing_controller"] = "magnum-auto-healer"
	labelMap["auto_healing_enabled"] = "true"
	labelMap["auto_scaling_enabled"] = "false"
	labelMap["calico_ipv4pool"] = "10.100.0.0/16"
	labelMap["calico_tag"] = "v3.13.4"
	labelMap["cinder_csi_enabled"] = "true"
	labelMap["cinder_csi_plugin_tag"] = "v1.19.0"
	labelMap["cloud_provider_enabled"] = "true"
	labelMap["cloud_provider_tag"] = "v1.18.0-catalyst"
	labelMap["container_infra_prefix"] = "docker.io/catalystcloud/"
	labelMap["coredns_tag"] = "1.6.6"
	labelMap["etcd_tag"] = "v3.3.20"
	labelMap["etcd_volume_size"] = "20"
	labelMap["heapster_enabled"] = "false"
	labelMap["heat_container_agent_tag"] = "train-stable"
	labelMap["ingress_controller"] = "octavia"
	labelMap["k8s_keystone_auth_tag"] = "v1.18.0"
	labelMap["keystone_auth_enabled"] = "true"
	labelMap["kube_dashboard_enabled"] = "true"
	labelMap["kube_dashboard_version"] = "v2.0.0"
	labelMap["kube_image_digest"] = "sha256:98e3d3d634e2e347fe5e45259e8c3a7c2e0db7372a94ce63a8d041270f25403f"
	labelMap["kube_tag"] = "v1.19.4"
	labelMap["magnum_auto_healer_tag"] = "v1.19.0-cpo.2020.5.25"
	labelMap["master_lb_floating_ip_enabled"] = "false"
	labelMap["monitoring_enabled"] = "true"
	labelMap["octavia_ingress_controller_tag"] = "v1.18.0-catalyst"
	labelMap["ostree_commit"] = "b51037798e93e5aae5123633fb596c80ddf30302b5110b0581900dbc5b2f0d24"
	labelMap["prometheus_adapter_enabled"] = "false"
	labelMap["prometheus_operator_chart_tag"] = "v8.2.2"
	labelMap["tiller_enabled"] = "true"
	labelMap["use_podman"] = "true"

	dockerVolumeSize := 20
	boolTrue := true
	boolFalse := false

	createOpts := clustertemplates.CreateOpts{
		Name: "go-test-cluster-template",
		Labels: labelMap,
		Public: &boolFalse,
		MasterFlavorID: "c1.c2r4",
		FloatingIPEnabled: &boolFalse,
		DockerVolumeSize: &dockerVolumeSize,
		TLSDisabled: &boolFalse,
		ServerType: "vm",
		ExternalNetworkID: "f10ad6de-a26d-4c29-8c64-2a7418d47f8f",
		ImageID: "a1a77ccb-a7e5-4773-8255-e92cfd4a5271",
		VolumeDriver: "cinder",
		RegistryEnabled: &boolFalse,
		DockerStorageDriver: "overlay2",
		NetworkDriver: "calico",
		COE: "kubernetes",
		FlavorID: "c1.c4r8",
		MasterLBEnabled: &boolTrue,
		DNSNameServer: "202.78.244.85,202.78.244.86,202.78.244.87",
	}

	clustertemplates.Create(client, createOpts).Extract()

	allPages, _ := clustertemplates.List(client, clustertemplates.ListOpts{}).AllPages()
	allTemplates, _ := clustertemplates.ExtractClusterTemplates(allPages)
	for _, template := range allTemplates {
		templateJSON, _ := json.MarshalIndent(template, "", " ")
		fmt.Printf("%s\n", templateJSON)
	}

}
