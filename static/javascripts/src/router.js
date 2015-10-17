'use strict';

define([
    'jquery',
    'underscore',
    'backbone',
    'views/home/HomeView',
    'views/projects/ProjectsListView',
    'views/resume/ResumeView'
], function($, _, Backbone, HomeView, ProjectsListView, ResumeView) {
    var AppRouter = Backbone.Router.extend({
        routes: {
            // Define some URL routes
            'projects': 'showProjectsList',
            'resume': 'showResume',
            // Default
            '*actions': 'defaultAction'
        },

        titles: {
            'showProjectsList': 'Projects',
            'showResume': 'Resume',
            'default': 'Rollie Ma - Polyglot Developer from Vancouver, BC'
        },

        initialize: function(options) {
            var that = this;

            this.on('route', function(router, route, params) {
                if (that.titles) {
                    document.title = (
                        that.titles[router] ? that.titles[router] + ' - ' : ''
                    ) + (
                        that.titles['default'] ? that.titles['default'] : ''
                    );
                }
            });
        },

        loadView: function(view) {
            if (this.view) {
                if (this.view.close) {
                    this.view.close();
                } else
                    this.view.remove();
            }
            this.view = view;

            return this.view;
        }
    });

    var initialize = function(options) {
        var router = new AppRouter();

        router
            .on('route:showProjectsList', function() {
                $('.tetris-main-container').append(
                    this.loadView(new ProjectsListView(options)).render().el
                );
            })
            .on('route:showResume', function() {
                $('.tetris-main-container').append(
                    this.loadView(new ResumeView(options)).render().el
                );
            })
            .on('route:defaultAction', function(actions) {
                // no matching route, fall back to HomeView
                $('.tetris-main-container').append(
                    this.loadView(new HomeView(options)).render().el
                );
            });

        try {
          Backbone.history.start({pushState: true});
        } catch(e) { }

        attachGlobalClickListener();
        attachGlobalScrollListener();
        attachGlobalNavMenuToggleListener();

        return router;
    };

    var attachGlobalClickListener = function() {
        // on every click, check if it's an href that can be handled by the router
        $(document).on('click', 'a', function(evt) {
            var a = $(this);
            var rel = a.attr('rel');
            var href = a.attr('href');
            var protocol = this.protocol + '//';

            if (rel == 'external') {
                a.attr('target', '_blank');
                return true;
            }

            if (href.slice(protocol.length) !== protocol &&
                protocol !== 'javascript://' &&
                href.substring(0, 1) !== '#'
            ) {
                evt.preventDefault();
                Backbone.history.navigate(href, {trigger: true});
            }
        });
    };

    var attachGlobalScrollListener = function() {
        $(window).scroll(function() {
            var smallLogoHeight = $('.small-logo').height();
            var bigLogoHeight = $('.big-logo-container').height();
            var navbarHeight = $('.navbar-header').height();

            var smallLogoEndPos = 0;
            var smallSpeed = (smallLogoHeight / bigLogoHeight);
            var ySmall = ($(window).scrollTop() * smallSpeed);
            var smallPadding = navbarHeight - ySmall;

            if (smallPadding > navbarHeight) {
                smallPadding = navbarHeight;
            }
            if (smallPadding < smallLogoEndPos) {
                smallPadding = smallLogoEndPos;
            }
            if (smallPadding < 0) {
                smallPadding = 0;
            }

            $('.small-logo-container').css({
                'width': (smallLogoHeight > smallPadding) ? '50px' : 0,
                'padding-top': smallPadding
            });

            var navOpacity = ySmall / smallLogoHeight;
            if (navOpacity > 1) {
                navOpacity = 1;
            }
            if (navOpacity < 0) {
                navOpacity = 0;
            }

            var shadowOpacity = navOpacity * 0.3;
            if (ySmall > 1) {
                $('.navbar').css({'box-shadow': '0 2px 3px rgba(0,0,0,' + shadowOpacity + ')'});
            } else {
                $('.navbar').css({'box-shadow': 'none'});
            }
        });
    };

    var attachGlobalNavMenuToggleListener = function() {
        var originalPadding = parseFloat($('.big-logo-row').css('padding-top'));

        if (_.isNaN(originalPadding)) {
            originalPadding = 0;
        }

        $('.navbar .collapse')
            .on('shown.bs.collapse hidden.bs.collapse', function() {
                $('.big-logo-row').css({
                    'padding-top': originalPadding + $(this).height() + 'px'
                });
            })
            .on('click', 'a', function(evt) {
                var collapse = $(evt.delegateTarget);
                if (collapse.is(".collapse.in")) {
                    collapse.collapse('hide');
                }
            });
    };

    return {
        initialize: initialize
    };
});
