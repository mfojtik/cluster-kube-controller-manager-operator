// Code generated by go-bindata.
// sources:
// bindata/v3.11.0/kube-controller-manager/cm.yaml
// bindata/v3.11.0/kube-controller-manager/defaultconfig.yaml
// bindata/v3.11.0/kube-controller-manager/kubeconfig-cert-syncer.yaml
// bindata/v3.11.0/kube-controller-manager/kubeconfig-cm.yaml
// bindata/v3.11.0/kube-controller-manager/leader-election-rolebinding.yaml
// bindata/v3.11.0/kube-controller-manager/ns.yaml
// bindata/v3.11.0/kube-controller-manager/pod-cm.yaml
// bindata/v3.11.0/kube-controller-manager/pod.yaml
// bindata/v3.11.0/kube-controller-manager/sa.yaml
// bindata/v3.11.0/kube-controller-manager/svc.yaml
// DO NOT EDIT!

package v311_00_assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _v3110KubeControllerManagerCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-kube-controller-manager
  name: config
data:
  config.yaml:
`)

func v3110KubeControllerManagerCmYamlBytes() ([]byte, error) {
	return _v3110KubeControllerManagerCmYaml, nil
}

func v3110KubeControllerManagerCmYaml() (*asset, error) {
	bytes, err := v3110KubeControllerManagerCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-controller-manager/cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeControllerManagerDefaultconfigYaml = []byte(`apiVersion: kubecontrolplane.config.openshift.io/v1
kind: KubeControllerManagerConfig
extendedArguments:
  enable-dynamic-provisioning:
  - "true"
  allocate-node-cidrs:
  - "true"
  configure-cloud-routes:
  - "false"
  cluster-cidr:
  - "10.2.0.0/16"
  service-cluster-ip-range:
  - "10.3.0.0/16"
  use-service-account-credentials:
  - "true"
  flex-volume-plugin-dir:
  - "/etc/kubernetes/kubelet-plugins/volume/exec" # created by machine-config-operator, owned by storage team/hekumar@redhat.com
  leader-elect:
  - "true"
  leader-elect-retry-period:
  - "3s"
  leader-elect-resource-lock:
  - "configmaps"
  controllers:
  - "*"
  - "-ttl" # TODO: this is excluded in kube-core, but not in #21092
  - "-bootstrapsigner"
  - "-tokencleaner"
  experimental-cluster-signing-duration:
  - "720h"
  secure-port:
  - "10257"
  port:
  - "0"
  cert-dir:
  - "/var/run/kubernetes"
  root-ca-file:
  - "/etc/kubernetes/static-pod-resources/configmaps/serviceaccount-ca/ca-bundle.crt"
  service-account-private-key-file:
  - "/etc/kubernetes/static-pod-resources/secrets/service-account-private-key/service-account.key"
  cluster-signing-cert-file:
  - "/etc/kubernetes/static-pod-resources/secrets/csr-signer/tls.crt"
  cluster-signing-key-file:
  - "/etc/kubernetes/static-pod-resources/secrets/csr-signer/tls.key"
  kube-api-qps:
  - "150" # this is a historical values
  kube-api-burst:
  - "300" # this is a historical values

`)

func v3110KubeControllerManagerDefaultconfigYamlBytes() ([]byte, error) {
	return _v3110KubeControllerManagerDefaultconfigYaml, nil
}

func v3110KubeControllerManagerDefaultconfigYaml() (*asset, error) {
	bytes, err := v3110KubeControllerManagerDefaultconfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-controller-manager/defaultconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeControllerManagerKubeconfigCertSyncerYaml = []byte(`# TODO provide distinct trust between this and the KCM itself
apiVersion: v1
kind: ConfigMap
metadata:
  name: kube-controller-cert-syncer-kubeconfig
  namespace: openshift-kube-controller-manager
data:
  kubeconfig: |
    apiVersion: v1
    clusters:
      - cluster:
          certificate-authority: /etc/kubernetes/static-pod-resources/configmaps/serviceaccount-ca/ca-bundle.crt
          server: https://localhost:6443
        name: loopback
    contexts:
      - context:
          cluster: loopback
          user: kube-controller-manager
        name: kube-controller-manager
    current-context: kube-controller-manager
    kind: Config
    preferences: {}
    users:
      - name: kube-controller-manager
        user:
          client-certificate: /etc/kubernetes/static-pod-resources/secrets/kube-controller-manager-client-cert-key/tls.crt
          client-key: /etc/kubernetes/static-pod-resources/secrets/kube-controller-manager-client-cert-key/tls.key
`)

func v3110KubeControllerManagerKubeconfigCertSyncerYamlBytes() ([]byte, error) {
	return _v3110KubeControllerManagerKubeconfigCertSyncerYaml, nil
}

func v3110KubeControllerManagerKubeconfigCertSyncerYaml() (*asset, error) {
	bytes, err := v3110KubeControllerManagerKubeconfigCertSyncerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-controller-manager/kubeconfig-cert-syncer.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeControllerManagerKubeconfigCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  name: controller-manager-kubeconfig
  namespace: openshift-kube-controller-manager
data:
  kubeconfig: |
    apiVersion: v1
    clusters:
      - cluster:
          certificate-authority: /etc/kubernetes/static-pod-resources/configmaps/serviceaccount-ca/ca-bundle.crt
          server: https://localhost:6443
        name: loopback
    contexts:
      - context:
          cluster: loopback
          user: kube-controller-manager
        name: kube-controller-manager
    current-context: kube-controller-manager
    kind: Config
    preferences: {}
    users:
      - name: kube-controller-manager
        user:
          client-certificate: /etc/kubernetes/static-pod-resources/secrets/kube-controller-manager-client-cert-key/tls.crt
          client-key: /etc/kubernetes/static-pod-resources/secrets/kube-controller-manager-client-cert-key/tls.key
`)

func v3110KubeControllerManagerKubeconfigCmYamlBytes() ([]byte, error) {
	return _v3110KubeControllerManagerKubeconfigCmYaml, nil
}

func v3110KubeControllerManagerKubeconfigCmYaml() (*asset, error) {
	bytes, err := v3110KubeControllerManagerKubeconfigCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-controller-manager/kubeconfig-cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeControllerManagerLeaderElectionRolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  namespace: kube-system
  name: system:openshift:leader-locking-kube-controller-manager
roleRef:
  kind: Role
  name: system::leader-locking-kube-controller-manager
subjects:
- kind: User
  name: system:kube-controller-manager
`)

func v3110KubeControllerManagerLeaderElectionRolebindingYamlBytes() ([]byte, error) {
	return _v3110KubeControllerManagerLeaderElectionRolebindingYaml, nil
}

func v3110KubeControllerManagerLeaderElectionRolebindingYaml() (*asset, error) {
	bytes, err := v3110KubeControllerManagerLeaderElectionRolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-controller-manager/leader-election-rolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeControllerManagerNsYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  annotations:
    openshift.io/node-selector: ""
  name: openshift-kube-controller-manager
  labels:
    openshift.io/run-level: "0"
`)

func v3110KubeControllerManagerNsYamlBytes() ([]byte, error) {
	return _v3110KubeControllerManagerNsYaml, nil
}

func v3110KubeControllerManagerNsYaml() (*asset, error) {
	bytes, err := v3110KubeControllerManagerNsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-controller-manager/ns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeControllerManagerPodCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-kube-controller-manager
  name: kube-controller-manager-pod
data:
  pod.yaml:
  forceRedeploymentReason:
  version:
`)

func v3110KubeControllerManagerPodCmYamlBytes() ([]byte, error) {
	return _v3110KubeControllerManagerPodCmYaml, nil
}

func v3110KubeControllerManagerPodCmYaml() (*asset, error) {
	bytes, err := v3110KubeControllerManagerPodCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-controller-manager/pod-cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeControllerManagerPodYaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  name: kube-controller-manager
  namespace: openshift-kube-controller-manager
  labels:
    app: kube-controller-manager
    kube-controller-manager: "true"
    revision: "REVISION"
spec:
  shareProcessNamespace: true
  initContainers:
  - name: wait-for-host-port
    terminationMessagePolicy: FallbackToLogsOnError
    image: ${IMAGE}
    imagePullPolicy: IfNotPresent
    command: ['/usr/bin/timeout', '30', "/bin/bash", "-c"]
    args:
    - |
      echo -n "Waiting for port :10257 to be released."
      while [ -n "$(lsof -ni :10257)" ]; do
        echo -n "."
        sleep 1
      done
  containers:
  - name: kube-controller-manager-REVISION
    image: ${IMAGE}
    imagePullPolicy: IfNotPresent
    terminationMessagePolicy: FallbackToLogsOnError
    command: ["hyperkube", "kube-controller-manager"]
    args:
    - --openshift-config=/etc/kubernetes/static-pod-resources/configmaps/config/config.yaml
    - --kubeconfig=/etc/kubernetes/static-pod-resources/configmaps/controller-manager-kubeconfig/kubeconfig
    - --authentication-kubeconfig=/etc/kubernetes/static-pod-resources/configmaps/controller-manager-kubeconfig/kubeconfig
    - --authorization-kubeconfig=/etc/kubernetes/static-pod-resources/configmaps/controller-manager-kubeconfig/kubeconfig
    - --client-ca-file=/etc/kubernetes/static-pod-certs/configmaps/client-ca/ca-bundle.crt
    - --requestheader-client-ca-file=/etc/kubernetes/static-pod-certs/configmaps/aggregator-client-ca/ca-bundle.crt
    resources:
      requests:
        memory: 200Mi
        cpu: 100m
    ports:
      - containerPort: 10257
    volumeMounts:
    - mountPath: /etc/kubernetes/static-pod-resources
      name: resource-dir
    - mountPath: /etc/kubernetes/static-pod-certs
      name: cert-dir
    livenessProbe:
      httpGet:
        scheme: HTTPS
        port: 10257
        path: healthz
      initialDelaySeconds: 45
      timeoutSeconds: 10
    readinessProbe:
      httpGet:
        scheme: HTTPS
        port: 10257
        path: healthz
      initialDelaySeconds: 10
      timeoutSeconds: 10
  - name: kube-controller-manager-cert-syncer-REVISION
    env:
      - name: POD_NAME
        valueFrom:
          fieldRef:
            fieldPath: metadata.name
      - name: POD_NAMESPACE
        valueFrom:
          fieldRef:
            fieldPath: metadata.namespace
    image: ${OPERATOR_IMAGE}
    imagePullPolicy: IfNotPresent
    volumeMounts:
      - mountPath: /etc/kubernetes/static-pod-resources
        name: resource-dir
      - mountPath: /etc/kubernetes/static-pod-certs
        name: cert-dir
    terminationMessagePolicy: FallbackToLogsOnError
    command: ["cluster-kube-controller-manager-operator", "cert-syncer"]
    args:
      - --kubeconfig=/etc/kubernetes/static-pod-resources/configmaps/kube-controller-cert-syncer-kubeconfig/kubeconfig
      - --namespace=$(POD_NAMESPACE)
      - --destination-dir=/etc/kubernetes/static-pod-certs
    resources:
      requests:
        memory: 50Mi
        cpu: 10m
  - name: kube-controller-manager-watchdog-REVISION
    securityContext:
      capabilities:
        add:
          - SYS_PTRACE
    env:
      - name: POD_NAME
        valueFrom:
          fieldRef:
            fieldPath: metadata.name
      - name: POD_NAMESPACE
        valueFrom:
          fieldRef:
            fieldPath: metadata.namespace
    image: ${OPERATOR_IMAGE}
    imagePullPolicy: IfNotPresent
    volumeMounts:
      - mountPath: /etc/kubernetes/static-pod-resources
        name: resource-dir
    terminationMessagePolicy: FallbackToLogsOnError
    command: ["cluster-kube-controller-manager-operator", "file-watcher-watchdog"]
    args:
      - --kubeconfig=/etc/kubernetes/static-pod-resources/configmaps/kube-controller-cert-syncer-kubeconfig/kubeconfig
      - --namespace=$(POD_NAMESPACE)
      - --process-name=hyperkube
      - --termination-grace-period=30s
      - --files=/etc/kubernetes/static-pod-certs/secrets/kube-controller-manager-client-cert-key/tls.crt,/etc/kubernetes/static-pod-/secrets/kube-controller-manager-client-cert-key/tls.key,/etc/kubernetes/static-pod-resources/secrets/csr-signer/tls.crt,/etc/kubernetes/static-pod-resources/secrets/csr-signer/tls.key,/etc/kubernetes/static-pod-resources/secrets/kube-scheduler-client-cert-key/tls.crt,/etc/kubernetes/static-pod-resources/secrets/csr-signer/tls.key
    resources:
      requests:
        memory: 50Mi
        cpu: 10m
  hostNetwork: true
  priorityClassName: system-node-critical
  tolerations:
  - operator: "Exists"
  volumes:
  - hostPath:
      path: /etc/kubernetes/static-pod-resources/kube-controller-manager-pod-REVISION
    name: resource-dir
  - hostPath:
      path: /etc/kubernetes/static-pod-resources/kube-controller-manager-certs
    name: cert-dir
`)

func v3110KubeControllerManagerPodYamlBytes() ([]byte, error) {
	return _v3110KubeControllerManagerPodYaml, nil
}

func v3110KubeControllerManagerPodYaml() (*asset, error) {
	bytes, err := v3110KubeControllerManagerPodYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-controller-manager/pod.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeControllerManagerSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: openshift-kube-controller-manager
  name: kube-controller-manager-sa
`)

func v3110KubeControllerManagerSaYamlBytes() ([]byte, error) {
	return _v3110KubeControllerManagerSaYaml, nil
}

func v3110KubeControllerManagerSaYaml() (*asset, error) {
	bytes, err := v3110KubeControllerManagerSaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-controller-manager/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110KubeControllerManagerSvcYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  namespace: openshift-kube-controller-manager
  name: kube-controller-manager
  annotations:
    service.alpha.openshift.io/serving-cert-secret-name: serving-cert
    prometheus.io/scrape: "true"
    prometheus.io/scheme: https
spec:
  selector:
    kube-controller-manager: "true"
  ports:
  - name: https
    port: 443
    targetPort: 10257
`)

func v3110KubeControllerManagerSvcYamlBytes() ([]byte, error) {
	return _v3110KubeControllerManagerSvcYaml, nil
}

func v3110KubeControllerManagerSvcYaml() (*asset, error) {
	bytes, err := v3110KubeControllerManagerSvcYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/kube-controller-manager/svc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"v3.11.0/kube-controller-manager/cm.yaml":                          v3110KubeControllerManagerCmYaml,
	"v3.11.0/kube-controller-manager/defaultconfig.yaml":               v3110KubeControllerManagerDefaultconfigYaml,
	"v3.11.0/kube-controller-manager/kubeconfig-cert-syncer.yaml":      v3110KubeControllerManagerKubeconfigCertSyncerYaml,
	"v3.11.0/kube-controller-manager/kubeconfig-cm.yaml":               v3110KubeControllerManagerKubeconfigCmYaml,
	"v3.11.0/kube-controller-manager/leader-election-rolebinding.yaml": v3110KubeControllerManagerLeaderElectionRolebindingYaml,
	"v3.11.0/kube-controller-manager/ns.yaml":                          v3110KubeControllerManagerNsYaml,
	"v3.11.0/kube-controller-manager/pod-cm.yaml":                      v3110KubeControllerManagerPodCmYaml,
	"v3.11.0/kube-controller-manager/pod.yaml":                         v3110KubeControllerManagerPodYaml,
	"v3.11.0/kube-controller-manager/sa.yaml":                          v3110KubeControllerManagerSaYaml,
	"v3.11.0/kube-controller-manager/svc.yaml":                         v3110KubeControllerManagerSvcYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"v3.11.0": {nil, map[string]*bintree{
		"kube-controller-manager": {nil, map[string]*bintree{
			"cm.yaml":                          {v3110KubeControllerManagerCmYaml, map[string]*bintree{}},
			"defaultconfig.yaml":               {v3110KubeControllerManagerDefaultconfigYaml, map[string]*bintree{}},
			"kubeconfig-cert-syncer.yaml":      {v3110KubeControllerManagerKubeconfigCertSyncerYaml, map[string]*bintree{}},
			"kubeconfig-cm.yaml":               {v3110KubeControllerManagerKubeconfigCmYaml, map[string]*bintree{}},
			"leader-election-rolebinding.yaml": {v3110KubeControllerManagerLeaderElectionRolebindingYaml, map[string]*bintree{}},
			"ns.yaml":                          {v3110KubeControllerManagerNsYaml, map[string]*bintree{}},
			"pod-cm.yaml":                      {v3110KubeControllerManagerPodCmYaml, map[string]*bintree{}},
			"pod.yaml":                         {v3110KubeControllerManagerPodYaml, map[string]*bintree{}},
			"sa.yaml":                          {v3110KubeControllerManagerSaYaml, map[string]*bintree{}},
			"svc.yaml":                         {v3110KubeControllerManagerSvcYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
