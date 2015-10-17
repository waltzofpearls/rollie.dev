'use strict';

define([
    'underscore',
    'backbone',
    'models/home/GithubContribModel'
], function(_, Backbone, GithubContribModel) {
    var GithubContribCollection = Backbone.Collection.extend({
        model: GithubContribModel,

        url: '/api/contributions'
    });

    return GithubContribCollection;
});
