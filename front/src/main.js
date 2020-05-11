import Vue from 'vue'
import App from './App.vue'
import '@/styles/normalize.css'
import '@/styles/index.scss'

import { library } from '@fortawesome/fontawesome-svg-core'
import { faCog, faTimesCircle, faTrashAlt, faFileUpload, faLongArrowAltRight, faLongArrowAltLeft, faArrowsAltH, faPlusCircle } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faCog, faTimesCircle, faTrashAlt, faFileUpload, faLongArrowAltRight, faLongArrowAltLeft, faArrowsAltH, faPlusCircle)

Vue.config.productionTip = false
Vue.component("fa-icon", FontAwesomeIcon)

new Vue({
  render: h => h(App),
}).$mount('#app')
