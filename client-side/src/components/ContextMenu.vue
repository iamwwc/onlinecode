<template>

  <ul
    v-show="showMenu"
    ref="context_ul"
    :style="css"
    class="context-menu">
    <li
      v-for="item in items"
      :key="item.action"
      @click="($event) => {showMenu = false;$emit(item.action)}"
      @contextmenu.prevent>{{ item.name }}</li>
  </ul>
</template>

<script>
export default {
  props: {
    items: {
      type: Array,
      default: undefined
    }
  },
  data () {
    return {
      showMenu: false,
      rightcss: {
        left: '0px',
        top: '0px'
      }
    }
  },
  computed: {
    css: function () {
      return {
        top: this.rightcss.top + 'px',
        left: this.rightcss.left + 'px'
      }
    }
  },
  mounted () {
    let contextMenu = document.getElementsByClassName('context-menu')[0]
    // 这里事件由body向下传播，直接停止，不允许进行事件冒泡
    document.body.addEventListener('click', function (e) {
      if (e.target !== contextMenu && !contextMenu.contains(e.target) && this.showMenu === true) {
        this.showMenu = false
        e.stopPropagation()
      }
    }.bind(this), true)
  },
  methods: {
    closeMenu () {
      this.showMenu = false
    },
    openMenu (e) {
      this.showMenu = true
      this.$refs.context_ul.focus()
      this.rightcss.top = e.clientY
      this.rightcss.left = e.clientX
    },
    itemClick (e) {
      let action = e.currentTarget.attributes['action'].value
      this.$emit('item-click', {action: action})
    }
  }
}
</script>

<style lang="scss">
    ul {
        background: #F9f7f7;
        border: 1px solid #BDBDBD;
        box-shadow: 0 2px 2px 0 rgba(0,0,0,.14),0 3px 1px -2px rgba(0,0,0,.2),0 1px 5px 0 rgba(0,0,0,.12);
        display: block;
        list-style: none;
        margin: 0;
        padding: 0;
        position: absolute;
        width: 150px;
        z-index: 999999;
        li {
            border-bottom: 1px solid #E0E0E0;
            margin: 0;
            padding: 5px 5px;
            font-size: 12px;
            font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
            &:last-child {
                border-bottom: none;
            }
            &:hover {
                background-color: #1e88e5;
                color: #FAFAFA;
            }
        }
    }

</style>
