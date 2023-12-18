package controller

import (
	"net/http"

	"github.com/ZooLearn/zoo/domain"
	"github.com/gin-gonic/gin"
)

type TagController struct {
	TagUsecase domain.TagUsecase
}

func (tc *TagController) Create(c *gin.Context) {
	var request domain.CreateTagRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err = tc.TagUsecase.Create(c.Request.Context(), domain.Tag{
		Key:      &request.Key,
		Alias:    &request.Alias,
		Color:    &request.Alias,
		ImageUrl: &request.ImageUrl,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Msg:  "success",
		Code: http.StatusOK,
	})
}

func (tc *TagController) Update(c *gin.Context) {
	var request domain.UpdateTagRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err = tc.TagUsecase.Update(c.Request.Context(), domain.Tag{
		ID:       &request.ID,
		Key:      request.Key,
		Alias:    request.Alias,
		Color:    request.Color,
		ImageUrl: request.ImageUrl,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Msg:  "success",
		Code: http.StatusOK,
	})
}

func (tc *TagController) Delete(c *gin.Context) {
	var request domain.DeleteTagRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err = tc.TagUsecase.Delete(c.Request.Context(), request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Msg:  "success",
		Code: http.StatusOK,
	})
}

func (tc *TagController) GetByID(c *gin.Context) {
	var request domain.GetTagByIDRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	result, err := tc.TagUsecase.GetByID(c.Request.Context(), request.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	response := domain.GetByIDResponse{
		Response: domain.Response{
			Msg:  "success",
			Code: http.StatusOK,
		},
		Data: domain.TagInfo{
			ID:        result.ID,
			Key:       result.Key,
			Alias:     result.Alias,
			Color:     result.Color,
			ImageUrl:  result.ImageUrl,
			CreatedAt: *result.CreatedAt,
			UpdateAt:  *result.UpdateAt,
		},
	}
	c.JSON(http.StatusOK, response)
}

func (tc *TagController) FetchAll(c *gin.Context) {
	result, err := tc.TagUsecase.FetchAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	respose := domain.FetchAllResponse{
		Response: domain.Response{
			Msg:  "success",
			Code: http.StatusOK,
		},
	}
	for _, v := range result {
		respose.Data = append(respose.Data, domain.TagInfo{
			ID:        v.ID,
			Alias:     v.Alias,
			Key:       v.Key,
			ImageUrl:  v.ImageUrl,
			Color:     v.Color,
			CreatedAt: *v.CreatedAt,
			UpdateAt:  *v.UpdateAt,
		})
	}
	c.JSON(http.StatusOK, respose)
}

func (tc *TagController) Fetch(c *gin.Context) {
	var request domain.FetchTagRequest
	err := c.ShouldBindQuery(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	result, paginationInfo, err := tc.TagUsecase.Fetch(c.Request.Context(), domain.FetchTagRequest{
		Page:     request.Page,
		PageSize: request.PageSize,
		Key:      request.Key,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	respose := domain.FetchTagResponse{
		Response: domain.Response{
			Msg:  "success",
			Code: http.StatusOK,
		},
		Pagination: domain.Pagination{
			Page:     paginationInfo.Page,
			PageSize: paginationInfo.PageSize,
			Total:    paginationInfo.Total,
		},
	}

	for _, v := range result {
		respose.Data = append(respose.Data, domain.TagInfo{
			ID:        v.ID,
			Alias:     v.Alias,
			Key:       v.Key,
			ImageUrl:  v.ImageUrl,
			Color:     v.Color,
			CreatedAt: *v.CreatedAt,
			UpdateAt:  *v.UpdateAt,
		})
	}
	c.JSON(http.StatusOK, respose)
}
