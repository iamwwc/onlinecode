<template>
  <div id="app">
    <el-container>
      <el-aside>
        <FileTree
          ref="filetree"/>
      </el-aside>
      <el-container>
        <el-header
          class="header"
          height="40px">
          <Controller
            @run="ide.run($refs.filetree.fileTreeForRunCode())"
            @share="ide.share()"/>
          <div class="editor-area">
            <div
              id="ace-code-area"
              class="ace-code-area"/>
          </div>
        </el-header>
        <Resizer/>
        <el-main class="main" >
          <Output/>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script >
import Output from './components/Output.vue'
import Resizer from './components/Resizer.vue'
import Controller from './components/Controller.vue'
import FileTree from './components/FileTree.vue'
import IDE from './editor'
import 'xterm/lib/xterm.css'
import Terminal from './terminal.js'
import {mapMutations} from 'vuex'

export default{
  components: {
    Output, Resizer, Controller, FileTree
  },
  data () {
    return {
    }
  },
  created () {
    if (typeof (window.EditorConfig) !== 'undefined') {
      this.setWorkSpace(JSON.parse(window.EditorConfig).data)
      this.$nextTick(() => {
        this.$refs.filetree.fromInjectedEnv()
      })
      return
    }
    this.$nextTick(() => {
      this.$refs.filetree.createNewEnv()
    })
  },
  mounted () {
    this.terminal = new Terminal()
    
    Object.freeze(this.terminal)
    this.terminal.init('output-terminal')
    
    this.setTerminal(this.terminal)
    this.ide = new IDE('ace-code-area', this)
    Object.freeze(this.ide)
    this.setIDE(this.ide)
  },
  methods: {
    ...mapMutations('global', ['setTerminal', 'setIDE', 'setWorkSpace'])
  }
}
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  width: 100%;
  height: 100%;
}

html,
body {
  height: 100%;
  width: 100%;
}

body {
  padding: 0px;
  margin: 0px;
}

.main {
  padding: 0px !important;
  margin: 0px !important;
  display: flex !important;
  flex-direction: column !important;
}

.editor-area {
  flex-grow: 1;
  display: flex;
  flex-direction: row;
  align-items: stretch;
}

.ace-code-area {
  flex-grow: 1;
}

.header {
  padding: 0px !important;
  margin: 0px;
  display: flex;
  flex-direction: column;
  flex-basis: 75%;
}

.el-button {
  padding: 0px !important;
  width: 60px;
}

.left-sidebar {
  width: 200px;
  height: 100%;
}

.el-container {
  height: 100%;
}

#output-area{
  flex-grow: 1;
}
</style>
