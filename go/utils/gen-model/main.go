package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panicln("Error loading .env file")
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:           "utils/gen-model/export",
		ModelPkgPath:      "utils/gen-model/entity",
		FieldWithTypeTag:  false,
		FieldWithIndexTag: false,
		FieldSignable:     false,
	})
	fmt.Println(os.Getenv("DSN_SICERDAS"))
	dbOpen, _ := gorm.Open(postgres.Open(os.Getenv("DSN_SICERDAS")))
	g.UseDB(dbOpen) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
