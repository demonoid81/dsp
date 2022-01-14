import Vue from 'vue'
import VueI18n from 'vue-i18n'
import enUsLocale from 'view-design/src/locale/lang/en-US'
import ruRULocale from 'view-design/src/locale/lang/en-US'
import customEnUS from './lang/en-US'
import customRuRU from './lang/ru-RU'
import store from "store/index";

Vue.use(VueI18n)

// Автоматически устанавливать язык в соответствии с языком системы браузера
const navLang = navigator.language
const localLang = (navLang === 'ru-RU' || navLang === 'en-US') ? navLang : false
let lang = localLang || store.getters['app.lang'] || 'ru-RU'

Vue.locale = () => {}
const messages = {
    'ru-RU': Object.assign(ruRULocale, customRuRU),
    'en-US': Object.assign(enUsLocale, customEnUS)
}


const i18n = new VueI18n({
    locale: lang,
    messages
})

export default i18n