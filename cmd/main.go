package main

import (
	"company-name/cmd/api"
	"company-name/configs"
	"company-name/pkg/database"
	"company-name/pkg/email"
	loc "company-name/pkg/localization"
	"company-name/pkg/validators"
	"fmt"
	validator2 "github.com/go-playground/validator/v10"
	"log"
	"os"
)

func main() {
	cfg := configs.GetConfig()

	db, err := database.NewDatabase(cfg.DB.ConnectionString, cfg.DB.Name)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		os.Exit(1)
	}
	log.Println("Connected to database")

	// Set the path to the localization file
	localizationFilePath := "assets/locales/localization.json"

	// Load localization messages
	if err := loc.LoadMessages(localizationFilePath); err != nil {
		log.Fatalf("Error loading localization messages: %v", err)
	}

	// Set default language
	loc.SetLang("en")

	// Retrieve localized messages
	fmt.Println(loc.L("success"))                      // Output: Success
	fmt.Println(loc.L("resource_created", "User"))     // Output: User created successfully
	fmt.Println(loc.L("not_found"))                    // Output: Resource not found
	fmt.Println(loc.L("non_existing_key", "Fallback")) // Output: non_existing_key

	validatorPkg := validator2.New()
	validators.RegisterTimeFormatValidators(validatorPkg)
	validator := validators.NewValidator(validatorPkg)

	emailService := email.NewEmailService(cfg.Email.Host, cfg.Email.Port, cfg.Email.Username, cfg.Email.Password, cfg.Email.From)

	apiInstance := api.NewAPIServer(db, cfg, emailService, validator)
	if err := apiInstance.Run(); err != nil {
		log.Fatalf("Error running API server: %v", err)
		os.Exit(1)
	}
}
