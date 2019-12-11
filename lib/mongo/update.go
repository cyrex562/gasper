package mongo

import (
	"context"
	"time"

	"github.com/sdslabs/gasper/types"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateOne updates a document in the mongoDB collection
func UpdateOne(collectionName string, filter types.M, data interface{}, option *options.FindOneAndUpdateOptions) error {
	collection := link.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOneAndUpdate(ctx, filter, types.M{"$set": data}, option)
	return res.Err()
}

// UpdateInstance is an abstraction over UpdateOne which updates an application in mongoDB
func UpdateInstance(filter types.M, data interface{}) error {
	return UpdateOne(InstanceCollection, filter, data, nil)
}

// UpsertInstance is an abstraction over UpdateOne which updates an application in mongoDB
// or inserts it if the corresponding document doesn't exist
func UpsertInstance(filter types.M, data interface{}) error {
	return UpdateOne(InstanceCollection, filter, data, options.FindOneAndUpdate().SetUpsert(true))
}

// UpdateUser is an abstraction over UpdateOne which updates an application in mongoDB
func UpdateUser(filter types.M, data interface{}) error {
	return UpdateOne(UserCollection, filter, data, nil)
}

// UpsertUser is an abstraction over UpdateOne which updates an application in mongoDB
// or inserts it if the corresponding document doesn't exist
func UpsertUser(filter types.M, data interface{}) error {
	return UpdateOne(UserCollection, filter, data, options.FindOneAndUpdate().SetUpsert(true))
}

// UpsertMetrics is an abstraction over UpdateOne which updates an metrics document in mongoDB
// or inserts it if the corresponding document doesn't exist
func UpsertMetrics(filter types.M, data interface{}) error {
	return UpdateOne(MetricsCollection, filter, data, options.FindOneAndUpdate().SetUpsert(true))
}

// UpdateMany updates multiple documents in the mongoDB collection
func UpdateMany(collectionName string, filter types.M, data interface{}) (interface{}, error) {
	collection := link.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.UpdateMany(ctx, filter, types.M{"$set": data}, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateInstances is an abstraction over UpdateMany which updates multiple applications in mongoDB
func UpdateInstances(filter types.M, data interface{}) (interface{}, error) {
	return UpdateMany(InstanceCollection, filter, data)
}
