package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"go.indent.com/indent-go/pkg/cliutil"
)

// Provider contains the root setup for all resources as well
// as global provider config
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"space": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INDENT_SPACE", nil),
				Description: "The Indent space name to work in",
			},
			"staging": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INDENT_STAGING", nil),
				Description: "Use Indent staging environment",
			},
			"json_key_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INDENT_JSON_KEY_FILE", nil),
				Description: "Path to Google Developers service account JSON key file",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"indent_webhook": resourceWebhook(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"indent_webhook": dataSourceWebhook(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	f, config := cliutil.New()
	var ok bool
	if config.Space, ok = d.Get("space").(string); !ok {
		return nil, diag.Errorf("failed to cast space to string")
	} else if config.Staging, ok = d.Get("staging").(bool); !ok {
		return nil, diag.Errorf("failed to cast staging to bool")
	} else if config.JSONKeyFile, ok = d.Get("json_key_file").(string); !ok {
		return nil, diag.Errorf("failed to cast json_key_file to string")
	} else if config.JSONKeyFile == "" && !f.IsLoggedIn(ctx) {
		return nil, diag.Errorf("user is not logged in")
	} else if _, err := cliutil.NewAPIClient(ctx, f); err != nil {
		return nil, diag.Errorf(err.Error())
	}
	return f, nil
}
