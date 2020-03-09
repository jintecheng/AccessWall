package routes

import (
	"context"
	"log"
	"strings"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

type FileStruct struct {
	Name  string   `json:"name" bson:"name"`
	User  string   `json:"user" bson:"user"`
	Guest []string `json:"guest" bson:"guest"`
}

type FileFindStruct struct {
	Name  string   `json:"name" bson:"name"`
	Guest []string `json:"guest" bson:"guest"`
}

func connDB() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), "mongodb://"+dbIp+":"+dbPort)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, err
}

func disConnDB(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	return err
}

func insertFileDBOne(client *mongo.Client, file *FileStruct) error {
	collection := client.Database("hqbfs").Collection("file")

	_, err := collection.InsertOne(context.TODO(), file)
	return err
}

//[]interface{}{FileStruct, FileStruct ...}
func insertFileDBMany(client *mongo.Client, files []interface{}) error {
	collection := client.Database("hqbfs").Collection("file")

	_, err := collection.InsertMany(context.TODO(), files)
	return err
}

func updateFileDB(client *mongo.Client, file *FileStruct) error {
	collection := client.Database("hqbfs").Collection("file")
	filter := bson.M{"name": file.Name, "user": file.User}
	update := bson.D{
		{"$set", bson.D{
			{"guest", file.Guest},
		}},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)

	return err
}

func findLinkFileDB(client *mongo.Client, file *FileStruct) ([]*FileStruct, error) {
	var results []*FileStruct
	var err error
	collection := client.Database("hqbfs").Collection("file")
	findOptions := options.Find()
	findOptions.SetLimit(2)
	//filename := strings.Replace(file.Name, userDir, "", -1)
	temp := strings.Split(file.Name, "/")
	filename := temp[len(temp)-1]
	filename = strings.Replace(filename, "/", "", -1)
	filter := bson.M{"name": filename, "guest": file.Guest[0]}

	cur, err := collection.Find(context.TODO(), filter, findOptions)
	for cur.Next(context.TODO()) {
		var fs FileStruct
		err = cur.Decode(&fs)
		results = append(results, &fs)
	}
	err = cur.Err()

	cur.Close(context.TODO())

	return results, err
}

func findFileDB(client *mongo.Client, file *FileStruct) ([]*FileStruct, error) {
	var results []*FileStruct
	var err error
	collection := client.Database("hqbfs").Collection("file")
	findOptions := options.Find()
	findOptions.SetLimit(2)
	filename := strings.Replace(file.Name, userDir, "", -1)
	filename = strings.Replace(filename, "/", "", -1)
	filter := bson.M{"name": filename, "user": gUser.Email}

	cur, err := collection.Find(context.TODO(), filter, findOptions)
	for cur.Next(context.TODO()) {
		var fs FileStruct
		err = cur.Decode(&fs)
		results = append(results, &fs)
	}
	err = cur.Err()

	cur.Close(context.TODO())

	return results, err
}

func deleteFileDB(client *mongo.Client, file *FileStruct) (bool, error) {
	collection := client.Database("hqbfs").Collection("file")
	filter := bson.D{{"name", file.Name}, {"user", file.User}}
	_, err := collection.DeleteOne(context.TODO(), filter, nil)

	if err != nil {
		return false, err
	}
	return true, err
}
