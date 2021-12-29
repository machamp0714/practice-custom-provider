package hashicups

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceOrder() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"coffee_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"coffee_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_teaser": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"coffee_price": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"coffee_image": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"quantity": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
