package templates

import (
	networking "k8s.io/api/networking/v1"
)

#Ingress: networking.#Ingress & {
	#config:    #Config
	apiVersion: "networking.k8s.io/v1"
	kind:       "Ingress"
	metadata:   #config.metadata

	metadata: {
		annotations: {
			"nginx.ingress.kubernetes.io/ssl-redirect":   "false"
			"nginx.ingress.kubernetes.io/use-regex":      "true"
			"nginx.ingress.kubernetes.io/rewrite-target": "/$2"
		}
	}
	spec: rules: [{
		http: paths: [{
			path:     "/\(#config.metadata.namespace)/\(#config.metadata.name)(/|$)(.*)"
			pathType: networking.#PathTypeImplementationSpecific
			backend: service: {
				name: #config.metadata.name
				port: number: 8080
			}
		}]
	}]
}
