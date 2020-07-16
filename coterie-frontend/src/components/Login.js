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
        <a href="http://localhost:3000/auth/google"><button className="btn btn-light">Google Login</button></a>
    </div>
  )
}

export default Login;