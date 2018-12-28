package actions

import "github.com/gobuffalo/buffalo"

// BandsIndex default implementation.
func BandsIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("bandas/index.html"))
}

// BandsCreate default implementation.
func BandsCreate(c buffalo.Context) error {
	return c.Render(200, r.HTML("bandas/create.html"))
}

// BandsEdit default implementation.
func BandsEdit(c buffalo.Context) error {
	return c.Render(200, r.HTML("bandas/edit.html"))
}

// BandsDelete default implementation.
func BandsDelete(c buffalo.Context) error {
	return c.Render(200, r.HTML("bandas/delete.html"))
}

// BandsDetail default implementation.
func BandsDetail(c buffalo.Context) error {
	return c.Render(200, r.HTML("bandas/detail.html"))
}
