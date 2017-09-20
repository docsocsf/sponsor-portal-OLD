import React from 'react';
import '../../../style/students.scss';
import fetchWithConfig from '../../fetch';
import getJWTHeader from '../../jwt';
import StudentProfile from './StudentProfile';

const fetchUser = async () => {
  const headers = await getJWTHeader();
  return await fetchWithConfig('/students/api/user', { headers, json: true })
}

const fetchCV = async () => {
  const headers = await getJWTHeader();
  return await fetchWithConfig('/students/api/cv', { headers, json: true })
}

export default () => <StudentProfile fetchUser={fetchUser} fetchCV={fetchCV} />
