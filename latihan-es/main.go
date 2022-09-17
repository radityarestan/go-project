package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"
)

type Student struct {
	Name         string  `json:"name"`
	Age          int     `json:"age"`
	AverageScore float64 `json:"average_score"`
}

func insertData(esclient *elastic.Client, ctx context.Context) {
	//creating student object
	newStudent := Student{
		Name:         "YAOS doe",
		Age:          10,
		AverageScore: 99.9,
	}

	_, err := esclient.Index().
		Index("students").
		BodyJson(newStudent).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful")
}

func bulk(esclient *elastic.Client, ctx context.Context) {
	//creating student object
	newStudents := []Student{
		{
			Name:         "Janaqy doe",
			Age:          10,
			AverageScore: 99.9,
		},
		{
			Name:         "Hajita doe",
			Age:          10,
			AverageScore: 99.9,
		},
	}

	bulk := esclient.Bulk()
	for _, student := range newStudents {
		req := elastic.NewBulkIndexRequest().Index("students").Doc(student)
		bulk.Add(req)
	}

	_, err := bulk.Do(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("[Elastic][InsertProducts]Insertion Successful")

}

func search(esclient *elastic.Client, ctx context.Context) {
	var students []Student

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("name", "qy"))

	/* this block will basically print out the es query */
	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	}
	fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))
	/* until this block */

	searchService := esclient.Search().Index("students").SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var student Student
		err := json.Unmarshal(hit.Source, &student)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}

		students = append(students, student)
	}

	if err != nil {
		fmt.Println("Fetching student fail: ", err)
	} else {
		for _, s := range students {
			fmt.Printf("Student found Name: %s, Age: %d, Score: %f \n", s.Name, s.Age, s.AverageScore)
		}
	}
}

func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err
}

func main() {
	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error creating the client: ", err)
	}

	// insertData(esclient, ctx)
	// bulk(esclient, ctx)
	search(esclient, ctx)
}
