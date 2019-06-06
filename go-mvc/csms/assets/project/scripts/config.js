;(function() {
  'use strict'

  var xa = window.xa || (window.xa = {})
  xa.CONSTS = {
    MOD_IFRAME: Symbol('MOD_IFRAME'),
    MOD_SINGLE_PAGE: Symbol('MOD_SINGLE_PAGE')
  }

  xa.$CONFIG = {
    mode: xa.CONSTS.MOD_IFRAME,
    debug: false,
    frameCheck: false,
    frame: './' + encodeURIComponent('框架') + '.html',
    collapsed: window.innerWidth < 1440
  }
})()
