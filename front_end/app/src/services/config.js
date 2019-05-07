import axios from 'axios'

export default {
    url_lambda03: "https://kiyyewrc3f.execute-api.us-east-1.amazonaws.com/prod/lambda_info",

    http: axios.create({
        baseURL: 'http://localhost:3000/api/v1/',
    })

}




