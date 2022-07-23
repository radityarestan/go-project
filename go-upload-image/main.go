package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", routeIndexGet)
	router.POST("/process", routeSubmitPost)

	fmt.Println("server started at localhost:9000")
	router.Run(":9000")
}

func routeIndexGet(c *gin.Context) {
	var tmpl = template.Must(template.ParseFiles("view.html"))
	var err = tmpl.Execute(c.Writer, nil)

	if err != nil {
		c.Error(err)
	}
}

func routeSubmitPost(c *gin.Context) {
	// if err := c.Request.ParseMultipartForm(1024); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// ...

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	files := form.File["files"]
	var wg sync.WaitGroup
	for i, file := range files {
		wg.Add(1)
		go func(file *multipart.FileHeader, i int) {
			defer wg.Done()

			uploadedFile, err := file.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			defer uploadedFile.Close()

			fileLocation := filepath.Join(dir, "files", file.Filename)
			targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			defer targetFile.Close()

			if _, err := io.Copy(targetFile, uploadedFile); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("File %d successfully uploaded", i)})

		}(file, i)
	}

	wg.Wait()

	// uploadedFile, handler, err := c.Request.FormFile("file")
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// defer uploadedFile.Close()

	// dir, err := os.Getwd()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// fileLocation := filepath.Join(dir, "files", handler.Filename)
	// targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// defer targetFile.Close()

	// if _, err := io.Copy(targetFile, uploadedFile); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
