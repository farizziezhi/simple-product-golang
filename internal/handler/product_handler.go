package handler

import (
	"net/http"

	"github.com/farizziezhi/layered/internal/domain"
	"github.com/farizziezhi/layered/internal/service"
	"github.com/farizziezhi/layered/pkg/helper"
	"github.com/go-chi/chi/v5"
)


type ProductHandler interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}

type ProductHandlerImpl struct {
	Service service.ProductService
}

func NewProductHandler(service service.ProductService) ProductHandler {
	return &ProductHandlerImpl{
		Service: service,
	}
}

func (handler *ProductHandlerImpl) GetProducts(w http.ResponseWriter, r *http.Request) {
	data, err := handler.Service.GetProducts(r.Context())

	if err != nil {
		helper.SendJSON(w, http.StatusInternalServerError, domain.Response{
			Message: "Gagal mengambil data produk",
			Data:    nil,
		})

		return
	}

	helper.SendJSON(w, http.StatusOK, domain.Response{
		Message: "Data produk berhasil diambil",
		Data:    data,
	})
}
func (handler *ProductHandlerImpl) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data, err := handler.Service.GetProduct(r.Context(), id)

	if err != nil {
		helper.SendJSON(w, http.StatusInternalServerError, domain.Response{
			Message: "Gagal mengambil data produk",
			Data:    nil,
		})

		return
	}

	helper.SendJSON(w, http.StatusOK, domain.Response{
		Message: "Data produk berhasil diambil",
		Data:    data,
	})
}
func (handler *ProductHandlerImpl) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := handler.Service.DeleteProduct(r.Context(), id)
	if err != nil {
		helper.SendJSON(w, http.StatusInternalServerError, domain.Response{
			Message: "Gagal menghapus data produk",
			Data:    nil,
		})

		return
	}

	helper.SendJSON(w, http.StatusOK, domain.Response{
		Message: "Data produk berhasil dihapus",
	})
}
func (handler *ProductHandlerImpl) CreateProduct(w http.ResponseWriter, r *http.Request) {
	data := domain.Product{}
	helper.ParseBody(r, &data)
	err := handler.Service.CreateProduct(r.Context(), data)
	if err != nil {
		helper.SendJSON(w, http.StatusInternalServerError, domain.Response{
			Message: "Gagal membuat data produk",
			Data:    nil,
		})

		return
	}

	helper.SendJSON(w, http.StatusOK, domain.Response{
		Message: "Data produk berhasil dibuat",
	})
}
func (handler *ProductHandlerImpl) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data := domain.Product{}
	helper.ParseBody(r, &data)
	err := handler.Service.UpdateProduct(r.Context(), id, data)
	if err != nil {
		helper.SendJSON(w, http.StatusInternalServerError, domain.Response{
			Message: "Gagal mengupdate data produk",
			Data:    nil,
		})

		return 
	}

	helper.SendJSON(w, http.StatusOK, domain.Response{
		Message: "Data produk berhasil diupdate",
	})
}