<template>
  <div class="controller">
    <el-button
      type="success"
      size="mini"
      icon="el-icon-caret-right"
      @click="setLoading(true);$emit('run')">Run</el-button>
    <el-button
      id="share"
      icon="el-icon-share"
      size="mini"
      @click="$emit('share')">Share</el-button>
    <div class="currentLang">
      <span>CurrentLanguage: </span>
      <el-select 
        v-model="currentNodeLang" 
        size="mini">
        <el-option
          v-for="lang in langitems"
          :key="lang.key"
          :value="lang.value"
        />

      </el-select>
    </div>

    <el-button
      id="menu"
      icon="el-icon-menu"
      size="mini"
      @click.stop="toggleSidebar"/>
    <Sidebar ref="sidebar"/>
  </div>
</template>

<script>
import Sidebar from './Sidebar.vue'
import {mapState, mapMutations, mapGetters} from 'vuex'

export default{
  components: {
    Sidebar
  },
  computed: {
    ...mapState({
      ide: state => state.global.ide,
      langitems: state => state.constvar.langitems,
      currentNode: state => state.global.workspace.filetree.currentNode
    }),
    ...mapGetters('constvar', [
      'valueOfLang',
      'keyOfLang'
    ]),

    currentNodeLang: {
      get () {
        return this.valueOfLang(this.currentNode.detail.lang)
      },
      set (newV) {
        this.setCurrentNodeLang(this.keyOfLang(newV))
      }
    }
  },
  methods: {
    toggleSidebar (e) {
      this.$refs.sidebar.toggle()
    },
    ideRunCode () {
      this.ide.run()
      this.setLoading(true)
    },
    ideShareCode () {

    },
    ...mapMutations('global', [
      'setLoading',
      'setCurrentNodeLang'
    ])
  }
}

</script>

<style scoped>
.controller {
  display: flex;
  flex-direction: row;
  padding: 0px;
  margin: 0px;
  height: 30px;
  flex-basis: 30px;

  background-color: #ddd;
}

#menu{
    margin-left: auto;
    background-color: #409eff;
    color: white;
}

#share{
    background-color: rgb(253, 188, 7);
    color: #fff;
}

.currentLang{
  border: 1px solid blue;
  display: flex;
  flex-direction: row;
  align-items: center;
}

.button{
    width:60px;
}
.el-select{
  width: 120px;
}

.controller > * {
  margin-left: 10px;
}

.currentLang > span{
  font-size: 12px;
  margin-right: 4px;
}
</style>
