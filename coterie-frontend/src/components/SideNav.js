import React from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import Organization from './Organization';
import Profile from './Profile';
import Announcements from './Announcements';
import Events from './Events';
import Holidays from './Holidays';
import Scriptures from './Scriptures';

const SideNav = (props) => {
  
  return (
    <Router>
    <div className="d-flex side-nav">
      <nav className="">
        <ul className="nav flex-column">
          <li className="nav-item nav-link text-decoration-none"><Link to={'/dashboard'} className="">Dashboard</Link></li>
          <li className="nav-item nav-link"><Link to={'/dashboard/profile'} className="">Profile</Link></li>
          <li className="nav-item nav-link"><Link to={'/dashboard/scriptures'} className="">Scriptures</Link></li>
          <li className="nav-item nav-link"><Link to={'/dashboard/holidays'} className="">Holidays</Link></li>
          <li className="nav-item nav-link"><Link to={'/dashboard/events'} className="">Events</Link></li>
          <li className="nav-item nav-link"><Link to={'/dashboard/announcements'} className="">Announcements</Link></li>
        </ul>
      </nav>
      <div className="dashboard-display">
        <Switch>
        <Route exact path='/dashboard' component={Organization} />
          <Route path='/dashboard/profile' component={Profile} />
          <Route path='/dashboard/announcements' component={Announcements} />
          <Route path='/dashboard/events' component={Events} />
          <Route path='/dashboard/holidays' component={Holidays} />
          <Route path='/dashboard/scriptures' component={Scriptures} />
            {/* <Route 
              path={'/'}
              render={() => (
              <Announcements 
              />
            )} /> */}
        </Switch>
      </div>
    </div>
  </Router>
  )
}

export default SideNav;