// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commercetools

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/providers/commercetools/connectivity"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/labd/commercetools-go-sdk/commercetools"
)

type ProductTypeGenerator struct {
	CommercetoolsService
}

// InitResources generates Terraform Resources from Commercetools API
func (g *ProductTypeGenerator) InitResources() error {
	cfg := connectivity.Config{
		ClientID:     g.GetArgs()["client_id"].(string),
		ClientSecret: g.GetArgs()["client_secret"].(string),
		ClientScope:  g.GetArgs()["client_scope"].(string),
		TokenURL:     g.GetArgs()["token_url"].(string) + "/oauth/token",
		BaseURL:      g.GetArgs()["base_url"].(string),
	}

	client := cfg.NewClient()

	productTypes, err := client.ProductTypeQuery(context.Background(), &commercetools.QueryInput{})
	if err != nil {
		return err
	}
	for _, productType := range productTypes.Results {
		resourceName := productType.Key
		if resourceName == "" {
			resourceName = normalizeResourceName(productType.Name)
		}
		g.Resources = append(g.Resources, terraformutils.NewResource(
			productType.ID,
			resourceName,
			"commercetools_product_type",
			"commercetools",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		))
	}
	return nil
}
