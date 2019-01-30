package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/imcsk8/korima-tour/korima_hishow/models"
	"github.com/pkg/errors"
)

// VenuesIndex default implementation.
func VenuesIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	venues := &models.Venues{}
	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())
	// Retrieve all venues from the database
	// TODO add ACL's
	if err := q.All(venues); err != nil {
		return errors.WithStack(err)
	}
	// Make posts available inside the html template
	c.Set("venues", venues)
	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
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
	return c.Redirect(302, "/venues/index")
}

// VenuesEdit default implementation.
func VenuesEdit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	venue := &models.Venue{}
	if err := tx.Find(venue, c.Param("id")); err != nil {
		return c.Error(404, err)
	}
	if err := c.Bind(venue); err != nil {
		return errors.WithStack(err)
	}
	verrs, err := tx.ValidateAndUpdate(venue)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("venue", venue)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("venues/detail.html"))
	}
	c.Flash().Add("success", "Venue was updated successfully.")
	return c.Redirect(302, "/venues/detail/%s", venue.ID)
}

// VenuesDelete default implementation.
func VenuesDelete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	venue := &models.Venue{}
	if err := tx.Find(venue, c.Param("id")); err != nil {
		return c.Error(404, err)
	}

	// Check if we own the venue before deleting
	current_user := c.Value("current_user").(*models.User)
	allowed, flashkey, msg := current_user.AuthorizeDelete(venue.OwnerID, c)
	if allowed {
		if err := tx.Destroy(venue); err != nil {
			return errors.WithStack(err)
		}
	}
	c.Flash().Add(flashkey, msg)
	return c.Redirect(302, "/venues/index")
}

// VenuesDetail default implementation.
func VenuesDetail(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	venue := &models.Venue{}
	if err := tx.Find(venue, c.Param("id")); err != nil {
		return c.Error(404, err)
	}
	owner := &models.User{}
	if err := tx.Find(owner, venue.OwnerID); err != nil {
		return c.Error(404, err)
	}
	// Get current user
	current_user := c.Value("current_user").(*models.User)
	venue.Photo = models.GetPhotoName(venue.ID.String(), venue.Photo)
	c.Set("venue", venue)
	// Choose template
	if venue.OwnerID == current_user.ID {
		c.Set("owner", owner)
		return c.Render(200, r.HTML("venues/detail.html"))
	} else {
		return c.Render(200, r.HTML("venues/show.html"))

	}
}
