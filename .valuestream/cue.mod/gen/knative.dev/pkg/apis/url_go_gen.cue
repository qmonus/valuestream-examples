// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go knative.dev/pkg/apis

package apis

// URL is an alias of url.URL.
// It has custom json marshal methods that enable it to be used in K8s CRDs
// such that the CRD resource will have the URL but operator code can can work with url.URL struct
// +kubebuilder:validation:Type=string
#URL: _