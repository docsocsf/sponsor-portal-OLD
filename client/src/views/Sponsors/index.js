import React from 'react';
import fetchWithConfig from '../../fetch';
import getJWTHeader from '../../jwt';
import SponsorProfile from './SponsorProfile';
import '../../../style/sponsors.scss';

const fetchCVs = async () => {
  const headers = await getJWTHeader();
  return await fetchWithConfig('/sponsors/api/cvs', { headers, json: true })
}

export default () => <SponsorProfile fetchCVs={fetchCVs} />;
