import './app.css'
import './lib/locale.svelte'
import { mount } from 'svelte'
import App from './App.svelte'
import { applyStartupSettings } from './lib/settings'

// Apply the command-line theme and language before the first render. If they
// can't be read, start with the defaults rather than not at all.
applyStartupSettings()
  .catch((error) => console.error('Could not apply startup settings:', error))
  .finally(() => mount(App, { target: document.getElementById('app')! }))
