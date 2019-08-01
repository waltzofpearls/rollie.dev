'use strict';

define([
    'jquery',
    'underscore',
    'backbone',
    'views/home/GithubContribView',
    'text!templates/home/homeTemplate.html'
], function($, _, Backbone, GithubContribView, homeTemplate) {
    var HomeView = Backbone.View.extend({
        tagName: 'div',
        className: 'tetris-page-home',

        xhr: null,
        subViews: {},

        initialize: function(options) {
            this.app = options.app;
            this.tube = options.tube;
            this.subViews.githubContribView = new GithubContribView(options);
        },

        render: function() {
            this.$el.html(_.template(homeTemplate)({
                avatarRandNum: this.getRandomInt(1, 5)
            }));
            this.$el.append(this.subViews.githubContribView.$el);
            this.subViews.githubContribView.render();

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
        },

        getRandomInt: function(min, max) {
            return Math.floor(Math.random() * (max+1 - min)) + min;
        }
    });

    return HomeView;
});
