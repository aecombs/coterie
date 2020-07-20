
import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import logo from '../images/logo.png';
import Homepage from './Homepage';
import Login from './Login';
import Dashboard from './Dashboard';

const Header = () => {

  const selectURL = window.location.href
  .split('?')
  .find(row => row.startsWith('id'))

  let userID = undefined;

  if (selectURL) {
    userID = selectURL.split('=')[1][0];
    sessionStorage.setItem('userID', userID);
  }

  userID = sessionStorage.getItem('userID');


  const logoutUser = () => {
    if (userID) {
      sessionStorage.removeItem('userID');
      window.location.assign(`${process.env.REACT_APP_API_BASE_URL}/logout`);
    }
  };

  return (
  <Router>
    <div className="w-100">
      <nav className="navbar navbar-expand-lg navbar-light bg-light">
        <div className=" navbar-brand float-left w-50">
          <a href="/"><img className="w-25" alt="logo" src={logo}/></a>
        </div>
        <ul className="navbar-nav mr-auto justify-content-end w-75">
          {/* TODO: Make dashboard only visible to auth'd users */}
          <li><Link to={'/dashboard'} className={userID ? "nav-link float-right" : "hidden"}>Dashboard</Link></li>


          <li><Link to={'/login'} className={userID ? "hidden" : "nav-link float-right"}>Login</Link></li>

          {/* <li><a href={process.env.REACT_APP_API_BASE_URL+"/logout"} className={userID ? "nav-link float-right" : "hidden"}>Logout</a></li> */}
          <li><button onClick={logoutUser} className={userID ? "nav-link float-right border-0 btn" : "hidden"}>Logout</button></li>
        </ul>
      </nav>
      
      <Switch>
        <Route exact path='/' component={Homepage} />
        <Route path='/dashboard' component={Dashboard} />
        <Route path='/login' component={Login} />
        <Route exact path={process.env.REACT_APP_API_BASE_URL+"/logout"}  component={Homepage} />
      </Switch>
    </div>
  </Router>
  );
}
  
  export default Header;