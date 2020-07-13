import React from 'react';
import auth from '../Auth';

const Login = (props) => {
  const loginCallback = () => {
    auth.login(() => {
      props.history.push("/dashboard");
    })
    //probably add some junk to add info to LocalStorage
  }

  return (
    <div>
      <p className="baloo-da">Login</p>
      <div className="card">
        <button className="btn btn-primary" onClick={loginCallback}>
        Login Here!
        </button>

      </div>
    </div>
  )
}

export default Login;