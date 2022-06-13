package controller

import (
	"allright-tiktok/models"
	"allright-tiktok/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	//获取视频id
	video_id := c.Query("video_id")

	action_type := c.Query("action_type")
	//判断用户登录
	userClaims, err := utils.AnalyseToken(token)
	//如果鉴权失败，则报错返回
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	var user models.User
	if rowsAffected := db.Where("username = ?", userClaims.Username).First(&user).RowsAffected; rowsAffected != 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
		var favorite models.Favorite
		if action_type == "1" {	
			fmt.Println(1)		
			err := db.Where("video_id = ? and user_id = ?", video_id, strconv.FormatInt(user.Id,10)).First(&favorite).Error
			if err != nil {
				newFavorite := models.Favorite{
					ViedoId: video_id,
					UserId: strconv.FormatInt(user.Id,10),
				}
				db.Create(&newFavorite)
			}else{
				db.Model(&models.Favorite{}).Where("video_id = ? and user_id = ?", video_id, strconv.FormatInt(user.Id,10)).Update("delete", 0)
			}
		}
		if action_type == "2" {			
			db.Model(&models.Favorite{}).Where("video_id = ? and user_id = ?", video_id, strconv.FormatInt(user.Id,10)).Update("delete", 1)
		}
		
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}




}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {

	//获取用户鉴权token
	// token := c.Query("token")
	//获取用户id

	id := c.Query("user_id")

	var user models.User
	if rowsAffected := db.Where("id = ?", id).First(&user).RowsAffected; rowsAffected == 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	// var favoriteList []models.Favorite //创建一个Video数组存起来
	// db.Where("user_id = ? and delete = ?", id, 0).Find(&favoriteList)
	// fmt.Print(favoriteList)
	var videoList []models.Video
	db.Table("favorite").Select("favorite.video_id, video.*").
	Where("favorite.user_id = ? and favorite.delete = ?", id, 0).
	Joins("LEFT JOIN video ON favorite.video_id = video.id").
	Find(&videoList)
	videoPoList := make([]Video, len(videoList))
	for i := range videoList {
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
