// Package REST provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package REST

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	DeviceAuthScopes = "deviceAuth.Scopes"
)

// Defines values for HealthResult.
const (
	Down HealthResult = "Down"
	Up   HealthResult = "Up"
)

// Device a device
type Device struct {
	// Id Device ID is the unique identifier for a remote device
	Id DeviceID `json:"id"`
}

// DeviceID Device ID is the unique identifier for a remote device
type DeviceID = openapi_types.UUID

// DeviceList list of devices
type DeviceList struct {
	// Count Amount of Items contained in List
	Count ListItemCount `json:"count"`

	// Items array of devices, it will always at least contain the device of the authenticated user
	Items []Device `json:"items"`
}

// HealthAggregation defines model for HealthAggregation.
type HealthAggregation struct {
	// Components The different Components of the Server
	Components *[]HealthAggregationComponent `json:"components,omitempty"`

	// Health A Health Check Result
	Health HealthResult `json:"health"`
}

// HealthAggregationComponent defines model for HealthAggregationComponent.
type HealthAggregationComponent struct {
	// Health A Health Check Result
	Health HealthResult `json:"health"`

	// Name The Name of the Component to be Health Checked
	Name string `json:"name"`
}

// HealthResult A Health Check Result
type HealthResult string

// ListItemCount Amount of Items contained in List
type ListItemCount = int

// ModifiedAtTimestamp A Timestamp indicating when a datum was last modified
type ModifiedAtTimestamp = time.Time

// ModuleDataStream Module Data Stream
type ModuleDataStream = openapi_types.File

// ModuleName Module Name
type ModuleName = string

// RegistrationResult defines model for RegistrationResult.
type RegistrationResult struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// ShareResponse defines model for ShareResponse.
type ShareResponse struct {
	ShareCode *string `json:"shareCode,omitempty"`
}

// DeviceIDQuery Device ID is the unique identifier for a remote device
type DeviceIDQuery = DeviceID

// ShareCode defines model for ShareCode.
type ShareCode = string

// XDeviceID Device ID is the unique identifier for a remote device
type XDeviceID = DeviceID

// DeviceListResponse list of devices
type DeviceListResponse = DeviceList

// ModuleDataAccepted An Empty JSON
type ModuleDataAccepted = interface{}

// ModuleDeletionAccepted An Empty JSON
type ModuleDeletionAccepted = interface{}

// RegisterParams defines parameters for Register.
type RegisterParams struct {
	// Share The Share Code from the Share API. If presented in combination with a new Device ID,
	// it can be used to add new devices to an account.
	Share *ShareCode `form:"share,omitempty" json:"share,omitempty"`

	// XDeviceID Unique Identifier of the calling Device. If calling Data endpoints, must be presented in order
	// to be properly authenticated.
	XDeviceID XDeviceID `json:"X-Device-ID"`
}

// ShareParams defines parameters for Share.
type ShareParams struct {
	// XDeviceID Unique Identifier of the calling Device. If calling Data endpoints, must be presented in order
	// to be properly authenticated.
	XDeviceID XDeviceID `json:"X-Device-ID"`
}

// GetDevicesParams defines parameters for GetDevices.
type GetDevicesParams struct {
	// XDeviceID Unique Identifier of the calling Device. If calling Data endpoints, must be presented in order
	// to be properly authenticated.
	XDeviceID XDeviceID `json:"X-Device-ID"`
}

// DeleteModulesParams defines parameters for DeleteModules.
type DeleteModulesParams struct {
	// DeviceId Device Identifier to use for the Query. If given, takes precedence over X-Device-ID or other hints.
	// Use to query data from devices in your account from another account.
	DeviceId *DeviceIDQuery `form:"device-id,omitempty" json:"device-id,omitempty"`

	// XDeviceID Unique Identifier of the calling Device. If calling Data endpoints, must be presented in order
	// to be properly authenticated.
	XDeviceID XDeviceID `json:"X-Device-ID"`
}

// GetModuleParams defines parameters for GetModule.
type GetModuleParams struct {
	// DeviceId Device Identifier to use for the Query. If given, takes precedence over X-Device-ID or other hints.
	// Use to query data from devices in your account from another account.
	DeviceId *DeviceIDQuery `form:"device-id,omitempty" json:"device-id,omitempty"`

	// XDeviceID Unique Identifier of the calling Device. If calling Data endpoints, must be presented in order
	// to be properly authenticated.
	XDeviceID XDeviceID `json:"X-Device-ID"`
}

// CreateModuleParams defines parameters for CreateModule.
type CreateModuleParams struct {
	// XDeviceID Unique Identifier of the calling Device. If calling Data endpoints, must be presented in order
	// to be properly authenticated.
	XDeviceID XDeviceID `json:"X-Device-ID"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Register A Device
	// (POST /auth/register)
	Register(ctx echo.Context, params RegisterParams) error
	// Share your Account
	// (POST /auth/share)
	Share(ctx echo.Context, params ShareParams) error
	// Get All registered Devices for your Account
	// (GET /devices)
	GetDevices(ctx echo.Context, params GetDevicesParams) error
	// Checks if the Service is Available for Processing Request
	// (GET /health)
	IsHealthy(ctx echo.Context) error
	// Clears Module Data for a Device
	// (DELETE /module)
	DeleteModules(ctx echo.Context, params DeleteModulesParams) error
	// Get Module Data
	// (GET /module/{name})
	GetModule(ctx echo.Context, name ModuleName, params GetModuleParams) error
	// Create/Update Module Data
	// (POST /module/{name})
	CreateModule(ctx echo.Context, name ModuleName, params CreateModuleParams) error
	// Checks if the Service is Operational
	// (GET /ready)
	IsReady(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Register converts echo context to params.
func (w *ServerInterfaceWrapper) Register(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params RegisterParams
	// ------------- Optional query parameter "share" -------------

	err = runtime.BindQueryParameter("form", true, false, "share", ctx.QueryParams(), &params.Share)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter share: %s", err))
	}

	headers := ctx.Request().Header
	// ------------- Required header parameter "X-Device-ID" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Device-ID")]; found {
		var XDeviceID XDeviceID
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for X-Device-ID, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Device-ID", runtime.ParamLocationHeader, valueList[0], &XDeviceID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter X-Device-ID: %s", err))
		}

		params.XDeviceID = XDeviceID
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Header parameter X-Device-ID is required, but not found"))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Register(ctx, params)
	return err
}

// Share converts echo context to params.
func (w *ServerInterfaceWrapper) Share(ctx echo.Context) error {
	var err error

	ctx.Set(DeviceAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShareParams

	headers := ctx.Request().Header
	// ------------- Required header parameter "X-Device-ID" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Device-ID")]; found {
		var XDeviceID XDeviceID
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for X-Device-ID, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Device-ID", runtime.ParamLocationHeader, valueList[0], &XDeviceID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter X-Device-ID: %s", err))
		}

		params.XDeviceID = XDeviceID
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Header parameter X-Device-ID is required, but not found"))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Share(ctx, params)
	return err
}

// GetDevices converts echo context to params.
func (w *ServerInterfaceWrapper) GetDevices(ctx echo.Context) error {
	var err error

	ctx.Set(DeviceAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetDevicesParams

	headers := ctx.Request().Header
	// ------------- Required header parameter "X-Device-ID" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Device-ID")]; found {
		var XDeviceID XDeviceID
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for X-Device-ID, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Device-ID", runtime.ParamLocationHeader, valueList[0], &XDeviceID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter X-Device-ID: %s", err))
		}

		params.XDeviceID = XDeviceID
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Header parameter X-Device-ID is required, but not found"))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetDevices(ctx, params)
	return err
}

// IsHealthy converts echo context to params.
func (w *ServerInterfaceWrapper) IsHealthy(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.IsHealthy(ctx)
	return err
}

// DeleteModules converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteModules(ctx echo.Context) error {
	var err error

	ctx.Set(DeviceAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteModulesParams
	// ------------- Optional query parameter "device-id" -------------

	err = runtime.BindQueryParameter("form", true, false, "device-id", ctx.QueryParams(), &params.DeviceId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter device-id: %s", err))
	}

	headers := ctx.Request().Header
	// ------------- Required header parameter "X-Device-ID" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Device-ID")]; found {
		var XDeviceID XDeviceID
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for X-Device-ID, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Device-ID", runtime.ParamLocationHeader, valueList[0], &XDeviceID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter X-Device-ID: %s", err))
		}

		params.XDeviceID = XDeviceID
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Header parameter X-Device-ID is required, but not found"))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteModules(ctx, params)
	return err
}

// GetModule converts echo context to params.
func (w *ServerInterfaceWrapper) GetModule(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "name" -------------
	var name ModuleName

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	ctx.Set(DeviceAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetModuleParams
	// ------------- Optional query parameter "device-id" -------------

	err = runtime.BindQueryParameter("form", true, false, "device-id", ctx.QueryParams(), &params.DeviceId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter device-id: %s", err))
	}

	headers := ctx.Request().Header
	// ------------- Required header parameter "X-Device-ID" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Device-ID")]; found {
		var XDeviceID XDeviceID
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for X-Device-ID, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Device-ID", runtime.ParamLocationHeader, valueList[0], &XDeviceID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter X-Device-ID: %s", err))
		}

		params.XDeviceID = XDeviceID
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Header parameter X-Device-ID is required, but not found"))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetModule(ctx, name, params)
	return err
}

// CreateModule converts echo context to params.
func (w *ServerInterfaceWrapper) CreateModule(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "name" -------------
	var name ModuleName

	err = runtime.BindStyledParameterWithLocation("simple", false, "name", runtime.ParamLocationPath, ctx.Param("name"), &name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter name: %s", err))
	}

	ctx.Set(DeviceAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params CreateModuleParams

	headers := ctx.Request().Header
	// ------------- Required header parameter "X-Device-ID" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("X-Device-ID")]; found {
		var XDeviceID XDeviceID
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for X-Device-ID, got %d", n))
		}

		err = runtime.BindStyledParameterWithLocation("simple", false, "X-Device-ID", runtime.ParamLocationHeader, valueList[0], &XDeviceID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter X-Device-ID: %s", err))
		}

		params.XDeviceID = XDeviceID
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Header parameter X-Device-ID is required, but not found"))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateModule(ctx, name, params)
	return err
}

// IsReady converts echo context to params.
func (w *ServerInterfaceWrapper) IsReady(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.IsReady(ctx)
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

	router.POST(baseURL+"/auth/register", wrapper.Register)
	router.POST(baseURL+"/auth/share", wrapper.Share)
	router.GET(baseURL+"/devices", wrapper.GetDevices)
	router.GET(baseURL+"/health", wrapper.IsHealthy)
	router.DELETE(baseURL+"/module", wrapper.DeleteModules)
	router.GET(baseURL+"/module/:name", wrapper.GetModule)
	router.POST(baseURL+"/module/:name", wrapper.CreateModule)
	router.GET(baseURL+"/ready", wrapper.IsReady)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
