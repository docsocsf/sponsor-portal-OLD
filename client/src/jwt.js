import memoize from 'promise-memoize';
import fetchWithConfig from './fetch';

export const generalToken = "/auth/jwt"
export const onetimeToken = "/auth/jwt?single"

export const getJWTHeader = async (hdr) => {
  const token = await getToken(generalToken);

  return {...hdr, Authorization: `Bearer ${token}`}
}

export const getToken = memoize(async (endpoint) => {
  const resp = await fetchWithConfig(endpoint);
  return resp.text;
}, {maxAge: 60000});

