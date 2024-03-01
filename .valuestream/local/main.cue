package local

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/yaml"
	"list"
	"strings"
)

_const: {
	#name: "nginx-demo"
	#port: 80
}

#Secret: {
	key:     string
	version: string
}

#ExternalSecret: {
	provider:   "kubernetes"
	apiVersion: "external-secrets.io/v1beta1"
	kind:       "ExternalSecret"
	output: [...string]
	...
}

DesignPattern: {
	name: "local"

	parameters: {
		k8sNamespace: string
		imageName:    string
		secrets: {
			SECRET01: #Secret
			SECRET02: #Secret
		}
	}

	resources: app: {
		deployment:     _deployment
		service:        _service
		externalSecret: _externalSecret
	}

	let _k8sNamespace = parameters.k8sNamespace
	let _secrets = parameters.secrets

	// This name is used for ExternalSecret metadata.name, target.name 
	let _k8sSecretName = "\(_k8sNamespace)-secrets-\(_hash)"

	let _esData = yaml.Marshal(_externalSecret.spec.data)

	let _hash = strings.SliceRunes(hex.Encode(sha256.Sum256(_esData)), 0, 10)

	let _selector = {
		app: _const.#name
	}
	let _container = {
		name:  _const.#name
		image: parameters.imageName
		ports: [{
			containerPort: _const.#port
		}]
		resources: {
			requests: {
				cpu:    "5m"
				memory: "32Mi"
			}
			limits: {
				cpu:    "10m"
				memory: "64Mi"
			}
		}
	}
	_deployment: {
		apiVersion: "apps/v1"
		kind:       "Deployment"
		metadata: name: _const.#name
		spec: {
			replicas: 1
			selector: matchLabels: _selector
			template: {
				metadata: labels: _selector
				spec: containers: [
					_container,
				]
			}
		}
	}

	_service: {
		apiVersion: "v1"
		kind:       "Service"
		metadata: name: _const.#name
		spec: {
			type: "LoadBalancer"
			ports: [{
				name:       "http"
				port:       8080
				targetPort: _const.#port
			}]
			selector: _selector
		}
	}

	_externalSecret: #ExternalSecret & {
		metadata: {
			name:      _k8sSecretName
			namespace: _k8sNamespace
		}
		spec: {
			refreshInterval: "0"
			secretStoreRef: {
				name: "gcp-secret-manager"
				kind: "ClusterSecretStore"
			}
			target: {
				name:           _k8sSecretName
				creationPolicy: "Owner"
			}

			// to determine order of elements of spec.data list uniquely, sort elements in ascending order based on secretKey
			// this determines value of _hash in _k8sSecretName uniquely
			let _dataUnsorted = [ for k, v in _secrets {
				secretKey: k
				remoteRef: {
					key:     v.key
					version: v.version
				}
			}]
			data: list.Sort(_dataUnsorted, {x: {}, y: {}, less: x.secretKey < y.secretKey})
		}
	}

	// apply namespace
	_deployment: metadata: namespace: parameters.k8sNamespace
	_service: metadata: namespace:    parameters.k8sNamespace
}
