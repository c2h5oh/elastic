// Copyright 2012-2014 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"testing"
)

func TestCount(t *testing.T) {
	client := setupTestClientAndCreateIndex(t)

	tweet1 := tweet{User: "olivere", Message: "Welcome to Golang and Elasticsearch."}
	tweet2 := tweet{User: "olivere", Message: "Another unrelated topic."}
	tweet3 := tweet{User: "sandrae", Message: "Cycling is fun."}

	// Add all documents
	_, err := client.Index().Index(testIndexName).Type("tweet").Id("1").BodyJson(&tweet1).Do()
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Index().Index(testIndexName).Type("tweet").Id("2").BodyJson(&tweet2).Do()
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Index().Index(testIndexName).Type("tweet").Id("3").BodyJson(&tweet3).Do()
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Flush().Index(testIndexName).Do()
	if err != nil {
		t.Fatal(err)
	}

	// Count documents
	count, err := client.Count(testIndexName).Do()
	if err != nil {
		t.Fatal(err)
	}
	if count != 3 {
		t.Errorf("expected Count = %d; got %d", 3, count)
	}

	// Count documents
	count, err = client.Count(testIndexName).Type("tweet").Do()
	if err != nil {
		t.Fatal(err)
	}
	if count != 3 {
		t.Errorf("expected Count = %d; got %d", 3, count)
	}

	// Count documents
	count, err = client.Count(testIndexName).Type("gezwitscher").Do()
	if err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Errorf("expected Count = %d; got %d", 0, count)
	}

	// Count with query
	query := NewTermQuery("user", "olivere")
	count, err = client.Count(testIndexName).Query(query).Do()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Errorf("expected Count = %d; got %d", 2, count)
	}

	// Count with query and type
	query = NewTermQuery("user", "olivere")
	count, err = client.Count(testIndexName).Type("tweet").Query(query).Do()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Errorf("expected Count = %d; got %d", 2, count)
	}
}
