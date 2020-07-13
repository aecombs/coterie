import React from 'react';
import Auth from '../Auth';

const Login = (props) => {
  const login = () => {
    Auth.login()
    //probably add some junk to add info to LocalStorage
  }
  return (
    <div>
      <p>Login</p>
      <div className="card">
        <button className="" onClick=""></button>

      </div>
    </div>
  )
}

export default Login;