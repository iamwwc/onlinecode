<template>
  <div
    :class="{open:isOpen}"
    class="sidebar">
    <div class="container" >
      <div class="lang-select">
        <span>Language: </span>
        <el-popover
          placement="top-start"
          width="200"
          trigger="hover"
          content="This is global language mode. It affects all files created later.
                  For special mode in single file, You can change in Controller bar">
          <i 
            slot="reference" 
            class="far fa-question-circle"/>
        </el-popover>
        
        <el-select
          v-model="langValue"
          class="el-select">
          <el-option
            v-for="item in langitems"
            :key="item.key"
            :value="item.value"/>
        </el-select>
      </div>

      <div class="theme-select">
        <span>Theme: </span>
        <el-select
          v-model="themeValue"
          class="el-select">
          <el-option
            v-for="item in themeitems"
            :key="item.key"
            :value="item.value"/>
        </el-select>
      </div>
      <div class="font-change">
        <span>FontSize</span>
        <el-input-number
          v-model="font_Size"
          size="small"/>
      </div>

    </div>
  </div>
</template>

<script>
import * as $ from 'jquery'
import {mapState, mapMutations, mapGetters} from 'vuex'
export default{
  data () {
    return {
      isOpen: false
    }
  },

  computed: {
    ...mapState('global', [
      'ide'
    ]),
    ...mapState({
      theme: state => state.global.workspace.ideConfig.theme,
      fontSize: state => state.global.workspace.ideConfig.fontSize,
      lang: state => state.global.workspace.ideConfig.lang,
      langitems: state => state.constvar.langitems,
      themeitems: state => state.constvar.themeitems
    }),
    ...mapGetters('constvar', [
      'valueOfLang',
      'valueOfTheme',
      'keyOfTheme',
      'keyOfLang'
    ]),
    themeValue: {
      get () {
        return this.valueOfTheme(this.theme)
      },
      set (newV) {
        this.setIdeConfig({
          theme: this.keyOfTheme(newV)
        })
      }
    },
    langValue: {
      get () {
        return this.valueOfLang(this.lang)
      },
      set (newV) {
        this.setIdeConfig({
          lang: this.keyOfLang(newV)
        })
      }
    },
    font_Size: {
      get () {
        return this.fontSize
      },
      set (newV) {
        this.setIdeConfig({
          fontSize: newV
        })
      }
    }
  },
  mounted () {
    this.$nextTick(() => {
      this.defaultFontSize = this.fontSize
      this.langvalue = this.valueOfLang(this.lang)
    })
  },
  methods: {
    toggle () {
      this.isOpen = true
      $('body').on('click', e => {
        if ($('.sidebar')[0] !== e.target && !$.contains($('.sidebar')[0], e.target)) {
          this.isOpen = false
          $('body').off('click')
          e.stopPropagation()
        }
      })
    },
    onFontSizeChange (value) {
      this.setIdeConfig({
        fontSize: value
      })
    },
    ...mapMutations('global', [
      'setIdeConfig'
    ])
  }
}
</script>

<style lang="scss">
.sidebar{
    position: fixed;
    right:0px;
    top:0px;
    bottom:0px;
    width:280px;
    transform:translate(280px,0px);
    transition: transform 0.4s cubic-bezier(0.4,0,0,1);
    background-color:#f9f9f9;
    z-index: 100;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.2)
}
.open{
    transform: translate(0px,0px);
    background-color: #f9f9f9;
    transition: transform 0.4s cubic-bezier(0.4,0,0,1);
    z-index: 100;
}

.container{
    margin: 2px 2px;
    padding: 10px 10px;
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: column;
}

.container > div{
    margin: 10px 0px 0px 0px;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    margin-right: 20px;
    border-bottom: 1px solid #ddd;
    padding: 2px 0px 0px 0px;
}

.lang-select{
    height:40px;
    float: left;
}

.el-select{
    width: 100px;
}
.el-scrollbar{
    overflow: initial !important;
}

.sidebar span{
    font-size:16px;
}
.fa-question-circle{
  width: 20px;
  height: 20px;
}
.el-select-dropdown__item{
  &:hover{
    color: aquamarine;
  }
}
</style>
