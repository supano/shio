package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"time"
)

type Person struct {
	Name string
	LastName string
}

type Asset struct {
	Id  string
	AclTag string
	MerchantId string
	DescribeURL string
	Price int64
	Meta AssetMetadata
	CreatedAt time.Time
}

type AssetMetadata struct {
	Kind string
	CoverImageURL string
	Title string
	Description string
	Teaser string
	Slug string
	EventURL string
}

func main()  {
	ctx := context.Background()

	//err := os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:5545")
	//if err != nil {
	//	panic(err)
	//}

	c, err := datastore.NewClient(ctx, "catcat-development")
	if err != nil {
		fmt.Println(err)
		return
	}


	//user := &Person{
	//	Name:     "สมชาย",
	//	LastName: "ชาติเจริญ",
	//}
	//
	//kind := "Person"
	//id := "person1"
	//
	//Personkey := datastore.NameKey(kind, id, nil)
	//
	//rowKey, err := c.Put(ctx, Personkey, user)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(rowKey)

	var dest []Asset

	q := datastore.NewQuery("kind").Namespace("namespace")

	keys, err := c.GetAll(ctx, q, &dest)
	if err != nil {
		panic(err)
	}

	fmt.Println("%+v\n",keys)
	for _, v := range dest {
		fmt.Println("%+v\n",v)

	}

}