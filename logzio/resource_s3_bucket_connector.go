package logzio

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

const (
	s3BucketConnectorAccessKey 					string = "access_key"
	s3BucketConnectorSecretKey 					string = "secret_key"
	s3BucketConnectorArn 						string = "arn"
	s3BucketConnectorBucket 					string = "bucket"
	s3BucketConnectorPrefix 					string = "prefix"
	s3BucketConnectorActive						string = "active"
	s3BucketConnectorAddS3ObjectKeyAsLogField 	string = "add_s3_object_key_as_log_field"
	s3BucketConnectorRegion 					string = "region"
	s3BucketConnectorLogsType					string = "logs_type"
)


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
			s3BucketConnectorAccessKey: {
				Type: schema.TypeString,
				Optional: true,
			},
			s3BucketConnectorSecretKey: {
				Type: schema.TypeString,
				Optional: true,
			},
			s3BucketConnectorArn: {
				Type: schema.TypeString,
				Optional: true,
			},
			s3BucketConnectorBucket: {
				Type: schema.TypeString,
				Required: true,
			},
			s3BucketConnectorPrefix: {
				Type: schema.TypeString,
				Optional: true,
			},
			s3BucketConnectorActive: {
				Type: schema.TypeBool,
				Optional: true,
				Default: true,
			},
			s3BucketConnectorAddS3ObjectKeyAsLogField: {
				Type: schema.TypeBool,
				Optional: true,
				Default: false,
			},
			s3BucketConnectorRegion: {
				Type: schema.TypeString,
				Required: true,
			},
			s3BucketConnectorLogsType: {
				Type: schema.TypeString,
				Required: true,
			},
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
