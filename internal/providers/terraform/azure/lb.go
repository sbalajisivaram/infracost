package azure

import (
	"fmt"
	"github.com/infracost/infracost/internal/schema"
	"github.com/shopspring/decimal"
)

func GetAzureRMLBRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_lb",
		RFunc: NewAzureRMLB,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

func NewAzureRMLB(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	var costComponents []*schema.CostComponent
	location := "Global"
	group := d.References("resource_group_name")

	if strings.HasPrefix(strings.ToLower(group[0].Get("location").String()), "usgov") {
		location = "US Gov"
	}

	if strings.Contains(strings.ToLower(group[0].Get("location").String()), "china") {
		location = "Сhina"
	}
	if sku == "Basic" {
		return &schema.Resource{
			NoPrice:   true,
			IsSkipped: true,
		}
	}

	if u != nil && u.Get("monthly_data_processed_gb").Type != gjson.Null {
		monthlyDataProcessedGb = decimalPtr(decimal.NewFromFloat(u.Get("monthly_data_processed_gb").Float()))
	}

	costComponents := []*schema.CostComponent{
		{
			Name:           "Data processed"
			Unit:           "GB",
			UnitMultiplier:  1,
			MonthlyQuantity: dataProcessed,
			ProductFilter: &schema.ProductFilter{
				VendorName:    strPtr("azure"),
				Region:        strPtr("location"),
				Service:       strPtr("Load Balancer"),
				ProductFamily: strPtr("Networking"),
				AttributeFilters: []*schema.AttributeFilter{
					{Key: "meterName", ValueRegex: strPtr(fmt.Sprintf("%s Data Processed", sku))},
					{Key: "skuName", Value: strPtr(sku)},
				},
			},
			PriceFilter: &schema.PriceFilter{
				PurchaseOption: strPtr("Consumption"),
			},
		},
	}

	return &schema.Resource{
		Name:           d.Address,
		CostComponents: costComponents,
	}
}