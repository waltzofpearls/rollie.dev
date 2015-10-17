'use strict';

define([
    'jquery',
    'underscore',
    'backbone',
], function($, _, Backbone) {
    return {
        bootstrap: function() {
            var img = $('img');

            _.each(this.preload, function(src) {
                img.clone().attr('src', src);
            });
        },

        preload: {
            loader: '/images/spinner1.gif',
        },

        template: {
            loader: '\
                <div class="col-md-12 tetris-loader">\
                    <img src="<%= loader %>"> Loading...\
                </div>\
            '
        }
    };
});
