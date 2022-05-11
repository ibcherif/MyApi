## Autheur : Ibrahima Cherif

# MyApi
Api RestFull développé en Go

## Mise en place de l'API
Pour mettre en place cette API, j'ai utilisé des 3 librairies ci-dessous.

Gin : le micro framework basé sur HttpRouter permettant de faire la gestion des routes => < go get github.com/gin-gonic/gin >

go-sqlite3 : le driver(pilote) SQLite3 pour la connexion en base de donnée => <go get github.com/mattn/go-sqlite3> 

Gorm : l'ORM (Object-Relational Mapping) pour les requête en base de donnée => <go get github.com/jinzhu/gorm>

## Tester l'API avec isomnia 
Pour tester cette API, j'ai utilisé le logiciel  [insomnia](https://insomnia.rest/)  pour faire les requêtes sur l'API :

docs := result.Group("MyApi/doc")

{

    docs.GET("/", getDocs) //recupérer tous les documents
    docs.POST("/addDoc", addDoc) //ajouter un document
    docs.GET("/getDocById/:id", getDocById) //recuperer un document par son Id
    docs.DELETE("/deleteDocById/:id", deleteDocById) //supprimer un document par son Id
}





