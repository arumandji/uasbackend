package achievements

import "go.mongodb.org/mongo-driver/bson/primitive"

type Achievement struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    StudentID   string             `bson:"student_id" json:"student_id"`
    Type        string             `bson:"type" json:"type"`
    Title       string             `bson:"title" json:"title"`
    Description string             `bson:"description" json:"description"`
    Details     any                `bson:"details,omitempty" json:"details"`
    Attachments []string           `bson:"attachments,omitempty" json:"attachments"`
}
