package repositories

import (
	"context"
	"errors"
	"time"
	"uas_backend/database"
	"uas_backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AchievementMongoRepository interface {
	Create(a *models.Achievement) (string, error)
	FindByID(id string) (*models.Achievement, error)
	Update(id string, update bson.M) error
	SoftDelete(id string) error
	ListByMahasiswaIDs(mahasiswaIDs []string, offset, limit int64) ([]models.Achievement, error)
	AggregateStats() (bson.M, error)
}

type achMongoRepo struct {
	coll *database.MongoCollection // note: we will access database.MongoDB in impl
}

func NewAchievementMongoRepository() AchievementMongoRepository {
	coll := database.MongoDB.Collection("achievements")
	return &achMongoRepo{coll: coll}
}

func (r *achMongoRepo) Create(a *models.Achievement) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	a.CreatedAt = time.Now()
	a.UpdatedAt = a.CreatedAt
	res, err := r.coll.InsertOne(ctx, a)
	if err != nil {
		return "", err
	}
	oid := res.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil
}

func (r *achMongoRepo) FindByID(id string) (*models.Achievement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var a models.Achievement
	if err := r.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&a); err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *achMongoRepo) Update(id string, update bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.coll.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": update})
	return err
}

func (r *achMongoRepo) SoftDelete(id string) error {
	// We'll mark deleted by setting a field "deleted": true
	return r.Update(id, bson.M{"deleted": true, "updatedAt": time.Now()})
}

func (r *achMongoRepo) ListByMahasiswaIDs(mahasiswaIDs []string, offset, limit int64) ([]models.Achievement, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"mahasiswaId": bson.M{"$in": mahasiswaIDs}, "deleted": bson.M{"$ne": true}}
	opts := options.Find().SetSkip(offset).SetLimit(limit).SetSort(bson.M{"createdAt": -1})
	cursor, err := r.coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	var out []models.Achievement
	if err := cursor.All(ctx, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *achMongoRepo) AggregateStats() (bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := mongoPipelineForStats()
	cursor, err := r.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	var result []bson.M
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return bson.M{}, nil
	}
	return result[0], nil
}

func mongoPipelineForStats() mongo.Pipeline {
	return mongo.Pipeline{
		{{"$match", bson.D{{"deleted", bson.D{{"$ne", true}}}}}},
		{{"$group", bson.D{{"_id", "$achievementType"}, {"count", bson.D{{"$sum", 1}}}}}},
	}
}
