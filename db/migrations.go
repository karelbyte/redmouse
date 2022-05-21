package db

import (
	"elpuertodigital/redmouse/models"
	"os"
)

func ExecMigrations() {

	env := os.Getenv("APP_ENV")
	if env == "local" {
		Conect().Migrator().DropTable("colors")
		Conect().Migrator().DropTable("sizes")
		Conect().Migrator().DropTable("measures")
		Conect().Migrator().DropTable("categories")
		Conect().Migrator().DropTable("providers")
		Conect().Migrator().DropTable("products")
		Conect().Migrator().DropTable("product_providers")
		Conect().Migrator().DropTable("product_variations")
		Conect().Migrator().DropTable("product_variation_prices")
		Conect().Migrator().DropTable("receptions")
		Conect().Migrator().DropTable("reception_products")
		Conect().Migrator().DropTable("outputs")
		Conect().Migrator().DropTable("output_products")
		Conect().Migrator().DropTable("adjustment")
		Conect().Migrator().DropTable("adjustment_products")
		Conect().Migrator().DropTable("clients")
		Conect().Migrator().DropTable("invoices")
		Conect().Migrator().DropTable("invoice_products")
		Conect().Migrator().DropTable("quotations")
		Conect().Migrator().DropTable("quotation_products")
		Conect().Migrator().DropTable("inventories")
		Conect().Migrator().DropTable("warehouses")
		Conect().Migrator().DropTable("user")
	}

	err := Conect().AutoMigrate(
		&models.Color{},
		&models.Measure{},
		&models.Category{},
		&models.Size{},
		&models.Provider{},
		&models.Product{},
		&models.ProductProvider{},
		&models.ProductVariation{},
		&models.ProductVariationPrice{},
		&models.Reception{},
		&models.ReceptionProduct{},
		&models.Output{},
		&models.OuputProduct{},
		&models.Adjustment{},
		&models.AjustmentProduct{},
		&models.Invoice{},
		&models.InvoiceProduct{},
		&models.Quotation{},
		&models.QuotationProduct{},
		&models.Client{},
		&models.Inventory{},
		&models.Warehouse{},
		&models.User{},
	)

	if err == nil && env == "local" {
    //    MeasureSeed()
	}

	println("All tables migrate!")
}
