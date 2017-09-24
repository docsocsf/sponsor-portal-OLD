import React from 'react';
import ReactDOM from 'react-dom';
import '../style/index.scss';
import Home from 'Views/Home';
import Students from 'Views/Students';
import Sponsors from 'Views/Sponsors';
import Login from 'Views/Login';
import {
  BrowserRouter as Router,
  Route
} from 'react-router-dom';

let root = (
  <Router>
    <div>
      <Route exact path="/" component={Home}/>
      <Route path="/login" component={Login}/>
      <Route exact path="/students/" component={Students}/>
      <Route exact path="/sponsors/" component={Sponsors}/>
    </div>
  </Router>
);

ReactDOM.render(root, document.getElementById("main"));
