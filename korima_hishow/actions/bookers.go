package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/imcsk8/korima-tour/korima_hishow/models"
	"github.com/pkg/errors"
)

// BookersIndex default implementation.
func BookersIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	bookers := &models.Bookers{}
	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())
	// Retrieve all bookers from the database
	// TODO add ACL's
	if err := q.All(bookers); err != nil {
		return errors.WithStack(err)
	}
	// Make posts available inside the html template
	c.Set("bookers", bookers)
	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	return c.Render(200, r.HTML("bookers/index.html"))
}

// BookersCreateShow shows the booker create page
func BookersCreateShow(c buffalo.Context) error {
	c.Set("booker", &models.Booker{})
	return c.Render(200, r.HTML("bookers/create.html"))
}

// BookersCreate adds a booker to the database
func BookersCreate(c buffalo.Context) error {
	// Allocate an empty Booker
	booker := &models.Booker{}
	user := c.Value("current_user").(*models.User)
	// Bind booker to the html form elements
	if err := c.Bind(booker); err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Validate the data from the html form
	booker.OwnerID = user.ID
	verrs, err := tx.ValidateAndCreate(booker)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("booker", booker)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("bookers/create"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "New Booker added successfully.")
	// and redirect to the index page
	return c.Redirect(302, "/bookers/index")
}

// BookersEdit default implementation.
func BookersEdit(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	booker := &models.Booker{}
	if err := tx.Find(booker, c.Param("id")); err != nil {
		return c.Error(404, err)
	}
	if err := c.Bind(booker); err != nil {
		return errors.WithStack(err)
	}
	verrs, err := tx.ValidateAndUpdate(booker)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("booker", booker)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("bookers/detail.html"))
	}
	c.Flash().Add("success", "Booker was updated successfully.")
	return c.Redirect(302, "/bookers/detail/%s", booker.ID)
}

// BookersDelete default implementation.
func BookersDelete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	booker := &models.Booker{}
	if err := tx.Find(booker, c.Param("id")); err != nil {
		return c.Error(404, err)
	}
	// Check if we own the booker before deleting
	current_user := c.Value("current_user").(*models.User)
	allowed, flashkey, msg := current_user.AuthorizeDelete(booker.OwnerID, c)
	if allowed {
		if err := tx.Destroy(booker); err != nil {
			return errors.WithStack(err)
		}
	}
	c.Flash().Add(flashkey, msg)
	return c.Redirect(302, "/bookers/index")
}

// BookersDetail default implementation.
func BookersDetail(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	booker := &models.Booker{}
	if err := tx.Find(booker, c.Param("id")); err != nil {
		return c.Error(404, err)
	}
	owner := &models.User{}
	if err := tx.Find(owner, booker.OwnerID); err != nil {
		return c.Error(404, err)
	}
	booker.Photo = models.GetPhotoName(booker.ID.String(), booker.Photo)
	c.Set("booker", booker)
	c.Set("owner", owner)
	return c.Render(200, r.HTML("bookers/detail.html"))
}
