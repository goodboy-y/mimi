import axios from "axios";
import baseUrl from '../config/config'


const instance = axios.create({
  baseURL: process.env.NODE_ENV === 'development' ? baseUrl.dev.BACKEND_API : baseUrl.prod.BACKEND_API,
  headers: {
    'Content-Type': 'application/json;charset=UTF-8',
  },
})


instance.interceptors.response.use(
  (response) => {
    const { data } = response
    return data
  },
  (error) => {
    console.log(error);
  }
)

export default instance