package handlers

import (
	"bytes"
	"htmxdemo/service"
	"htmxdemo/types"
	"htmxdemo/views/components"
	"htmxdemo/views/pages"
	"net/http"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	service     service.IService
}

func New(service service.IService) *Handlers {
	return &Handlers{
		service:     service,
	}
}

func (h *Handlers) HandleHomePage(c echo.Context) error {
	return render(c, pages.Home())
}

func (h *Handlers) HandleInsertTransaction(c echo.Context) error {
	req := types.Transaction{}
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	err = h.service.InsertTransaction(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handlers) HandleTransactions(c echo.Context) (err error) {
	transactions, err := h.service.ListTransactions(c.Request().Context())
	if err != nil {
		return err
	}

	return render(c, pages.Transactions(transactions))
}

func (h *Handlers) HandleTransactionsSearch(c echo.Context) (err error) {
	searchVal := strings.ToLower(c.FormValue("search"))
	searchValInt, err := strconv.ParseInt(searchVal, 10, 64)
	if err != nil {
		return err
	}

	final, err := h.service.SearchTransactions(c.Request().Context(), searchValInt)
	if err != nil {
		return err
	}

	return render(c, components.TransactionsTableBody(final))
}

func (h *Handlers) HandleUpdateTransaction(c echo.Context) (err error) {
	idQueryParam := c.QueryParam("id")
	id, err := strconv.ParseInt(idQueryParam, 10, 64)
	if err != nil {
		return err
	}

	txn, err := h.service.UpdateTransaction(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return render(c, components.TransactionsTableRow(txn))
}

func render(c echo.Context, component templ.Component) error {
	var buffer bytes.Buffer
	err := component.Render(c.Request().Context(), &buffer)
	if err != nil {
		return err
	}

	return c.HTML(http.StatusOK, buffer.String())
}
