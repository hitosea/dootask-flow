import utils from './utils'

const languageList = utils.languageTypes
const languageType = utils.getLanguage()
const languageRege = {}

// 添加语言
async function addLocales(locale) {
    locale = locale.toLowerCase()
    try {
        await import('./locales/key')
        await import(`./locales/${locale}.js`)
    } catch (error) {
        console.log(`${locale}.js的语言包不存在`)
    }
}

/**
 * 添加语言数据
 * @param data
 */
function addLanguage(data) {
    if (!$A.isArray(data)) {
        return
    }
    const keys = Object.assign(Object.keys(languageList))
    data.some(item => {
        let index = -1;
        item.key && keys.some(key => {
            const value = item[key] || item['general'] || null
            if (value && typeof window.LANGUAGE_DATA[key] !== "undefined") {
                index = window.LANGUAGE_DATA[key].push(value) - 1
            }
        })
        if (index > -1) {
            window.LANGUAGE_DATA['key'][item.key] = index
        }
    })
}

/**
 * 设置语言
 * @param language
 */
function setLanguage(language) {
    if (language === undefined) {
        return
    }
    $A.modalConfirm({
        content: '切换语言需要刷新后生效，是否确定刷新？',
        cancelText: '取消',
        okText: '确定',
        onOk: () => {
            window.localStorage.setItem("__language:type__", language)
            $A.reloadUrl()
        }
    })
}

/**
 * 转换语言
 * @param text
 * @returns {string|*}
 */
function switchLanguage(text) {
    if (typeof arguments[1] !== "undefined") {
        return switchLanguage(utils.replaceArgumentsLanguage(text, arguments))
    }
    if (typeof text !== "string" || !text) {
        return text
    }
    //
    if (typeof window.LANGUAGE_DATA === "undefined"
        || typeof window.LANGUAGE_DATA["key"] === "undefined"
        || typeof window.LANGUAGE_DATA[languageType] === "undefined") {
        return text
    }
    const index = window.LANGUAGE_DATA["key"][text] || -1
    if (index > -1) {
        return window.LANGUAGE_DATA[languageType][index] || text
    }
    if (typeof languageRege[text] === "undefined") {
        languageRege[text] = false
        for (let key in window.LANGUAGE_DATA["key"]) {
            if (key.indexOf("(*)") > -1) {
                const rege = new RegExp("^" + utils.replaceEscape(key) + "$", "g")
                if (rege.test(text)) {
                    let j = 0
                    const index = window.LANGUAGE_DATA["key"][key]
                    const value = (window.LANGUAGE_DATA[languageType][index] || key)?.replace(/\(\*\)/g, function () {
                        return "$" + (++j)
                    })
                    languageRege[text] = {rege, value}
                    break
                }
            }
        }
    }
    if (languageRege[text]) {
        return text.replace(languageRege[text].rege, languageRege[text].value)
    }
    return text
}

export { languageType, languageList, addLanguage, setLanguage, switchLanguage, addLocales}
