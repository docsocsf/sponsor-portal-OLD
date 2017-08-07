import React from 'react';
import ReactDOM from 'react-dom';
import '../style/students.scss';
import Students from 'Views/Students';
import fetchWithConfig from './fetch';
import getJWTHeader from './jwt';

const endpoint = "/students/auth/jwt/token";

const fetchUser = async () => {
  const headers = await getJWTHeader(endpoint);
  const resp = await fetchWithConfig('/students/api/user', { headers })
  return resp.json()
}

ReactDOM.render(<Students fetchUser={fetchUser}/>, document.getElementById("main"));
