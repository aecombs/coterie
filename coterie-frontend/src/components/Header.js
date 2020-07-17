
import React from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import logo from '../images/logo.png';
import Homepage from './Homepage';
import Login from './Login';
import Dashboard from './Dashboard';
import ProtectedRoute from '../ProtectedRoute';

const Header = (props) => {

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

  const userID = grabSessionValue()

  return (
  <Router>
    <div className="">
      <nav className="navbar navbar-expand-lg navbar-light bg-light">
        <div className=" navbar-brand float-left w-50">
          <a href="/"><img className="w-25" alt="logo" src={logo}/></a>
        </div>
        <ul className="navbar-nav mr-auto justify-content-end w-75">
          {/* TODO: Make dashboard only visible to auth'd users */}
          <li><Link to={'/dashboard'} className={userID ? "nav-link float-right" : "hidden"}>Dashboard</Link></li>


          <li><Link to={'/login'} className={userID ? "hidden" : "nav-link float-right"}>Login</Link></li>
          <li><a href="http://localhost:3000/logout" className={userID ? "nav-link float-right" : "hidden"}>Logout</a></li>
        </ul>
      </nav>
      
      <Switch> 
        <Route exact path='/' component={Homepage} />

        <ProtectedRoute 
        exact path="/dashboard"
        component={Dashboard}/>
        {/* <Route path='/dashboard' component={Dashboard} /> */}

        <Route path='/login' component={Login} />
        {/* <Route path='http://localhost:3000/logout' component={Homepage} /> */}
        <ProtectedRoute 
        exact path='http://localhost:3000/logout'
        component={Homepage}/>

        {/* <Route 
          path={'/'+loginButtonText.toLowerCase()}
          component=
          { loginButtonText === "Login" ? Login: Homepage } 
        /> */}
        {/* <Route 
        path={'/'+props.buttonText.toLowerCase()}
        render={ props.buttonText === "Login" ? (props) => (
        <Login {...props} 
          loginButtonCallback={props.buttonTextCallback}
          /> )
          : (props) => (
            <Homepage {...props} 
            />)
          }
        /> */}
      </Switch>
    </div>
  </Router>
  );
}
  
  export default Header;