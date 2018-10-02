'use strict'
import * as Ace from 'ace-builds'
import 'ace-builds/src-noconflict/ext-language_tools'
import store from './store'
import Vue from 'vue'
export default class IDE {
  constructor (id, vm) {
    this.$store = store
    Ace.config.set('basePath', '/static/js/ace')

    // 直接将Vue实例注册上去
    this.$vm = vm

    this.global = this.$store.state.global
    this.lang = this.global.workspace.ideConfig.lang
    this.fontSize = this.global.workspace.ideConfig.fontSize
    this.theme = this.global.workspace.ideConfig.theme

    this.aceEditor = Ace.edit(id)
    this.aceEditor.setTheme('ace/theme/' + this.theme)
    this.aceEditor.setShowPrintMargin(false)
    this.$http = Vue.prototype.$http
    this.$terminal = store.state.global.terminal

    this.compilePath = '/api/compile'
    this.sharePath = '/api/share'

    this.aceEditor.setFontSize(14)
    this.aceEditor.setOptions({
      enableBasicAutocompletion: true,
      enableSnippets: true,
      enableLiveAutocompletion: true
    })
    this.aceEditor.commands.bindKey('enter', null)

    this.addWatch()
  }

  resize () {
    this.aceEditor.resize()
  }
  run (value) {
    let request = this.packRequest(value)
    this.$http.post(this.compilePath, request, {
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(res => {
      this.runCompleted()
      let response = res.data
      this.$terminal.renderResult(response)
    }).catch(() => {
      let e = {
        Error: 'Server Error',
        Events: null
      }
      this.runCompleted()
      this.$terminal.renderResult(e)
    })
  }

  share () {
    let workspace = {...this.$store.state.global.workspace}
    let request = {
      version: '1',
      data: workspace
    }
    this.$http.post(this.sharePath, request, {
      headers: {
        'Content-type': 'application/json'
      }
    }).then(res => {
      let data = res.data
      if (data['error'] !== '') {
        return Promise.reject(data['error'])
      }
      let id = data['id']
      if (history.state) {
        history.replaceState('share', '', id.toString())
      } else {
        history.replaceState('share', '', 'share/' + id)
      }
      this.$vm.$confirm(location.href, 'Share',
        {
          confirmButtonText: 'OK',
          showCancelButton: false,
          type: 'success'
        })
    }).catch(err => {
      let msg = typeof (err) === 'string' ? err : 'Error when share'
      this.$vm.$notify.error({
        title: 'error',
        message: msg
      })
    })
  }

  runCompleted () {
    this.$store.commit('global/setLoading', false)
  }

  getEditor () {
    return this.aceEditor
  }

  getDocument () {
    return this.aceSession.getDocument()
  }

  disableEditorSelect () {
    this.aceEditor.setOption('dragEnabled', false)
  }
  enableEditorSelect () {
    this.aceEditor.setOption('dragEnabled', true)
  }

  setMode (mode) {
    this.aceSession.setMode(mode)
  }

  setTheme (theme) {
    this.aceEditor.setTheme(theme)
  }

  setFontSize (value) {
    this.aceEditor.setFontSize(value)
  }

  test () {
    console.log('THIS IS A TEST FUNCTION TO TEST REFERENCE')
  }

  packRequest (value) {
    let request = {
      Version: '1',
      Files: value,
      Lang: this.lang
    }
    return request
  }

  addWatch () {
    let ideConfig = this.$store.state.global.workspace.ideConfig
    this.$store.watch(
      function () {
        return ideConfig.theme
      }
      , (newV) => {
        this.setTheme('ace/theme/' + newV)
      }, {deep: false}
    )

    this.$store.watch(
      function () {
        return ideConfig.fontSize
      },
      (newV) => {
        this.setFontSize(newV)
      }
    )
  }
}

/*

    response json format
    from golang struct

    type response strcut{
        Error string
        Events []Event
    }

    type Event struct{
        Message string
        Kind string
    }

    */
