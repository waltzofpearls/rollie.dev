'use strict';

require(['config'], function() {
    require(['app', 'ga'], function(App, Ga) {
        App.initialize();
        Ga('send', 'pageview');
    });
});
