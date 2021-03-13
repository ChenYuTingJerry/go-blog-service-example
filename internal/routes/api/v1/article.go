package v1

import (
	"github.com/gin-gonic/gin"
	"github/ChenYuTingJerry/blog-service/pkg/app"
	"github/ChenYuTingJerry/blog-service/pkg/errcode"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (t Article) Get(c *gin.Context)    {}
func (t Article) List(c *gin.Context)   {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
func (t Article) Create(c *gin.Context) {}
func (t Article) Update(c *gin.Context) {}
func (t Article) Delete(c *gin.Context) {}
