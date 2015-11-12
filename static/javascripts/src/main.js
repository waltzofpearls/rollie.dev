'use strict';

requirejs(['config'], function() {
    requirejs(['app', 'ga'], function(App, Ga) {
        App.initialize();
        Ga('send', 'pageview');
    });
});
