package appDoc

import (
	"github.com/gin-gonic/gin"
)

// Doc - Définition de la structure du document avec les champs en
type Doc struct {
	Id          int    `gorm:"not null" form:"id" json:"id"`
	Name        string `gorm:"not null" form:"name" json:"name"`
	Description string `gorm:"not null" form:"description" json:"description"`
}

// AddDoc Ajouter un document dans la dataBase
func addDoc(c *gin.Context) {
	//Ouverture de la dataBase pour stocker le document et ensuite la fermer plus tard après l'excution de la fonction
	db := InitDb()
	defer db.Close()

	//extraction des donnée envoyé par la requête
	var json Doc
	err := c.Bind(&json)
	if err != nil {
		return
	}

	// verifier Si tous les  champs sont saisies
	if json.Name != "" && json.Id > 0 && json.Description != "" {
		db.Create(&json)
		// Affichage des données saisies en json
		c.JSON(201, gin.H{"success": json})
	} else {
		// Affichage de l'erreur
		c.JSON(422, gin.H{"error": "un ou plusieurs champs sont vides ou incorrect"})
	}
}

// getDocs - Récuperer tous les documents
func getDocs(c *gin.Context) {
	//Ouverture de la dataBase pour prendre les documents et ensuite la fermer plus tard après l'excution de la fonction
	db := InitDb()
	defer db.Close()

	//Recupérations des documents
	var docs []Doc
	db.Find(&docs)
	// Affichage des documents
	c.JSON(200, docs)
}

// getDocById - Recuperer un document  son id
func getDocById(c *gin.Context) {
	//Ouverture de la dataBase pour la recherche de document et ensuite la fermer plus tard après l'excution de la fonction
	db := InitDb()
	defer db.Close()

	// Récupérer l'id dans une variable
	id := c.Params.ByName("id")
	var doc Doc
	db.First(&doc, id)

	//verifier si l'id à été trouvé
	if doc.Id != 0 {
		// Affichage du document correspondant à l'id
		c.JSON(200, doc)
	} else {
		// Affichage de l'erreur
		c.JSON(404, gin.H{"error": "Document non trouvé ou id incorrect"})
	}
}

// deleteDocById - Supprimer un document par son id
func deleteDocById(c *gin.Context) {
	//Ouverture de la dataBase pour la recherche de document et ensuite la fermer plus tard après l'excution de la fonction
	db := InitDb()
	defer db.Close()

	// Récupération de l'id dans une variable
	id := c.Params.ByName("id")
	var doc Doc
	db.First(&doc, id)

	//verifier si l'id à été trouvé
	if doc.Id != 0 {
		db.Delete(&doc)
		// Affichage d'un message de suppression
		c.JSON(200, gin.H{"success": "Document " + id + " supprimé"})
	} else {
		// Affichage de l'erreur
		c.JSON(404, gin.H{"error": "Document non trouvé ou id incorrect"})
	}
}
