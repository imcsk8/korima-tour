package actions

import "github.com/gobuffalo/buffalo"

// BandsIndex default implementation.
func BandsIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("bands/index.html"))
}

// BandsCreateShow shows the band creation form
func BandsCreateShow(c buffalo.Context) error {
	return c.Render(200, r.HTML("bands/create.html"))
}

// BandsCreate default implementation.
func BandsCreate(c buffalo.Context) error {
	return c.Render(200, r.HTML("bands/create.html"))
}

// BandsEdit default implementation.
func BandsEdit(c buffalo.Context) error {
	return c.Render(200, r.HTML("bands/edit.html"))
}

// BandsDelete default implementation.
func BandsDelete(c buffalo.Context) error {
	return c.Render(200, r.HTML("bands/delete.html"))
}

// BandsDetail default implementation.
func BandsDetail(c buffalo.Context) error {
	return c.Render(200, r.HTML("bands/detail.html"))
}
