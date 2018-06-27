// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetUsersIDBackgroundimageJpegURL generates an URL for the get users ID backgroundimage jpeg operation
type GetUsersIDBackgroundimageJpegURL struct {
	ID string

	Height *float64
	Width  *float64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetUsersIDBackgroundimageJpegURL) WithBasePath(bp string) *GetUsersIDBackgroundimageJpegURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetUsersIDBackgroundimageJpegURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetUsersIDBackgroundimageJpegURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/users/{id}/backgroundimage.jpeg"

	id := o.ID
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("ID is required on GetUsersIDBackgroundimageJpegURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v2/"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var height string
	if o.Height != nil {
		height = swag.FormatFloat64(*o.Height)
	}
	if height != "" {
		qs.Set("height", height)
	}

	var width string
	if o.Width != nil {
		width = swag.FormatFloat64(*o.Width)
	}
	if width != "" {
		qs.Set("width", width)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetUsersIDBackgroundimageJpegURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetUsersIDBackgroundimageJpegURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetUsersIDBackgroundimageJpegURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetUsersIDBackgroundimageJpegURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetUsersIDBackgroundimageJpegURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetUsersIDBackgroundimageJpegURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
