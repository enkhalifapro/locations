package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("location", func() {
	Title("Location Service")
	Description("Service for manipulate locations")
	Server("location", func() {
		Description("location hosts the Locations Service.")

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

// Location type
var Location = Type("Location", func() {
	Description("Describe location properties.")
	Field(1, "country", String, "User country", func() {
		Example("Egypt")
	})
	Field(2, "date", String, "User current date", func() {
		Example("dd-mm-yyyy")
	})
	Field(3, "time", String, "User current time", func() {
		Example("hh:mm:ss")
	})
	Required("country", "date", "time")
})

var _ = Service("locations", func() {
	Description("The location service gets user current country, date and time.")

	HTTP(func() {
		Path("/location")
	})

	Method("now", func() {
		Description("Get client IP location")
		Payload(func() {
			Attribute("X-Forwarded-For", String)
		})
		Result(Location)
		HTTP(func() {
			GET("/now")
			Headers(func() {
				Header("X-Forwarded-For", String)
			})
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
