package main




import (
	."github.com/ali2210/verbose-octo-palm-tree/firestorestorage"
	"fmt"
)

var(
	firestore_client Person = Person{}
)

func main() {
	client , err := firestore_client.NewClient("userbase-c4c78"); if err != nil {
		fmt.Println("Connection not secure:", err)
		return 
	}
	fmt.Println("Client:", client)

	doc, result, err:= firestore_client.Add(Person{Id: "0", Name: "Ali"}, client);if err != nil {
		fmt.Println("No Collection :", err)
		return 
	}

	fmt.Println("Doc_Ref:", doc, "Result:", result)

	dataset, err := firestore_client.Find(doc.ID, "Users", client); if err != nil {
		fmt.Println("No Collection exist:", err)
		return
	} 
	fmt.Println("DataSet:", dataset)
	firestore_client.CloseConnection(client)
}