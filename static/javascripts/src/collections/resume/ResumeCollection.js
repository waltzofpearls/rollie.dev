'use strict';

define([
    'underscore',
    'backbone',
    'models/resume/ResumeModel'
], function(_, Backbone, ResumeModel) {
    var ResumeCollection = Backbone.Collection.extend({
        model: ResumeModel,

        url: '/static/data/resume.json'
    });

    return ResumeCollection;
});
