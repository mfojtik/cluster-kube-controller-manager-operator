package operator

import (
	"fmt"
	"os"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	appsclientv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	coreclientv1 "k8s.io/client-go/kubernetes/typed/core/v1"

	operatorsv1alpha1 "github.com/openshift/api/operator/v1alpha1"
	"github.com/openshift/cluster-kube-controller-manager-operator/pkg/apis/kubecontrollermanager/v1alpha1"
	"github.com/openshift/cluster-kube-controller-manager-operator/pkg/operator/v311_00_assets"
	"github.com/openshift/library-go/pkg/operator/resource/resourceapply"
	"github.com/openshift/library-go/pkg/operator/resource/resourcemerge"
	"github.com/openshift/library-go/pkg/operator/resource/resourceread"
)

// syncKubeControllerManager_v311_00_to_latest takes care of synchronizing (not upgrading) the thing we're managing.
// most of the time the sync method will be good for a large span of minor versions
func syncKubeControllerManager_v311_00_to_latest(c KubeControllerManagerOperator, operatorConfig *v1alpha1.KubeControllerManagerOperatorConfig, previousAvailability *operatorsv1alpha1.VersionAvailablity) (operatorsv1alpha1.VersionAvailablity, []error) {
	versionAvailability := operatorsv1alpha1.VersionAvailablity{
		Version: operatorConfig.Spec.Version,
	}

	errors := []error{}
	var err error

	requiredNamespace := resourceread.ReadNamespaceV1OrDie(v311_00_assets.MustAsset("v3.11.0/kube-controller-manager/ns.yaml"))
	_, _, err = resourceapply.ApplyNamespace(c.corev1Client, requiredNamespace)
	if err != nil {
		errors = append(errors, fmt.Errorf("%q: %v", "ns", err))
	}

	requiredPublicRole := resourceread.ReadRoleV1OrDie(v311_00_assets.MustAsset("v3.11.0/kube-controller-manager/public-info-role.yaml"))
	_, _, err = resourceapply.ApplyRole(c.rbacv1Client, requiredPublicRole)
	if err != nil {
		errors = append(errors, fmt.Errorf("%q: %v", "svc", err))
	}

	requiredPublicRoleBinding := resourceread.ReadRoleBindingV1OrDie(v311_00_assets.MustAsset("v3.11.0/kube-controller-manager/public-info-rolebinding.yaml"))
	_, _, err = resourceapply.ApplyRoleBinding(c.rbacv1Client, requiredPublicRoleBinding)
	if err != nil {
		errors = append(errors, fmt.Errorf("%q: %v", "svc", err))
	}

	requiredService := resourceread.ReadServiceV1OrDie(v311_00_assets.MustAsset("v3.11.0/kube-controller-manager/svc.yaml"))
	_, _, err = resourceapply.ApplyService(c.corev1Client, requiredService)
	if err != nil {
		errors = append(errors, fmt.Errorf("%q: %v", "svc", err))
	}

	requiredSA := resourceread.ReadServiceAccountV1OrDie(v311_00_assets.MustAsset("v3.11.0/kube-controller-manager/sa.yaml"))
	_, saModified, err := resourceapply.ApplyServiceAccount(c.corev1Client, requiredSA)
	if err != nil {
		errors = append(errors, fmt.Errorf("%q: %v", "sa", err))
	}

	controllerManagerConfig, configMapModified, err := manageKubeControllerManagerConfigMap_v311_00_to_latest(c.corev1Client, operatorConfig)
	if err != nil {
		errors = append(errors, fmt.Errorf("%q: %v", "configmap", err))
	}

	forceDeployment := operatorConfig.ObjectMeta.Generation != operatorConfig.Status.ObservedGeneration
	if saModified { // SA modification can cause new tokens
		forceDeployment = true
	}
	if configMapModified {
		forceDeployment = true
	}

	// our configmaps and secrets are in order, now it is time to create the DS
	// TODO check basic preconditions here
	actualDeployment, _, err := manageKubeControllerManagerDeployment_v311_00_to_latest(c.appsv1Client, operatorConfig, previousAvailability, forceDeployment)
	if err != nil {
		errors = append(errors, fmt.Errorf("%q: %v", "deployment", err))
	}

	configData := ""
	if controllerManagerConfig != nil {
		configData = controllerManagerConfig.Data["config.yaml"]
	}
	_, _, err = manageKubeControllerManagerPublicConfigMap_v311_00_to_latest(c.corev1Client, configData, operatorConfig)
	if err != nil {
		errors = append(errors, fmt.Errorf("%q: %v", "configmap/public-info", err))
	}

	return resourcemerge.ApplyGenerationAvailability(versionAvailability, actualDeployment, errors...), errors
}

func manageKubeControllerManagerConfigMap_v311_00_to_latest(client coreclientv1.ConfigMapsGetter, operatorConfig *v1alpha1.KubeControllerManagerOperatorConfig) (*corev1.ConfigMap, bool, error) {
	configMap := resourceread.ReadConfigMapV1OrDie(v311_00_assets.MustAsset("v3.11.0/kube-controller-manager/cm.yaml"))
	defaultConfig := v311_00_assets.MustAsset("v3.11.0/kube-controller-manager/defaultconfig.yaml")
	requiredConfigMap, _, err := resourcemerge.MergeConfigMap(configMap, "config.yaml", nil, defaultConfig, operatorConfig.Spec.KubeControllerManagerConfig.Raw)
	if err != nil {
		return nil, false, err
	}
	return resourceapply.ApplyConfigMap(client, requiredConfigMap)
}

func manageKubeControllerManagerDeployment_v311_00_to_latest(client appsclientv1.DeploymentsGetter, options *v1alpha1.KubeControllerManagerOperatorConfig, previousAvailability *operatorsv1alpha1.VersionAvailablity, forceDeployment bool) (*appsv1.Deployment, bool, error) {
	required := resourceread.ReadDeploymentV1OrDie(v311_00_assets.MustAsset("v3.11.0/kube-controller-manager/deployment.yaml"))
	required.Spec.Template.Spec.Containers[0].Image = options.Spec.ImagePullSpec
	required.Spec.Template.Spec.Containers[0].Args = append(required.Spec.Template.Spec.Containers[0].Args, fmt.Sprintf("-v=%d", options.Spec.Logging.Level))

	return resourceapply.ApplyDeployment(client, required, resourcemerge.ExpectedDeploymentGeneration(required, previousAvailability), forceDeployment)
}

func manageKubeControllerManagerPublicConfigMap_v311_00_to_latest(client coreclientv1.ConfigMapsGetter, controllerManagerConfigString string, operatorConfig *v1alpha1.KubeControllerManagerOperatorConfig) (*corev1.ConfigMap, bool, error) {
	uncastUnstructured, err := runtime.Decode(unstructured.UnstructuredJSONScheme, []byte(controllerManagerConfigString))
	if err != nil {
		return nil, false, err
	}
	controllerManagerConfig := uncastUnstructured.(runtime.Unstructured)

	configMap := resourceread.ReadConfigMapV1OrDie(v311_00_assets.MustAsset("v3.11.0/kube-controller-manager/public-info.yaml"))
	if operatorConfig.Status.CurrentAvailability != nil {
		configMap.Data["version"] = operatorConfig.Status.CurrentAvailability.Version
	} else {
		configMap.Data["version"] = ""
	}
	configMap.Data["imagePolicyConfig.internalRegistryHostname"], _, err = unstructured.NestedString(controllerManagerConfig.UnstructuredContent(), "imagePolicyConfig", "internalRegistryHostname")
	if err != nil {
		return nil, false, err
	}
	configMap.Data["imagePolicyConfig.externalRegistryHostname"], _, err = unstructured.NestedString(controllerManagerConfig.UnstructuredContent(), "imagePolicyConfig", "externalRegistryHostname")
	if err != nil {
		return nil, false, err
	}
	configMap.Data["projectConfig.defaultNodeSelector"], _, err = unstructured.NestedString(controllerManagerConfig.UnstructuredContent(), "projectConfig", "defaultNodeSelector")
	if err != nil {
		return nil, false, err
	}

	return resourceapply.ApplyConfigMap(client, configMap)
}
