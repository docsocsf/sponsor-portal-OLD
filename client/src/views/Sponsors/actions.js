import fetchWithConfig from '../../fetch';
import { onetimeToken, getToken, getJWTHeader } from '../../jwt';

export const fetchCVs = async () => {
  const headers = await getJWTHeader();
  return await fetchWithConfig('/sponsors/api/cvs', { headers, json: true })
}

export const downloadCV = async id => {
    const endpoint = `/sponsors/api/cv/${id}/download?token=`;
    const token = await getToken(onetimeToken);
    window.location.href = `${endpoint}${token}`;
}

