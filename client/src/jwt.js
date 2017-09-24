import memoize from 'promise-memoize';
import fetchWithConfig from './fetch';

export const generalToken = "/jwt/token"
export const onetimeToken = "/jwt/onetime-token"

export const getJWTHeader = async (hdr) => {
  const token = await getToken(generalToken);

  return {...hdr, Authorization: `Bearer ${token}`}
}

export const getToken = memoize(async (endpoint) => {
  const resp = await fetchWithConfig(endpoint);

  return resp;
}, {maxAge: 60000});

