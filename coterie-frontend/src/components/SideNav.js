import React from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import Profile from './Profile';
import Announcements from './Announcements';
import Events from './Events';
import Holidays from './Holidays';
import Scriptures from './Scriptures';

const SideNav = () => {
  return (
    <Router>
    <div className="">
      <nav className="">
        <ul className="">
          <li><Link to={'/dashboard'} className="">Dashboard</Link></li>
          <li><Link to={'/profile'} className="">Profile</Link></li>
          <li><Link to={'/scriptures'} className="">Scriptures</Link></li>
          <li><Link to={'/holidays'} className="">Holidays</Link></li>
          <li><Link to={'/events'} className="">Events</Link></li>
          <li><Link to={'/announcements'} className="">Announcements</Link></li>
        </ul>
      </nav>
      
      <Switch> 
        <Route exact path='/profile' component={Profile} />
        <Route exact path='/announcements' component={Announcements} />
        <Route exact path='/events' component={Events} />
        <Route exact path='/holidays' component={Holidays} />
        <Route exact path='/scriptures' component={Scriptures} />
        {/* <Route 
          path={'/'}
          render={() => (
          <Announcements 
          />
        )} /> */}
      </Switch>
    </div>
  </Router>
  )
}

export default SideNav;