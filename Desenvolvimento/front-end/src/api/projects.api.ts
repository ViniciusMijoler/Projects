import axios from 'axios';
import { backend } from '../util/constants.util';

export const postProject = (data: any) => {
    return axios.request({ method: 'post', baseURL: backend, url: `/project`, data });
};

export const getProjects = (params: any, companyId?: number) => {
    if (companyId) {
        return axios.request({ method: 'get', baseURL: backend, url: `/project/company/${companyId}`, params });
    } else {
        return axios.request({ method: 'get', baseURL: backend, url: `/project`, params });
    }
};

export const getProject = (projectId: number, personId: number) => {
    return axios.request({ method: 'get', baseURL: backend, url: `/project/${projectId}/person/${personId}` });
}