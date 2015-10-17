'use strict';

define([
    'underscore',
    'backbone'
], function(_, Backbone) {
    var ProjectsListModel = Backbone.Model.extend({
        defaults: {}
    });

    return ProjectsListModel;
});
