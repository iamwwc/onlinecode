import * as Term from 'xterm'
import * as fit from 'xterm/lib/addons/fit/fit'

export default class Terminal {
  constructor () {
    Term.Terminal.applyAddon(fit)
    this.terminal = new Term.Terminal()
    window.addEventListener('resize', () => {
      this.fit()
    })
  }

  init (id) {
    this.terminal.open(document.getElementById(id))
    this.terminal.fit()
  }
  fit () {
    this.terminal.fit()
  }

  renderResult (response) {
    this.terminal.reset()
    if (response['Error'] !== '') {
      this.terminal.write(response['Error'])
      return
    }
    let i = 0
    for (i = 0; i < response['Events'].length; ++i) {
      if (response['Events'][i]['Message'] !== '') {
        this.terminal.writeln(response['Events'][i]['Message'])
      }
    }
  }
}
