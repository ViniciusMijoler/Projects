import axios from 'axios';
import { backend } from '../util/constants.util';

// const headers = {
//     'Content-Type': 'application/x-www-form-urlencoded',
//     'Accept': 'application/json'
// }

export const postProject = (data: any) => {
    return axios.request({ method: 'post', baseURL: backend, url: `/project`, data });
};

export const getProjects = (params, companyId) => {
    return axios.request({ method: 'get', baseURL: backend, url: `/project/company/${companyId}`, params });
};