import  config  from './config'
import axios from 'axios'

export default {

    

    Vote: (participant) => {
        return config.http.post('participants/vote', JSON.stringify(participant))
    },

    Get: () => {
        const instance = axios
        return instance.get(config.url_lambda03)
    }


}