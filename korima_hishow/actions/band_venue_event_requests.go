package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/imcsk8/korima-tour/korima_hishow/models"
	"github.com/pkg/errors"
)

// BandVenueEventRequestsIndex default implementation.
func BandVenueEventRequestsIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	band_venue_event_requests := &models.BandVenueEventRequests{}
	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())
	// Retrieve all band_venue_event_requests from the database
	// TODO add ACL's
	if err := q.All(band_venue_event_requests); err != nil {
		return errors.WithStack(err)
	}
	// Make posts available inside the html template
	c.Set("band_venue_event_requests", band_venue_event_requests)
	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	return c.Render(200, r.HTML("band_venue_event_requests/index.html"))
}

// BandVenueEventRequestsCreateShow shows the band_venue_event_request create page
func BandVenueEventRequestsCreateShow(c buffalo.Context) error {
	c.Set("band_venue_event_request", &models.BandVenueEventRequest{})
	return c.Render(200, r.HTML("band_venue_event_requests/create.html"))
}

// BandVenueEventRequestsCreate adds a band_venue_event_request to the database
func BandVenueEventRequestsCreate(c buffalo.Context) error {
	// Allocate an empty BandVenueEventRequest
	band_venue_event_request := &models.BandVenueEventRequest{}
	user := c.Value("current_user").(*models.User)
	// Bind band_venue_event_request to the html form elements
	if err := c.Bind(band_venue_event_request); err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Validate the data from the html form
	band_venue_event_request.OwnerID = user.ID
	verrs, err := tx.ValidateAndCreate(band_venue_event_request)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("band_venue_event_request", band_venue_event_request)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("band_venue_event_requests/create"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "Booking has been successfully requested .")
	// and redirect to the index page
	return c.Redirect(302, "/venues/index")
}

// BandVenueEventRequestsEdit default implementation.
func BandVenueEventRequestsEdit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	band_venue_event_request := &models.BandVenueEventRequest{}
	if err := tx.Find(band_venue_event_request, c.Param("id")); err != nil {
		return c.Error(404, err)
	}
	if err := c.Bind(band_venue_event_request); err != nil {
		return errors.WithStack(err)
	}
	verrs, err := tx.ValidateAndUpdate(band_venue_event_request)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("band_venue_event_request", band_venue_event_request)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("band_venue_event_requests/detail.html"))
	}
	c.Flash().Add("success", "BandVenueEventRequest was updated successfully.")
	return c.Redirect(302, "/band_venue_event_requests/detail/%s", band_venue_event_request.ID)
}

// BandVenueEventRequestsDelete default implementation.
func BandVenueEventRequestsDelete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	band_venue_event_request := &models.BandVenueEventRequest{}
	if err := tx.Find(band_venue_event_request, c.Param("id")); err != nil {
		return c.Error(404, err)
	}

	// Check if we own the band_venue_event_request before deleting
	current_user := c.Value("current_user").(*models.User)
	allowed, flashkey, msg := current_user.AuthorizeDelete(band_venue_event_request.OwnerID, c)
	if allowed {
		if err := tx.Destroy(band_venue_event_request); err != nil {
			return errors.WithStack(err)
		}
	}
	c.Flash().Add(flashkey, msg)
	return c.Redirect(302, "/band_venue_event_requests/index")
}

// BandVenueEventRequestsDetail default implementation.
func BandVenueEventRequestsDetail(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	band_venue_event_request := &models.BandVenueEventRequest{}
	if err := tx.Find(band_venue_event_request, c.Param("id")); err != nil {
		return c.Error(404, err)
	}
	owner := &models.User{}
	if err := tx.Find(owner, band_venue_event_request.OwnerID); err != nil {
		return c.Error(404, err)
	}
	// Get current user
	//current_user := c.Value("current_user").(*models.User)
	c.Set("band_venue_event_request", band_venue_event_request)
	// Choose template
	/* FIX LATER if band_venue_event_request.OwnerID == current_user.ID {
		c.Set("owner", owner)
		return c.Render(200, r.HTML("band_venue_event_requests/detail.html"))
	} else {
		return c.Render(200, r.HTML("band_venue_event_requests/show.html"))

	}*/
	return c.Render(200, r.HTML("band_venue_event_requests/detail.html"))
}

// BandVenueEventRequestsApprove approves the request for booking
func BandVenueEventRequestsApprove(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	band_venue_event_request := &models.BandVenueEventRequest{}
	if err := tx.Find(band_venue_event_request, c.Param("request_id")); err != nil {
		return c.Error(404, err)
	}
	band_venue_event_request.Status = 1
	verrs, err := tx.ValidateAndUpdate(band_venue_event_request)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("band_venue_event_request", band_venue_event_request)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("band_venue_event_requests/detail.html"))
	}
	return nil
}
