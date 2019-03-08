package v3

import (
	"context"
	"fmt"
	"net/http"
	//"reflect"

	"github.com/kumarsarath588/karbon/client"
	mainv3 "github.com/terraform-providers/terraform-provider-nutanix/client/v3"
	//"github.com/terraform-providers/terraform-provider-nutanix/utils"
)

const (
	getCookieUrl = "https://%s:9440/api/nutanix/v3/clusters/list"
	taskStatusUrl = "https://%s:9440/api/nutanix/v3/tasks/%s"
)

// Operations ...
type Operations struct {
	client *client.Client
}

// Service ...
type Service interface {
	CreateKarbonCluster(createRequest *KarbonClusterIntentInput) (*KarbonClusterCreateDeleteResponse, error)
	DeleteKarbonCluster(UUID string) (*KarbonClusterCreateDeleteResponse, error)
	GetKarbonCluster(UUID string) (*KarbonClusterIntentResponse, error)
	GetKarbonClusterKubeconfig(UUID string) (*KarbonClusterKubeconfigIntentResponse, error)
	ListKarbonClusters(getEntitiesRequest *mainv3.DSMetadata) (*KarbonClusterListIntentResponse, error)
	TaskStatus(UUID string) (*mainv3.TasksResponse, error)
}

func (op Operations) DeleteKarbonCluster(UUID string) (*KarbonClusterCreateDeleteResponse, error) {
	ctx := context.TODO()
	path := "/cluster/" + UUID

	gurl := fmt.Sprintf(getCookieUrl, op.client.Credentials.Endpoint)
	body := &mainv3.DSMetadata{}
	cookies, err := op.client.GetCookies(ctx, http.MethodPost, gurl, body)
	if err != nil {
		return nil, err
	}

	req, err := op.client.NewRequest(ctx, http.MethodDelete, path, cookies, nil)
	karbonClusterDeleteResponse := new(KarbonClusterCreateDeleteResponse)
	if err != nil {
		return nil, err
	}
	return karbonClusterDeleteResponse, op.client.Do(ctx, req, karbonClusterDeleteResponse)
}

func (op Operations) TaskStatus(UUID string) (*mainv3.TasksResponse, error) {
	ctx := context.TODO()
	tsurl := fmt.Sprintf(taskStatusUrl, op.client.Credentials.Endpoint, UUID)
	taskStatus := new(mainv3.TasksResponse)
	err := op.client.GetTaskStatus(ctx, http.MethodGet, tsurl, taskStatus)
	if err != nil {
		return nil, err
	}
	return taskStatus, nil
}

func (op Operations) CreateKarbonCluster(createRequest *KarbonClusterIntentInput) (*KarbonClusterCreateDeleteResponse, error) {
	ctx := context.TODO()
	path := "/cluster"

	gurl := fmt.Sprintf(getCookieUrl, op.client.Credentials.Endpoint)
	body := &mainv3.DSMetadata{}
	cookies, err := op.client.GetCookies(ctx, http.MethodPost, gurl, body)
	if err != nil {
		return nil, err
	}

	req, err := op.client.NewRequest(ctx, http.MethodPost, path, cookies, createRequest)
	karbonClusterCreateResponse := new(KarbonClusterCreateDeleteResponse)
	if err != nil {
		return nil, err
	}
	return karbonClusterCreateResponse, op.client.Do(ctx, req, karbonClusterCreateResponse)
}
/*ListKarbonClusters Get a list of Kubenetes clusters created in PC
 * @param getEntitiesRequest {} body @return *
 */
func (op Operations) ListKarbonClusters(getEntitiesRequest *mainv3.DSMetadata) (*KarbonClusterListIntentResponse, error) {
	ctx := context.TODO()
	path := "/cluster/list"

	gurl := fmt.Sprintf(getCookieUrl, op.client.Credentials.Endpoint)
	body := &mainv3.DSMetadata{}
	cookies, err := op.client.GetCookies(ctx, http.MethodPost, gurl, body)
	if err != nil {
		return nil, err
	}

	req, err := op.client.NewRequest(ctx, http.MethodPost, path, cookies, getEntitiesRequest)
	karbonClusterListIntentResponse := new(KarbonClusterListIntentResponse)
	if err != nil {
		return nil, err
	}
	return karbonClusterListIntentResponse, op.client.Do(ctx, req, karbonClusterListIntentResponse)
}

//GetKarbonCluster Get a list of Kubenetes clusters created in PC
func (op Operations) GetKarbonCluster(UUID string) (*KarbonClusterIntentResponse, error) {
	ctx := context.TODO()
	path := "/cluster/" + UUID

	gurl := fmt.Sprintf(getCookieUrl, op.client.Credentials.Endpoint)
	body := &mainv3.DSMetadata{}
	cookies, err := op.client.GetCookies(ctx, http.MethodPost, gurl, body)
	if err != nil {
		return nil, err
	}
	req, err := op.client.NewRequest(ctx, http.MethodGet, path, cookies, nil)
	karbonClusterIntentResponse:= new(KarbonClusterIntentResponse)
	if err != nil {
		return nil, err
	}
	return karbonClusterIntentResponse, op.client.Do(ctx, req, karbonClusterIntentResponse)
}
//GetKarbonClusterKubeconfig
func (op Operations) GetKarbonClusterKubeconfig(UUID string) (*KarbonClusterKubeconfigIntentResponse, error) {
	ctx := context.TODO()
	path := "/cluster/" + UUID + "/kubeconfig"

	gurl := fmt.Sprintf(getCookieUrl, op.client.Credentials.Endpoint)
	body := &mainv3.DSMetadata{}
	cookies, err := op.client.GetCookies(ctx, http.MethodPost, gurl, body)
	if err != nil {
		return nil, err
	}
	req, err := op.client.NewRequest(ctx, http.MethodGet, path, cookies, nil)
	karbonClusterKubeconfigIntentResponse := new(KarbonClusterKubeconfigIntentResponse)
	if err != nil {
		return nil, err
	}
	return karbonClusterKubeconfigIntentResponse, op.client.Do(ctx, req,karbonClusterKubeconfigIntentResponse)
}
