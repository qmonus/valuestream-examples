package local

_const: {
	#name: "nginx-demo"
	#port: 80
}

DesignPattern: {
	name: "local"

	parameters: {
		k8sNamespace: string
		imageName:    string
	}

	resources: app: {
		deployment: _deployment
		service:    _service
	}

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

	// apply namespace
	_deployment: metadata: namespace: parameters.k8sNamespace
	_service: metadata: namespace:    parameters.k8sNamespace
}
