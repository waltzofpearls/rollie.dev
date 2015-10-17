'use strict';

window.GoogleAnalyticsObject = '__ga__';
window.__ga__ = {
    q: [['create', 'UA-64152368-2', 'auto']],
    l: Date.now()
};

require.config({
    shim: {
        'bootstrap': {deps: ['jquery']},
        'cal-heatmap': {deps: ['d3']},
        'ga': {exports: '__ga__'}
    },
    paths: {
        'backbone': '../bower/backbone/backbone',
        'bootstrap': '../bower/bootstrap/dist/js/bootstrap',
        'cal-heatmap': 'libs/cal-heatmap',
        'd3': '../bower/d3/d3',
        'domReady': '../bower/requirejs-domready/domReady',
        'ga': '//www.google-analytics.com/analytics',
        'jquery': '../bower/jquery/dist/jquery',
        'marked': '../bower/marked/lib/marked',
        'templates': '../../templates',
        'text': '../bower/requirejs-text/text',
        'underscore': '../bower/underscore/underscore'
    }
});
