package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/imcsk8/korima-tour/korima_hishow/models"
	"github.com/pkg/errors"
)

// BandsIndex default implementation.
func BandsIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	bands := &models.Bands{}
	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())
	// Retrieve all bands from the database
	// TODO add ACL's
	if err := q.All(bands); err != nil {
		return errors.WithStack(err)
	}
	// Make posts available inside the html template
	c.Set("bands", bands)
	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	return c.Render(200, r.HTML("bands/index.html"))
}

// BandsCreateShow shows the band create page
func BandsCreateShow(c buffalo.Context) error {
	c.Set("band", &models.Band{})
	return c.Render(200, r.HTML("bands/create.html"))
}

// BandsCreate adds a band to the database
func BandsCreate(c buffalo.Context) error {
	// Allocate an empty Band
	band := &models.Band{}
	user := c.Value("current_user").(*models.User)
	// Bind band to the html form elements
	if err := c.Bind(band); err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Validate the data from the html form
	band.OwnerID = user.ID
	verrs, err := tx.ValidateAndCreate(band)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("band", band)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("bands/create"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "New Band added successfully.")
	// and redirect to the index page
	return c.Redirect(302, "/bands/index")
}

// BandsEdit default implementation.
func BandsEdit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	band := &models.Band{}
	if err := tx.Find(band, c.Param("id")); err != nil {
		return c.Error(404, err)
	}
	if err := c.Bind(band); err != nil {
		return errors.WithStack(err)
	}
	verrs, err := tx.ValidateAndUpdate(band)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("band", band)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("bands/detail.html"))
	}
	c.Flash().Add("success", "Band was updated successfully.")
	return c.Redirect(302, "/bands/detail/%s", band.ID)
}

// BandsDelete default implementation.
func BandsDelete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	band := &models.Band{}
	if err := tx.Find(band, c.Param("id")); err != nil {
		return c.Error(404, err)
	}
	// Check if we own the band before deleting
	current_user := c.Value("current_user").(*models.User)
	allowed, flashkey, msg := current_user.AuthorizeDelete(band.OwnerID, c)
	if allowed {
		if err := tx.Destroy(band); err != nil {
			return errors.WithStack(err)
		}
	}
	c.Flash().Add(flashkey, msg)
	return c.Redirect(302, "/bands/index")
}

// BandsDetail default implementation.
func BandsDetail(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	band := &models.Band{}
	if err := tx.Find(band, c.Param("id")); err != nil {
		return c.Error(404, err)
	}
	owner := &models.User{}
	if err := tx.Find(owner, band.OwnerID); err != nil {
		return c.Error(404, err)
	}
	band.Photo = models.GetPhotoName(band.ID.String(), band.Photo)
	c.Set("band", band)
	c.Set("owner", owner)
	return c.Render(200, r.HTML("bands/detail.html"))
}
