/* global _: true  */
;(function() {
  'use strict'

  var window = new Function('return this')() // eslint-disable-line no-new-func
  var xa = window.xa || (window.xa = {})

  var config = xa.$CONFIG
  var Vue = window.Vue
  var document = window.document
  var body = document.body
  var parent = window.parent
  var location = window.location
  var top = window.top

  if (!body) {
    var scripts = document.scripts
    window.alert(
      '请在 body 中 加载此脚本\n' + scripts[scripts.length - 1].src + ''
    )
    return
  }

  var normalizeUrl = (function(a) {
    return function(url) {
      a.href = url
      return a.href
    }
  })(document.createElement('a'))

  // ie normalizeUrl didn't encode path
  // this function simply decode url to compare,
  // but not correct!
  function isSameUrl(url1, url2) {
    return (
      url1 === url2 ||
      decodeURIComponent(normalizeUrl(url1)) ===
        decodeURIComponent(normalizeUrl(url2))
    )
  }

  var isFrameWindow = (function(frameUrl) {
    frameUrl = normalizeUrl(frameUrl)

    return function(win) {
      try {
        if (win.xa && win.xa.IS_FRAME) {
          return true
        }
        return isSameUrl(win.location.href.split('#')[0], frameUrl)
      } catch (e) {
        return false
      }
    }
  })(config.frame)

  var isFrameMode = config.mode === xa.CONSTS.MOD_IFRAME
  var frame = (function() {
    if (config.mode !== xa.CONSTS.MOD_IFRAME) {
      return
    }
    var win = window.self

    while (win !== top) {
      if (isFrameWindow(win)) {
        return win
      }
      win = win.parent
    }

    return isFrameWindow(win) ? win : null
  })()
  var isFrame = window === frame

  /* eslint strict: 0, no-unused-vars: 0, no-empty-function: 0 */
  /* global frame: true, normalizeUrl: true  */
  var _ = (function() {
    // utils
    var _ = xa.utils || (xa.utils = {})

    var _toString = Object.prototype.toString

    // 获取对象类型 typeof
    _.type = function(obj) {
      return _toString.call(obj).slice(8, -1)
    }

    _.reject = Promise.reject.bind(Promise)
    _.resolve = Promise.resolve.bind(Promise)

    // 延迟
    _.delay = function(time) {
      return new Promise(function(resolve, reject) {
        window.setTimeout(function() {
          resolve(time)
        }, time)
      })
    }

    // 减号分割大写字母
    _.kebabCase = function(s) {
      return s.replace(/[A-Z]/g, function($0) {
        return '-' + $0.toLowerCase()
      })
    }

    // on事件
    function on(el, type, fn, config) {
      el.addEventListener(type, fn, config || false)
    }

    // off事件
    function off(el, type, fn) {
      el.removeEventListener(type, fn)
    }

    _.on = on
    _.off = off

    // vue使用组件
    'loading msgbox alert confirm prompt prompt notify message'
      .split(' ')
      .forEach(function(prop) {
        _[prop] = (frame || window).Vue.prototype['$' + prop]
      })

    _.Message = (frame || window).Message
    _.MessageBox = (frame || window).MessageBox

    function _forIn(obj, fn) {
      obj = obj || {}
      for (var key in obj) {
        if (obj.hasOwnProperty(key)) {
          fn.call(obj, obj[key], key, obj)
        }
      }
    }

    function isPrimitive(obj) {
      var type = _.type(obj)
      return (
        type === 'String' ||
        type === 'Number' ||
        type === 'Boolean' ||
        type === 'Symbol' ||
        typeof obj === 'undefined' ||
        obj === null
      )
    }

    // 与obj2Param功能相同
    _.param = function(data) {
      if (!data || isPrimitive(data)) {
        return '' + data
      }

      var param = []
      _forIn(data, function(value, key) {
        param.push(encodeURIComponent(key) + '=' + encodeURIComponent(value))
      })

      return param.join('&')
    }

    // 设置默认颜色组
    var defaultChartColors = (xa.defaultChartColors = [
      '#19d4ae',
      '#5867c3',
      '#1d8ce0',
      '#f7ba2a',
      '#13ce66',
      '#8acc6d',
      '#20a0ff',
      '#f6744c',
      '#fa6e86',
      '#3de1eb',
      '#8f74f4',
      '#ec5ea8'
    ])

    // vue使用v-chart,echarts
    _.setDefaultChartColor = function(colors) {
      var VeIndex = window.VeIndex

      if (!VeIndex) {
        return
      }

      colors = colors || defaultChartColors
      var components = 'Bar Histogram Line Pie Ring Waterfall Funnel Radar Chart Map Bmap'.split(
        ' '
      )

      VeIndex.VeBar.mixins[0].computed.chartColor = function() {
        return this.colors || (this.theme && this.theme.color) || colors
      }

      components.forEach(function(component) {
        VeIndex['Ve' + component]._Ctor = null
      })

      Vue.use(VeIndex)
    }

    _.false = function() {
      return false
    }

    _.true = function() {
      return true
    }

    _.noop = function() {}

    _.identity = function(v) {
      return v
    }

    _.normalize = normalizeUrl

    // 封装vuenexttick
    _.nextTick = function(fn) {
      var app = this
      var args = arguments
      return function() {
        Vue.nextTick(function() {
          fn.apply(app, args)
        })
      }
    }

    // low-hign之间随机数
    _.random = function(low, high) {
      if (arguments.length === 1) {
        high = low
        low = 0
      }
      return Math.floor(low + Math.random() * (high - low))
    }

    // xa.utils.random.int(1ow, hign)
    _.random.int = _.random

    // xa.utils.random.text(length)
    _.random.text = function randomText(length) {
      length = length || 8
      var str = Math.random()
        .toString(36)
        .slice(2)

      while (str.length < length) {
        str += randomText(length - str.length)
      }

      return str.slice(0, length)
    }

    // xa.utils.random.boolean()
    _.random.boolean = function() {
      return Math.random() > 0.5
    }

    // xa.utils.random.array(length)
    _.random.array = function(length) {
      return Array.from({length: length})
    }

    // 获取数组随机一个
    _.randomOne = function(arr) {
      return arr[_.random(0, arr.length)]
    }

    // ???
    _.importSnippets = (function() {
      var cache = {}
      function getSnippetHTML(selector) {
        if (cache[selector]) {
          return cache[selector]
        }

        var el = document.querySelector(selector)
        var html = el && el.innerHTML

        return (cache[selector] = html)
      }

      function getSelector(el) {
        return (
          (el.dataset && el.dataset.snippet) ||
          el.getAttribute('[data-snippet]')
        )
      }

      return function() {
        var placeholders = document.querySelectorAll('[data-snippet]')
        Array.prototype.forEach.call(placeholders, function(placeholder) {
          var snippetSelector = getSelector(placeholder)
          var snippetHTML = getSnippetHTML(snippetSelector)

          placeholder.insertAdjacentHTML('beforebegin', snippetHTML)
          placeholder.parentNode.removeChild(placeholder)
        })
      }
    })()

    // 时间处理
    _.formatDate = function(date, fmt) {
      date = new Date(date)
      var o = {
        'M+': date.getMonth() + 1, // 月份
        'd+': date.getDate(), // 日
        'h+': date.getHours() % 12 === 0 ? 12 : date.getHours() % 12, // 小时
        'H+': date.getHours(), // 小时
        'm+': date.getMinutes(), // 分
        's+': date.getSeconds(), // 秒
        'q+': Math.floor((date.getMonth() + 3) / 3), // 季度
        S: date.getMilliseconds() // 毫秒
      }
      var week = '日一二三四五六'.split('')
      if (/(y+)/.test(fmt)) {
        fmt = fmt.replace(
          RegExp.$1,
          (date.getFullYear() + '').substr(4 - RegExp.$1.length)
        )
      }
      // /u661f/u671f -> 星期
      // /u5468 -> 周
      if (/(E+)/.test(fmt)) {
        var weekPrefix = ''
        if (RegExp.$1.length > 1) {
          weekPrefix = '周'
        }
        if (RegExp.$1.length > 2) {
          weekPrefix = '星期'
        }
        fmt = fmt.replace(RegExp.$1, weekPrefix + week[date.getDay() + ''])
      }
      for (var k in o) {
        if (new RegExp('(' + k + ')').test(fmt)) {
          fmt = fmt.replace(
            RegExp.$1,
            RegExp.$1.length === 1
              ? o[k]
              : ('00' + o[k]).substr(('' + o[k]).length)
          )
        }
      }
      return fmt
    }

    var Vue = (frame || window).Vue
    var createElement = new Vue().$createElement
    var messageBox = Vue.prototype.$msgbox

    // vue封装的prompt框
    _.promptTextArea = function(message, title, options) {
      if (typeof title === 'object') {
        options = title
        title = ''
      } else if (typeof title === 'undefined') {
        title = ''
      }

      options = options || {}
      options.inputType = 'textarea'

      return _.prompt(message, title, options)
    }

    _.TableResizer = (function(root) {
      function TableResizer(table) {
        this.table = table

        xa.utils.on(window, 'resize load click', this.resize.bind(this))
        this.resize()
      }

      TableResizer.prototype.resize = function() {
        var winHeight = window.innerHeight
        var resizer = this
        resizer.table.height = winHeight

        return new Promise(function(resolve) {
          Vue.nextTick(function() {
            var newHeight = winHeight - (root.scrollHeight - winHeight)
            resizer.table.height = newHeight < 0 ? 'auto' : newHeight
            resolve(newHeight)
          })
        })
      }

      return TableResizer
    })(document.documentElement)

    // 替换特殊符号
    _.escape = (function() {
      var htmlEscapes = {
        '&': '&amp;',
        '<': '&lt;',
        '>': '&gt;',
        '"': '&quot;',
        "'": '&#39;'
      }
      var reUnescapedHtml = /[&<>"']/g
      var reHasUnescapedHtml = RegExp(reUnescapedHtml.source)

      function replaceFn(chr) {
        return htmlEscapes[chr]
      }

      return function(s) {
        s = '' + s
        return s && reHasUnescapedHtml.test(s)
          ? s.replace(reUnescapedHtml, replaceFn)
          : s
      }
    })()

    _.createTheadTip = function(name, tip) {
      var createElement = this.$createElement
      var icon = createElement('i', {
        class: 'el-icon-question p-thead__icon'
      })
      var slot = createElement('div', {
        class: 'p-thead__tooltip',
        slot: 'content',
        domProps: {
          innerHTML: tip
        }
      })
      var tooltip = createElement(
        'el-tooltip',
        {
          attrs: {
            // effect: 'light',
            placement: 'bottom'
            // content: tip
          }
        },
        [icon, slot]
      )

      return createElement(
        'span',
        {
          class: 'p-thead'
        },
        [name, tooltip]
      )
    }

    // 笛卡尔积
    _.descartes = (function() {
      function append(current) {
        return this.concat(current)
      }

      function compile(compiled, list) {
        return compiled.reduce(function reducer(acc, values) {
          var current = list.map(append, values)
          return acc.concat(current)
        }, [])
      }

      function filter(values) {
        return Array.isArray(values) && values.length
      }

      function toArray(value) {
        return [value]
      }

      function transform(values) {
        return values.map(toArray)
      }

      function descartes(collection) {
        collection = collection || []
        collection = collection.filter(filter)

        if (!collection.length) {
          return []
        }

        collection = collection.map(transform)

        return collection.reduce(compile)
      }

      return descartes
    })()

    _.not = function(fn) {
      return function() {
        return !fn.apply(this, arguments)
      }
    }

    _.prop = function(prop, value) {
      var checker = {
        is: function(value) {
          return function(item) {
            return item[prop] === value
          }
        },
        not: function(value) {
          return function(item) {
            return item[prop] !== value
          }
        }
      }

      return arguments.length === 2 ? checker.is(value) : checker
    }

    _.download = (function(url) {
      var iframe

      function createDownloadFrame() {
        if (!iframe) {
          var body = document.body
          iframe = document.createElement('iframe')
          iframe.src = 'about:blank'
          iframe.hidden = true
          body.appendChild(iframe)
        }

        return iframe
      }

      function addNoCacheParam(url) {
        return (
          url + (url.indexOf('?') === -1 ? '?' : '&') + '_=' + Math.random()
        )
      }

      return function(url) {
        createDownloadFrame().contentWindow.location = addNoCacheParam(url)
      }
    })()

    _.once = function(fn) {
      return (function() {
        var run = false
        return function() {
          if (run) {
            return
          }

          run = true

          return fn.apply(this, arguments)
        }
      })()
    }

    // 结果与getQueryObject相同
    _.param2Obj = function(url) {
      const search = url.split('?')[1]
      if (!search) return {}
      return JSON.parse(
        '{"' +
          decodeURIComponent(search)
            .replace(/"/g, '\\"')
            .replace(/&/g, '","')
            .replace(/=/g, '":"') +
          '"}'
      )
    }

    // 将形如{page:'1',sort:'4'}JSON转换为(page=1&sort=4)param
    _.obj2Param = function(json) {
      if (!json) return ''
      return Object.keys(json)
        .map(key => {
          if (json[key] === undefined) return ''
          return encodeURIComponent(key) + '=' + encodeURIComponent(json[key])
        })
        .join('&')
    }

    // 根据url获取？后面的参数，以对象的形式返回
    _.getQueryObject = function(url) {
      url = url == null ? window.location.href : url
      const search = url.substring(url.lastIndexOf('?') + 1)
      const obj = {}
      const reg = /([^?&=]+)=([^?&=]*)/g
      //注意这里的 $1 $2 与上文的区别
      search.replace(reg, (rs, $1, $2) => {
        const name = decodeURIComponent($1)
        let val = decodeURIComponent($2)
        val = String(val)
        obj[name] = val
        return rs
      })
      return obj
    }

    _.html2Text = function(val) {
      const div = document.createElement('div')
      div.innerHTML = val
      return div.textContent || div.innerText
    }

    // 将传入的数字以3位添加逗号返回 1,234元
    _.toThousandslsFilter = function(num) {
      return (+num || 0)
        .toString()
        .replace(/^-?\d+/g, m => m.replace(/(?=(?!\b)(\d{3})+$)/g, ','))
    }

    // 获取树
    _.getTree = function(nodes) {
      let map = {},
        node,
        roots = []
      for (let i = 0; i < nodes.length; i += 1) {
        node = nodes[i]
        node.children = []
        map[node.id] = i
        if (node.pid != 0) {
          nodes[map[node.pid]].children.push(node)
        } else {
          roots.push(node)
        }
      }
      return roots
    }

    function DataTransfer(data) {
      if (!(this instanceof DataTransfer)) {
        return new DataTransfer(data, null, null)
      }
    }

    DataTransfer.treeToArray = function(data, parent, level, expandedAll) {
      let tmp = []
      Array.from(data).forEach(function(record) {
        if (record._expanded === undefined) {
          Vue.set(record, '_expanded', expandedAll)
        }
        if (parent) {
          Vue.set(record, '_parent', parent)
        }
        let _level = 0
        if (level !== undefined && level !== null) {
          _level = level + 1
        }
        Vue.set(record, '_level', _level)
        tmp.push(record)
        if (record.children && record.children.length > 0) {
          let children = DataTransfer.treeToArray(
            record.children,
            record,
            _level,
            expandedAll
          )
          tmp = tmp.concat(children)
        }
      })
      return tmp
    }

    /**
     * [treeToArray description]
     * @param  data []
     * @param  parent      null
     * @param  level       number null
     * @param  expandedAll boolean
     * @return
     */
    _.treeToArray = function(data, parent, level, expandedAll) {
      return DataTransfer.treeToArray(data, parent, level, expandedAll)
    }

    // 将两个参数合并为一个object
    _.objectMerge = function(target, source) {
      if (typeof target !== 'object') {
        target = {}
      }
      if (Array.isArray(source)) {
        return source.slice()
      }
      for (const property in source) {
        if (source.hasOwnProperty(property)) {
          const sourceProperty = source[property]
          if (typeof sourceProperty === 'object') {
            target[property] = objectMerge(target[property], sourceProperty)
            continue
          }
          target[property] = sourceProperty
        }
      }
      return target
    }

    // 合法密码
    _.validatePassword = function(rule, value, callback) {
      if (value.length < 6) {
        callback(new Error('密码不能小于6位'))
      } else {
        callback()
      }
    }

    // 合法手机号码
    _.validateMobile = function(rule, value, callback) {
      const reg = /^1[34578]\d{9}$/
      if (!reg.test(value)) {
        callback(new Error('输入的手机号码不正确'))
      } else {
        callback()
      }
    }

    // 合法uri
    _.validateURL = function(str) {
      const reg = /^(https?|ftp):\/\/([a-zA-Z0-9.-]+(:[a-zA-Z0-9.&%$-]+)*@)*((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3}|([a-zA-Z0-9-]+\.)*[a-zA-Z0-9-]+\.(com|edu|gov|int|mil|net|org|biz|arpa|info|name|pro|aero|coop|museum|[a-zA-Z]{2}))(:[0-9]+)*(\/($|[a-zA-Z0-9.,?'\\+&%$#=~_-]+))*$/
      return reg.test(textval)
    }

    // 小写字母
    _.validateLowerCase = function(str) {
      const reg = /^[a-z]+$/
      return reg.test(str)
    }

    // 大写字母
    _.validateUpperCase = function(str) {
      const reg = /^[A-Z]+$/
      return reg.test(str)
    }

    //大小写字母
    _.validateAlphabets = function(str) {
      const reg = /^[A-Za-z]+$/
      return reg.test(str)
    }

    // 数字
    _.validateNumber = function(rule, value, callback) {
      const reg = /^[0-9]*$/
      if (!reg.test(value)) {
        callback(new Error('请输入整数'))
      } else {
        callback()
      }
    }

    // 正整数
    _.validateInt = function(rule, value, callback) {
      const reg = /^\+?[1-9][0-9]*$/
      if (!reg.test(value)) {
        callback(new Error('请输入正整数'))
      } else {
        callback()
      }
    }

    // 折扣
    _.validateDiscount = function(rule, value, callback) {
      const reg = /^0\.{0,1}[0-9]{0,2}|1$/
      if (!reg.test(value)) {
        callback(new Error('请输入0-1之间的数值'))
      } else {
        callback()
      }
    }

    // money
    _.validateMoney = function(rule, value, callback) {
      const reg = /(^[1-9]([0-9]+)?(\.[0-9]{1,2})?$)|(^(0){1}$)|(^[0-9]\.[0-9]([0-9])?$)/
      if (!reg.test(value)) {
        callback(new Error('请输入金钱'))
      } else {
        callback()
      }
    }

    return _
  })()

  if (!config) {
    _.alert('请配置 project/scripts/config.js')
    return
  }

  var isDebug = config.debug

  if (isDebug) {
    Vue.config.devtools = isDebug
    Vue.config.silent = !isDebug
  }

  if (!isFrame) {
    initNProgress()
    showFrameLink()
  }

  checkFrame()

  function loadHash() {
    var hash = location.hash.slice(1)
    if (!hash) {
      return
    }

    this.frame.src = hash
  }

  function setActiveMenu() {
    var url = _.normalize(this.frame.src.split('#')[0])
    var currentMenu = flattenMenus(this.menus).find(function(menu) {
      return _.normalize(menu.link.split('#')[0]) === url
    })

    var index = currentMenu ? currentMenu.index : '0'
    this.menu.active = index
  }

  function flattenMenus(menus, prefix) {
    prefix = prefix || ''
    return menus.reduce(function(acc, current, index) {
      var fullIndex = prefix ? prefix + '-' + index : '' + index
      var children =
        current.children && current.children.length
          ? flattenMenus(current.children, fullIndex)
          : null

      return acc.concat(
        children ||
          Object.assign(current, {
            index: fullIndex
          })
      )
    }, [])
  }

  function saveHash() {
    var currentHash = frame.location.hash
    var hash =
      '#' +
      (frame.location.origin === location.origin ? '' : location.origin) +
      location.pathname +
      location.search +
      location.hash
    if (currentHash !== hash) {
      try {
        frame.history.replaceState(null, document.title, hash)
      } catch (e) {
        frame.location.replace(hash)
      }
    }
    try {
      frame.window.document.title = document.title
    } catch (e) {}
  }

  function checkFrame() {
    // normal in frame
    if (!isFrame && frame) {
      _.on(window, 'hashchange', saveHash)
      saveHash()
      return
    }

    if (location.protocol === 'file:' || !isFrameMode || !config.frameCheck) {
      return
    }

    if (isFrame && parent !== window && isFrameWindow(frame.parent)) {
      parent.location.hash = location.hash.replace(/^#+/, '#')
      parent.location.reload()
      return
    }

    if (
      isFrame ||
      !document.documentElement.classList.contains('xa-page--inner')
    ) {
      return
    }

    if (!frame) {
      location.href =
        config.frame + '#' + location.pathname + location.search + location.hash
    }
  }

  function initNProgress() {
    var NProgress = window.NProgress
    NProgress.start()

    if (!xa.holdProgress) {
      _.on(window, 'load', function() {
        NProgress.done()
      })
    }
  }

  function showFrameLink() {
    if (!isDebug || config.frameCheck || isFrame) {
      return
    }

    var href = location.href
    var text = '移除框架'

    if (!frame) {
      href = config.frame + '#' + href
      text = '在框架中查看'
    }

    var link = Object.assign(document.createElement('a'), {
      className: 'view-in-frame-link',
      href: href,
      target: '_top',
      textContent: text,
      title: decodeURIComponent(location.pathname)
    })

    _.on(document, 'DOMContentLoaded', function() {
      body.appendChild(link)
    })
  }

  // __inline('_xa-element-hack.js')

  // explore
  xa.loadHash = loadHash
  xa.setActiveMenu = setActiveMenu
  xa.frame = frame

  // alias
  xa.nextTick = _.nextTick
})()
