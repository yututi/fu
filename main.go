package main

import (
	"io"
	"os"
	"path"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"google.golang.org/api/iterator"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var (
	// bucket name of google cloud storage.
	bucketName string
)

func main() {

	bucketName = os.Getenv("BUCKET_NAME")

	r := gin.Default()

	r.LoadHTMLGlob("ginTemplates/*")

	r.Static("/uploader", "./vueDist")

	r.POST("/file/upload", func(c *gin.Context) {

		ctx := appengine.NewContext(c.Request)

		f, fh, err := c.Request.FormFile("file")

		client, err := storage.NewClient(ctx)
		if err != nil {
			c.JSON(501, gin.H{"error": err.Error()})
			return
		}

		fname := uuid.Must(uuid.NewV4()).String() + path.Ext(fh.Filename)

		o := client.Bucket(bucketName).Object(fname)

		// if err := o.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		// 	c.JSON(501, gin.H{"error": err.Error()})
		// 	return
		// }

		w := o.NewWriter(ctx)
		w.ContentType = fh.Header.Get("Content-Type")
		w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}

		// w.CacheControl = "public, max-age=86400"

		if _, err := io.Copy(w, f); err != nil {
			c.JSON(501, gin.H{"error": err.Error()})
			return
		}

		if err := w.Close(); err != nil {
			log.Warningf(ctx, err.Error())
		}

		c.JSON(200, gin.H{
			"object": fname,
			"bucket": bucketName,
		})
	})

	r.GET("/media/:name", func(c *gin.Context) {
		name := c.Param("name")

		client, err := storage.NewClient(c.Request.Context())
		if err != nil {
			c.String(501, err.Error())
			return
		}
		obj := client.Bucket(bucketName).Object(name)

		attrs, err := obj.Attrs(c.Request.Context())
		if err != nil {
			c.String(404, err.Error())
			return
		}

		c.HTML(200, "image.tmpl", gin.H{
			"mediaLink": attrs.MediaLink,
		})
	})

	r.GET("/media", func(c *gin.Context) {

		ctx := appengine.NewContext(c.Request)

		client, err := storage.NewClient(ctx)
		if err != nil {
			c.String(501, err.Error())
			return
		}

		it := client.Bucket(bucketName).Objects(ctx, nil)

		olist := []string{}
		for {
			o, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				c.String(501, err.Error())
				return
			}
			olist = append(olist, o.Name)
		}

		c.JSON(200, gin.H{
			"bucket":  bucketName,
			"objects": olist,
		})
	})

	// history api fallback
	r.NoRoute(func(c *gin.Context) {
		c.File("/vueDist/index.html")
	})

	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/uploader")
	})

	r.Run()
}
