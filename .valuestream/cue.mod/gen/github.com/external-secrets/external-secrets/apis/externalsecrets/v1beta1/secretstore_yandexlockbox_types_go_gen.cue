// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1

package v1beta1

import esmeta "github.com/external-secrets/external-secrets/apis/meta/v1"

#YandexLockboxAuth: {
	// The authorized key used for authentication
	// +optional
	authorizedKeySecretRef?: esmeta.#SecretKeySelector @go(AuthorizedKey)
}

#YandexLockboxCAProvider: {
	certSecretRef?: esmeta.#SecretKeySelector @go(Certificate)
}

// YandexLockboxProvider Configures a store to sync secrets using the Yandex Lockbox provider.
#YandexLockboxProvider: {
	// Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443')
	// +optional
	apiEndpoint?: string @go(APIEndpoint)

	// Auth defines the information necessary to authenticate against Yandex Lockbox
	auth: #YandexLockboxAuth @go(Auth)

	// The provider for the CA bundle to use to validate Yandex.Cloud server certificate.
	// +optional
	caProvider?: null | #YandexLockboxCAProvider @go(CAProvider,*YandexLockboxCAProvider)
}
