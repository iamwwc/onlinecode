<template>
  <div
    class="split-pane-resizer"
    @mousedown="onMouseDown"/>
</template>

<script>
import * as $ from 'jquery'
import {mapState} from 'vuex'
import eventBus from '../eventbus.js'

export default{
  data () {
    return {
    }
  },
  computed: {
    ...mapState({
      terminal: state => state.global.terminal,
      ide: state => state.global.ide
    })
  },
  methods: {
    onMouseDown () {
      let resizer = $('.split-pane-resizer')
      let prev = resizer.prev()
      let body = $('body')
      body.on('mousemove', (e) => {
        eventBus.$emit('resize', $(window).height() - e.pageY)
        prev.css({
          height: e.pageY,
          'flex-basis': e.pageY
        })
        this.ide.resize()
        this.ide.disableEditorSelect()
        this.terminal.fit()
        this.ide.aceEditor.setReadOnly(true)
      }).on('mouseup', () => {
        $('body').off('mousemove')
        this.ide.aceEditor.setReadOnly(false)
      })
    }
  }
}

</script>

<style scoped>
.split-pane-resizer {
  height: 5px;
  width: 100%;
  background-color: #837f7f;
  cursor: ns-resize;
  z-index: 10;
}
</style>
