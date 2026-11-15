package handler

import (
	"encoding/json"
	"net/http"
	domain "url_shortener_api/internal/domains"
	"url_shortener_api/internal/response"
	"url_shortener_api/internal/usecase"
)

type Handler interface {
	CreateShortUrl(w http.ResponseWriter, r *http.Request)
}
type handler struct {
	urlUsecase usecase.UrlUsecase
}

func NewUrlHandler(urlUsecase usecase.UrlUsecase) Handler {
	return &handler{
		urlUsecase: urlUsecase,
	}
}

func (h *handler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var reqShortUrl domain.Url
	err := json.NewDecoder(r.Body).Decode(&reqShortUrl)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "The body request is invalid!!", err.Error())
		return
	}
	url, err := h.urlUsecase.CreateUrl(r.Context(), reqShortUrl.Url, reqShortUrl.Code)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	response.Success(w, http.StatusOK, "Short URL created successfully", url)

}
