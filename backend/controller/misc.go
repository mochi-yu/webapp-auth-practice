package controller

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mochi-yu/webapp-auth-practice/config"
)

type IMiscController interface {
	SignIn(c *gin.Context)
}

type MiscController struct {
	cfg *config.Config
}

func NewMiscController(cfg *config.Config) IMiscController {
	return &MiscController{cfg: cfg}
}

func (mc *MiscController) SignIn(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code is required"})
		return
	}

	// "https://github.com/login/oauth/access_token"へPOSTリクエストを送信
	// リクエストボディにはcode, client_id, client_secretを含める
	// レスポンスはJSON形式
	resp, err := http.PostForm("https://github.com/login/oauth/access_token", url.Values{
		"code":          {code},
		"client_id":     {os.Getenv("GITHUB_CLIENT_ID")},
		"client_secret": {os.Getenv("GITHUB_CLIENT_SECRET")},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := resp.Body.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// レスポンスボディのform形式をパース
	values, err := url.ParseQuery(string(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("%+v", values)

	// access_tokenを取得
	accessToken := values.Get("access_token")
	if accessToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "access_token is required"})
		return
	}

	// Bearerトークンにaccess_tokenをセットしてGetリクエストを送信
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ユーザ情報のレスポンスを処理
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	log.Printf("%+v", string(body))

	c.JSON(http.StatusOK, gin.H{"message": "sign in"})
}
