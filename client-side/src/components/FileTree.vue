<template>
  <div
    class="file-tree">
    <div class="controller-header">
      <div
        class="new-file"
        title="AddFile"
        @click="addFile">
        <i class="fas fa-file-medical"/>
      </div>
      <div
        class="add-dir"
        title="AddDirectory"
        @click="addDir">
        <i class="fas fa-folder-plus"/>
      </div>

    </div>
    <el-tree
      ref="maintree"
      :data="Files"
      :allow-drop="allowDrop"
      :allow-drag="allowDrag"
      :render-content="renderContent"
      highlight-current
      node-key="id"
      default-expand-all
      @node-drag-start="handleDragStart"
      @node-drag-enter="handleDragEnter"
      @node-drag-end="handleDragEnd"
      @node-drag-leave="handleDragLeave"
      @node-click="handleNodeClick"
      @current-change="handleCurrentChange"
      @node-contextmenu="handleContextMenu"
    />
    <keep-alive>
      <component
        ref="context_menu"
        :is="currentContextMenu"
        :items="currentProperties"
        @delete="deleteNode"
        @rename="rename"/>
    </keep-alive>

  </div>
</template>

<script>
import ContextMenu from './ContextMenu.vue'
import * as Ace from 'ace-builds'
import {mapMutations, mapState} from 'vuex'
import * as CONST from '../constvar.js'
export default {
  components: {
    'fileContextMenu': ContextMenu,
    'folderContextMenu': ContextMenu
  },
  data () {
    return {
      currentContextMenu: 'fileContextMenu',
      Document: undefined,
      EditSession: undefined,
      sessionFilesMap: {}
    }
  },
  computed: {
    ...mapState({
      files: state => state.global.workspace.filetree.files,
      workspace: state => state.global.workspace,
      maxId: state => state.global.workspace.filetree.maxId,
      currentNodeId: state => state.global.workspace.filetree.currentNodeId,
      currentRightNodeId: state => state.global.workspace.filetree.currentRightNodeId,
      ideConfig: state => state.global.workspace.ideConfig
    }),
    ...mapState('global', ['ide']),

    Files: {
      get () {
        return [...this.files]
      },
      set (newV) {
        this.setFiles(newV)
      }
    },

    currentProperties () {
      if (this.currentContextMenu === 'fileContextMenu') {
        return [
          {
            name: 'DeleteFile',
            action: 'delete'
          },
          {
            name: 'Rename',
            action: 'rename'
          }
        ]
      } else if (this.currentContextMenu === 'folderContextMenu') {
        return [
          {
            name: 'NewFolder',
            action: 'addFolder'
          },
          {
            name: 'NewFile',
            action: 'addFile'
          },
          {
            name: 'DeleteFolder',
            action: 'delete'
          },
          {
            name: 'Rename',
            action: 'rename'
          }
        ]
      }
    }
  },
  created () {
    this.Document = Ace.require('ace/document').Document
    this.EditSession = Ace.require('ace/edit_session').EditSession
  },
  methods: {
    fileTreeForRunCode () {
      let fileTree = {}
      let rootDir = this.$refs.maintree.getNode(CONST.DEFAULT_ROOT_ID)
      let start = rootDir.data.children
      
      start.forEach(value => {
        this.getSingleFileTree(fileTree, '', value, undefined)
      })
      return fileTree
    },
    getSingleFileTree (fileTree, prefix, value, callback) {
      if (value.detail.type === 'dir') {
        if (value.children === undefined || value.children === null) {
          return
        }
        value.children.forEach(v => {
          this.getSingleFileTree(fileTree, prefix + value.label + '/', v, callback)
        })
      } else {
        if (typeof (callback) !== 'undefined') {
          callback(value)
          return
        }
        let fileName = prefix + value.label
        let src = btoa(this.sessionFilesMap[value.id].getValue())
        fileTree[fileName] = {
          src: src,
          type: 'file'
        }
      }
    },

    createNewEnv () {
      let session = this.createNewEditSession('')

      let d = {detail: {type: 'file', src: '', lang: this.ideConfig.lang}, id: this.maxId, label: 'main.go'}
      let parent = this.getNode(CONST.DEFAULT_ROOT_ID)
      this.fileTreeAppendNode(d, session, parent)
    },

    fromInjectedEnv () {
      let rebuild = data => {
        let session = this.createNewEditSession(data.detail.src)
        this.sessionFilesMap[data.id] = session
      }

      this.files.forEach(v => {
        this.getSingleFileTree(undefined, '', v, rebuild)
      })
      
      let currentNodeId = this.workspace.filetree.currentNodeId
      // this.$refs.maintree.setCurrentNode(currentNodeId)
      this.$nextTick(() => {
        let node = this.$refs.maintree.getNode(currentNodeId)
        this.handleCurrentChange(node.data, node)
      })
    },

    addSessionOnChange (data, session) {
      session.on('change', (d) => {
        this.setFileSrc([data, session.getValue()])
      })
    },
    renderContent (h, {node, data, store}) {
      if (node.data.detail.type === 'dir') {
        return (
          <span class="custom-tree-node el-tree-el-tree-node__content">
            <i class="fa fa-folder-open"></i>
            <span class="el-tree-node__label">{node.label}</span>
          </span>
        )
      }
      return (
        <span class="custom-tree-node el-tree-el-tree-node__label">
          <span class="el-tree-node__label">{node.label}</span>
        </span>
      )
    },

    handleDragStart (node, env) {

    },

    handleDragEnd (draggedNode, finalEnterNode, pos, event) {

    },

    handleDragOver (draggedNode, overNode, event) {

    },

    handleDragLeave (draggedNode, leaveNode, event) {

    },

    handleDragEnter () {

    },

    handleContextMenu (event, o, node, nodecomponent) {
      if (o.detail.type === 'dir') {
        this.currentContextMenu = 'folderContextMenu'
      } else {
        this.currentContextMenu = 'fileContextMenu'
      }
      this.$nextTick(() => {
        this.$refs.context_menu.openMenu(event)
      })
      this.setCurrentRightNodeId(o.id)
    },

    handleNodeClick (o, node, nodecomponent) {
      this.setCurrentNodeId(node.data.id)
      let body = document.getElementsByTagName('body')[0]
      body.click()
    },

    allowDrop (draggingNode, dropNode, type) {
      if (dropNode.id === CONST.DEFAULT_ROOT_ID || dropNode.data.detail.type === 'file') {
        return false
      }
      if (dropNode.id === CONST.DEFAULT_ROOT_ID) {
        return false
      }
      return true
    },

    allowDrag (node) {
      if (node.data.id === CONST.DEFAULT_ROOT_ID) {
        return false
      }
      return true
    },

    rename () {
      this.$prompt('Input New Name', 'Rename', {
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel'
      }).then(({value}) => {
        let currentRightNode = this.getNode(this.currentRightNodeId)
        this.renameNode([currentRightNode, value])    
      })
    },

    openAddFilePrompt () {
      return this.$prompt('Input New File Name', 'Add New File', {
        confirmButtonText: 'Ok',
        cancelButtonText: 'Cancel',
        inputPattern: /^[0-9a-zA-Z._]+$/,
        inputErrorMessage: 'Only allow Aa-Zz, 1-9, _.'
      })
    },

    fileTreeRemoveNode (node) {
      let id = node.data.id
      if (node.data.detail.type === 'file') {
        delete this.sessionFilesMap[id]
      }
      this.$refs.maintree.remove(node)
    },

    fileTreeAppendNode (data, session, parent) {
      this.$refs.maintree.append(data, parent)
      if (data.detail.type !== 'file') {
        return
      }
      this.sessionFilesMap[data.id] = session
      this.configCurrentSession(data, session)
      this.addSessionOnChange(data, session)
    },

    addFile (e) {
      let tree = this.$refs.maintree
      let currentNode = this.getNode(this.currentNodeId)
      // 不再使用CurrentNode,统一使用CurrentNodeId
      // 可以使用 this.currentNode == undefined || null
      if (currentNode === undefined || currentNode === null) {
        currentNode = tree.getNode(1)
      }
      this.openAddFilePrompt().then(({value}) => {
        this.incrementMaxId()
        let session = this.createNewEditSession('')
        let add = {detail: {type: 'file', src: '', lang: this.ideConfig.lang}, id: this.maxId, label: value}
        let parent

        if (currentNode.data.detail.type === 'file') {
          parent = currentNode.parent
        } else {
          parent = currentNode
        }
        this.fileTreeAppendNode(add, session, parent)
      })
    },
    addFileFinal () {

    },

    handleCurrentChange (data, node) {
      if (data.detail.type === 'file') {
        this.$refs.maintree.setCurrentNode(data.id)
        // 如果没有对应的session那么新建(come from share, session not attached yet)
        let session = this.sessionFilesMap[data.id]
        if (session === undefined) {
          session = this.createNewEditSession('')
          session.setValue(data.detail.src)
          this.sessionFilesMap[data.id] = session
          this.addSessionOnChange(data, session)
        }
        this.configCurrentSession(data, session)
      }
      this.setCurrentNodeId(data.id)
    },

    configCurrentSession (data, session) {
      this.ide.getEditor().setSession(session)
      this.ide.getEditor().focus()
      session.setMode('ace/mode/' + data.detail.lang)
      this.setCurrentNode(data)
    },

    createNewEditSession (c) {
      let document = new this.Document(c)
      let langMode = 'ace/mode/' + this.workspace.ideConfig.lang
      return new this.EditSession(document, langMode)
    },

    deleteNode () {
      if (this.currentRightNodeId === CONST.DEFAULT_ROOT_ID) {
        this.$message({
          message: 'Not allow to remove root folder',
          type: 'warning',
          duration: 2000
        })
        return
      }
      this.$confirm('Are you sure to remove? ', 'Attention', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel'
      }).then(() => {
        let currentNode = this.getNode(this.currentRightNodeId)
        this.setCurrentNodeId(currentNode.parent.id) 
        this.fileTreeRemoveNode(currentNode)        
      })
    },

    addDir () {
      this.incrementMaxId()
      let currentNode = this.getNode(this.currentNodeId)
      let tree = this.$refs.maintree
      if (currentNode === undefined || currentNode === null) {
        currentNode = tree.getNode(1)
      }
      this.$prompt('Input New Directory Name', 'Add New Directory', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        inputPattern: /^[0-9a-zA-Z._]+$/,
        inputErrorMessage: 'Only allow Aa-Zz, 1-9, _.'
      }).then(({value}) => {
        let parent
        let add = {detail: {type: 'dir'}, id: this.maxId, label: value}
        if (currentNode.data.detail.type === 'file') {
          parent = currentNode.parent
        } else {
          parent = currentNode
        }
        this.fileTreeAppendNode(add, undefined, parent)
      })
    },
    getNode (id) {
      return this.$refs.maintree.getNode(id)
    },

    ...mapMutations('global', [
      'incrementMaxId',
      'setFiles',
      'setCurrentNodeId',
      'setCurrentRightNodeId',
      'setWorkSpace',
      'setFileSrc',
      'setCurrentNode',
      'renameNode'
    ])
  }
}
</script>

<style lang="scss">
  $back-color: #f9f7f7;
  $left-width: 200px;
  $max-width:100%;
  $max-height:100%;

  .file-tree{
    background-color: $back-color;
    display: flex;
    flex-direction: column;
    height: $max-height;
    user-select: none;
    overflow: auto;
    el-tree{
      flex-grow: 1;
    }
    .el-tree-node__label{
      padding-left: 3px;
    }
    .fa-folder-open{
      color: #4596de;
      &:hover{
        color: #4596de;
      }
    }

    .fa-file-medical{
      color: #6bbf75;
      &:hover{
        color: #6bbf75
      }
    }

    .fa-folder-plus{
      color: #bd8352;
      &:hover{
        color: #bd8352;
      }
    }

    .controller-header{
      flex-basis: 20px;
      display: flex;
      flex-direction: row-reverse;
      div:hover{
        cursor: pointer;
      }

      div{
        margin-left: 5px;
      }

      .new-file{
        order:4;

      }

      .remove-file{
        order: 3;
      }
      .remove-dir{
        order:1;
      }

      .add-dir{
        order: 2
      }
    }
  }

</style>
