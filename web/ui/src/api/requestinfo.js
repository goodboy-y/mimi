import request from '@/utils/request'


export function fetchRequestInfos(data) {
    return request({
        url: '/requestInfo',
        method: 'get',
        data,
    })
}