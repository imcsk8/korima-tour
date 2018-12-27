package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/imcsk8/korima-tour/korima_hishow/models"
	"github.com/pkg/errors"
)

// VenuesIndex default implementation.
func VenuesIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("venues/index.html"))
}

// VenuesCreateShow shows the venue create page
func VenuesCreateShow(c buffalo.Context) error {
	c.Set("venue", &models.Venue{})
	return c.Render(200, r.HTML("venues/create.html"))
}

// VenuesCreate adds a venue to the database
func VenuesCreate(c buffalo.Context) error {
	// Allocate an empty Venue
	venue := &models.Venue{}
	user := c.Value("current_user").(*models.User)
	// Bind venue to the html form elements
	if err := c.Bind(venue); err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Validate the data from the html form
	venue.OwnerID = user.ID
	verrs, err := tx.ValidateAndCreate(venue)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("venue", venue)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("venues/create"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "New Venue added successfully.")
	// and redirect to the index page
	return c.Redirect(302, "/")
}

// VenuesEdit default implementation.
func VenuesEdit(c buffalo.Context) error {
	return c.Render(200, r.HTML("venues/edit.html"))
}

// VenuesDelete default implementation.
func VenuesDelete(c buffalo.Context) error {
	return c.Render(200, r.HTML("venues/delete.html"))
}

// VenuesDetail default implementation.
func VenuesDetail(c buffalo.Context) error {
	return c.Render(200, r.HTML("venues/detail.html"))
}
