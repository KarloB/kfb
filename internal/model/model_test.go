package model

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestSampleDB(t *testing.T) {
	f := FoodDB{
		Good: []Food{
			{Name: "Pljeskavica", Where: "Zdravljak", Why: "Massive and tasty"},
			{Name: "Zagrebacki", Where: "Tramontana", Why: "Love the smell of burned plastic in the morning"},
			{Name: "Rucak vulgaris", Where: "Dioniz", Why: "Cheap and sort of ok"},
		},
		Bad: []Food{
			{Name: "Pizza", Where: "OhYeah", Why: "Loaf of bread with 2 snjitas of salama"},
		},
	}

	b, err := json.Marshal(f)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("db.json", b, 0644)
	if err != nil {
		t.Fatal(err)
	}
}
