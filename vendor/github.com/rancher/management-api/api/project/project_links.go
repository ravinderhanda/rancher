package project

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/convert"
	clusterSchema "github.com/rancher/types/apis/cluster.cattle.io/v3/schema"
	projectSchema "github.com/rancher/types/apis/project.cattle.io/v3/schema"
	"github.com/rancher/types/client/management/v3"
	"github.com/rancher/types/config"
)

func ClusterLinks(apiContext *types.APIContext, resource *types.RawResource) {
	if resource.Type == client.ClusterType {
		for _, schema := range clusterSchema.Schemas.Schemas() {
			if !schema.CanList() {
				continue
			}
			resource.Links[schema.PluralName] = apiContext.URLBuilder.Link(schema.PluralName, resource)
		}

		resource.Links["namespaces"] = apiContext.URLBuilder.Link("namespaces", resource)
		resource.Links["schemas"] = apiContext.URLBuilder.Link("schemas", resource)

		for _, schema := range projectSchema.Schemas.Schemas() {
			if !schema.CanList() {
				continue
			}
			if _, ok := schema.ResourceFields["projectId"]; ok {
				continue
			}
			resource.Links[schema.PluralName] = apiContext.URLBuilder.Link(schema.PluralName, resource)
		}
	} else if resource.Type == client.ProjectType {
		for _, schema := range projectSchema.Schemas.Schemas() {
			if !schema.CanList() {
				continue
			}
			if _, ok := schema.ResourceFields["projectId"]; ok {
				resource.Links[schema.PluralName] = apiContext.URLBuilder.Link(schema.PluralName, resource)
			}
		}

		for _, subType := range config.ProjectTypes {
			name := "namespaced" + convert.Capitalize(subType) + "s"
			resource.Links[name] = apiContext.URLBuilder.Link(name, resource)
		}

		resource.Links["schemas"] = apiContext.URLBuilder.Link("schemas", resource)
	}
}