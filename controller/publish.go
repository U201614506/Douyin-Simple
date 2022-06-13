package controller

import (
	"allright-tiktok/models"
	"allright-tiktok/utils"
	"fmt"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	title := c.PostForm("title")

	var user models.User

	userClaims, err := utils.AnalyseToken(token)
	//如果鉴权失败，则报错返回
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	if rowsAffected := db.Where("username = ?", userClaims.Username).First(&user).RowsAffected; rowsAffected == 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}


	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	//利用ffmpeg找第1帧作为封面
	cmd := exec.Command(
		"ffmpeg", "-i", "../public/" + finalName,
		"-vf", "select=eq(n\\, 0)", "-vframes", "1",
		"../public/image/"+ strings.Replace(finalName, ".mp4", ".jpg", 1),
	)
	if err := cmd.Run(); err != nil {
		fmt.Println("视频封面截取失败")
	}
	playUrl := "/static/" + finalName
	coverUrl := "/static/image/" + strings.Replace(finalName, ".mp4", ".jpg", 1) + ".jpg"
	video :=  models.Video{
		Author:   user.UserName,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    title,
	}

	db.Select("Author", "PlayUrl", "CoverUrl", "Title").Create(&video)
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {

	//获取用户鉴权token
	// token := c.Query("token")
	//获取用户id


	id := c.Query("user_id")

	var user models.User
	if rowsAffected := db.Where("id = ?", id).First(&user).RowsAffected; rowsAffected == 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	var videoList []models.Video //创建一个Video数组存起来
	db.Where("author = ?", user.UserName).Find(&videoList)
	fmt.Print(videoList)
	videoPoList := make([]Video, len(videoList))
	for i := range videoList {
		fmt.Println(i)
		videoPoList[i].CoverUrl = "http://172.33.170.41:8080" + videoList[i].CoverUrl
		videoPoList[i].PlayUrl = "http://172.33.170.41:8080" + videoList[i].PlayUrl
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videoPoList,
	})
}
