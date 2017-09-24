import React from 'react';
import ReactDOM from 'react-dom';
import '../style/students.scss';
import Students from 'Views/Students';
import fetchWithConfig from './fetch';
import { getJWTHeader } from './jwt';

const endpoint = "/students/auth/jwt/token";

const fetchUser = async () => {
  const headers = await getJWTHeader(endpoint);
  const resp = await fetchWithConfig('/students/api/user', { headers })

  return resp.body
}

const fetchCV = async () => {
  const headers = await getJWTHeader(endpoint);
  const resp = await fetchWithConfig('/students/api/cv', { headers })

  return resp.body
}

ReactDOM.render(<Students fetchUser={fetchUser} fetchCV={fetchCV}/>, document.getElementById("main"));
