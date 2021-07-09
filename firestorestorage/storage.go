package firestorestorage

import (
	"firebase.google.com/go"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
	"context"
	"google.golang.org/api/iterator"
	"fmt"
)


type Person struct {
	Name string `json:"name", omitempty`
	Id string `json:"id", omitempty`
}

var(
	filename string = "userbase-c4c78-firebase-adminsdk-8veu9-5e6fd0f3d3.json"
)

type FirebaseOperations interface {
	Add(user Person, client *firestore.Client) (*firestore.DocumentRef, *firestore.WriteResult, error)
	Find(name_tofind, collection_name string, client *firestore.Client)(map[string]interface{}, error)
	CloseConnection(client *firestore.Client)
	NewClient(project_name string) (*firestore.Client, error)
}

func (p *Person) Add(user Person, client *firestore.Client) (*firestore.DocumentRef, *firestore.WriteResult, error) {
		
	doc , result , err := client.Collection("Users").Add(context.Background(),map[string]interface{}{
		"name" : user.Name,
		"id" : user.Id,
	})
	if err != nil {
		fmt.Println("Collection fail:", err)
		return doc, result, err
	}
	return doc, result, nil	
}

func (p *Person) Find(name_tofind, collection_name string, client *firestore.Client)(map[string]interface{}, error) { 
	var user_data map[string]interface{} 
	var errs error
	it := client.Collection(collection_name).Documents(context.Background())
	for{
		doc , err := it.Next()
			if err == iterator.Done{
				break
			}
			user_data = doc.Data()
			errs = err
	}
	return user_data, errs
}
func (p *Person) NewClient(project_name string) (*firestore.Client, error) {
	
	firebase_client, err := firebase.NewApp(context.Background(),&firebase.Config{ProjectID: project_name},option.WithCredentialsFile("credentials/"+filename)); if err != nil {
		return &firestore.Client{}, err
	}

	client, err := firebase_client.Firestore(context.Background());if err != nil {
		return &firestore.Client{}, err
	}
	return client, err
}
func NewInstance() FirebaseOperations{return &Person{}}

func (p *Person) CloseConnection(client *firestore.Client){ defer client.Close()}