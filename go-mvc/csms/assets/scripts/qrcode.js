;(function() {
  'use strict'

  var Vue = window.Vue
  var VueQrcode = window.VueQrcode
  var Promise = window.Promise

  function loadImg(src) {
    var img = new Image()
    img.src = src
    return new Promise(function(resolve, reject) {
      img.onload = function() {
        resolve(img)
      }
      img.onerror = reject
    })
  }

  function drawLogo(logo, size) {
    var canvas = document.createElement('canvas')
    canvas.width = canvas.height = size
    var context = canvas.getContext('2d')
    context.arc(size / 2, size / 2, size / 2 - 2, 0, Math.PI * 2)
    context.clip()
    context.fillStyle = '#fff'
    context.fillRect(0, 0, size, size)

    return loadImg(logo).then(function(logo) {
      context.drawImage(logo, 0, 0, size, size)
      return canvas
    })
  }

  function canvasToImg(canvas) {
    return canvas.toDataURL()
  }

  var XaQrcode = {
    name: 'XaQrcode',
    template: [
      '<div>',
      '<qrcode :value="path" :options="qrOptions" hidden></qrcode>',
      '<img :src="image">',
      '</div>'
    ].join(''),
    components: {
      qrcode: VueQrcode
    },
    props: {
      value: String,
      options: {
        type: Object,
        default: function() {
          return {}
        }
      },
      logo: {
        type: Object,
        default: function() {
          return {}
        }
      }
    },
    data: function() {
      return {
        qrOptions: {},
        path: '',
        image: ''
      }
    },
    watch: {
      value: function(value, oldValue) {
        this.draw()
      },
      options: {
        handler: function(value, oldValue) {
          this.draw()
        },
        deep: true
      },
      logo: {
        handler: function(value, oldValue) {
          this.draw()
        },
        deep: true
      }
    },
    mounted: function() {
      this.draw()
    },
    methods: {
      triggerRedraw: function() {
        var component = this
        component.path = ''
        component.qrOptions = {}

        return new Promise(function(resolve) {
          Vue.nextTick(function() {
            component.path = component.value
            component.qrOptions = component.options

            Vue.nextTick(resolve)
          })
        })
      },
      draw: function() {
        this.triggerRedraw().then(
          function() {
            var canvas = this.$el.children[0]
            var context = canvas.getContext('2d')
            var size = canvas.width
            var logo = this.logo.src
            var logoSize = this.logo.size

            return Promise.resolve()
              .then(function() {
                if (logo) {
                  return drawLogo(logo, logoSize).then(function(logo) {
                    context.drawImage(
                      logo,
                      size / 2 - logoSize / 2,
                      size / 2 - logoSize / 2
                    )
                  })
                }
                return null
              })
              .then(function() {
                return canvasToImg(canvas)
              })
              .then(
                function(url) {
                  this.image = url
                }.bind(this)
              )
          }.bind(this)
        )
      }
    }
  }

  Vue.component(XaQrcode.name, XaQrcode)
})()
