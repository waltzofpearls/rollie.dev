'use strict';

define([
    'jquery',
    'underscore',
    'backbone',
    'cal-heatmap',
    'utils/Partial',
    'collections/home/GithubContribCollection',
    'text!templates/home/githubContribTemplate.html'
], function(
    $, _, Backbone, CalHeatMap, Partial, GithubContribCollection, githubContribTemplate
) {
    var ResumeView = Backbone.View.extend({
        tagName: 'div',
        className: 'tetris-view-github-contrib',

        xhr: null,

        cal: null,
        data: {},

        initialize: function(options) {
            this.app = options.app;
            this.tube = options.tube;

            $(window).on('resize.github-contrib-view', _.bind(this.handleWindowResize, this));
        },

        handleWindowResize: function() {
            this.createNewCalHeatMap();
        },

        createNewCalHeatMap: function() {
            if (_.isEmpty(this.data)) {
                return;
            }

            if (this.cal !== null) {
                this.cal = this.cal.destroy();
            }

            this.cal = new CalHeatMap();
            this.cal.init({
                itemSelector: this.$('.tetris-heatmap')[0],
                domain: 'month',
                data: this.data,
                start: this.getApplicableStart(),
                cellSize: 15,
                range: this.getApplicableRange(),
                legend: [2, 4, 6, 8],
                legendCellSize: 15,
                legendCellPadding: 4,
                legendHorizontalPosition: 'right',
                itemName: ['contribution', 'contributions'],
                tooltip: true,
                considerMissingDataAsZero: true,
                animationDuration: 0
            });
        },

        getApplicableRange: function() {
            return Math.ceil(
                this.$el.width() / 104
            );
        },

        getApplicableStart: function() {
            var date = new Date();
            date.setMonth(
                date.getMonth() - this.getApplicableRange() + 1
            );
            return date;
        },

        render: function() {
            var that = this;

            this.collection = new GithubContribCollection();
            this.xhr = this.collection.fetch({
                success: function(collection, response) {
                    that.$el.html(_.template(githubContribTemplate)({
                        _: _
                    }));

                    that.data = collection.models[0].toJSON();
                    that.createNewCalHeatMap();

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

            $(window).off('window.github-contrib-view');
        }
    });

    return ResumeView;
});
