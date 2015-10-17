'use strict';

define([
    'underscore',
    'backbone'
], function(_, Backbone) {
    var ResumeModel = Backbone.Model.extend({
        defaults: {}
    });

    return ResumeModel;
});
