
import React, { useState } from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import logo from '../logo.png';
import Homepage from './Homepage';
import Login from './Login';
import Dashboard from './Dashboard';
// import axios from 'axios'

const Header = () => {
  const [loginButtonText, setLoginButtonText] = useState("Login");



  return (
  <Router>
    <div>
      <nav className="navbar navbar-expand-lg navbar-light bg-light">
        <div className=" navbar-brand float-left w-50">
          <a href="/"><img className="w-25" alt="logo" src={logo}/></a>
        </div>
        <ul className=" navbar-nav mr-auto justify-content-end w-75">
          {/* TODO: Make dashboard only visible to auth'd users */}
          <li><Link to={'/dashboard'} className="nav-link float-right">Dashboard</Link></li>
          <li><Link to={ loginButtonText === "login" ? "/login" : "/logout" } className="nav-link float-right">{ loginButtonText }</Link></li>
        </ul>
      </nav>
      
      <Switch> 
        <Route exact path='/' component={Homepage} />
        <Route path='/dashboard' component={Dashboard} />
        <Route path='/login' component={Login} />
        {/* <Route 
          path={'/'}
          render={() => (
          <Login 
          />
        )} /> */}
      </Switch>
    </div>
  </Router>
  );
}
  
  export default Header;