'use strict';

define([
    'jquery',
    'underscore',
    'backbone',
    'router',
    'bootstrap'
], function($, _, Backbone, Router) {
    var app = {};
    var tube = _.extend({}, Backbone.Events);

    var initialize = function() {
        return Router.initialize({app: app, tube: tube});
    };

    return {
        app: app,
        tube: tube,
        initialize: initialize
    };
});
