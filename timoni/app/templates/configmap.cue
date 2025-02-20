package templates

import (
	timoniv1 "timoni.sh/core/v1alpha1"
)

#ConfigMap: timoniv1.#ImmutableConfig & {
	#config: #Config
	#Kind:   timoniv1.#ConfigMapKind

	#Meta: #config.metadata
	#Data: {
		"\(#config.metadata.name).txt":
			"""
		\(#config.message)
		"""
	}

}

#ConfigMapEnvs: timoniv1.#ImmutableConfig & {
	#config: #Config
	#Kind:   timoniv1.#ConfigMapKind
	#Meta:   #config.metadata

	#Suffix: "-envs"
	#Data: {
		MESSAGE_FILE:      "/message-file/\(#config.metadata.name).txt"
		NEXT_HOP_HOSTNAME: #config.nextHopHostname
		NEXT_HOP_PORT:     "8080"
	}
}
