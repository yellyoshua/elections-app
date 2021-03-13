package repository

import (
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Demo struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name  string             `bson:"name" json:"name"`
	Value string             `bson:"value" json:"value"`
}

var samples = []Demo{
	{Name: "Yoshua1", Value: "Lopez1"},
	{Name: "Yoshua2", Value: "Lopez2"},
	{Name: "Yoshua3", Value: "Lopez3"},
}

// TODO: Test UpdateMany repository func!
func TestRepository(t *testing.T) {
	os.Setenv("DATABASE_NAME", "golangtest")
	os.Setenv("DATABASE_URI", "mongodb://root:dbpwd@localhost:27017")
	var collection string = "demo_no_indexes"
	var indexes bool = false

	Initialize(indexes)
	repo := NewRepository(collection)
	loadSampleData(t, repo, samples)

	var sample1 Demo
	var sample1ByID Demo
	var sample2 Demo
	var sample2ByID Demo
	var sample3 Demo
	var sample3ByID Demo
	var sampleNotFound Demo
	var samplesDatabase []Demo

	repo.FindOne(bson.M{"name": "Yoshua1"}, &sample1)
	repo.FindByID(sample1.ID, &sample1ByID)
	repo.FindOne(bson.M{"name": "Yoshua2"}, &sample2)
	repo.FindByID(sample2.ID, &sample2ByID)
	repo.FindOne(bson.M{"name": "Yoshua3"}, &sample3)
	repo.FindByID(sample3.ID, &sample3ByID)
	repo.FindOne(bson.M{"name": "YoshuaNo"}, &sampleNotFound)

	if sample1.Name != samples[0].Name {
		t.Fatal("Sample1 not found")
	}

	if sample1ByID.Name != samples[0].Name {
		t.Fatal("Sample1 not found by ID")
	}

	if sample2.Name != samples[1].Name {
		t.Fatal("sample2 not found")
	}

	if sample2ByID.Name != samples[1].Name {
		t.Fatal("Sample2 not found by ID")
	}

	if sample3.Name != samples[2].Name {
		t.Fatal("Sample3 not found")
	}

	if sample3ByID.Name != samples[2].Name {
		t.Fatal("Sample3 not found by ID")
	}

	if len(sampleNotFound.Name) != 0 {
		t.Fatal("SampleNotFound has founded ???")
	}

	repo.Find(bson.D{}, &samplesDatabase)

	if len(samplesDatabase) != len(samples) {
		t.Fatalf("Samples length not equal to samples length alocated in the database %v - %v", len(samplesDatabase), len(samples))
	}

	sample1Update := sample1
	sample1Update.Name = "UpdatedYoshua1"
	sample1Filter := bson.M{"_id": sample1.ID}
	updateSample1 := map[string]interface{}{"name": sample1Update.Name}
	if err := repo.UpdateOne(sample1Filter, updateSample1); err != nil {
		t.Fatalf("Error updating sample1 %v", err)
	}
	if repo.FindOne(bson.M{"name": sample1Update.Name}, &sample1); sample1.Name != sample1Update.Name {
		t.Fatal("Sample1 not updated")
	}

	sample2Update := sample2
	sample2Update.Name = "UpdatedYoshua2"
	sample2Filter := bson.M{"_id": sample2.ID}
	updateSample2 := map[string]interface{}{"name": sample2Update.Name}
	if err := repo.UpdateOne(sample2Filter, updateSample2); err != nil {
		t.Fatalf("Error updating sample2 %v", err)
	}
	if repo.FindOne(bson.M{"name": sample2Update.Name}, &sample2); sample2.Name != sample2Update.Name {
		t.Fatal("Sample2 not updated")
	}

	sample3Update := sample3
	sample3Update.Name = "UpdatedYoshua3"
	sample3Filter := bson.M{"_id": sample3.ID}
	updateSample3 := map[string]interface{}{"name": sample3Update.Name}
	if err := repo.UpdateOne(sample3Filter, updateSample3); err != nil {
		t.Fatalf("Error updating sample3 %v", err)
	}
	if repo.FindOne(bson.M{"name": sample3Update.Name}, &sample3); sample3.Name != sample3Update.Name {
		t.Fatal("Sample3 not updated")
	}

	if err := repo.UpdateOne(map[string]interface{}{"name": "SampleNotFound"}, map[string]interface{}{"name": "SampleNotFound"}); err.Error() != "No matched documents" {
		t.Fatal("Error should not returned a error ???")
	}
}

func loadSampleData(t *testing.T, repo Repository, samples []Demo) {
	var samplesInterface []interface{}

	for _, sample := range samples {
		samplesInterface = append(samplesInterface, sample)
	}

	if err := repo.Drop(); err != nil {
		t.Fatalf("Err droping collection %v", err)
	}

	if err := repo.InsertMany(samplesInterface); err != nil {
		t.Fatalf("Err inserting many samples %v", err)
	}

	if err := repo.Drop(); err != nil {
		t.Fatalf("Err droping collection %v", err)
	}

	for _, sample := range samplesInterface {
		if _, err := repo.InsertOne(sample); err != nil {
			t.Fatalf("Err insert sample %v", err)
		}
	}
}
