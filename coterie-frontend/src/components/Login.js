import React from 'react';
import auth from '../auth';

const Login = (props) => {
  // const login = () => {
  //   auth.login(() => {
  //     this.props.history.push("/dashboard");
  //   })
  //   //probably add some junk to add info to LocalStorage
  // }

  return (
    <div>
      <p>Login</p>
      <div className="card">
        <button className="btn btn-primary" onClick={() => {
          auth.login(() => {
            props.buttonTextCallback("Logout")
            props.history.push("/dashboard");
          })
          //probably add some junk to add info to LocalStorage
        }}>
        Login Here!
        </button>

      </div>
    </div>
  )
}

export default Login;