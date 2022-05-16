package db

import "elpuertodigital/redmouse/models"


func ExecMigrations(env string) {

	if (env == "local") {
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
		Conect().Migrator().DropTable("user")
	}

	Conect().AutoMigrate(
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
		&models.User{},
	)

    println("All tables migrate!")
}
