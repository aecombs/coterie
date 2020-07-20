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
        <a href="/"><p>Careers</p></a>
        <a href="/"><p>Blog</p></a>
        <a href="/"><p>News</p></a>
        <a href="/"><p>Company Info</p></a>
      </div>
      <div className="container text-left col">
        <h5 className="pb-1">Workshops</h5>
        <a href="/"><p>Getting New Members</p></a>
        <a href="/"><p>asdf</p></a>
        <a href="/"><p>Rentals Policy</p></a>
        <a href="/"><p>Contact Us</p></a>
      </div>
      <div className="container text-left col">
        <h5 className="pb-1">Community</h5>
        <a href="/"><p>Announcements</p></a>
        <a href="/"><p>Picnic Schedule</p></a>
        <a href="/"><p>Employee of the Month</p></a>
        <a href="/"><p>Videobook</p></a>
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