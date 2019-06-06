;(function() {
  'use strict'

  var UE = window.UE
  var Vue = window.Vue
  var xa = window.xa

  UE.registerUI('autosave', xa.utils.noop)

  var XaEditor = {
    name: 'XaEditor',
    template: [
      '<textarea v-model="currentValue" :id="id" style="width: 100%">',
      '</textarea>'
    ].join(''),
    props: {
      value: String,
      editorConfig: {
        // UEditor 配置项
        type: Object,
        default: function() {
          return {}
        }
      }
    },
    model: {
      prop: 'value',
      event: 'contentchange'
    },
    computed: {
      currentValue: {
        get: function() {
          return this.value
        },
        set: function(value) {
          return (this.value = value)
        }
      }
    },
    watch: {
      value: function(value, oldValue) {
        try {
          if (value !== this.instance.getContent()) {
            this.instance.setContent(value)
          }
        } catch (_) {}
      }
    },
    data: function() {
      return {
        id:
          'editor_' +
          Math.random()
            .toString(36)
            .slice(2, 6),
        instance: null
      }
    },
    created: function() {
      this.initEditor()
    },
    beforeDestroy: function() {
      if (this.instance !== null && this.instance.destroy) {
        this.instance.destroy()
      }
    },
    methods: {
      initEditor: function() {
        this.$nextTick(
          function() {
            this.instance = UE.getEditor(this.id, this.editorConfig)

            this.instance.addListener(
              'ready',
              function() {
                this.$emit('ready', this.instance)
                fixSimpleUploadAccept(this.instance)
                fixMutiImageUploader(this.instance)
              }.bind(this)
            )

            this.instance.addListener(
              'contentChange',
              function() {
                if (this.currentValue !== this.instance.getContent()) {
                  this.currentValue = this.instance.getContent()
                  this.$emit('contentchange', this.currentValue)
                }
              }.bind(this)
            )
          }.bind(this)
        )
      }
    }
  }

  function fixSimpleUploadAccept(editor) {
    var btn = editor.container.querySelector('.edui-for-simpleupload')

    if (!btn) {
      return
    }

    var timer = window.setInterval(function() {
      try {
        var frame = btn.querySelector('iframe')
        var input = frame.contentWindow.document.getElementsByTagName(
          'input'
        )[0]
        input.accept = 'image/png,image/jpeg,image/gif'
        window.clearInterval(timer)
      } catch (_) {}
    }, 500)
  }

  function fixMutiImageUploader(editor) {
    var dialog = editor.getDialog('insertimage')

    if (!dialog) {
      return
    }

    dialog.iframeUrl = dialog.iframeUrl.replace(
      '/image.html',
      '/image-modified.html'
    )
  }

  Vue.component(XaEditor.name, XaEditor)
})()
