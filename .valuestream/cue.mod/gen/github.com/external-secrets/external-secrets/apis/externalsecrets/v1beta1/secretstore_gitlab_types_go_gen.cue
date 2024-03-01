// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1

package v1beta1

import esmeta "github.com/external-secrets/external-secrets/apis/meta/v1"

// Configures a store to sync secrets with a GitLab instance.
#GitlabProvider: {
	// URL configures the GitLab instance URL. Defaults to https://gitlab.com/.
	url?: string @go(URL)

	// Auth configures how secret-manager authenticates with a GitLab instance.
	auth: #GitlabAuth @go(Auth)

	// ProjectID specifies a project where secrets are located.
	projectID?: string @go(ProjectID)
}

#GitlabAuth: SecretRef: #GitlabSecretRef

#GitlabSecretRef: {
	// AccessToken is used for authentication.
	accessToken?: esmeta.#SecretKeySelector @go(AccessToken)
}