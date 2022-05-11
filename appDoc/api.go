package appDoc

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// InitDb - initialisation de la Base de donnée
func InitDb() *gorm.DB {
	// Ouverture du fichier dataDoc.db avec sqlite
	db, err := gorm.Open("sqlite3", "dataDoc.db")
	db.LogMode(true)

	// Création de la table en utilisant le ORM de gorm
	if !db.HasTable(&Doc{}) {
		db.CreateTable(&Doc{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Doc{})
	}

	// Erreur de chargement
	if err != nil {
		panic(err)
	}

	return db
}

// Handlers - événement d'écoutes pour les requetes  entrants
func Handlers() *gin.Engine {
	result := gin.Default()

	docs := result.Group("MyApi/doc")
	{
		docs.GET("/", getDocs)
		docs.POST("/addDoc", addDoc)
		docs.GET("/getDocById/:id", getDocById)
		docs.DELETE("/deleteDocById/:id", deleteDocById)
	}

	return result
}
