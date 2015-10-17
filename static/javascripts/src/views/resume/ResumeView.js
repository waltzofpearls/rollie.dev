'use strict';

define([
    'jquery',
    'underscore',
    'backbone',
    'marked',
    'utils/Partial',
    'collections/resume/ResumeCollection',
    'text!templates/resume/resumeTemplate.html'
], function($, _, Backbone, Marked, Partial, ResumeCollection, resumeTemplate) {
    var ResumeView = Backbone.View.extend({
        tagName: 'div',
        className: 'tetris-page-resume',

        xhr: null,

        initialize: function(options) {
            this.app = options.app;
            this.tube = options.tube;
        },

        render: function() {
            var that = this;

            this.collection = new ResumeCollection();
            this.xhr = this.collection.fetch({
                success: function(collection, response) {
                    that.$el.html(_.template(resumeTemplate)({
                        _: _,
                        Marked: Marked,
                        resume: collection.models[0]
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

    return ResumeView;
});
