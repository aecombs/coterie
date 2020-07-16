
import React from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import logo from '../logo.png';
import Homepage from './Homepage';
import Login from './Login';
import Dashboard from './Dashboard';
// import ProtectedRoute from '../ProtectedRoute';
// import auth from './Auth';
// import axios from 'axios'

const Header = (props) => {

  

  return (
  <Router>
    <div className="">
      <nav className="navbar navbar-expand-lg navbar-light bg-light">
        <div className=" navbar-brand float-left w-50">
          <a href="/"><img className="w-25" alt="logo" src={logo}/></a>
        </div>
        <ul className=" navbar-nav mr-auto justify-content-end w-75">
          {/* TODO: Make dashboard only visible to auth'd users */}
          <li><Link to={'/dashboard'} className="nav-link float-right">Dashboard</Link></li>


          <li><Link to={'/login'} className="nav-link float-right">Login</Link></li>
          <li><Link to={'http://localhost:3000/logout'} className="nav-link float-right">Logout</Link></li>

          {/* using props */}
          {/* <li><Link to={'/'+props.buttonText.toLowerCase()} className="nav-link float-right">{props.buttonText}</Link></li> */}

          {/* Not using the fucking router holy fuck */}
          {/* <li><button className="nav-link float-right btn btn-info" onClick={props.buttonTextCallback}><a className="text-white text-decoration-none" href={props.buttonText === "Login" ? "/login" : "/logout"}>{props.buttonText}</a></button></li> */}

        </ul>
      </nav>
      
      <Switch> 
        <Route exact path='/' component={Homepage} />
        <Route path='/dashboard' component={Dashboard} />

        <Route path='/login' component={Login} />
        
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