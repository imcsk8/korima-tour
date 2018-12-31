package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/pop/popmw"
	"github.com/gobuffalo/envy"
	csrf "github.com/gobuffalo/mw-csrf"
	forcessl "github.com/gobuffalo/mw-forcessl"
	i18n "github.com/gobuffalo/mw-i18n"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/gobuffalo/packr"
	"github.com/imcsk8/korima-tour/korima_hishow/models"
	"github.com/unrolled/secure"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_korima_hishow_session",
		})

		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))

		// Setup and use translations:
		app.Use(translations())

		// Set current user to the context
		app.Use(SetCurrentUser)

		app.GET("/", HomeHandler)
		app.GET("/routes", RouteHandler)

		//app.Resource("/users", UsersResource{&buffalo.BaseResource{}})
		auth := app.Group("/users")
		auth.GET("/", UsersResource{&buffalo.BaseResource{}}.List)
		auth.GET("/list", UsersResource{&buffalo.BaseResource{}}.List)
		auth.GET("/index", UsersResource{&buffalo.BaseResource{}}.List)
		auth.GET("/new", UsersResource{&buffalo.BaseResource{}}.New)
		auth.POST("/create", UsersResource{&buffalo.BaseResource{}}.Create)
		auth.GET("/detail/{id}", UsersResource{&buffalo.BaseResource{}}.Edit)
		auth.GET("/edit/{id}", UsersResource{&buffalo.BaseResource{}}.Edit)
		auth.POST("/update/{id}", UsersResource{&buffalo.BaseResource{}}.Update)
		auth.GET("/delete/{id}", UsersResource{&buffalo.BaseResource{}}.Destroy)
		auth.DELETE("/delete/{id}", UsersResource{&buffalo.BaseResource{}}.Destroy)
		//auth.GET("/index", UsersResource{&buffalo.BaseResource{}})
		auth.GET("/register", UsersRegisterGet)
		auth.POST("/register", UsersRegisterPost)
		auth.GET("/login", UsersLoginShow)
		auth.POST("/login", UsersLogin)
		auth.GET("/logout", UsersLogout)

		venuesGroup := app.Group("/venues")
		venuesGroup.GET("/index", VenuesIndex)
		venuesGroup.GET("/create", VenuesCreateShow)
		venuesGroup.POST("/create", VenuesCreate)
		venuesGroup.GET("/edit/{id}", VenuesDetail)
		venuesGroup.POST("/edit/{id}", VenuesEdit)
		venuesGroup.GET("/delete/{id}", VenuesDelete)
		venuesGroup.GET("/detail/{id}", VenuesDetail)

		bandsGroup := app.Group("/bands")
		bandsGroup.GET("/index", BandsIndex)
		bandsGroup.GET("/create", BandsCreateShow)
		bandsGroup.POST("/create", BandsCreate)
		bandsGroup.GET("/edit/{id}", BandsDetail)
		bandsGroup.POST("/edit/{id}", BandsEdit)
		bandsGroup.GET("/delete/{id}", BandsDelete)
		bandsGroup.GET("/detail/{id}", BandsDetail)

		bookersGroup := app.Group("/bookers")
		bookersGroup.GET("/index", BookersIndex)
		bookersGroup.GET("/create", BookersCreateShow)
		bookersGroup.POST("/create", BookersCreate)
		bookersGroup.GET("/edit/{id}", BookersDetail)
		bookersGroup.POST("/edit/{id}", BookersEdit)
		bookersGroup.GET("/delete/{id}", BookersDelete)
		bookersGroup.GET("/detail/{id}", BookersDetail)

		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
