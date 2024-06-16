// Package ocatalog provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package ocatalog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/krisch/crm-backend/domain"
	"github.com/krisch/crm-backend/dto"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// CatalogCreateRequest defines model for CatalogCreateRequest.
type CatalogCreateRequest struct {
	CompanyUuid openapi_types.UUID `json:"company_uuid" validate:"uuid"`
	Name        string             `json:"name" validate:"trim,name,min=1,max=200"`
}

// CatalogDTO defines model for CatalogDTO.
type CatalogDTO = dto.CatalogDTO

// CatalogFieldCreateRequest defines model for CatalogFieldCreateRequest.
type CatalogFieldCreateRequest struct {
	DataType domain.FieldDataType `json:"data_type" validate:"min=0,max=8"`
	DataUuid *openapi_types.UUID  `json:"data_uuid,omitempty" validate:"omitempty,uuid"`
	Name     string               `json:"name" validate:"trim,name,min=1,max=50"`
}

// CatalogFieldDTO defines model for CatalogFieldDTO.
type CatalogFieldDTO = dto.CatalogFieldDTO

// CatalogFieldPutRequest defines model for CatalogFieldPutRequest.
type CatalogFieldPutRequest struct {
	Name string `json:"name" validate:"trim,name,min=1,max=50"`
}

// CatalogNamedFieldCreateRequest defines model for CatalogNamedFieldCreateRequest.
type CatalogNamedFieldCreateRequest struct {
	DataType domain.FieldDataType `json:"data_type" validate:"min=0,max=8"`
	DataUuid *openapi_types.UUID  `json:"data_uuid,omitempty" validate:"omitempty,uuid"`
	Hash     string               `json:"hash" validate:"trim,name,min=3,max=20"`
	Name     string               `json:"name" validate:"trim,name,min=1,max=50"`
}

// CatalogSearchRequest defines model for CatalogSearchRequest.
type CatalogSearchRequest struct {
	CompanyUuid openapi_types.UUID `json:"company_uuid" validate:"uuid"`
}

// NameRequest defines model for NameRequest.
type NameRequest struct {
	Name string `json:"name" validate:"trim,name,min=0,max=100"`
}

// EntityUUID defines model for entityUUID.
type EntityUUID = openapi_types.UUID

// Uuid defines model for uuid.
type Uuid = openapi_types.UUID

// GetCatalogUUIDDataParams defines parameters for GetCatalogUUIDData.
type GetCatalogUUIDDataParams struct {
	Offset *int    `form:"offset,omitempty" json:"offset,omitempty"`
	Limit  *int    `form:"limit,omitempty" json:"limit,omitempty"`
	Fields *string `form:"fields,omitempty" json:"fields,omitempty"`
	Order  *string `form:"order,omitempty" json:"order,omitempty"`
	By     *string `form:"by,omitempty" json:"by,omitempty"`
}

// PostCatalogUUIDDataJSONBody defines parameters for PostCatalogUUIDData.
type PostCatalogUUIDDataJSONBody struct {
	Fields map[string]interface{} `json:"fields"`
}

// GetCatalogJSONRequestBody defines body for GetCatalog for application/json ContentType.
type GetCatalogJSONRequestBody = CatalogSearchRequest

// PostCatalogJSONRequestBody defines body for PostCatalog for application/json ContentType.
type PostCatalogJSONRequestBody = CatalogCreateRequest

// PostCatalogUUIDDataJSONRequestBody defines body for PostCatalogUUIDData for application/json ContentType.
type PostCatalogUUIDDataJSONRequestBody PostCatalogUUIDDataJSONBody

// PostCatalogUUIDFieldsJSONRequestBody defines body for PostCatalogUUIDFields for application/json ContentType.
type PostCatalogUUIDFieldsJSONRequestBody = CatalogFieldCreateRequest

// PostCatalogUUIDFieldsNamedJSONRequestBody defines body for PostCatalogUUIDFieldsNamed for application/json ContentType.
type PostCatalogUUIDFieldsNamedJSONRequestBody = CatalogNamedFieldCreateRequest

// PutCatalogUUIDFieldsEntityUUIDJSONRequestBody defines body for PutCatalogUUIDFieldsEntityUUID for application/json ContentType.
type PutCatalogUUIDFieldsEntityUUIDJSONRequestBody = CatalogFieldPutRequest

// PatchCatalogUUIDNameJSONRequestBody defines body for PatchCatalogUUIDName for application/json ContentType.
type PatchCatalogUUIDNameJSONRequestBody = NameRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /catalog)
	GetCatalog(ctx echo.Context) error

	// (POST /catalog)
	PostCatalog(ctx echo.Context) error

	// (DELETE /catalog/{UUID})
	DeleteCatalogUUID(ctx echo.Context, uUID Uuid) error

	// (GET /catalog/{UUID})
	GetCatalogUUID(ctx echo.Context, uUID Uuid) error

	// (GET /catalog/{UUID}/data)
	GetCatalogUUIDData(ctx echo.Context, uUID openapi_types.UUID, params GetCatalogUUIDDataParams) error

	// (POST /catalog/{UUID}/data)
	PostCatalogUUIDData(ctx echo.Context, uUID Uuid) error

	// (GET /catalog/{UUID}/fields)
	GetCatalogUUIDFields(ctx echo.Context, uUID Uuid) error

	// (POST /catalog/{UUID}/fields)
	PostCatalogUUIDFields(ctx echo.Context, uUID Uuid) error

	// (POST /catalog/{UUID}/fields/named)
	PostCatalogUUIDFieldsNamed(ctx echo.Context, uUID Uuid) error

	// (DELETE /catalog/{UUID}/fields/{entityUUID})
	DeleteCatalogUUIDFieldsEntityUUID(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error

	// (PUT /catalog/{UUID}/fields/{entityUUID})
	PutCatalogUUIDFieldsEntityUUID(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error

	// (PATCH /catalog/{UUID}/name)
	PatchCatalogUUIDName(ctx echo.Context, uUID Uuid) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetCatalog converts echo context to params.
func (w *ServerInterfaceWrapper) GetCatalog(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetCatalog(ctx)
	return err
}

// PostCatalog converts echo context to params.
func (w *ServerInterfaceWrapper) PostCatalog(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostCatalog(ctx)
	return err
}

// DeleteCatalogUUID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteCatalogUUID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "UUID" -------------
	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteCatalogUUID(ctx, uUID)
	return err
}

// GetCatalogUUID converts echo context to params.
func (w *ServerInterfaceWrapper) GetCatalogUUID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "UUID" -------------
	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetCatalogUUID(ctx, uUID)
	return err
}

// GetCatalogUUIDData converts echo context to params.
func (w *ServerInterfaceWrapper) GetCatalogUUIDData(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "UUID" -------------
	var uUID openapi_types.UUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetCatalogUUIDDataParams
	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "fields" -------------

	err = runtime.BindQueryParameter("form", true, false, "fields", ctx.QueryParams(), &params.Fields)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fields: %s", err))
	}

	// ------------- Optional query parameter "order" -------------

	err = runtime.BindQueryParameter("form", true, false, "order", ctx.QueryParams(), &params.Order)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter order: %s", err))
	}

	// ------------- Optional query parameter "by" -------------

	err = runtime.BindQueryParameter("form", true, false, "by", ctx.QueryParams(), &params.By)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter by: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetCatalogUUIDData(ctx, uUID, params)
	return err
}

// PostCatalogUUIDData converts echo context to params.
func (w *ServerInterfaceWrapper) PostCatalogUUIDData(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "UUID" -------------
	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostCatalogUUIDData(ctx, uUID)
	return err
}

// GetCatalogUUIDFields converts echo context to params.
func (w *ServerInterfaceWrapper) GetCatalogUUIDFields(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "UUID" -------------
	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetCatalogUUIDFields(ctx, uUID)
	return err
}

// PostCatalogUUIDFields converts echo context to params.
func (w *ServerInterfaceWrapper) PostCatalogUUIDFields(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "UUID" -------------
	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostCatalogUUIDFields(ctx, uUID)
	return err
}

// PostCatalogUUIDFieldsNamed converts echo context to params.
func (w *ServerInterfaceWrapper) PostCatalogUUIDFieldsNamed(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "UUID" -------------
	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostCatalogUUIDFieldsNamed(ctx, uUID)
	return err
}

// DeleteCatalogUUIDFieldsEntityUUID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteCatalogUUIDFieldsEntityUUID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "UUID" -------------
	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	// ------------- Path parameter "entityUUID" -------------
	var entityUUID EntityUUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "entityUUID", runtime.ParamLocationPath, ctx.Param("entityUUID"), &entityUUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter entityUUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteCatalogUUIDFieldsEntityUUID(ctx, uUID, entityUUID)
	return err
}

// PutCatalogUUIDFieldsEntityUUID converts echo context to params.
func (w *ServerInterfaceWrapper) PutCatalogUUIDFieldsEntityUUID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "UUID" -------------
	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	// ------------- Path parameter "entityUUID" -------------
	var entityUUID EntityUUID

	err = runtime.BindStyledParameterWithLocation("simple", false, "entityUUID", runtime.ParamLocationPath, ctx.Param("entityUUID"), &entityUUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter entityUUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutCatalogUUIDFieldsEntityUUID(ctx, uUID, entityUUID)
	return err
}

// PatchCatalogUUIDName converts echo context to params.
func (w *ServerInterfaceWrapper) PatchCatalogUUIDName(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "UUID" -------------
	var uUID Uuid

	err = runtime.BindStyledParameterWithLocation("simple", false, "UUID", runtime.ParamLocationPath, ctx.Param("UUID"), &uUID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter UUID: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchCatalogUUIDName(ctx, uUID)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/catalog", wrapper.GetCatalog)
	router.POST(baseURL+"/catalog", wrapper.PostCatalog)
	router.DELETE(baseURL+"/catalog/:UUID", wrapper.DeleteCatalogUUID)
	router.GET(baseURL+"/catalog/:UUID", wrapper.GetCatalogUUID)
	router.GET(baseURL+"/catalog/:UUID/data", wrapper.GetCatalogUUIDData)
	router.POST(baseURL+"/catalog/:UUID/data", wrapper.PostCatalogUUIDData)
	router.GET(baseURL+"/catalog/:UUID/fields", wrapper.GetCatalogUUIDFields)
	router.POST(baseURL+"/catalog/:UUID/fields", wrapper.PostCatalogUUIDFields)
	router.POST(baseURL+"/catalog/:UUID/fields/named", wrapper.PostCatalogUUIDFieldsNamed)
	router.DELETE(baseURL+"/catalog/:UUID/fields/:entityUUID", wrapper.DeleteCatalogUUIDFieldsEntityUUID)
	router.PUT(baseURL+"/catalog/:UUID/fields/:entityUUID", wrapper.PutCatalogUUIDFieldsEntityUUID)
	router.PATCH(baseURL+"/catalog/:UUID/name", wrapper.PatchCatalogUUIDName)

}

type GetCatalogRequestObject struct {
	Body *GetCatalogJSONRequestBody
}

type GetCatalogResponseObject interface {
	VisitGetCatalogResponse(w http.ResponseWriter) error
}

type GetCatalog200JSONResponse struct {
	Count int          `json:"count"`
	Items []CatalogDTO `json:"items"`
}

func (response GetCatalog200JSONResponse) VisitGetCatalogResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostCatalogRequestObject struct {
	Body *PostCatalogJSONRequestBody
}

type PostCatalogResponseObject interface {
	VisitPostCatalogResponse(w http.ResponseWriter) error
}

type PostCatalog200JSONResponse struct {
	Uuid openapi_types.UUID `json:"uuid"`
}

func (response PostCatalog200JSONResponse) VisitPostCatalogResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteCatalogUUIDRequestObject struct {
	UUID Uuid `json:"UUID"`
}

type DeleteCatalogUUIDResponseObject interface {
	VisitDeleteCatalogUUIDResponse(w http.ResponseWriter) error
}

type DeleteCatalogUUID200Response struct {
}

func (response DeleteCatalogUUID200Response) VisitDeleteCatalogUUIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type GetCatalogUUIDRequestObject struct {
	UUID Uuid `json:"UUID"`
}

type GetCatalogUUIDResponseObject interface {
	VisitGetCatalogUUIDResponse(w http.ResponseWriter) error
}

type GetCatalogUUID200JSONResponse CatalogDTO

func (response GetCatalogUUID200JSONResponse) VisitGetCatalogUUIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetCatalogUUIDDataRequestObject struct {
	UUID   openapi_types.UUID `json:"UUID"`
	Params GetCatalogUUIDDataParams
}

type GetCatalogUUIDDataResponseObject interface {
	VisitGetCatalogUUIDDataResponse(w http.ResponseWriter) error
}

type GetCatalogUUIDData200ResponseHeaders struct {
	CacheControl string
}

type GetCatalogUUIDData200JSONResponse struct {
	Body struct {
		Count int                      `json:"count"`
		Items []map[string]interface{} `json:"items"`
		Total int64                    `json:"total"`
	}
	Headers GetCatalogUUIDData200ResponseHeaders
}

func (response GetCatalogUUIDData200JSONResponse) VisitGetCatalogUUIDDataResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("cache-control", fmt.Sprint(response.Headers.CacheControl))
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response.Body)
}

type PostCatalogUUIDDataRequestObject struct {
	UUID Uuid `json:"UUID"`
	Body *PostCatalogUUIDDataJSONRequestBody
}

type PostCatalogUUIDDataResponseObject interface {
	VisitPostCatalogUUIDDataResponse(w http.ResponseWriter) error
}

type PostCatalogUUIDData200JSONResponse struct {
	Uuid openapi_types.UUID `json:"uuid"`
}

func (response PostCatalogUUIDData200JSONResponse) VisitPostCatalogUUIDDataResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetCatalogUUIDFieldsRequestObject struct {
	UUID Uuid `json:"UUID"`
}

type GetCatalogUUIDFieldsResponseObject interface {
	VisitGetCatalogUUIDFieldsResponse(w http.ResponseWriter) error
}

type GetCatalogUUIDFields200JSONResponse struct {
	Count int               `json:"count"`
	Items []CatalogFieldDTO `json:"items"`
}

func (response GetCatalogUUIDFields200JSONResponse) VisitGetCatalogUUIDFieldsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostCatalogUUIDFieldsRequestObject struct {
	UUID Uuid `json:"UUID"`
	Body *PostCatalogUUIDFieldsJSONRequestBody
}

type PostCatalogUUIDFieldsResponseObject interface {
	VisitPostCatalogUUIDFieldsResponse(w http.ResponseWriter) error
}

type PostCatalogUUIDFields200JSONResponse struct {
	CatalogUuid     openapi_types.UUID   `json:"catalog_uuid"`
	Hash            string               `json:"hash"`
	Type            domain.FieldDataType `json:"type"`
	TypeDescription string               `json:"type_description"`
	TypeUuid        *openapi_types.UUID  `json:"type_uuid,omitempty"`
	Uuid            openapi_types.UUID   `json:"uuid"`
}

func (response PostCatalogUUIDFields200JSONResponse) VisitPostCatalogUUIDFieldsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostCatalogUUIDFieldsNamedRequestObject struct {
	UUID Uuid `json:"UUID"`
	Body *PostCatalogUUIDFieldsNamedJSONRequestBody
}

type PostCatalogUUIDFieldsNamedResponseObject interface {
	VisitPostCatalogUUIDFieldsNamedResponse(w http.ResponseWriter) error
}

type PostCatalogUUIDFieldsNamed200JSONResponse struct {
	CatalogUuid     *openapi_types.UUID  `json:"catalog_uuid,omitempty"`
	Hash            string               `json:"hash"`
	Type            domain.FieldDataType `json:"type"`
	TypeDescription string               `json:"type_description"`
	TypeUuid        *openapi_types.UUID  `json:"type_uuid,omitempty"`
	Uuid            string               `json:"uuid"`
}

func (response PostCatalogUUIDFieldsNamed200JSONResponse) VisitPostCatalogUUIDFieldsNamedResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteCatalogUUIDFieldsEntityUUIDRequestObject struct {
	UUID       Uuid       `json:"UUID"`
	EntityUUID EntityUUID `json:"entityUUID"`
}

type DeleteCatalogUUIDFieldsEntityUUIDResponseObject interface {
	VisitDeleteCatalogUUIDFieldsEntityUUIDResponse(w http.ResponseWriter) error
}

type DeleteCatalogUUIDFieldsEntityUUID200Response struct {
}

func (response DeleteCatalogUUIDFieldsEntityUUID200Response) VisitDeleteCatalogUUIDFieldsEntityUUIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PutCatalogUUIDFieldsEntityUUIDRequestObject struct {
	UUID       Uuid       `json:"UUID"`
	EntityUUID EntityUUID `json:"entityUUID"`
	Body       *PutCatalogUUIDFieldsEntityUUIDJSONRequestBody
}

type PutCatalogUUIDFieldsEntityUUIDResponseObject interface {
	VisitPutCatalogUUIDFieldsEntityUUIDResponse(w http.ResponseWriter) error
}

type PutCatalogUUIDFieldsEntityUUID200Response struct {
}

func (response PutCatalogUUIDFieldsEntityUUID200Response) VisitPutCatalogUUIDFieldsEntityUUIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PatchCatalogUUIDNameRequestObject struct {
	UUID Uuid `json:"UUID"`
	Body *PatchCatalogUUIDNameJSONRequestBody
}

type PatchCatalogUUIDNameResponseObject interface {
	VisitPatchCatalogUUIDNameResponse(w http.ResponseWriter) error
}

type PatchCatalogUUIDName200Response struct {
}

func (response PatchCatalogUUIDName200Response) VisitPatchCatalogUUIDNameResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /catalog)
	GetCatalog(ctx context.Context, request GetCatalogRequestObject) (GetCatalogResponseObject, error)

	// (POST /catalog)
	PostCatalog(ctx context.Context, request PostCatalogRequestObject) (PostCatalogResponseObject, error)

	// (DELETE /catalog/{UUID})
	DeleteCatalogUUID(ctx context.Context, request DeleteCatalogUUIDRequestObject) (DeleteCatalogUUIDResponseObject, error)

	// (GET /catalog/{UUID})
	GetCatalogUUID(ctx context.Context, request GetCatalogUUIDRequestObject) (GetCatalogUUIDResponseObject, error)

	// (GET /catalog/{UUID}/data)
	GetCatalogUUIDData(ctx context.Context, request GetCatalogUUIDDataRequestObject) (GetCatalogUUIDDataResponseObject, error)

	// (POST /catalog/{UUID}/data)
	PostCatalogUUIDData(ctx context.Context, request PostCatalogUUIDDataRequestObject) (PostCatalogUUIDDataResponseObject, error)

	// (GET /catalog/{UUID}/fields)
	GetCatalogUUIDFields(ctx context.Context, request GetCatalogUUIDFieldsRequestObject) (GetCatalogUUIDFieldsResponseObject, error)

	// (POST /catalog/{UUID}/fields)
	PostCatalogUUIDFields(ctx context.Context, request PostCatalogUUIDFieldsRequestObject) (PostCatalogUUIDFieldsResponseObject, error)

	// (POST /catalog/{UUID}/fields/named)
	PostCatalogUUIDFieldsNamed(ctx context.Context, request PostCatalogUUIDFieldsNamedRequestObject) (PostCatalogUUIDFieldsNamedResponseObject, error)

	// (DELETE /catalog/{UUID}/fields/{entityUUID})
	DeleteCatalogUUIDFieldsEntityUUID(ctx context.Context, request DeleteCatalogUUIDFieldsEntityUUIDRequestObject) (DeleteCatalogUUIDFieldsEntityUUIDResponseObject, error)

	// (PUT /catalog/{UUID}/fields/{entityUUID})
	PutCatalogUUIDFieldsEntityUUID(ctx context.Context, request PutCatalogUUIDFieldsEntityUUIDRequestObject) (PutCatalogUUIDFieldsEntityUUIDResponseObject, error)

	// (PATCH /catalog/{UUID}/name)
	PatchCatalogUUIDName(ctx context.Context, request PatchCatalogUUIDNameRequestObject) (PatchCatalogUUIDNameResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetCatalog operation middleware
func (sh *strictHandler) GetCatalog(ctx echo.Context) error {
	var request GetCatalogRequestObject

	var body GetCatalogJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetCatalog(ctx.Request().Context(), request.(GetCatalogRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetCatalog")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetCatalogResponseObject); ok {
		return validResponse.VisitGetCatalogResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostCatalog operation middleware
func (sh *strictHandler) PostCatalog(ctx echo.Context) error {
	var request PostCatalogRequestObject

	var body PostCatalogJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostCatalog(ctx.Request().Context(), request.(PostCatalogRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostCatalog")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostCatalogResponseObject); ok {
		return validResponse.VisitPostCatalogResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteCatalogUUID operation middleware
func (sh *strictHandler) DeleteCatalogUUID(ctx echo.Context, uUID Uuid) error {
	var request DeleteCatalogUUIDRequestObject

	request.UUID = uUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteCatalogUUID(ctx.Request().Context(), request.(DeleteCatalogUUIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteCatalogUUID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteCatalogUUIDResponseObject); ok {
		return validResponse.VisitDeleteCatalogUUIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetCatalogUUID operation middleware
func (sh *strictHandler) GetCatalogUUID(ctx echo.Context, uUID Uuid) error {
	var request GetCatalogUUIDRequestObject

	request.UUID = uUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetCatalogUUID(ctx.Request().Context(), request.(GetCatalogUUIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetCatalogUUID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetCatalogUUIDResponseObject); ok {
		return validResponse.VisitGetCatalogUUIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetCatalogUUIDData operation middleware
func (sh *strictHandler) GetCatalogUUIDData(ctx echo.Context, uUID openapi_types.UUID, params GetCatalogUUIDDataParams) error {
	var request GetCatalogUUIDDataRequestObject

	request.UUID = uUID
	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetCatalogUUIDData(ctx.Request().Context(), request.(GetCatalogUUIDDataRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetCatalogUUIDData")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetCatalogUUIDDataResponseObject); ok {
		return validResponse.VisitGetCatalogUUIDDataResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostCatalogUUIDData operation middleware
func (sh *strictHandler) PostCatalogUUIDData(ctx echo.Context, uUID Uuid) error {
	var request PostCatalogUUIDDataRequestObject

	request.UUID = uUID

	var body PostCatalogUUIDDataJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostCatalogUUIDData(ctx.Request().Context(), request.(PostCatalogUUIDDataRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostCatalogUUIDData")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostCatalogUUIDDataResponseObject); ok {
		return validResponse.VisitPostCatalogUUIDDataResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetCatalogUUIDFields operation middleware
func (sh *strictHandler) GetCatalogUUIDFields(ctx echo.Context, uUID Uuid) error {
	var request GetCatalogUUIDFieldsRequestObject

	request.UUID = uUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetCatalogUUIDFields(ctx.Request().Context(), request.(GetCatalogUUIDFieldsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetCatalogUUIDFields")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetCatalogUUIDFieldsResponseObject); ok {
		return validResponse.VisitGetCatalogUUIDFieldsResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostCatalogUUIDFields operation middleware
func (sh *strictHandler) PostCatalogUUIDFields(ctx echo.Context, uUID Uuid) error {
	var request PostCatalogUUIDFieldsRequestObject

	request.UUID = uUID

	var body PostCatalogUUIDFieldsJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostCatalogUUIDFields(ctx.Request().Context(), request.(PostCatalogUUIDFieldsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostCatalogUUIDFields")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostCatalogUUIDFieldsResponseObject); ok {
		return validResponse.VisitPostCatalogUUIDFieldsResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostCatalogUUIDFieldsNamed operation middleware
func (sh *strictHandler) PostCatalogUUIDFieldsNamed(ctx echo.Context, uUID Uuid) error {
	var request PostCatalogUUIDFieldsNamedRequestObject

	request.UUID = uUID

	var body PostCatalogUUIDFieldsNamedJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostCatalogUUIDFieldsNamed(ctx.Request().Context(), request.(PostCatalogUUIDFieldsNamedRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostCatalogUUIDFieldsNamed")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostCatalogUUIDFieldsNamedResponseObject); ok {
		return validResponse.VisitPostCatalogUUIDFieldsNamedResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteCatalogUUIDFieldsEntityUUID operation middleware
func (sh *strictHandler) DeleteCatalogUUIDFieldsEntityUUID(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error {
	var request DeleteCatalogUUIDFieldsEntityUUIDRequestObject

	request.UUID = uUID
	request.EntityUUID = entityUUID

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteCatalogUUIDFieldsEntityUUID(ctx.Request().Context(), request.(DeleteCatalogUUIDFieldsEntityUUIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteCatalogUUIDFieldsEntityUUID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteCatalogUUIDFieldsEntityUUIDResponseObject); ok {
		return validResponse.VisitDeleteCatalogUUIDFieldsEntityUUIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PutCatalogUUIDFieldsEntityUUID operation middleware
func (sh *strictHandler) PutCatalogUUIDFieldsEntityUUID(ctx echo.Context, uUID Uuid, entityUUID EntityUUID) error {
	var request PutCatalogUUIDFieldsEntityUUIDRequestObject

	request.UUID = uUID
	request.EntityUUID = entityUUID

	var body PutCatalogUUIDFieldsEntityUUIDJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PutCatalogUUIDFieldsEntityUUID(ctx.Request().Context(), request.(PutCatalogUUIDFieldsEntityUUIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutCatalogUUIDFieldsEntityUUID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PutCatalogUUIDFieldsEntityUUIDResponseObject); ok {
		return validResponse.VisitPutCatalogUUIDFieldsEntityUUIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchCatalogUUIDName operation middleware
func (sh *strictHandler) PatchCatalogUUIDName(ctx echo.Context, uUID Uuid) error {
	var request PatchCatalogUUIDNameRequestObject

	request.UUID = uUID

	var body PatchCatalogUUIDNameJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchCatalogUUIDName(ctx.Request().Context(), request.(PatchCatalogUUIDNameRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchCatalogUUIDName")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchCatalogUUIDNameResponseObject); ok {
		return validResponse.VisitPatchCatalogUUIDNameResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}