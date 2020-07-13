
import React, { useState } from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import logo from '../logo.png';
import Homepage from './Homepage';
import Login from './Login';
import Dashboard from './Dashboard';
// import axios from 'axios'

const Header = () => {
  const [loginButtonText, setLoginButtonText] = useState("Login");

  const toggleButtonText = () => {
    if (loginButtonText === "Login") {
      setLoginButtonText("Logout");
    } else {
      setLoginButtonText("Login");
    }
  }

  const loginButtonCallback = (text) => {
    setLoginButtonText(text);
  }

  return (
  <Router>
    <div>
      <nav className="navbar navbar-expand-lg navbar-light bg-light">
        <div className=" navbar-brand float-left w-50">
          <a href="/"><img className="w-25" alt="logo" src={logo}/></a>
        </div>
        <ul className=" navbar-nav mr-auto justify-content-end w-75">
          {/* TODO: Make dashboard only visible to auth'd users */}
          {/* <li><button onClick={ setLoginButtonText("Logout") }><Link to={'/dashboard'} className="nav-link float-right">Dashboard</Link></button></li> */}
          {/* { loginButtonText === "Login" ? <li><Link to={ loginButtonText === "Login" ? '/login' : '/' } className="nav-link float-right">{ loginButtonText }</Link></li> : <button onClick={toggleButtonText}>Logout</button>} */}

          <li><Link to={'/login'} className={ loginButtonText === "Login" ? "nav-link float-right" : "hidden"}>Login</Link></li>
          <li><Link to={'/'} className={ loginButtonText === "Logout" ? "nav-link float-right" : "hidden"}>Logout</Link></li>
        </ul>
      </nav>
      
      <Switch> 
        <Route exact path='/' component={Homepage} />
        <Route path='/dashboard' component={Dashboard} />
        {/* <Route path='/login' component={Login} /> */}
        
        {/* <Route 
          path={'/'+loginButtonText.toLowerCase()}
          component=
          { loginButtonText === "Login" ? Login: Homepage } 
        /> */}
        <Route 
        path={'/'+loginButtonText.toLowerCase()}
        render={ loginButtonText === "Login" ? (props) => (
        <Login {...props} 
          loginButtonCallback={loginButtonCallback}
          /> ) 
          : (props) => (
            <Homepage {...props} 
            />)
          }
        />
      </Switch>
    </div>
  </Router>
  );
}
  
  export default Header;