import Vue from 'vue'
import { TimePicker } from 'element-ui'
import locale from 'element-ui/lib/locale'
import ja from 'element-ui/lib/locale/lang/ja'
import 'element-ui/lib/theme-chalk/index.css';

locale.use(ja)
Vue.use(TimePicker)
