// 受现代 JavaScript 的限制 (而且 Object.observe 也已经被废弃)，
// Vue 不能检测到对象属性的添加或删除。
// files一直会动态变化，Vue不能追踪其响应过程,不能在其上直接删除添加object，
// 可以新创建一个然后赋值给files
'use strict'
import {
  DEFAULT_ACE_THEME, 
  DEFAULT_FONT_SIZE,
  DEFAULT_LANGGUAGE,
  DEFAULT_ID,
  DEFAULT_ROOT_ID
} from '../../constvar.js'

export default{
  namespaced: true,
  state: {
    ide: undefined,
    terminal: undefined,
    loading: false, // Terminal Loading 
    workspace: {
      filetree: {
        maxId: 3,
        currentNodeId: DEFAULT_ID,
        currentRightNodeId: DEFAULT_ROOT_ID,
        files: [
          {
            detail: {
              type: 'dir',
              lang: 'text'
            },
            id: 1,
            label: 'source-codes'
          }
        ],
        currentNode: {
          id: 0,
          label: '',
          detail: {
            src: '',
            lang: 'java', // 随便一个有的语言就可以，最后FileTree.vue初始化完成之后会自己更改
            type: ''
          }
        }
      },
      ideConfig: {
        lang: DEFAULT_LANGGUAGE,
        theme: DEFAULT_ACE_THEME,
        fontSize: DEFAULT_FONT_SIZE
      }
    }
  },
  mutations: {
    setIDE (state, myide) {
      state.ide = myide
    },
    setTerminal (state, term) {
      state.terminal = term
    },
    setWorkSpace (state, workspace) {
      state.workspace = workspace
    },
    incrementMaxId (state) {
      state.workspace.filetree.maxId++
    },
    setCurrentNode (state, node) {
      state.workspace.filetree.currentNode = node
    },
    setCurrentNodeLang (state, lang) {
      state.workspace.filetree.currentNode.detail.lang = lang
    },
    getWorkSpace (state) {
      return state.workspace
    },
    setFiles (state, newfiles) {
      state.workspace.filetree.files = newfiles
    },
    setCurrentNodeId (state, id) {
      state.workspace.filetree.currentNodeId = id
    },
    setCurrentRightNodeId (state, id) {
      state.workspace.filetree.currentRightNodeId = id
    },
    setFileSrc (state, o) {
      let [d, v] = o
      d.detail.src = v
    },
    renameNode (state, o) {
      let [node, newName] = o
      node.data.label = newName
    },

    // 局部更新json,检测config有的项来更新，其余的不修改
    setIdeConfig (state, config) {
      let c = state.workspace.ideConfig
      for (var i in config) {
        var v = config[i]
        if (typeof (v) === 'undefined') {
          continue
        }
        c[i] = v
      }
    },
    setLoading (state, l) {
      state.loading = l
    }
  }
}
