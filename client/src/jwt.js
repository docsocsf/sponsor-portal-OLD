import memoize from 'promise-memoize';
import fetchWithConfig from './fetch';

const getJWTHeader = async (endpoint, hdr) => {
  const token = await getToken(endpoint);

  const headers = hdr || new Headers();
  headers.append('Authorization', `Bearer ${token}`);
  return headers;
}

const getToken = memoize(async (endpoint) => {
  const response = await fetchWithConfig(endpoint);

  return response.text();
}, {maxAge: 60000});

export default getJWTHeader;
