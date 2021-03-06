package controllers

import (
	"net/http"
	"strconv"

	log "github.com/cihub/seelog"
	"github.com/marconi/devfeed/core"
	"github.com/marconi/devfeed/db"
	"github.com/marconi/devfeed/utils"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
)

type ProjController struct{}

func (c *ProjController) ReadMany(ctx context.Context) error {
	user, isLoggedIn := IsLoggedIn(ctx)
	if !isLoggedIn {
		return goweb.Respond.WithStatus(ctx, http.StatusUnauthorized)
	}
	projects, err := user.GetProjects()
	if err != nil {
		log.Error("Unable to get projects: ", err)
		return goweb.Respond.WithStatus(ctx, http.StatusInternalServerError)
	}
	return goweb.API.RespondWithData(ctx, projects)
}

func (c *ProjController) Read(id string, ctx context.Context) error {
	_, isLoggedIn := IsLoggedIn(ctx)
	if !isLoggedIn {
		return goweb.Respond.WithStatus(ctx, http.StatusUnauthorized)
	}
	projectId, _ := strconv.Atoi(id)
	project, err := db.GetProjectById(projectId)
	if err != nil {
		log.Error("Unable to get project: ", err)
		return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
	}

	// redirect to home page if project being access is not yet synced
	if !project.IsSynced {
		r := ctx.HttpRequest()
		session, _ := core.GetSession(r)
		session.AddFlash("The project you're trying to access is not yet synced")
		if err := session.Save(r, ctx.HttpResponseWriter()); err != nil {
			log.Error("Unable to save session: ", err)
		}
		data := struct {
			RedirectTo string `json:"redirect_to"`
		}{"/"}
		return goweb.API.RespondWithData(ctx, data)
	}

	stories, err := project.GetStories(utils.NewPaging(0, core.Config.App.StoriesPaging))
	if err != nil {
		log.Error("Unable to get stories for project ", id, " : ", err)
	}

	data := struct {
		*db.Project
		Stories []*db.Story `json:"stories"`
	}{
		project,
		stories,
	}
	return goweb.API.RespondWithData(ctx, data)
}
