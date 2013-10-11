define [
  "devfeed",
  "tpl!apps/project/list/templates/list.tpl",
  "tpl!apps/project/list/templates/empty.tpl",
  "tpl!apps/project/list/templates/project.tpl"
], (Devfeed, listTpl, emptyTpl, projectTpl) ->

  Devfeed.module "ProjectApp.List.View", (View, Devfeed, Backbone, Marionette, $, _) ->

    class View.Empty extends Marionette.ItemView
      className: "empty"
      template: emptyTpl

    class View.Project extends Marionette.ItemView
      tagName: "li"
      template: projectTpl
      events:
        "click a": "projectClicked"

      projectClicked: (e) ->
        e.preventDefault()
        Devfeed.trigger("project:show", @model.get("id"))

    class View.List extends Marionette.CompositeView
      id: "project-list"
      className: "row collapse"
      template: listTpl
      emptyView: View.Empty
      itemView: View.Project
      itemViewContainer: ".projects"
      events:
        "click .setup-msg a": "setupClicked"

      setupClicked: (e) ->
        e.preventDefault()
        Devfeed.trigger("settings:general")

      onCompositeRendered: ->
        userSession = Devfeed.request("user:session")
        if not userSession.get("apitoken")
          @$(".empty-msg").addClass("hide")
          @$(".setup-msg").removeClass("hide")

  return Devfeed.ProjectApp.List.View