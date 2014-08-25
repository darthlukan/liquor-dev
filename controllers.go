package main

import (
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"github.com/stretchr/objx"
	"net/http"
	"path"
)

func handleTemplate(ctx context.Context, template string) error {
	if err := templates.ExecuteTemplate(ctx.HttpResponseWriter(), template, ""); err != nil {
		messages.AddErrorMessage("Could not load template " + template + ": " + err.Error())
		return goweb.Respond.With(ctx, http.StatusInternalServerError, []byte(err.Error()))
	}
	return nil
}

func htmlFileHandler(ctx context.Context) error {
	template := path.Join(ctx.Path().Segments()[1:]...) + ctx.FileExtension()
	return handleTemplate(ctx, template)
}

func indexHandler(ctx context.Context) error {
	return handleTemplate(ctx, "index.html")
}

func pingHandler(ctx context.Context) error {
	return goweb.API.RespondWithData(ctx, objx.MSI("message", "pong!"))
}
