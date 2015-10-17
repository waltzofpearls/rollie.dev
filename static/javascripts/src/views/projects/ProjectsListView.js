'use strict';

define([
    'jquery',
    'underscore',
    'backbone',
    'utils/Partial',
    'collections/projects/ProjectsListCollection',
    'text!templates/projects/projectsListTemplate.html'
], function($, _, Backbone, Partial, ProjectsListCollection, projectsListTemplate) {
    var ProjectsListView = Backbone.View.extend({
        tagName: 'div',
        className: 'tetris-page-projects',

        xhr: null,

        initialize: function(options) {
            this.app = options.app;
            this.tube = options.tube;
        },

        render: function() {
            var that = this;

            this.collection = new ProjectsListCollection();
            this.xhr = this.collection.fetch({
                success: function(collection, response) {
                    that.$el.html(_.template(projectsListTemplate)({
                        _: _,
                        projects: collection.models
                    }));
                    that.xhr = null;
                }
            });

            this.$el.html(_.template(Partial.template.loader)(Partial.preload));

            return this;
        },

        close: function() {
            if (this.xhr !== null) {
                this.xhr.abort();
            }
            _.each(this.subViews, function(view) {
                view.remove();
            });
            this.remove();
        }
    });

    return ProjectsListView;
});
