'use strict';

define([
    'underscore',
    'backbone',
    'models/projects/ProjectsListModel'
], function(_, Backbone, ProjectsListModel) {
    var ProjectsListCollection = Backbone.Collection.extend({
        model: ProjectsListModel,

        url: '/api/projects'
    });

    return ProjectsListCollection;
});
