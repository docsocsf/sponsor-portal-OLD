import React from 'react';
import '../../../style/students.scss';
import fetchWithConfig from '../../fetch';
import { getJWTHeader } from '../../jwt';
import StudentProfile from './StudentProfile';

const fetchUser = async () => {
  const headers = await getJWTHeader();
  let resp = await fetchWithConfig('/api/students/user', { headers })
  return resp.body
}

const fetchCV = async () => {
  const headers = await getJWTHeader();
  let resp = await fetchWithConfig('/api/students/cv', { headers })
  return resp.body
}

export default () => <StudentProfile fetchUser={fetchUser} fetchCV={fetchCV} />
