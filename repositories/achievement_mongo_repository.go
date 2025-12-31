package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AchievementMongoRepository struct {
	collection *mongo.Collection
}

func NewAchievementMongoRepository(col *mongo.Collection) *AchievementMongoRepository {
	return &AchievementMongoRepository{collection: col}
}

// CREATE
func (r *AchievementMongoRepository) Create(ctx context.Context, data interface{}) error {
	_, err := r.collection.InsertOne(ctx, data)
	return err
}

// FIND ALL
func (r *AchievementMongoRepository) FindAll(ctx context.Context) ([]bson.M, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

// FIND BY ID
func (r *AchievementMongoRepository) FindByID(ctx context.Context, id string) (bson.M, error) {
	var result bson.M
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	return result, err
}

// AGGREGATE STATS (SEDERHANA)
func (r *AchievementMongoRepository) AggregateStats(ctx context.Context) ([]bson.M, error) {
	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id":   "$kategori",
				"total": bson.M{"$sum": 1},
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []bson.M
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// UPDATE
func (r *AchievementMongoRepository) Update(ctx context.Context, data interface{}) error { {
	achievement, ok := data.(bson.M)
	if !ok {
		return mongo.ErrNilDocument
	}	
	id, ok := achievement["_id"].(string)
	if !ok {
		return mongo.ErrNilDocument
	}	
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": id}, achievement)
	return err
}
}

// DELETE
func (r *AchievementMongoRepository) Delete(ctx context.Context, id uint) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})			
	return err
}