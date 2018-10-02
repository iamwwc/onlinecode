<template>
  <div class="output-area">
    <i
      class="up el-icon-arrow-up"/>
    <div class="output-controller-bar">
      <span>
        Input
      </span>
    </div>
    <div
      v-loading="loading"
      class="terminal-container">
      <div
        id="output-terminal"
        class="output-terminal"/>
    </div>

  </div>
</template>

<script>
import {mapState} from 'vuex'
export default {
  data () {
    return {
    }
  },
  computed: {
    ...mapState({
      loading: state => state.global.workspace.loading
    })
  },
  mounted  () {
  }
}
</script>

<style>
.output-controller-bar {
  background-color: #fff;
  flex-basis: 13px;
  padding: 2px 4px 2px 2px;
  position: relative;
}
.output-area {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  position: relative;
}

.up{
  position: absolute;
  height: 20px;
  width: 20px;
  z-index: 1000;
}

.output-controller-bar > span {
  margin: 0 auto;
  font-size: 12px;
  display: block;
  color: #555;
  text-decoration: none;
  user-select: none;
}

.terminal-container {
  flex-grow: 1;
  flex-shrink: 1;
  display: flex;
}

.output-terminal {
  flex-grow: 1;
}
</style>
