import React from 'react';
import favicon from '../coterie-favicon.png';
import googleLogo from '../google-logo.png';
import githubLogo from '../github-logo.png';
import fbLogo from '../fb-logo.png';
import auth from '../Auth';

const Login = (props) => {
  const loginCallback = () => {
    auth.login(() => {
      props.history.push("/dashboard");
    })
    //probably add some junk to add info to LocalStorage
  }

  return (
    <div className="d-flex text-center justify-content-center login-page">
      <div className="border-light card login-card">
        <img src={favicon} className="card-img-top favicon" alt="Coterie Logo"></img>
        <div className="card-body">
          <h2 className="card-title">Sign in</h2>
          <a href="http://localhost:3000/auth/google"><button className="btn btn-light w-100 open-sans my-3 text-muted"><img src={googleLogo} alt="Google Logo" className="pr-2"/>Continue with Google</button></a>
          <a href="/"><button className="btn btn-light w-100 open-sans my-3 text-muted"><img src={fbLogo} alt="Facebook Logo" className="pr-2"/>Continue with Facebook</button></a>
          <a href="/"><button className="btn btn-light w-100 open-sans my-3 text-muted"><img src={githubLogo} alt="Github Logo" className="pr-2"/>Continue with Github</button></a>
        </div>
      </div>
    </div>
  )
}

export default Login;