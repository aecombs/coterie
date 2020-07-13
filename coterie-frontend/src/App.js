import React, { useState } from 'react';
import Header from './components/Header';
import Footer from './components/Footer';
import ReactNotification from 'react-notifications-component';
import 'react-notifications-component/dist/theme.css';

// function App() {
const App = () => {
  //change to check the Header or something.
  const [loginButtonText, setLoginButtonText] = useState("Login");

  const toggleLoginButtonText = () => {
    if (loginButtonText === "Login") {
      setLoginButtonText("Logout");
    } else {
      setLoginButtonText("Login");
    }
  }
  return (
    <section className="">
      <ReactNotification />
      <Header 
      buttonText={loginButtonText}
      buttonTextCallback={toggleLoginButtonText} />
      <Footer />
    </section>
  );
}

export default App;
