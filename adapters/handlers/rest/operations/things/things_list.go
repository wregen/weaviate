//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// ThingsListHandlerFunc turns a function with the right signature into a things list handler
type ThingsListHandlerFunc func(ThingsListParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ThingsListHandlerFunc) Handle(params ThingsListParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ThingsListHandler interface for that can handle valid things list params
type ThingsListHandler interface {
	Handle(ThingsListParams, *models.Principal) middleware.Responder
}

// NewThingsList creates a new http.Handler for the things list operation
func NewThingsList(ctx *middleware.Context, handler ThingsListHandler) *ThingsList {
	return &ThingsList{Context: ctx, Handler: handler}
}

/*ThingsList swagger:route GET /things things thingsList

Get a list of Things.

Lists all Things in reverse order of creation, owned by the user that belongs to the used token.

*/
type ThingsList struct {
	Context *middleware.Context
	Handler ThingsListHandler
}

func (o *ThingsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewThingsListParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}