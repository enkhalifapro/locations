package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("location", func() {
	Title("Location Service")
	Description("Service for manipulate user location")
	Server("location", func() {
		Description("user hosts the User Service.")

		// List the services hosted by this server.
		Services("locations", "openapi")

		// List the Hosts and their transport URLs.
		Host("development", func() {
			Description("Development hosts.")
			URI("http://localhost:8000")
		})

		Host("production", func() {
			Description("Production hosts.")
			// URIs can be parameterized using {param} notation.
			URI("https://{version}.goa.design")
			URI("grpcs://{version}.goa.design")

			// Variable describes a URI variable.
			Variable("version", String, "API version", func() {
				// URL parameters must have a default value and/or an
				// enum validation.
				Default("v1")
			})
		})
	})
})

// User type
var Location = Type("Location", func() {
	Description("Describe location properties.")
	Field(1, "country", String, "User country", func() {
		Pattern(`.+@.+\..{1,6}`)
		Example("ehabterra@hotmail.com")
	})
	Field(2, "date", String, "User current date", func() {
		MaxLength(100)
		Example("Ehab")
	})
	Field(3, "time", String, "User current time", func() {
		MaxLength(100)
		Example("Terra")
	})
	Required("country", "date", "time")
})

// NotFound type
var NotFound = Type("NotFound", func() {
	Description("NotFound is the type returned when attempting to show or delete a user that does not exist.")
	Attribute("message", String, "Message of error", func() {
		Meta("struct:error:name")
		Example("user 1 not found")
		Meta("rpc:tag", "1")
	})
	Field(2, "id", String, "ID of missing user")
	Required("message", "id")
})

var _ = Service("locations", func() {
	Description("The location service gets user current country, date and time.")

	HTTP(func() {
		Path("/location")
	})

	Method("now", func() {
		Description("List all stored users")
		/*Payload(func() {
			Field(1, "view", String, "View to render", func() {
				Enum("default", "tiny")
			})
		})*/
		Result(Location)
		HTTP(func() {
			GET("/now")
			Response(StatusOK)
		})
	})

})

var _ = Service("openapi", func() {
	Meta("swagger:generate", "false")
	HTTP(func() {
		Path("/")
	})
	// Serve the file with relative path ../../gen/http/openapi.json for
	// requests sent to /swagger.json.
	Files("/swagger/{*filepath}", "./public/swagger/")
	Files("/swagger.json", "./gen/http/openapi.json")
})
