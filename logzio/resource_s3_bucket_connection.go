package logzio

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func resourceS3BucketConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceS3BucketConnectorCreate,
		Read: resourceS3BucketConnectorRead,
		Update: resourceS3BucketConnectorUpdate,
		Delete: resourceS3BucketConnectorDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{

		},
	}
}

func resourceS3BucketConnectorCreate(d *schema.ResourceData, m interface{}) error {

}

func resourceS3BucketConnectorRead(d *schema.ResourceData, m interface{}) error {

}

func resourceS3BucketConnectorUpdate(d *schema.ResourceData, m interface{}) error {

}

func resourceS3BucketConnectorDelete(d *schema.ResourceData, m interface{}) error {

}
