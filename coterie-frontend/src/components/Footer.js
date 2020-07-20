import React from 'react';
import logo from '../images/coterie-favicon.png';

const Footer = () => {
  return (
<footer className="w-100 footer text-white pt-5">
  <div className="container">
    <div className="row">
      <div className="d-flex col">
        <div className="footer-logo text-left">
          <a href="/"><img className="w-75" src={logo} alt = "Coterie Logo"/></a>
        </div>
      </div>
      <div className="container text-left col">
        <h5 className="pb-1">About Coterie</h5>
        <a href="/"><p className="open-sans">Why Coterie</p></a>
        <a href="/"><p className="open-sans">What is a Cult Anyway?</p></a>
        <a href="/"><p className="open-sans">News</p></a>
        <a href="/"><p className="open-sans">Additional Reading</p></a>
      </div>
      <div className="container text-left col">
        <h5 className="pb-1">Workshops</h5>
        <a href="/"><p className="open-sans">Organizing the Organizers</p></a>
        <a href="/"><p className="open-sans">Contact Us</p></a>
        <a href="/"><p className="open-sans">Getting New Members</p></a>
        <a href="/"><p className="open-sans">Mixing Your Own Kool-Aid</p></a>
      </div>
      <div className="container text-left col">
        <h5 className="pb-1">Resources</h5>
        <a href="/"><p className="open-sans">How to Get Started</p></a>
        <a href="/"><p className="open-sans">So You've Decided to Be a Cult Leader</p></a>
        <a href="/"><p className="open-sans">Employee of the Month</p></a>
        <a href="/"><p className="open-sans">Videobook</p></a>
      </div>

      <div className="w-100"></div>
      <div className="col text-center pt-5">
        <p className="subtle-text px-2 small">© 2020 Coterie, All Rights Reserved</p>
      </div>
    </div>
  </div>
</footer>
  )
}

export default Footer;