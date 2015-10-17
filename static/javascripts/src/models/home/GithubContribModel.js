'use strict';

define([
    'underscore',
    'backbone'
], function(_, Backbone) {
    var GithubContribModel = Backbone.Model.extend({
        defaults: {}
    });

    return GithubContribModel;
});
