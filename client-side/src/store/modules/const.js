'use strict'
export default{
  namespaced: true,
  state: {
    langitems: [
      {
        key: 'java',
        value: 'Java'
      },
      {
        key: 'golang',
        value: 'Go'
      },
      {
        key: 'javascript',
        value: 'Javascript'
      },
      {
        key: 'python',
        value: 'Python'
      }
    ],
    themeitems: [
      {
        key: 'monokai',
        value: 'Dark'
      },
      {
        key: 'chrome',
        value: 'Light'
      }
    ]
  },
  getters: {
    // 得到value的key
    // 为了valueOfLang和keyOfLang都能使用objectOfLang， 只能使用 || 连接
    // fontSize没有使用到array， 不需要比较value和key，直接在Sidebar修改
    objectOfLang: state => key => state.langitems.find(item => item.value === key || item.key === key),
    valueOfLang: (state, getters) => key => getters.objectOfLang(key).value,
    keyOfLang: (state, getters) => key => getters.objectOfLang(key).key,

    // 得到key的value
    objectOfTheme: state => key => state.themeitems.find(item => item.key === key || item.value === key),
    valueOfTheme: (state, getters) => key => getters.objectOfTheme(key).value,
    keyOfTheme: (state, getters) => key => getters.objectOfTheme(key).key 
  }
}
