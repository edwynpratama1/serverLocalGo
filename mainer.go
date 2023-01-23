package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

// type iPassportData struct {
// 	ID      primitive.ObjectID `bson:"_id,omitempty"`
// 	CifCode string             `bson:"cifCode,omitempty"`
// 	Data    bson.A             `bson:"data,inline"`
// }

// func main() {
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	filter := bson.D{{"cifCode", "2210319"}}

//		var result iPassportData
//		err = client.Database("mongosimasbankingprod").Collection("iPassportData").FindOne(context.TODO(), filter).Decode(&result)
//		if err != nil {
//			fmt.Println("Error calling FindOne():", err)
//		} else {
//			fmt.Println("FindOne() result:", result)
//		}
//	}
func main() {
	db, err := sql.Open("godror", "ssldb/ssldb123@10.22.103.79:1521/ORCL")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select name from target_account ta where id = 226 ")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var cif string

	for rows.Next() {

		rows.Scan(&cif)
	}
	fmt.Printf("cif : %s\n", cif)
}
