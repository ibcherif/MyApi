package appDoc_test

import (
	"cherif.com/myApi/appDoc"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//initialisation des variables essentiels pour les tests
var (
	server           *httptest.Server
	reader           io.Reader
	docUrl, docUrlId string
)

// initialisation de la connexion et création des documents
func init() {
	// Ouverture de la connexion vers la BDD SQLite
	db := appDoc.InitDb()
	// Fermeture de la connexion vers la BDD SQLite
	defer db.Close()

	var doc appDoc.Doc

	// Suppression de la table
	db.DropTable(doc)
	// Création de la table
	db.CreateTable(doc)

	// Création des documents
	db.Create(&appDoc.Doc{Id: 1, Name: "BackEnd", Description: "C'est moi la logique"})
	db.Create(&appDoc.Doc{Id: 2, Name: "FrontEnd", Description: "C'est moi votre interlocuteur"})
	db.Create(&appDoc.Doc{Id: 3, Name: "DataBase", Description: "C'est moi votre dataDB"})

	// Démarrage du serveur HTTP
	server = httptest.NewServer(appDoc.Handlers())

	// URL sans paramêtre
	docUrl = server.URL + "/MyApi/doc"
	//URL avec parametre
	docUrlId = "/2"
}

// Tester l'ajout d'un document
func TestAddDoc(t *testing.T) {
	// Contenu à soumettre
	docJson := `{"Id":4, "Name": "Ibrahima Cherif", "Description":"I Like coding"}`

	// Contenu à soumettre au bon format
	reader = strings.NewReader(docJson)

	// Déclaration de la requête : type, URL, contenu
	request, err := http.NewRequest("POST", docUrl+"/addDoc", reader)
	// Requête de type JSON
	request.Header.Set("Content-Type", "application/json")

	// Exécution de la requête
	response, err := http.DefaultClient.Do(request)

	// Erreur si route inacessible
	if err != nil {
		t.Error(err)
	}
	// Erreur si code HTTP différent de 201
	if response.StatusCode != 201 {
		t.Errorf("Success expected: %d", response.StatusCode)
	}
}

//Tester la récupération de tous les documents
func TestGetDocs(t *testing.T) {
	// Contenu à soumettre
	reader = strings.NewReader("")

	// Déclaration de la reqûête : type, URL, contenu
	request, err := http.NewRequest("GET", docUrl+"/", reader)

	// Exécution de la requête
	response, err := http.DefaultClient.Do(request)

	// Erreur si route inacessible
	if err != nil {
		t.Error(err)
	}

	// Erreur si code HTTP différent de 200
	if response.StatusCode != 200 {
		t.Errorf("Success expected: %d", response.StatusCode)
	}
}

//Tester la récupération d'un document en utilisant l'id
func TestGetDocById(t *testing.T) {
	// Contenu à soumettre vide
	reader = strings.NewReader("")

	// Déclaration de la requête : type, URL, contenu
	request, err := http.NewRequest("GET", docUrl+"/getDocById"+docUrlId, reader)

	// Exécution de la requête
	response, err := http.DefaultClient.Do(request)

	// Erreur si route inacessible
	if err != nil {
		t.Error(err)
	}

	// Erreur si code HTTP différent de 200
	if response.StatusCode != 200 {
		t.Errorf("Success expected: %d", response.StatusCode)
	}
}

//Tester la suppression d'un document en utilisant l'id
func TestDeleteDoc(t *testing.T) {
	// Contenu à soumettre vide
	reader = strings.NewReader("")

	// Déclaration de la requête : type, URL, contenu
	request, err := http.NewRequest("DELETE", docUrl+"/deleteDocById"+docUrlId, reader)

	// Exécution de la requête
	response, err := http.DefaultClient.Do(request)

	// Erreur si route inacessible
	if err != nil {
		t.Error(err)
	}

	// Erreur si code HTTP différent de 200
	if response.StatusCode != 200 {
		t.Errorf("Success expected: %d", response.StatusCode)
	}
}
