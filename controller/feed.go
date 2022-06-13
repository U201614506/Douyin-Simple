package controller

import (
	"allright-tiktok/models"
	"allright-tiktok/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// lastTime := c.Query("latest_time")

	token := c.Query("token")
	var user models.User
	//判断登录状态
	isVisitor := true
	if token != "" {
		userClaims, err := utils.AnalyseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 1, StatusMsg: err.Error()},
			})
			return
		}
		db.Where("username = ?", userClaims.Username).First(&user)
		isVisitor = false
	}

	var videoList []models.Video //创建一个Video数组存起来
	// //按时间倒序
	// db.Order("create_time desc").Where("create_time < ?", lastTime).Limit(30).Find(&videoList)
	db.Order("create_time desc").Limit(30).Find(&videoList)
	fmt.Println(videoList)
	fmt.Println(len(videoList))
	videoPoList := make([]Video, len(videoList))
	for i := range videoList {
		var author models.User
		db.Where("username = ?", videoList[i].Author).First(&author)
		var followCount int64
		var followerCount int64
		db.Table("relation").Where("follower_id = ?", author.Id).Count(&followCount)
		db.Table("relation").Where("follow_id = ?", author.Id).Count(&followerCount)
		var relation models.Relation
		isFollow := false
		if !isVisitor {
			rowsAffected := db.Where("follower_id = ? and follow_id = ?", user.Id, author.Id).First(&relation).RowsAffected
			isFollow = rowsAffected != 0
		}
		Author :=  User{
			Id:   author.Id,
			Name:  author.UserName,
			FollowCount: followCount,
			FollowerCount: followerCount,
			IsFollow: isFollow,
		}

		videoPoList[i].Id = videoList[i].Id
		videoPoList[i].Author = Author
		videoPoList[i].PlayUrl = "http://10.141.93.139:8080" + videoList[i].PlayUrl
		videoPoList[i].CoverUrl = "http://10.141.93.139:8080" + videoList[i].CoverUrl

		var favoriteCount int64
		var commentCount int64
		db.Table("favorite").Where("video_id = ?", videoList[i].Id).Count(&favoriteCount)
		db.Table("comment").Where("video_id = ?", videoList[i].Id).Count(&commentCount)
		isFavorite := false
		var favorite models.Favorite
		favoriteRowsAffected := db.Where("video_id = ? and user_id = ?", videoList[i].Id, user.Id).First(&favorite).RowsAffected
		isFavorite = favoriteRowsAffected != 0

		fmt.Println(isFavorite)
		videoPoList[i].FavoriteCount = favoriteCount
		videoPoList[i].CommentCount = commentCount
		videoPoList[i].IsFavorite = isFavorite
		videoPoList[i].Title = videoList[i].Title

	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoPoList,
		NextTime:  time.Now().Unix(),
	})
}
