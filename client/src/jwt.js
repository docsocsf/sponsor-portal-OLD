import memoize from 'promise-memoize';
import fetchWithConfig from './fetch';

const getJWTHeader = async (endpoint, hdr) => {
  const token = await getToken(endpoint);

  return {...hdr, Authorization: `Bearer ${token}`}
}

const getToken = memoize(async (endpoint) => {
  const resp = await fetchWithConfig(endpoint);

  return resp.text;
}, {maxAge: 60000});

export default getJWTHeader;
