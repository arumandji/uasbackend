package achievements

import (
    "context"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterAchievementRoutes(rg *gin.RouterGroup) {
    svc := NewService()
    rg.POST("/", func(c *gin.Context) {
        var req Achievement
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        id, err := svc.Create(c.Request.Context(), &req)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, gin.H{"id": id.Hex()})
    })

    rg.POST("/:id/submit", func(c *gin.Context) {
        id := c.Param("id")
        if err := svc.Submit(c.Request.Context(), id); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "submitted"})
    })

    rg.POST("/:id/verify", func(c *gin.Context) {
        id := c.Param("id")
        verifier, _ := c.Get("user_id")
        if verifier == nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "no verifier"})
            return
        }
        if err := svc.Verify(c.Request.Context(), id, verifier.(string)); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "verified"})
    })

    // upload attachment
    rg.POST("/:id/attachments", func(c *gin.Context) {
        id := c.Param("id")
        f, err := c.FormFile("file")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "file required"})
            return
        }
        uploadDir := os.Getenv("UPLOAD_DIR")
        _ = os.MkdirAll(uploadDir, 0755)
        dst := filepath.Join(uploadDir, f.Filename)
        if err := c.SaveUploadedFile(f, dst); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        // append url/path to Mongo doc
        oid, _ := primitive.ObjectIDFromHex(id)
        _, err = NewService().coll.UpdateOne(context.Background(),
            primitive.M{"_id": oid},
            primitive.M{"$push": primitive.M{"attachments": dst}})
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"attachment": dst})
    })
}
