package k8s

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "k8s"

// Configure adds configurations for k8s resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("scaleway_k8s_cluster", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Cluster"
	})

	p.AddResourceConfigurator("scaleway_k8s_pool", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = shortGroup
		r.Kind = "Pool"

		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
	})
}
