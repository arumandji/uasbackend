package achievements

import (
    "context"
    "fmt"
    "time"

    "prestasi-api/database"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
    coll *mongo.Collection
    pg   *database.DBType // placeholder - we'll use database.DB directly
}

func NewService() *Service {
    return &Service{
        coll: database.MongoDB.Collection("achievements"),
    }
}

func (s *Service) Create(ctx context.Context, a *Achievement) (primitive.ObjectID, error) {
    res, err := s.coll.InsertOne(ctx, a)
    if err != nil {
        return primitive.NilObjectID, err
    }
    id := res.InsertedID.(primitive.ObjectID)
    // create reference in Postgres
    _, err = database.DB.Exec(`
        INSERT INTO achievement_references(student_id, mongo_achievement_id, status)
        VALUES ($1, $2, 'draft')`, a.StudentID, id.Hex())
    if err != nil {
        // best-effort: rollback mongo insert (simple compensating action)
        s.coll.DeleteOne(ctx, bson.M{"_id": id})
        return primitive.NilObjectID, fmt.Errorf("pg insert err: %w", err)
    }
    return id, nil
}

func (s *Service) Submit(ctx context.Context, mongoID string) error {
    _, err := database.DB.Exec(`
        UPDATE achievement_references SET status='submitted', submitted_at=NOW()
        WHERE mongo_achievement_id=$1`, mongoID)
    return err
}

func (s *Service) Verify(ctx context.Context, mongoID, verifierID string) error {
    _, err := database.DB.Exec(`
        UPDATE achievement_references SET status='verified', verified_at=NOW(), verified_by=$1
        WHERE mongo_achievement_id=$2`, verifierID, mongoID)
    return err
}
