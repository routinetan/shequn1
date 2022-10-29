import config from './config.js'
import axios from 'axios'
axios.defaults.baseURL = config.baseURL+"/apiv1/"
function Request(options){
    const instance = axios.create({
        url: options.url,
        method:options.method || 'GET',
        data:options.data,
        header: {
            "Content-Type": "application/x-www-form-urlencoded"
        }, // 请求头
        success(data) {
            res(data)
        },
        fail() {
            rej()
        }
    })
    return instance(options)
}
export default Request
