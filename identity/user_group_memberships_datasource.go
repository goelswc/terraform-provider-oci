// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package identity

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func UserGroupMembershipDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readUserGroupMemberships,
		Schema: map[string]*schema.Schema{
			"compartment_id":{
				Type:     schema.TypeString,
				Required: true,
			},
			"user_id":{
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_id":{
				Type:     schema.TypeString,
				Optional: true,
			},
			"memberships": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     UserGroupMembershipResource(),
			},
		},
	}
}

func readUserGroupMemberships(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &UserGroupMembershipDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}
