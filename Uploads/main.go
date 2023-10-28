package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"github.com/gin-gonic/gin"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
)

type Configuration struct {
	LogFileName   string
	Port          string
	UploadFolder  string
	MaxFileSizeMB int
}

type FileData struct {
	ID         string    `json:"id"`
	FileName   string    `json:"file_name"`
	SizeBytes  int64     `json:"size_bytes"`
	UploadDate time.Time `json:"upload_date"`
}

func generateSecureHexDec() string {
	currentNanoTime := time.Now().UnixNano()
	pid := os.Getpid()
	uniqueValue := uint64(currentNanoTime) << 16 | uint64(pid&0xFFFF)
	uniqueBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(uniqueBytes, uniqueValue)
	secureHex := hex.EncodeToString(uniqueBytes)
	return secureHex
}

func main() {
	config := Configuration{
		LogFileName:   "upload_log.json",
		Port:          "8099",
		UploadFolder:  "uploads",
		MaxFileSizeMB: 16,

	}

	router := gin.Default()

	maxFileSize := int64(config.MaxFileSizeMB) << 20
	router.MaxMultipartMemory = maxFileSize


router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Index of uploads api")
	})



	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if file.Size > maxFileSize {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds the maximum allowed size"})
			return
		}

		secureHex := generateSecureHexDec()
		fileExt := filepath.Ext(file.Filename)
		newFileName := secureHex + fileExt

		if err := c.SaveUploadedFile(file, filepath.Join(config.UploadFolder, newFileName)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fileData := FileData{
			ID:         secureHex,
			FileName:   newFileName,
			SizeBytes:  file.Size,
			UploadDate: time.Now(),
		}

		logFile, logErr := os.OpenFile(config.LogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if logErr != nil {
			log.Fatal(logErr)
		}
		defer logFile.Close()

		encoder := json.NewEncoder(logFile)
		encodeErr := encoder.Encode(fileData)
		if encodeErr != nil {
			log.Fatal(encodeErr)
		}

		c.JSON(http.StatusOK, fileData)
	})

	router.Static("/uploads", config.UploadFolder)

/*	router.GET("/info", func(c *gin.Context) {
		fileID := c.DefaultQuery("fid", "") // Get the "fid" query parameter

		if fileID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File ID is required"})
			return
		}

		logFile, logErr := os.Open(config.LogFileName)
		if logErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": logErr.Error()})
			return
		}
		defer logFile.Close()

		decoder := json.NewDecoder(logFile)
		for decoder.More() {
			var fileData FileData
			if decodeErr := decoder.Decode(&fileData); decodeErr == nil {
				if fileData.ID == fileID {
					c.JSON(http.StatusOK, fileData)
					return
				}
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
	})*/

router.GET("/info", func(c *gin.Context) {
    fileID := c.DefaultQuery("file", "") // Get the "file" query parameter

    if fileID == "" {
        // If no specific file is requested, return all lists
        logFile, logErr := os.Open(config.LogFileName)
        if logErr != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": logErr.Error()})
            return
        }
        defer logFile.Close()

        var allFiles []FileData
        decoder := json.NewDecoder(logFile)
        for decoder.More() {
            var fileData FileData
            if decodeErr := decoder.Decode(&fileData); decodeErr == nil {
                allFiles = append(allFiles, fileData)
            }
        }
        c.JSON(http.StatusOK, allFiles)
        return
    }

    // If a specific file is requested, return information for that file
    logFile, logErr := os.Open(config.LogFileName)
    if logErr != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": logErr.Error()})
        return
    }
    defer logFile.Close()

    decoder := json.NewDecoder(logFile)
    for decoder.More() {
        var fileData FileData
        if decodeErr := decoder.Decode(&fileData); decodeErr == nil {
            if fileData.ID == fileID {
                c.JSON(http.StatusOK, fileData)
                return
            }
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
})




	router.Run(":" + config.Port)
}

