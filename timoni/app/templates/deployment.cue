package templates

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

#Deployment: appsv1.#Deployment & {
	#config:    #Config
	#cmName:    string
	#cmeName:   string
	apiVersion: "apps/v1"
	kind:       "Deployment"
	metadata:   #config.metadata
	spec: appsv1.#DeploymentSpec & {
		replicas: #config.replicas
		selector: matchLabels: #config.selector.labels
		template: {
			metadata: {
				labels: #config.selector.labels
				if #config.podAnnotations != _|_ {
					annotations: #config.podAnnotations
				}
			}
			spec: corev1.#PodSpec & {
				serviceAccountName: #config.metadata.name
				containers: [
					{
						name:            #config.metadata.name
						image:           #config.image.reference
						imagePullPolicy: #config.image.pullPolicy
						ports: [
							{
								containerPort: 8080
								protocol:      "TCP"
							},
						]
						livenessProbe: {
							httpGet: {
								path: "/.well-known/health"
								port: 8080
							}
						}
						readinessProbe: {
							httpGet: {
								path: "/.well-known/health"
								port: 8080
							}
						}
						volumeMounts: [
							{
								mountPath: "/message-file"
								name:      "config"
							},
						]
						resources:       #config.resources
						securityContext: #config.securityContext
						env: [{
							name: "CURRENT_NAMESPACE"
							valueFrom: fieldRef: fieldPath: "metadata.namespace"
						}]
						envFrom: [{
							configMapRef: name: #cmeName
						}]

					},
				]
				volumes: [
					{
						name: "config"
						configMap: {
							name: #cmName
						}
					},
				]
				if #config.podSecurityContext != _|_ {
					securityContext: #config.podSecurityContext
				}
				if #config.topologySpreadConstraints != _|_ {
					topologySpreadConstraints: #config.topologySpreadConstraints
				}
				if #config.affinity != _|_ {
					affinity: #config.affinity
				}
				if #config.tolerations != _|_ {
					tolerations: #config.tolerations
				}
				if #config.imagePullSecrets != _|_ {
					imagePullSecrets: #config.imagePullSecrets
				}
			}
		}
	}
}
