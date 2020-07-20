
import React from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import logo from '../images/logo.png';
import Homepage from './Homepage';
import Login from './Login';
import Dashboard from './Dashboard';

const Header = () => {

  const grabSessionValue = () => {
    const sessionCookie = document.cookie
    .split('; ')
    .find(row => row.startsWith('session'))
    console.log(sessionCookie)
    if (sessionCookie) {
      return sessionCookie.split('=')[1];
    } else {
      return undefined
    }
  }

  const sessionID = grabSessionValue()

  return (
  <Router>
    <div className="">
      <nav className="navbar navbar-expand-lg navbar-light bg-light">
        <div className=" navbar-brand float-left w-50">
          <a href="/"><img className="w-25" alt="logo" src={logo}/></a>
        </div>
        <ul className="navbar-nav mr-auto justify-content-end w-75">
          {/* TODO: Make dashboard only visible to auth'd users */}
          <li><Link to={'/dashboard'} className={sessionID ? "nav-link float-right" : "hidden"}>Dashboard</Link></li>


          <li><Link to={'/login'} className={sessionID ? "hidden" : "nav-link float-right"}>Login</Link></li>
          <li><a href={process.env.REACT_APP_BASE_URL+"/logout"} className={sessionID ? "nav-link float-right" : "hidden"}>Logout</a></li>
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