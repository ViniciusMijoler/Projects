import axios from 'axios';
import { backend } from '../util/constants.util';

export const postProject = (data: any) => {
    console.log(data)
    return axios.request({ method: 'post', baseURL: backend, url: `/project`, data, responseType: 'json' });
};