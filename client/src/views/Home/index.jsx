import React from 'react'
import {
  BrowserRouter as Router,
  Route
} from 'react-router-dom'

import HomeView from './HomeView';
import Login from './Login';

export default () => (
  <Router>
    <div>
      <Route exact path="/" component={HomeView}/>
      <Route path="/login" component={Login}/>
    </div>
  </Router>
)
